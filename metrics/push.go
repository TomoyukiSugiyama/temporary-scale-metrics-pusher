package metrics

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type Metrics interface {
	Push()
}

type TemporaryScaleMetrics struct {
	ConditionId   string
	ConditionType string
	Duration      string
	MetricValue   float64
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
	currentTime := time.Now()

	temporaryScaleGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "temporary_scale",
		Help:        "temporary scale",
		ConstLabels: prometheus.Labels{"id": tsm.ConditionId, "type": tsm.ConditionType},
	})

	jobName := strings.Join([]string{jobNamePrefix, tsm.ConditionId, tsm.ConditionType}, "_")

	p := &pusher{
		currentTime:    currentTime,
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

func WithDate(year int, month time.Month, day int, hour int) PusherOption {
	return func(p *pusher) {
		p.currentTime = time.Date(year, month, day, hour, 0, 0, 0, time.Local)
	}
}

func (p *pusher) calcurateMetricValue() error {
	timeRenge := strings.Split(p.tsm.Duration, "-")
	if len(timeRenge) != 2 {
		return errors.New("id is invalid format")
	}
	min, err := strconv.Atoi(timeRenge[0])
	if err != nil {
		return errors.New("duration is invalid format")
	}
	max, err := strconv.Atoi(timeRenge[1])
	if err != nil {
		return errors.New("duration is invalid format")
	}
	if min <= p.currentTime.Hour() && p.currentTime.Hour() <= max {
		p.tsm.MetricValue = 1
		return nil
	}
	p.tsm.MetricValue = 0
	return nil
}

func (p *pusher) Push() {
	if err := push.New(p.pushgatewayUrl, p.jobName).
		Collector(p.gauge).
		Push(); err != nil {
		fmt.Printf(err.Error())
	}
}
