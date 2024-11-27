package handlers

import (
	"encoding/json"
	"home-auto/internal/db"
	"net/http"
)

type Measurement struct {
	RoomID      int     `json:"roomId"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Timestamp   int64   `json:"timestamp"` // Linux epoch timestamp
}

func AddMeasurement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var measurement Measurement
	if err := json.NewDecoder(r.Body).Decode(&measurement); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if measurement.RoomID <= 0 || measurement.Timestamp <= 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `
	INSERT INTO measurements (room_id, temperature, humidity, timestamp) 
	VALUES ($1, $2, $3, $4)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		measurement.RoomID,
		measurement.Temperature,
		measurement.Humidity,
		measurement.Timestamp,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Measurement added successfully"})
}
