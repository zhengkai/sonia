package metrics

import "github.com/prometheus/client_golang/prometheus"

func newCounter(name, help string) prometheus.Counter {
	return prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
	)
}
func newSummary(name, help string) prometheus.Summary {
	return prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       name,
		Help:       help,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
}

func newGauge(name, help string) prometheus.Gauge {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	})
}
