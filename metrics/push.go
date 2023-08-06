package metrics

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
	pushgatewayUrl        string
	temporaryScaleMetrics TemporaryScaleMetrics
}

func NewMetrics(pushgatewayUrl string, tsm TemporaryScaleMetrics) Metrics {

	p := &pusher{
		pushgatewayUrl:        pushgatewayUrl,
		temporaryScaleMetrics: tsm,
	}

	return p
}

func (p *pusher) Push() {

}

func (p *pusher) jobName() string {
	return "temporary_scale_job_" + p.temporaryScaleMetrics.ConditionId + "_" + p.temporaryScaleMetrics.ConditionType
}
