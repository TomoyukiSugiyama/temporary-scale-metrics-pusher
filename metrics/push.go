package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	Push()
}

type TemporaryScaleMetrics struct {
	ConditionId   string
	ConditionType string
	Duration      string
	MetricValue   string
}

type pusher struct {
	pushgatewayUrl string
	tsm            TemporaryScaleMetrics
	gauge          prometheus.Gauge
}

func NewMetrics(pushgatewayUrl string, tsm TemporaryScaleMetrics) Metrics {
	temporaryScaleGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "temporary_scale",
		Help:        "temporary scale",
		ConstLabels: prometheus.Labels{"id": tsm.ConditionId, "type": tsm.ConditionType},
	})

	p := &pusher{
		pushgatewayUrl: pushgatewayUrl,
		tsm:            tsm,
		gauge:          temporaryScaleGauge,
	}

	return p
}

func (p *pusher) Push() {

}

func (p *pusher) jobName() string {
	return "temporary_scale_job_" + p.tsm.ConditionId + "_" + p.tsm.ConditionType
}
