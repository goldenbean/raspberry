package main

//https://medium.com/@nabil.servais/monitoring-the-temperature-of-your-room-with-prometheus-and-timescaledb-4b67eb9a193e
import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/yryz/ds18b20"
)

var (
	gatewayUrl = "http://push-gateway:9091"
	tempGauge  = prometheus.NewGauge(prometheus.GaugeOpts{Name: "temperature", Help: "Température en °C"})
)

func publish(temps chan float64) {
	var t float64
	t = <-temps
	fmt.Printf("temperature channel: %.2f°C\n", t)
	tempGauge.Set(t)
	err := push.Collectors("chambre_bebe_job", push.HostnameGroupingKey(), gatewayUrl, tempGauge)
	if err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}

func sensors() {

	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sensor IDs: %v\n", sensors)

	for _, sensor := range sensors {
		if err == nil {
			var temps = make(chan float64, 10)
			for {
				t, err := ds18b20.Temperature(sensor)
				if err != nil {
					fmt.Println(err)
				}
				temps <- t
				go publish(temps)
				time.Sleep(time.Minute)
			}
		}
	}
}

// 获取温度
// cpu=$(echo /sys/class/thermal/thermal_zone0/temp)
// echo "$(date) @ $(hostname)"
// echo "-------------------------------------------"
// echo "GPU => $(/opt/vc/bin/vcgencmd measure_temp)"
// echo "CPU => $((cpu/1000))'C"

func main() {

}
