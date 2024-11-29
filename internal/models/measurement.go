package models

import (
	"home-auto/internal/db"
	"time"
)

type Measurement struct {
	ID           int64
	RoomID       int64
	TemperatureC float64
	TemperatureF float64
	Humidity     float64
	Timestamp    int64
	TimestampStr string
	RoomName     string
}

// Get all measurements over the lastSeconds provided
func GetMeasurements(lastSeconds int64) ([]Measurement, error) {
	query := `
	SELECT measurements.*, rooms.name as room_name
	FROM measurements
	LEFT JOIN rooms ON measurements.room_id = rooms.id
	WHERE measurements.timestamp > $1
	ORDER by measurements.timestamp;
	`
	includeAllAfter := time.Now().Unix() - lastSeconds
	rows, err := db.DB.Query(query, includeAllAfter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var measurements []Measurement

	for rows.Next() {
		var measurement Measurement
		err := rows.Scan(
			&measurement.ID,
			&measurement.RoomID,
			&measurement.TemperatureC,
			&measurement.Humidity,
			&measurement.Timestamp,
			&measurement.RoomName,
		)
		if err != nil {
			return nil, err
		}

		measurement.TimestampStr = time.Unix(measurement.Timestamp, 0).Format(time.RFC3339)
		measurement.TemperatureF = (measurement.TemperatureC * 9 / 5) + 32

		measurements = append(measurements, measurement)
	}

	return measurements, nil
}
