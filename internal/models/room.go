package models

import (
	"home-auto/internal/db"
)

type Room struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func GetAllRooms() ([]Room, error) {
	query := `SELECT * FROM rooms;`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []Room

	for rows.Next() {
		var room Room
		err := rows.Scan(
			&room.ID,
			&room.Name,
		)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}
