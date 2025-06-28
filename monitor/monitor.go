package monitor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Target struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

var CurrentStatus = make(map[string]string)

func StartMonitoring(configPath string) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	var targets []Target
	err = json.Unmarshal(data, &targets)
	if err != nil {
		log.Fatalf("Invalid JSON config: %v", err)
	}

	for {
		for _, target := range targets {
			status := checkHTTP(target.URL) // From ping.go
			CurrentStatus[target.Name] = status
			logResult(target.Name, status)
		}
		time.Sleep(30 * time.Second)
	}
}

func logResult(name, status string) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Name:      name,
		Status:    status,
	}

	jsonData, _ := json.Marshal(entry)
	f, _ := os.OpenFile("logs/events.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(string(jsonData) + "\n")
}
