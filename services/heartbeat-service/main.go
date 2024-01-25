package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	serverID := getEnvOrDefault("HEARTBEAT_SERVER_ID", "Heartbeat Service")

	// Define the HTTP handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the Server header with the value of HEARTBEAT_SERVER_ID
		w.Header().Set("Server", serverID)

		// Send a simple response
		w.Write([]byte("Heartbeat OK"))
	})

	// Start the server on port 8081
	fmt.Println("Heartbeat Service is running on :8000")
	http.ListenAndServe(":8000", nil)
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
