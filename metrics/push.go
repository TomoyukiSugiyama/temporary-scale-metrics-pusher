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
	jobName        string
}

func NewMetrics(pushgatewayUrl string, tsm TemporaryScaleMetrics) Metrics {
	temporaryScaleGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "temporary_scale",
		Help:        "temporary scale",
		ConstLabels: prometheus.Labels{"id": tsm.ConditionId, "type": tsm.ConditionType},
	})
	jobName := "temporary_scale_job_" + tsm.ConditionId + "_" + tsm.ConditionType
	p := &pusher{
		pushgatewayUrl: pushgatewayUrl,
		tsm:            tsm,
		gauge:          temporaryScaleGauge,
		jobName:        jobName,
	}

	return p
}

func (p *pusher) Push() {

}
