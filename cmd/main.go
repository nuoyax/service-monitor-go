package main

import (
	"log"
	"service-monitor/internal/web"
)

func main() {
	r := web.SetupRouter()
	log.Println("Server started on :8080")
	r.Run(":8080")
}
