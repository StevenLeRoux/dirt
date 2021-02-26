package dirt

import "github.com/prometheus/client_golang/prometheus"

var (
	rttMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "rtt",
		Name:      "min",
		Help:      "Round Trip time latency",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	rttMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "rtt",
		Name:      "mean",
		Help:      "Round Trip time latency",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	rttMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "rtt",
		Name:      "median",
		Help:      "Round Trip time latency",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	rttMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "rtt",
		Name:      "max",
		Help:      "Round Trip time latency",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	rttStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "rtt",
		Name:      "stddev",
		Help:      "Round Trip time latency",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	senddelayMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_delay",
		Name:      "min",
		Help:      "Delay to send the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	senddelayMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_delay",
		Name:      "mean",
		Help:      "Delay to send the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	senddelayMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_delay",
		Name:      "median",
		Help:      "Delay to send the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	senddelayMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_delay",
		Name:      "max",
		Help:      "Delay to send the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	senddelayStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_delay",
		Name:      "stddev",
		Help:      "Delay to send the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	receivedelayMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_delay",
		Name:      "min",
		Help:      "Delay to receive the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	receivedelayMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_delay",
		Name:      "mean",
		Help:      "Delay to receive the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	receivedelayMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_delay",
		Name:      "median",
		Help:      "Delay to receive the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	receivedelayMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_delay",
		Name:      "max",
		Help:      "Delay to receive the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	receivedelayStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_delay",
		Name:      "stddev",
		Help:      "Delay to receive the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	ipdvMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "ipdv",
		Name:      "min",
		Help:      "Round Trip time latency",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	ipdvMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "ipdv",
		Name:      "mean",
		Help:      "Round Trip time latency",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	ipdvMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "ipdv",
		Name:      "median",
		Help:      "Round Trip time latency",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	ipdvMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "ipdv",
		Name:      "max",
		Help:      "Round Trip time latency",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	ipdvStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "ipdv",
		Name:      "stddev",
		Help:      "Round Trip time latency",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	sendipdvMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_ipdv",
		Name:      "min",
		Help:      "ipdv to send the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	sendipdvMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_ipdv",
		Name:      "mean",
		Help:      "ipdv to send the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	sendipdvMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_ipdv",
		Name:      "median",
		Help:      "ipdv to send the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	sendipdvMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_ipdv",
		Name:      "max",
		Help:      "ipdv to send the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	sendipdvStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "send_ipdv",
		Name:      "stddev",
		Help:      "ipdv to send the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	receiveipdvMin = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_ipdv",
		Name:      "min",
		Help:      "ipdv to receive the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	receiveipdvMean = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_ipdv",
		Name:      "mean",
		Help:      "ipdv to receive the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	receiveipdvMedian = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_ipdv",
		Name:      "median",
		Help:      "ipdv to receive the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	receiveipdvMax = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_ipdv",
		Name:      "max",
		Help:      "ipdv to receive the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})

	receiveipdvStdDev = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dirt",
		Subsystem: "receive_ipdv",
		Name:      "stddev",
		Help:      "ipdv to receive the payload",
	}, []string{"src", "dst", "srcRack", "dstRack", "srcGroup", "dstGroup"})
)

func init() {
	prometheus.MustRegister(rttMin)
	prometheus.MustRegister(rttMean)
	prometheus.MustRegister(rttMedian)
	prometheus.MustRegister(rttMax)
	prometheus.MustRegister(rttStdDev)
	prometheus.MustRegister(senddelayMin)
	prometheus.MustRegister(senddelayMean)
	prometheus.MustRegister(senddelayMedian)
	prometheus.MustRegister(senddelayMax)
	prometheus.MustRegister(senddelayStdDev)
	prometheus.MustRegister(receivedelayMin)
	prometheus.MustRegister(receivedelayMean)
	prometheus.MustRegister(receivedelayMedian)
	prometheus.MustRegister(receivedelayMax)
	prometheus.MustRegister(receivedelayStdDev)
	prometheus.MustRegister(ipdvMin)
	prometheus.MustRegister(ipdvMean)
	prometheus.MustRegister(ipdvMedian)
	prometheus.MustRegister(ipdvMax)
	prometheus.MustRegister(ipdvStdDev)
	prometheus.MustRegister(sendipdvMin)
	prometheus.MustRegister(sendipdvMean)
	prometheus.MustRegister(sendipdvMedian)
	prometheus.MustRegister(sendipdvMax)
	prometheus.MustRegister(sendipdvStdDev)
	prometheus.MustRegister(receiveipdvMin)
	prometheus.MustRegister(receiveipdvMean)
	prometheus.MustRegister(receiveipdvMedian)
	prometheus.MustRegister(receiveipdvMax)
	prometheus.MustRegister(receiveipdvStdDev)
	//prometheus.MustRegister()
}
