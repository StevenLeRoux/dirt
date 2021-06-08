package dirt

import "github.com/prometheus/client_golang/prometheus"

var (
	registry *prometheus.Registry

	rttMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "rtt",
		Name:      "min",
		Help:      "Round Trip time latency",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	rttMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "rtt",
		Name:      "mean",
		Help:      "Round Trip time latency",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	rttMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "rtt",
		Name:      "median",
		Help:      "Round Trip time latency",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	rttMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "rtt",
		Name:      "max",
		Help:      "Round Trip time latency",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	rttStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "rtt",
		Name:      "stddev",
		Help:      "Round Trip time latency",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	senddelayMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_delay",
		Name:      "min",
		Help:      "Delay to send the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	senddelayMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_delay",
		Name:      "mean",
		Help:      "Delay to send the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	senddelayMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_delay",
		Name:      "median",
		Help:      "Delay to send the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	senddelayMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_delay",
		Name:      "max",
		Help:      "Delay to send the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	senddelayStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_delay",
		Name:      "stddev",
		Help:      "Delay to send the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	receivedelayMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_delay",
		Name:      "min",
		Help:      "Delay to receive the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	receivedelayMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_delay",
		Name:      "mean",
		Help:      "Delay to receive the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	receivedelayMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_delay",
		Name:      "median",
		Help:      "Delay to receive the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	receivedelayMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_delay",
		Name:      "max",
		Help:      "Delay to receive the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	receivedelayStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_delay",
		Name:      "stddev",
		Help:      "Delay to receive the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	ipdvMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "ipdv",
		Name:      "min",
		Help:      "Round Trip time latency",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	ipdvMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "ipdv",
		Name:      "mean",
		Help:      "Round Trip time latency",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	ipdvMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "ipdv",
		Name:      "median",
		Help:      "Round Trip time latency",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	ipdvMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "ipdv",
		Name:      "max",
		Help:      "Round Trip time latency",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	ipdvStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "ipdv",
		Name:      "stddev",
		Help:      "Round Trip time latency",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	sendipdvMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_ipdv",
		Name:      "min",
		Help:      "ipdv to send the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	sendipdvMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_ipdv",
		Name:      "mean",
		Help:      "ipdv to send the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	sendipdvMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_ipdv",
		Name:      "median",
		Help:      "ipdv to send the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	sendipdvMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_ipdv",
		Name:      "max",
		Help:      "ipdv to send the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	sendipdvStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_ipdv",
		Name:      "stddev",
		Help:      "ipdv to send the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	receiveipdvMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_ipdv",
		Name:      "min",
		Help:      "ipdv to receive the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	receiveipdvMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_ipdv",
		Name:      "mean",
		Help:      "ipdv to receive the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	receiveipdvMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_ipdv",
		Name:      "median",
		Help:      "ipdv to receive the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	receiveipdvMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_ipdv",
		Name:      "max",
		Help:      "ipdv to receive the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	receiveipdvStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_ipdv",
		Name:      "stddev",
		Help:      "ipdv to receive the payload",
	}, []string{"nodeSrc", "nodeDst", "rackSrc", "rackDst", "groupSrc", "groupDst"})

	clientError = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "dirt",
		Subsystem: "client",
		Name:      "error",
		Help:      "Errors encountered by clients",
	}, []string{"nodeSrc", "nodeDst"})

	clientRunCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "dirt",
		Subsystem: "client",
		Name:      "run_count",
		Help:      "Number of IRTT sessions run between a client and a server",
	}, []string{"nodeSrc", "nodeDst"})

	clientRunPkts = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "dirt",
		Subsystem: "client",
		Name:      "run_packets",
		Help:      "Number of IRTT packets per session",
	}, []string{"nodeSrc", "nodeDst"})
)

func registerMetrics() {
	registry.MustRegister(rttMin)
	registry.MustRegister(rttMean)
	registry.MustRegister(rttMedian)
	registry.MustRegister(rttMax)
	registry.MustRegister(rttStdDev)
	registry.MustRegister(senddelayMin)
	registry.MustRegister(senddelayMean)
	registry.MustRegister(senddelayMedian)
	registry.MustRegister(senddelayMax)
	registry.MustRegister(senddelayStdDev)
	registry.MustRegister(receivedelayMin)
	registry.MustRegister(receivedelayMean)
	registry.MustRegister(receivedelayMedian)
	registry.MustRegister(receivedelayMax)
	registry.MustRegister(receivedelayStdDev)
	registry.MustRegister(ipdvMin)
	registry.MustRegister(ipdvMean)
	registry.MustRegister(ipdvMedian)
	registry.MustRegister(ipdvMax)
	registry.MustRegister(ipdvStdDev)
	registry.MustRegister(sendipdvMin)
	registry.MustRegister(sendipdvMean)
	registry.MustRegister(sendipdvMedian)
	registry.MustRegister(sendipdvMax)
	registry.MustRegister(sendipdvStdDev)
	registry.MustRegister(receiveipdvMin)
	registry.MustRegister(receiveipdvMean)
	registry.MustRegister(receiveipdvMedian)
	registry.MustRegister(receiveipdvMax)
	registry.MustRegister(receiveipdvStdDev)

	registry.MustRegister(clientError)
	registry.MustRegister(clientRunCount)
	registry.MustRegister(clientRunPkts)
	//registry.MustRegister()
}
