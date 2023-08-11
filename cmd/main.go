package main

import (
	"flag"
	"fmt"
	"net/url"

	metricspkg "github.com/TomoyukiSugiyama/temporary-scale-metrics-pusher/metrics"
)

func main() {
	var pushgatewayAddr string
	var pushgatewayPort string
	var conditionId string
	var conditionType string
	var duration string
	const metricValue = "0"
	flag.StringVar(&pushgatewayAddr, "pushgateway-address", "localhost", "ip address for pushgateway")
	flag.StringVar(&pushgatewayPort, "pushgateway-port", "9091", "port number for pushgateway")
	flag.StringVar(&conditionId, "condition-id", "9-22", "unique id for the temporary metric")
	flag.StringVar(&conditionType, "condition-type", "training", "condition type")
	flag.StringVar(&duration, "duration", "9-22", "duration")
	flag.Parse()

	tsm := metricspkg.TemporaryScaleMetrics{
		ConditionId:   conditionId,
		ConditionType: conditionType,
		Duration:      duration,
		MetricValue:   metricValue,
	}
	pusher := metricspkg.NewMetrics(pushgatewayUrl(pushgatewayAddr, pushgatewayPort), tsm)
	pusher.Push()
}

func pushgatewayUrl(address string, port string) string {
	httpURL := &url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s:%s", address, port),
	}
	return httpURL.String()
}
