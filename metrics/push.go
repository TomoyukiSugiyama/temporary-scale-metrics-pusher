package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	Push()
}

type metrics struct {
	temporaryScaleGauge prometheus.GaugeVec
}

func NewMetrics(prometheusGauge prometheus.GaugeVec) Metrics {

	m := &metrics{temporaryScaleGauge: prometheusGauge}
	return m
}

func (m *metrics) Push() {

}
