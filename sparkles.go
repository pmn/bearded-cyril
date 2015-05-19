package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Sparkle struct {
	Sparkler string    `json:"sparkler"`
	Sparklee string    `json:"sparklee"`
	Reason   string    `json:"reason,omitempty"`
	Time     time.Time `json:"time,omitempty"`
}

type Leader struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// Return the data in JSON format. This is the default return method.
func returnJson(obj interface{}, w http.ResponseWriter, h *http.Request) {
	// Don't cache json returns. This is to work around ie's weird caching behavior
	w.Header().Set("Cache-Control", "no-cache")
	// Set the content type to json
	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, string(j))
}

// Placeholder in case anyone calls the root of the service.
// Perhaps change this to 404.
func defaultHandler(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Default sparkles")
}

// Add a sparkle!
func addSparkle(w http.ResponseWriter, h *http.Request) {
	var s Sparkle
	b := json.NewDecoder(h.Body)
	b.Decode(&s)

	result := db.AddSparkle(s)
	returnJson(result, w, h)
}

// Get the entire data set
func getSparkles(w http.ResponseWriter, h *http.Request) {
	returnJson(db.Sparkles, w, h)
}

// This only goes back 60 days
func topGiven(w http.ResponseWriter, h *http.Request) {
	after_date := time.Now().AddDate(0,0,-60)
	result := db.TopGiven(after_date)
	returnJson(result, w, h)
}

// This only goes back 60 days
func topReceived(w http.ResponseWriter, h *http.Request) {
	after_date := time.Now().AddDate(0,0,-60)
	result := db.TopReceived(after_date)
	returnJson(result, w, h)
}

// Get all the sparkles for someone in particular
func getSparklesForRecipient(w http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	rcpt := vars["recipient"]
	sparkles := db.SparklesForUser(rcpt)
	returnJson(sparkles, w, h)
}
