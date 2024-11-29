package handlers

import (
	"home-auto/internal/models"
	"html/template"
	"log"
	"net/http"
)

func HistoryPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(tplFolder, "templates/history.html")
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
