package main

import (
	"flag"

	metricspkg "github.com/TomoyukiSugiyama/temporary-scale-metrics-pusher/metrics"
)

func main() {
	var pushgatewayAddr string
	var pushgatewayPort string
	flag.StringVar(&pushgatewayAddr, "pushgateway-address", "localhost", "ip address for pushgateway")
	flag.StringVar(&pushgatewayPort, "pushgateway-port", "9091", "port number for pushgateway")
	flag.Parse()

	tsm := metricspkg.TemporaryScaleMetrics{
		ConditionId:   "9-22",
		ConditionType: "training",
		Duration:      "9-22",
		MetricValue:   "0",
	}
	pusher := metricspkg.NewMetrics(pushgatewayUrl(pushgatewayAddr, pushgatewayPort), tsm)
	pusher.Push()
}

func pushgatewayUrl(address string, port string) string {
	return "http://" + address + ":" + port
}
