package handlers

import (
    "fmt"
    "net/http"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        fmt.Fprintf(w, "Webhook received!")
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}
