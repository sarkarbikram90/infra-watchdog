package api

import (
	"encoding/json"
	"infra-watchdog/monitor"
	"io/ioutil"
	"net/http"
	"strings"
)

func StartServer() {
	http.HandleFunc("/status", getStatus)
	http.HandleFunc("/logs", getLogs)
	http.ListenAndServe(":8081", nil)
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monitor.CurrentStatus)
}

func getLogs(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("logs/events.log")
	logs := strings.Split(string(data), "\n")
	if len(logs) > 100 {
		logs = logs[len(logs)-100:] // last 100
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}
