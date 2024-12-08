package main

import (
	"fmt"
	"home-auto/internal/sensor"
)

func main() {
	measurement, err := sensor.ReadShtc3("")
	if err != nil {
		fmt.Printf("reading sensor: %v", err)
	}

	fmt.Println("temperature: ", measurement.TemperatureF)
	fmt.Println("humidity: ", measurement.Humidity)
	fmt.Println("timestamp: ", measurement.Timestamp)
}
