package main

import (
	"flag"

	metricspkg "github.com/TomoyukiSugiyama/temporary-scale-metrics-pusher/metrics"
	prometheusapi "github.com/prometheus/client_golang/api"
	prometheusv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func main() {
	var pushgatewayAddr string
	var pushgatewayPort string
	flag.StringVar(&pushgatewayAddr, "pushgateway-address", "localhost", "ip address for pushgateway")
	flag.StringVar(&pushgatewayPort, "pushgateway-port", "9091", "port number for pushgateway")
	flag.Parse()

	client, err := prometheusapi.NewClient(prometheusapi.Config{Address: pushgatewayUrl(pushgatewayAddr, pushgatewayPort)})
	if err != nil {

	}
	api := prometheusv1.NewAPI(client)
	pusher := metricspkg.NewMetrics(api)
	pusher.Push()
}

func pushgatewayUrl(address string, port string) string {
	return "http://" + address + ":" + port
}
