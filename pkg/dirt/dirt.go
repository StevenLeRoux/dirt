package dirt

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"

	"github.com/StevenLeRoux/dirt/pkg/cluster"
	mod "github.com/StevenLeRoux/dirt/pkg/mod"
	"github.com/StevenLeRoux/dirt/pkg/utils"
	"github.com/StevenLeRoux/promflush"
	"github.com/hashicorp/memberlist"
	"github.com/heistp/irtt"
)

type Node struct {
	Config   *mod.Config
	sessions map[string]chan struct{}
	Events   chan cluster.Event
	Quit     chan struct{}
}

func Create(r *prometheus.Registry, c *mod.Config) (*Node, error) {
	
	registry = r
	registerMetrics()

	ch := make(chan cluster.Event)

	n := &Node{}
	n.Config = c
	n.Events = ch
	n.sessions = make(map[string]chan struct{})
	n.Quit = make(chan struct{})

	return n, nil

}

func (n *Node) register(e cluster.Event) {
	log.Debug("Dirt: registering node: " + e.Node.Name)
	c := make(chan struct{})
	n.sessions[e.Node.Name] = c
	go n.runClient(e.Node, c)
}

func (n *Node) unregister(e cluster.Event) {
	log.Debug("Dirt: unregistering  node: " + e.Node.Name)
	n.sessions[e.Node.Name] <- struct{}{}
	delete(n.sessions, e.Node.Name)
}

func (n *Node) Run() {

	if !n.Config.Bootstrap {
		//Each Dirt instance run a server which can be reached by any client.
		go func() {
			listen := fmt.Sprintf("%s:%d", n.Config.Server.Bind, n.Config.Server.Port)

			log.Debug("Start IRTT Server")
			serverConfig := irtt.NewServerConfig()
			serverConfig.Addrs = []string{listen}
			serverConfig.AllowStamp = irtt.DualStamps
			serverConfig.MinInterval = 0
			serverConfig.MaxLength = 0
			serverConfig.IPVersion = irtt.IPv4

			server := irtt.NewServer(serverConfig)
			server.ListenAndServe()
		}()

		go func() {
			for {
				select {
				case e := <-n.Events:
					if e.Event == cluster.Join {
						n.register(e)
					} else if e.Event == cluster.Leave {
						n.unregister(e)
					}
					break
				}
			}
		}()
	}

	cluster.Run(registry, n.Quit, n.Config, n.Events)

}

func (n *Node) Close() {
	n.Quit <- struct{}{}
}

func (n *Node) runClient(server *memberlist.Node, quit chan struct{}) {
	ctx, _ := context.WithCancel(context.Background())

	m, err := cluster.DecodeMeta(server.Meta)
	if err != nil {
		log.Info("Dirt: runClient: Unable to decode Meta for " + server.Name)
		return
	}

	laddr := utils.LookupOutboundIP(server.Address())

	// Loop over client to server IRTT sessions
	for {
		select {
		case <-quit:
			return
		default:
			log.Debug("Dirt: irtt session start with: " + n.Config.Name + " -> " + server.Name)
			clientConfig := irtt.NewClientConfig()
			clientConfig.Duration = 5 * time.Second
			clientConfig.LocalAddress = fmt.Sprintf("%s:0", laddr.String())
			clientConfig.RemoteAddress = fmt.Sprintf("%s:%d", server.Addr.String(), m.Port)

			// Run IRTT session
			c := irtt.NewClient(clientConfig)

			// Get results
			r, err := c.Run(ctx)
			if err != nil {
				log.Info(err)
				clientError.WithLabelValues(n.Config.Name, server.Name).Inc()
			} else {
				log.Debug("Dirt: irtt session finished" + n.Config.Name + " -> " + server.Name)
				ts := time.Now().UnixNano()
				n.handleStats(ts, server.Name, m, r)
			}
		}
	}
	// clientRunCount
}

func (n *Node) handleStats(t int64, serverName string, m cluster.Meta, r *irtt.Result) {

	//ss := r.SendCallStats
	//tes := r.TimerErrorStats
	//sps := r.ServerProcessingTimeStats

	clientRunCount.WithLabelValues(n.Config.Name, serverName).Inc()
	clientRunPkts.WithLabelValues(n.Config.Name, serverName).Add(float64(r.RTTStats.N))

	rttMin.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.RTTStats.Min))
	rttMean.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.RTTStats.Mean()))
	rttMedian.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(utils.Ok(r.RTTStats.Median())))
	rttMax.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.RTTStats.Max))
	rttStdDev.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.RTTStats.Stddev()))

	senddelayMin.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.SendDelayStats.Min))
	senddelayMean.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.SendDelayStats.Mean()))
	senddelayMedian.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(utils.Ok(r.SendDelayStats.Median())))
	senddelayMax.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.SendDelayStats.Max))
	senddelayStdDev.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.SendDelayStats.Stddev()))

	receivedelayMin.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.ReceiveDelayStats.Min))
	receivedelayMean.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.ReceiveDelayStats.Mean()))
	receivedelayMedian.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(utils.Ok(r.ReceiveDelayStats.Median())))
	receivedelayMax.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.ReceiveDelayStats.Max))
	receivedelayStdDev.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.ReceiveDelayStats.Stddev()))

	ipdvMin.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.RoundTripIPDVStats.Min))
	ipdvMean.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.RoundTripIPDVStats.Mean()))
	ipdvMedian.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(utils.Ok(r.RoundTripIPDVStats.Median())))
	ipdvMax.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.RoundTripIPDVStats.Max))
	ipdvStdDev.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.RoundTripIPDVStats.Stddev()))

	sendipdvMin.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.SendIPDVStats.Min))
	sendipdvMean.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.SendIPDVStats.Mean()))
	sendipdvMedian.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(utils.Ok(r.SendIPDVStats.Median())))
	sendipdvMax.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.SendIPDVStats.Max))
	sendipdvStdDev.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.SendIPDVStats.Stddev()))

	receiveipdvMin.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.ReceiveIPDVStats.Min))
	receiveipdvMean.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.ReceiveIPDVStats.Mean()))
	receiveipdvMedian.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(utils.Ok(r.ReceiveIPDVStats.Median())))
	receiveipdvMax.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.ReceiveIPDVStats.Max))
	receiveipdvStdDev.WithLabelValues(n.Config.Name, serverName, n.Config.Rack, m.Rack, n.Config.Group, m.Group).Set(float64(r.ReceiveIPDVStats.Stddev()))

	ts := strconv.Itoa(int(t / 1e3))
	f := "/opt/beamium/sources/dirt-" + ts + ".metrics"

	err := promflush.WriteToTextfile(ts, f, registry)
	if err != nil {
		log.Error("Unable to flush metrics to file: ", err)
	}

	log.Debug("Dirt: irtt metrics flush for node: " + n.Config.Name + " -> " + serverName)
}
