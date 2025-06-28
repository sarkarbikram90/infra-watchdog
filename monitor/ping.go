package monitor

import (
	"net/http"
	"time"
)

func checkHTTP(url string) string {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return "DOWN"
	}
	return "UP"
}
