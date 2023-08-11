package metrics

import (
	"strings"
	"time"

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
	currentTime    time.Time
	pushgatewayUrl string
	tsm            TemporaryScaleMetrics
	gauge          prometheus.Gauge
	jobName        string
}

type PusherOption func(*pusher)

func NewMetrics(pushgatewayUrl string, tsm TemporaryScaleMetrics, opts ...PusherOption) Metrics {
	const jobNamePrefix = "temporary_scale_job"
	temporaryScaleGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "temporary_scale",
		Help:        "temporary scale",
		ConstLabels: prometheus.Labels{"id": tsm.ConditionId, "type": tsm.ConditionType},
	})

	jobName := strings.Join([]string{jobNamePrefix, tsm.ConditionId, tsm.ConditionType}, "_")

	p := &pusher{
		pushgatewayUrl: pushgatewayUrl,
		tsm:            tsm,
		gauge:          temporaryScaleGauge,
		jobName:        jobName,
	}

	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *pusher) WithDate(year int, month time.Month, day int) PusherOption {
	return func(p *pusher) {
		p.currentTime = time.Date(year, month, day, 9, 0, 0, 0, time.Local)
	}
}

func (p *pusher) Push() {

}
