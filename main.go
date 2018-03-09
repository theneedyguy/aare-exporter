package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Create aare struct
type aare struct {
	AareTempLast float64 `json:"aaretemp_last"`
	AareTempMax  float64 `json:"aaretemp_max24"`
	AareTempAvg  float64 `json:"aaretemp_mid24"`
}

// Define constants
const (
	url      = "https://api.purpl3.net/aare/v1/aare.json"
	version  = "1.0"
	interval = 30
)

// Define variables (prometheus metrics)
var (
	aareTempCelsius = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "aare_temp_celsius",
		Help: "Displays the last measured aare temperature.",
	})
	aareTempCelsiusMaxDay = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "aare_temp_celsius_max_day",
		Help: "Displays the maximum temperature of this day.",
	})
	aareTempCelsiusAvgDay = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "aare_temp_celsius_avg_day",
		Help: "Displays the average temperature throughout this day.",
	})
)

// Register prometheus metrics
func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(aareTempCelsius)
	prometheus.MustRegister(aareTempCelsiusMaxDay)
	prometheus.MustRegister(aareTempCelsiusAvgDay)
}

// Make reqauest, get json, unmarshal into struct,
// set prometheus gauges with value and wait
func doRequest() {
	for {
		aareClient := http.Client{
			Timeout: time.Second * 2, // Maximum of 2 secs
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "Aare-Exporter")

		res, getErr := aareClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		aareJSON := aare{}
		jsonErr := json.Unmarshal(body, &aareJSON)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		aareTempCelsius.Set(aareJSON.AareTempLast)
		aareTempCelsiusMaxDay.Set(aareJSON.AareTempMax)
		aareTempCelsiusAvgDay.Set(aareJSON.AareTempAvg)

		// Wait for the next interval
		time.Sleep(interval * time.Second)
	}
}

func main() {
	// Run in background
	go doRequest()

	// Handle prometheus
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":3005", nil))
}
