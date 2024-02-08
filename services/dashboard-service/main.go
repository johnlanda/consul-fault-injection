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
	Timestamp       string
	StatusCode      int
	StatusCodeColor string
	Latency         time.Duration
	LatencyColor    string
	ServerHeader    string
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

	fmt.Printf("Dashboard Service is running on :%s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Errorf("error starting server: %w", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)

	// Display the last 10 requests in the page
	tmpl, err := template.New("index").Parse(`
<html lang="en"> 
<head>
    <title>Dashboard Service</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <style type="text/css">
        #content {
            background: linear-gradient(to bottom, #ffffff, #DC447D);
            color: #000000;
            font-family: 'Arial', sans-serif;
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
    </style>
</head>
<body id="content" class="overflow-x-hidden">
<div class="max-w-full max-h-full h-screen">
    <div class="grid grid-rows-1">
        <div class="grid grid-cols-12">
            <div class="col-span-2">
                <svg id="LOGOS" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 301 132">
                    <defs>
                        <style>.cls-1 {
                            fill: #e03875;
                        }</style>
                    </defs>
                    <path d="M117.19,62.75c0-7.38,4.3-11.68,14.34-11.68a46,46,0,0,1,11,1.33l-.85,6.6a59,59,0,0,0-9.86-1c-5.27,0-7,1.81-7,6.11V79.57c0,4.29,1.69,6.11,7,6.11a59,59,0,0,0,9.86-1l.85,6.6a46,46,0,0,1-11,1.33c-10,0-14.34-4.3-14.34-11.68Z"/>
                    <path d="M158.82,92.58C148.71,92.58,146,87,146,81V73.52c0-6,2.72-11.62,12.83-11.62s12.82,5.57,12.82,11.62V81C171.64,87,168.92,92.58,158.82,92.58Zm0-24.39c-3.94,0-5.45,1.76-5.45,5.09V81.2c0,3.33,1.51,5.08,5.45,5.08s5.44-1.75,5.44-5.08V73.28C164.26,70,162.75,68.19,158.82,68.19Z"/>
                    <path d="M194.09,92V71.4c0-1.57-.67-2.36-2.36-2.36s-5,1.09-7.69,2.48V92h-7.38V62.51h5.63L183,65a29.49,29.49,0,0,1,11.8-3.09c4.9,0,6.66,3.45,6.66,8.71V92Z"/>
                    <path d="M216.65,92.58a34.66,34.66,0,0,1-10.16-1.7l1-5.62a33.18,33.18,0,0,0,8.77,1.27c3.27,0,3.75-.73,3.75-3,0-1.81-.36-2.72-5.14-3.87-7.2-1.75-8.05-3.57-8.05-9.26,0-5.93,2.6-8.53,11-8.53a38.51,38.51,0,0,1,8.83,1L226,68.8a52.36,52.36,0,0,0-8.11-.85c-3.21,0-3.75.73-3.75,2.54,0,2.36.18,2.54,4.18,3.57,8.22,2.18,9,3.27,9,9.32C227.3,89.07,225.55,92.58,216.65,92.58Z"/>
                    <path d="M239.52,62.51V83.08c0,1.57.66,2.36,2.36,2.36s5-1.09,7.68-2.48V62.51h7.38V92h-5.62l-.73-2.48a29.44,29.44,0,0,1-11.8,3.09c-4.9,0-6.65-3.45-6.65-8.72V62.51Z"/>
                    <path d="M263.29,92V50.47l7.38-1V92Z"/>
                    <path class="cls-1"
                          d="M61.5,103.59A31.76,31.76,0,1,1,83,48.47h0l-7.51,7.88h0a20.87,20.87,0,1,0,0,31h0L83,95.2h0A31.66,31.66,0,0,1,61.5,103.59Z"/>
                    <path class="cls-1" d="M87.06,87.7a2.6,2.6,0,1,1,2.6-2.6A2.6,2.6,0,0,1,87.06,87.7Z"/>
                    <path class="cls-1" d="M61.3,78.71a6.88,6.88,0,1,1,6.88-6.88A6.88,6.88,0,0,1,61.3,78.71Z"/>
                    <path class="cls-1" d="M90.05,78.81a2.6,2.6,0,1,1,2.6-2.6A2.61,2.61,0,0,1,90.05,78.81Z"/>
                    <path class="cls-1" d="M82.31,78.51a2.6,2.6,0,1,1,2.6-2.6A2.6,2.6,0,0,1,82.31,78.51Z"/>
                    <path class="cls-1" d="M90.05,70.06a2.6,2.6,0,1,1,2.6-2.6A2.61,2.61,0,0,1,90.05,70.06Z"/>
                    <path class="cls-1" d="M82.31,70.36a2.6,2.6,0,1,1,2.6-2.6A2.6,2.6,0,0,1,82.31,70.36Z"/>
                    <path class="cls-1" d="M87.21,61.31a2.6,2.6,0,1,1,2.6-2.6A2.6,2.6,0,0,1,87.21,61.31Z"/>
                    <path d="M125.44,41.74V36.09h-5.17v5.65h-2.64V28.16h2.64v5.7h5.17v-5.7h2.64V41.74Zm12.33,0h-2.1l-.19-.66a5.78,5.78,0,0,1-3,.87c-1.86,0-2.66-1.23-2.66-2.92,0-2,.9-2.76,3-2.76h2.45v-1c0-1.08-.31-1.46-2-1.46a14.94,14.94,0,0,0-2.83.3l-.31-1.87a13.46,13.46,0,0,1,3.5-.47c3.2,0,4.15,1.09,4.15,3.54ZM135.21,38h-1.89c-.83,0-1.06.22-1.06,1s.23,1,1,1a4.15,4.15,0,0,0,1.93-.5ZM143.12,42a12.59,12.59,0,0,1-3.52-.57l.36-1.87a11.94,11.94,0,0,0,3,.42c1.13,0,1.3-.24,1.3-1s-.13-.91-1.78-1.29c-2.5-.58-2.79-1.19-2.79-3.08s.9-2.84,3.81-2.84a14,14,0,0,1,3.06.34l-.25,2a18.84,18.84,0,0,0-2.81-.28c-1.11,0-1.3.24-1.3.84,0,.79.07.85,1.45,1.19,2.85.72,3.12,1.08,3.12,3.1S146.2,42,143.12,42Zm11.71-.21V34.9c0-.52-.23-.78-.82-.78a7.27,7.27,0,0,0-2.66.82v6.8h-2.56V28l2.56.38v4.34a9.36,9.36,0,0,1,3.73-.95c1.7,0,2.3,1.15,2.3,2.9v7.1Zm4.7-11.18v-2.4h2.56v2.4Zm0,11.18v-9.8h2.56v9.8Zm4.6-9.72c0-2.45,1.49-3.88,5-3.88a16.47,16.47,0,0,1,3.79.44l-.29,2.2a21.57,21.57,0,0,0-3.42-.35c-1.82,0-2.41.61-2.41,2v5.15c0,1.43.59,2,2.41,2a21.49,21.49,0,0,0,3.42-.34l.29,2.19a16.47,16.47,0,0,1-3.79.45c-3.48,0-5-1.43-5-3.89ZM178.54,42c-3.5,0-4.44-1.86-4.44-3.87V35.61c0-2,.94-3.87,4.44-3.87S183,33.59,183,35.61v2.47C183,40.09,182,42,178.54,42Zm0-8.11c-1.36,0-1.89.58-1.89,1.69v2.63c0,1.11.53,1.69,1.89,1.69s1.89-.58,1.89-1.69V35.53C180.43,34.42,179.9,33.84,178.54,33.84Zm11.64.16a19.4,19.4,0,0,0-2.7,1.43v6.31h-2.56v-9.8h2.16l.17,1.09a11.54,11.54,0,0,1,2.68-1.29Zm10.22,4.48c0,2.18-1,3.47-3.37,3.47a14.83,14.83,0,0,1-2.73-.29v4l-2.55.38V31.94h2l.25.83a5.55,5.55,0,0,1,3.23-1c2.05,0,3.14,1.17,3.14,3.4Zm-6.1,1.11a12,12,0,0,0,2.27.26c.92,0,1.28-.42,1.28-1.31V35.08c0-.8-.32-1.24-1.26-1.24a3.73,3.73,0,0,0-2.29.88Z"/>
                </svg>
            </div>
            <div class="col-span-8">
                <div class="flex items-center justify-center">
                    <div class="grid grid-rows-2">
                        <div class="row-span-1 text-center justify-center">
                            <h1 class="text-3xl mt-5 font-semibold leading-6 text-gray-900">Dashboard Service</h1>
                        </div>
                        <div class="row-span-1 text-center justify-center">
                            <p class="mt-2 text-sm text-gray-700">The last ten HTTP requests to the heartbeat service</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="flex items-center justify-center">
        <div class="w-2/3">
            <div class="mt-8 flow-root">
                <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                    <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
                        <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 sm:rounded-lg">
                            <table class="min-w-full divide-y divide-gray-300">
                                <thead class="bg-gray-50">
                                <tr>
                                    <th scope="col"
                                        class="py-3.5 text-center justify-center text-sm font-semibold text-gray-900 sm:pl-6">
                                        Timestamp
                                    </th>
                                    <th scope="col" class="text-center justify-center text-sm font-semibold text-gray-900">
                                        Status Code
                                    </th>
                                    <th scope="col" class="text-center justify-center text-sm font-semibold text-gray-900">
                                        Latency
                                    </th>
                                    <th scope="col" class="text-center justify-center text-sm font-semibold text-gray-900">
                                        Server
                                    </th>
                                </tr>
                                </thead>
                                <tbody class="divide-y divide-gray-200 bg-white">
                                {{range .}}
                                <tr class="even:bg-gray-50">
                                    <td class="whitespace-nowrap py-4 px-2 text-sm font-medium text-gray-900 sm:pl-6 justify-center text-center">
                                        {{.Timestamp}}
                                    </td>
                                    <td class="whitespace-nowrap py-4 px-2 text-sm text-gray-500">
                                        <div class="rounded-md {{.StatusCodeColor}} py-1.5 text-sm font-semibold text-white shadow-sm justify-center text-center">
                                            {{.StatusCode}}
                                        </div>
                                    </td>
                                    <td class="whitespace-nowrap py-4 px-2 text-sm text-gray-500">
                                        <div class="rounded-md {{.LatencyColor}} py-1.5 text-sm font-semibold text-white shadow-sm justify-center text-center">
                                            {{.Latency}}
                                        </div>
                                    </td>
                                    <td class="whitespace-nowrap py-4 px-2 text-sm text-gray-500 justify-center text-center">
                                        {{.ServerHeader}}
                                    </td>
                                </tr>
                                {{end}}
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</body>
<script type="application/javascript">
    setTimeout(function () {
        location.reload()
    }, 2000)
</script>
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

	err = tmpl.Execute(w, reversedHistory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func makePeriodicRequests() {
	for {
		fmt.Printf("making heartbeat request to %s\n", heartbeatServiceURL)

		t0 := time.Now()
		// Make an HTTP GET request to the heartbeat service
		resp, err := http.Get(heartbeatServiceURL)
		if err != nil {
			fmt.Println("error making request:", err)
			continue
		}

		// Extract relevant information from the response
		statusCode := resp.StatusCode
		var statusCodeColor string
		switch statusCode {
		case 200:
			statusCodeColor = "bg-green-500"
		case 500:
			statusCodeColor = "bg-red-500"
		default:
			statusCodeColor = "bg-yellow-500"
		}
		latency := time.Since(t0)
		var latencyColor string
		if latency.Milliseconds() < 2 {
			latencyColor = "bg-green-500"
		} else if latency.Seconds() < 3 {
			latencyColor = "bg-yellow-500"
		} else {
			latencyColor = "bg-red-500"
		}
		serverHeader := resp.Header.Get("Server")

		// Update request history
		requestHistory = append(requestHistory, PageData{
			Timestamp:       time.Now().Format("Monday, 02-Jan-2006 15:04:05 MST"),
			StatusCode:      statusCode,
			StatusCodeColor: statusCodeColor,
			Latency:         latency,
			LatencyColor:    latencyColor,
			ServerHeader:    serverHeader,
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
