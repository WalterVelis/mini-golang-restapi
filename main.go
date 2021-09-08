package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Types
type team struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allTeams []team

// Persistence
var teams = allTeams{
	{
		ID:      1,
		Name:    "Team backend",
		Content: "Content",
	},
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/teams", getTeams)
	log.Println("server running")
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Team!")
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}
