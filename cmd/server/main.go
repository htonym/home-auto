package main

import (
	"fmt"
	"home-auto/internal/handlers"
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
	router := http.NewServeMux()

	router.HandleFunc("GET /home", handlers.HomePage)

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
