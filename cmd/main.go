package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Billy278/assignment_project/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	go func() {
		// Membuat time ticker dengan interval 1 menit
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		// Memanggil fungsi untuk mengakses endpoint setiap kali ticker berdenyut
		for range ticker.C {
			fmt.Println("detik")
		}
	}()
	server.NewServer()
}
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
