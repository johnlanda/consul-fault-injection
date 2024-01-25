package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

// PageData struct to hold data for rendering the homepage
type PageData struct {
	Timestamp    string
	StatusCode   int
	Latency      time.Duration
	ServerHeader string
}

var (
	port                = getEnvOrDefault("PORT", "8080")
	heartbeatServiceURL = getEnvOrDefault("HEARTBEAT_SERVICE_URL", "http://localhost:8000")
	requestHistory      []PageData
	maxHistorySize      = 10
)

func main() {
	http.HandleFunc("/", homeHandler)
	go makePeriodicRequests()

	fmt.Printf("Dashboard Service is running on :%d\n", port)
	http.ListenAndServe(port, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Display the last 10 requests in the page
	tmpl, err := template.New("index").Parse(`
		<html>
		<head>
			<title>Dashboard Service</title>
			<style>
				body {
					background: linear-gradient(to bottom, #ffffff, #DC447D);
					color: #000000;
					font-family: 'Arial', sans-serif;
					margin: 0;
					padding: 0;
					box-sizing: border-box;
				}
		
				h1 {
					color: #DC447D;
					text-align: center;
					margin-top: 20px;
				}
		
				table {
					width: 80%;
					margin: 20px auto;
					border-collapse: collapse;
					box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
					background-color: #ffffff;
					border-radius: 8px;
				}
		
				th, td {
					padding: 15px;
					text-align: left;
					border: 1px solid #DC447D;
				}
		
				th {
					background-color: #DC447D;
					color: #ffffff;
				}
			</style>
		</head>
		<body>
			<h1>Last 10 Requests</h1>
			<table>
				<tr>
					<th>Timestamp</th>
					<th>Status Code</th>
					<th>Latency</th>
					<th>Server</th>
				</tr>
				{{range .}}
					<tr>
						<td>{{.Timestamp}}</td>
						<td>{{.StatusCode}}</td>
						<td>{{.Latency}} milliseconds</td>
						<td>{{.ServerHeader}}</td>
					</tr>
				{{end}}
			</table>
		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Reverse the requestHistory to display the latest requests first
	reversedHistory := make([]PageData, len(requestHistory))
	copy(reversedHistory, requestHistory)
	for i, j := 0, len(reversedHistory)-1; i < j; i, j = i+1, j-1 {
		reversedHistory[i], reversedHistory[j] = reversedHistory[j], reversedHistory[i]
	}

	tmpl.Execute(w, reversedHistory)
}

func makePeriodicRequests() {
	for {

		t0 := time.Now()
		// Make an HTTP GET request to the heartbeat service
		resp, err := http.Get(heartbeatServiceURL)
		if err != nil {
			fmt.Println("Error making request:", err)
			continue
		}

		// Extract relevant information from the response
		timestamp := time.Now().Format(time.RFC3339)
		statusCode := resp.StatusCode
		latency := time.Since(t0)
		serverHeader := resp.Header.Get("Server")

		// Update request history
		requestHistory = append(requestHistory, PageData{
			Timestamp:    timestamp,
			StatusCode:   statusCode,
			Latency:      latency,
			ServerHeader: serverHeader,
		})

		// Keep only the last 10 requests
		if len(requestHistory) > maxHistorySize {
			requestHistory = requestHistory[len(requestHistory)-maxHistorySize:]
		}

		// Close the response body
		resp.Body.Close()

		// Sleep for 1 second before making the next request
		time.Sleep(1 * time.Second)
	}
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
