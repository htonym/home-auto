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

func GetAllMeasurements() ([]Measurement, error) {
	query := `
	SELECT measurements.*, rooms.name as room_name
	FROM measurements
	LEFT JOIN rooms ON measurements.room_id = rooms.id;
	`
	rows, err := db.DB.Query(query)
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

		t := time.Unix(measurement.Timestamp, 0)
		measurement.TimestampStr = t.Format("2006-01-02 3:04 PM")

		measurement.TemperatureF = (measurement.TemperatureC * 9 / 5) + 32

		measurements = append(measurements, measurement)
	}

	return measurements, nil
}
