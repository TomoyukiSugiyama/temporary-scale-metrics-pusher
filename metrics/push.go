package metrics

import (
	prometheusv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type Metrics interface {
	Push()
}

type TemporaryScaleMetrics struct {
	conditionId   string
	conditionType string
	duration      string
}

type pusher struct {
	prometheusApi         prometheusv1.API
	temporaryScaleMetrics TemporaryScaleMetrics
}

func NewMetrics(prometheusApi prometheusv1.API) Metrics {
	tsm := TemporaryScaleMetrics{
		conditionId:   "9-22",
		conditionType: "training",
		duration:      "9-22",
	}

	p := &pusher{prometheusApi: prometheusApi,
		temporaryScaleMetrics: tsm,
	}
	return p
}

func (p *pusher) Push() {

}
