package clients

import (
	"net/http"
	"time"
)

// Why using time.Duration?
func TimeoutClient(d time.Duration) *http.Client {
	client := http.Client{Timeout: d}
	return &client
}
