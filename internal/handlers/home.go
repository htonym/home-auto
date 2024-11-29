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

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(tplFolder, "templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	measurements, err := models.GetAllMeasurements()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, measurements)
	if err != nil {
		log.Printf("writing template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
