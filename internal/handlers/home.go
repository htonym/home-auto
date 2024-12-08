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
	Rooms []models.Room
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(tplFolder, "templates/room-list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data HomeData

	data.Rooms, err = models.GetAllRooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("writing template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
