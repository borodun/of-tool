package src

import (
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordTicks() {
	go func() {
		for {
			percent, _ = cpu.Percent(time.Second, false)
			opsProcessed.Set(percent[0])
		}
	}()
}

var (
	percent, _ = cpu.Percent(time.Second, true)

	opsProcessed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "info_cpu_usage_total",
		Help: "The total cpu usage",
	})
)

func Run() {
	recordTicks()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
