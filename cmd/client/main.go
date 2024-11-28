package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"home-auto/internal/sensor"
	"log"
	"net/http"
	"time"
)

type Payload struct {
	RoomID      int64   `json:"roomId"`
	Timestamp   int64   `json:"timestamp"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func main() {
	interval := flag.Int64("interval", 10, "Number of seconds between measurements.")
	host := flag.String("host", "http://rpi4.local:8080", "Number of seconds between measurements.")
	roomId := flag.Int64("room-id", 1, "Number of seconds between measurements.")
	flag.Parse()

	timer := time.NewTicker(time.Duration(*interval) * time.Second)
	defer timer.Stop()

	// Create a channel to signal when we want to stop
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-timer.C:
				err := RecordMeasurement(*host, *roomId)
				if err != nil {
					log.Printf("Failed to record measurement: %v", err)
				} else {
					log.Println("Successfully recorded measurement")
				}
			case <-done:
				return
			}
		}
	}()

	// Runs the program indefinitely
	select {}
}

func RecordMeasurement(host string, roomId int64) error {
	url := host + "/measurement"

	measurement, err := sensor.ReadShtc3("")
	if err != nil {
		return fmt.Errorf("reading sensor: %w", err)
	}

	payload := Payload{
		RoomID:      roomId,
		Timestamp:   measurement.Timestamp,
		Temperature: measurement.TemperatureC,
		Humidity:    measurement.Humidity,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshaling json: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
