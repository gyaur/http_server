package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var state State

const magicnumber = 20 // number of states to remember

type State struct {
	State     bool      `json:"state"`
	Timestamp time.Time `json:"timestamp"`
	History   []State   `json:"history"`
}

type Body struct {
	State bool
}

func main() {
	fmt.Println("Starting...")
	http.HandleFunc("/api/state", setState)
	http.ListenAndServe(":7890", nil)
}

func setState(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		var b Body
		err := json.NewDecoder(req.Body).Decode(&b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		state.State = b.State
		state.Timestamp = time.Now()
		tempstate := state
		tempstate.History = nil // not proud of this
		state.History = append(state.History, tempstate)
		if len(state.History) > magicnumber {
			state.History = state.History[1:]
		}
		return
	} else if req.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(state)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}
