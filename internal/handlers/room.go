package handlers

import (
	"encoding/json"
	"fmt"
	"home-auto/internal/models"
	"log"
	"net/http"
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
