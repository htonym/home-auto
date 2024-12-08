package handlers

import (
	"encoding/json"
	"fmt"
	"home-auto/internal/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := models.GetAllRooms()
	if err != nil {
		returnedErr := fmt.Sprintf("Could not retrieve rooms: %v", err)
		log.Println(returnedErr)
		http.Error(w, returnedErr, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(rooms); err != nil {
		returnedErr := fmt.Sprintf("Could not encode json: %v", err)
		log.Println(returnedErr)
		http.Error(w, returnedErr, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type RoomData struct {
	Measurements    []models.Measurement
	LastMeasurement models.Measurement
}

func ViewRoom(w http.ResponseWriter, req *http.Request) {
	roomId, err := strconv.Atoi(req.PathValue("id")) // Hardcoded to master bedroom
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFS(tplFolder, "templates/room.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data RoomData

	data.Measurements, err = models.GetMeasurements(measurementSpan, int64(roomId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count := len(data.Measurements)
	if count > 0 {
		data.LastMeasurement = data.Measurements[count-1]
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("writing template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
