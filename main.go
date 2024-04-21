package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/events", eventsHandler)
	http.ListenAndServe(":8080", nil)
}
func generateRandomInt(maxnumber int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	return r.Intn(maxnumber)

}
func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Simulate sending events (you can replace this with real data)
	for i := 0; i < 1000; i++ {
		_, err := fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
		if err != nil {
			return
		}
		time.Sleep(time.Duration(generateRandomInt(i+500)) * time.Millisecond)
		w.(http.Flusher).Flush()
	}

	//// Simulate closing the connection
	//closeNotify := w.(http.CloseNotifier).CloseNotify()
	//<-closeNotify
}
