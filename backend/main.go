package main

import (
    "log"
    "net/http"
    "MiniCI/handlers"
)

func main() {
    http.HandleFunc("/webhook", handlers.WebhookHandler)
	http.HandleFunc("/api/build-status", handlers.BuildStatusHandler)
    log.Println("Server started on :9090")
    log.Fatal(http.ListenAndServe(":9090", nil))
}
