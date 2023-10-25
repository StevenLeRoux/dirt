package cluster

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	"github.com/hashicorp/memberlist"
	"github.com/prometheus/client_golang/prometheus"

	mod "github.com/StevenLeRoux/dirt/pkg/mod"
	log "github.com/sirupsen/logrus"
)

var (
	numMembers = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "cluster",
		Name:      "members",
		Help:      "Number of members in the memberlist cluster",
	}, []string{"group"})
)

type EventType int

const (
	Join EventType = iota
	Leave
)

type Event struct {
	Event EventType
	Node  *memberlist.Node
}

type Meta struct {
	Name  string
	Group string
	Rack  string
	Port  uint16
}

type cluster struct {
	meta          *Meta
	allowedGroups []string
	members       map[string]*memberlist.Node
	events        chan Event
}

func (c *cluster) addMember(n *memberlist.Node) {
	log.Debug("Dirt: Cluster: addMember: " + n.Name)
	c.members[n.Name] = n
	ev := Event{}
	ev.Event = Join
	ev.Node = n
	c.events <- ev
}

func (c *cluster) delMember(n *memberlist.Node) {
	log.Debug("Dirt: Cluster: delMember: " + n.Name)
	delete(c.members, n.Name)
	ev := Event{}
	ev.Event = Leave
	ev.Node = n
	c.events <- ev
}

func (c *cluster) isAllowed(s string) bool {
	for _, g := range c.allowedGroups {
		if g == s {
			return true
		}
	}
	return false
}

// Cluster struct implements the Memberlist Delegate interface
func (c *cluster) NodeMeta(limit int) []byte {
	log.Debug("Dirt: Cluster: NodeMeta: ")
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&c.meta)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	log.Debug("Dirt: Cluster: NodeMeta: buf:" + buf.String())
	return buf.Bytes()
}

func (c *cluster) NotifyMsg(b []byte) {}

func (c *cluster) NotifyJoin(n *memberlist.Node) {
	log.Debug("Dirt: Cluster: Join: " + n.Name)
	m, err := DecodeMeta(n.Meta)
	if err != nil {
		log.Info("Dirt: Cluster: Update: Unable to decode Meta for " + n.Name)
		return
	}
	_, member := c.members[n.Name]
	g := string(m.Group)
	log.Debug("Dirt: Cluster: Join: meta.Group: " + g)

	if n.Name == c.meta.Name {
		log.Debug("Dirt: self ignoring")
		return
	}

	if !member && c.isAllowed(g) {
		log.Debug("Dirt: Cluster: Accepted: " + n.Name)
		c.addMember(n)
	} else {
		log.Debug("Dirt: Cluster: Rejected: " + n.Name)
	}
}

func (c *cluster) NotifyLeave(n *memberlist.Node) {
	log.Debug("Dirt: Cluster: Leave: " + n.Name)
	_, member := c.members[n.Name]

	if member {
		log.Debug("Dirt: Cluster: isMember: " + n.Name)
		c.delMember(n)
	} else {
		log.Debug("Dirt: Cluster: ignored: " + n.Name)
	}
}

func (c *cluster) NotifyUpdate(n *memberlist.Node) {
	log.Debug("Dirt: Cluster: Update: " + n.Name)
	m, err := DecodeMeta(n.Meta)
	if err != nil {
		log.Info("Dirt: Cluster: Update: Unable to decode Meta for " + n.Name)
		return
	}
	_, member := c.members[n.Name]
	g := string(m.Group)
	log.Debug("Dirt: Cluster: Join: meta.Group: " + g)

	if n.Name == c.meta.Name {
		log.Debug("Dirt: self ignoring")
		return
	}

	if member {
		// Node was member but not allowed anymore. Remove it
		if !c.isAllowed(g) {
			log.Debug("Dirt: Cluster: member not allowed anymore: " + n.Name)
			c.delMember(n)
		}
	} else {
		// Node not member yet, but allowed now. Add it
		if c.isAllowed(g) {
			log.Debug("Dirt: Cluster: member newly allowed: " + n.Name)
			c.addMember(n)
		}
	}
}

func (c *cluster) GetBroadcasts(overhead, limit int) [][]byte { return [][]byte{} }

func (c *cluster) LocalState(join bool) []byte { return []byte{} }

func (c *cluster) MergeRemoteState(buf []byte, join bool) {}

func Run(r *prometheus.Registry, quit chan struct{}, cfg *mod.Config, ch chan Event) {

	r.MustRegister(numMembers)

	m := &Meta{}
	m.Group = cfg.Group
	m.Rack = cfg.Rack
	m.Name = cfg.Name
	m.Port = cfg.Server.Port

	c := &cluster{}
	c.meta = m

	c.allowedGroups = cfg.Peers
	c.members = make(map[string]*memberlist.Node)
	c.events = ch

	config := memberlist.DefaultLANConfig()
	config.Name = cfg.Name
	config.AdvertiseAddr = cfg.Discovery.AdvAddress
	config.AdvertisePort = cfg.Discovery.AdvPort
	config.BindAddr = cfg.Discovery.Bind
	config.BindPort = cfg.Discovery.Port
	config.Delegate = c

	if !cfg.Bootstrap {
		config.Events = c
	}

	ml, err := memberlist.Create(config)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}
	defer ml.Shutdown()

	go func(g string) {
		for {
			n := ml.NumMembers()
			numMembers.WithLabelValues(g).Set(float64(n))
			time.Sleep(5 * time.Second)
		}
	}(m.Group)

	if !cfg.Bootstrap {
		// Join the bootstrap node
		_, err = ml.Join([]string{cfg.Join})
		if err != nil {
			log.Panic("Failed to join cluster: " + err.Error())
		}
	}

	<-quit
}

func DecodeMeta(b []byte) (Meta, error) {
	r := bytes.NewReader(b)
	dec := gob.NewDecoder(r)
	var m Meta
	err := dec.Decode(&m)
	if err != nil {
		return m, errors.New("Dirt: Cluster: Meta decode error")
	}
	return m, nil
}
