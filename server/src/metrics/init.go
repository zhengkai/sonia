package metrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(connectorReq)
	prometheus.MustRegister(connectorRsp)
	prometheus.MustRegister(connectorFail)
}
