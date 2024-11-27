package main

import (
	"fmt"
	"home-auto/internal/db"
	"home-auto/internal/handlers"
	"log"
	"net/http"
)

const AppPort = "8080"

type RoomConditions struct {
	Name         string
	Temperature  float64
	Humidity     float64
	Timestamp    int64
	TimestampStr string
}

func main() {
	db.InitDB()

	if db.DB == nil {
		log.Fatal("DB is nil")
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /home", handlers.HomePage)
	router.HandleFunc("POST /measurement", handlers.AddMeasurement)

	server := &http.Server{
		Addr:    ":" + AppPort,
		Handler: router,
	}

	fmt.Printf("Server is running on http://rpi4.local:%s\n", AppPort)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
