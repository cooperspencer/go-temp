package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yryz/ds18b20"
	"gitlab.com/buddyspencer/chameleon"
	"net/http"
	"time"
)

var (
	tempProcessed = prometheus.NewGauge(prometheus.GaugeOpts{Name: "Temperature"})
)

func main() {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sensor IDs: %v\n", chameleon.BCyan(sensors))

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)

	for {
		for _, sensor := range sensors {
			t, err := ds18b20.Temperature(sensor)
			if err == nil {
				tempProcessed.Set(t)
			}
		}
		time.Sleep(15 * time.Minute)
	}
}