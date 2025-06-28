package main

import (
	"infra-watchdog/api"
	"infra-watchdog/monitor"
	"log"
)

func main() {
	log.Println("Starting infra-watchdog 🚀")

	// Start monitoring in a goroutine
	go monitor.StartMonitoring("config/targets.json")

	// Start REST API
	api.StartServer()
}
