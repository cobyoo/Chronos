package handlers

import (
    "encoding/json"
    "net/http"
)

type BuildStatus struct {
    Status string `json:"status"`
}

func BuildStatusHandler(w http.ResponseWriter, r *http.Request) {
    status := BuildStatus{Status: "Success"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(status)
}
