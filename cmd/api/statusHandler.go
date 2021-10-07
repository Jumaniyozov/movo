package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter,r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	js, err := json.Marshal(currentStatus)
	if err != nil {
		log.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}