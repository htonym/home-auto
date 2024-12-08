package handlers

import (
	"embed"
	"home-auto/internal/models"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*
var tplFolder embed.FS

const measurementSpan = 48 * 60 * 60

type HomeData struct {
	Measurements    []models.Measurement
	LastMeasurement models.Measurement
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(tplFolder, "templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data HomeData

	var roomId int64 = 1 // Hardcoded to master bedroom

	data.Measurements, err = models.GetMeasurements(measurementSpan, roomId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.LastMeasurement = data.Measurements[len(data.Measurements)-1]

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("writing template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
