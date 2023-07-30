package main

import "flag"

func main() {
	var pushgatewayAddr string
	var pushgatewayPort string
	flag.StringVar(&pushgatewayAddr, "pushgateway-address", "localhost", "ip address for pushgateway")
	flag.StringVar(&pushgatewayPort, "pushgateway-port", "9091", "port number for pushgateway")
	flag.Parse()
}
