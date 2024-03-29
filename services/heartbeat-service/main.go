package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	port     = getEnvOrDefault("PORT", "8080")
	serverID = getEnvOrDefault("HEARTBEAT_SERVER_ID", "Heartbeat Service")
)

func main() {
	// Define the HTTP handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the Server header with the value of HEARTBEAT_SERVER_ID
		w.Header().Set("Server", serverID)

		// Send a simple response
		w.Write([]byte("Heartbeat OK"))
	})

	// Start the server on port 8081
	fmt.Println("Heartbeat Service is running on :8000")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Errorf("error starting server: %w", err)
	}
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
