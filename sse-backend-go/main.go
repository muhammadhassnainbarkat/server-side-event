package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/event/subscribe", subscribe)
	http.HandleFunc("/event/unsubscribe", unsubscribe)
	go eventGenerator()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

type Client struct {
	id    string
	event chan string
}

var clients = make(map[string]*Client)

func eventGenerator() {
	counter := 0
	for {
		time.Sleep(5 * time.Second)
		counter++
		for _, client := range clients {
			client.event <- fmt.Sprintf("Message %d from Server", counter)
		}
	}
}

func subscribe(w http.ResponseWriter, r *http.Request) {

	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	clientID := r.URL.Query().Get("clientId")

	if clientID == "" {
		http.Error(w, "Client Id Missing", http.StatusBadRequest)
		return
	}

	client := &Client{id: clientID, event: make(chan string)}

	clients[clientID] = client

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for event := range client.event {

		n, err := fmt.Fprintf(w, "data: Message %s from Server to Client %s \n\n", event, clientID)

		if err != nil {
			fmt.Printf("Error: %s", err)
		} else {
			fmt.Printf("Written: %d\n", n)
		}
		flusher.Flush()
	}

}

func unsubscribe(w http.ResponseWriter, r *http.Request) {
	clientId := r.URL.Query().Get("clientId")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if clientId == "" {
		http.Error(w, "Client id is required", http.StatusBadRequest)
		return
	}

	if client, ok := clients[clientId]; ok {
		close(client.event)
		delete(clients, clientId)
	}
}
