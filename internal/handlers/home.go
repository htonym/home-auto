package handlers

import (
	"embed"
	"home-auto/internal/sensor"
	"html/template"
	"log"
	"net/http"
	"time"
)

//go:embed templates/*
var tplFolder embed.FS

type RoomConditions struct {
	Name         string
	Temperature  float64
	Humidity     float64
	Timestamp    int64
	TimestampStr string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(tplFolder, "templates/home.html", "templates/room.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	measurement, err := sensor.ReadShtc3("Master Bedroom")
	if err != nil {
		log.Printf("ERROR")
	}

	t := time.Unix(measurement.Timestamp, 0)
	timestampStr := t.Format("2006-01-02 3:04 PM")

	data := RoomConditions{
		Name:         measurement.Location,
		Temperature:  measurement.TemperatureF,
		Humidity:     measurement.Humidity,
		Timestamp:    measurement.Timestamp,
		TimestampStr: timestampStr,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
