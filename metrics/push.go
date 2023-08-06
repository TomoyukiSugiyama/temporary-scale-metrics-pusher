package metrics

import (
	prometheusv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	Push()
}

type metrics struct {
	temporaryScaleGauge prometheus.GaugeVec
	prometheusApi       prometheusv1.API
}

func NewMetrics(prometheusGauge prometheus.GaugeVec) Metrics {

	m := &metrics{temporaryScaleGauge: prometheusGauge}
	return m
}

func (m *metrics) Push() {

}
