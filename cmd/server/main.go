package main

import (
	"flag"
	"fmt"
	"home-auto/internal/db"
	"home-auto/internal/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
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
	var envPath string
	flag.StringVar(&envPath, "env", ".env", "provide path to .env file")
	flag.Parse()

	initEnv(envPath)
	db.InitDB()

	if db.DB == nil {
		log.Fatal("DB is nil")
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /home", handlers.HomePage)
	router.HandleFunc("POST /measurement", handlers.AddMeasurement)
	router.HandleFunc("GET /rooms", handlers.GetAllRooms)

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

func initEnv(envPath string) {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
