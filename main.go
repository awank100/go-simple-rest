package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Goal :
type Goal struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  string `json:"year"`
}

var goals []Goal

func main() {
	router := mux.NewRouter()

	goals = append(goals,
		Goal{ID: 1, Title: "Beli Motor", Year: "2017"},
		Goal{ID: 2, Title: "Menikah", Year: "2018"},
		Goal{ID: 3, Title: "Bangun Rumah", Year: "2019"})

	router.HandleFunc("/goals", getGoals).Methods("GET")
	router.HandleFunc("/goals/{id}", getGoal).Methods("GET")
	router.HandleFunc("/goals", addGoal).Methods("POST")
	router.HandleFunc("/goals/{id}", updateGoal).Methods("PUT")
	router.HandleFunc("/goals/{id}", removeGoal).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":7000", router))
}

func getGoals(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(goals)
}

func getGoal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for _, goal := range goals {
		if goal.ID == id {
			json.NewEncoder(w).Encode(&goal)
		}
	}
}

func addGoal(w http.ResponseWriter, r *http.Request) {
	var goal Goal
	json.NewDecoder(r.Body).Decode(&goal)

	goals = append(goals, goal)

	json.NewEncoder(w).Encode(goals)
	log.Println(goal)
}

func updateGoal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for i, goal := range goals {
		if goal.ID == id {
			var goal Goal
			json.NewDecoder(r.Body).Decode(&goal)

			goal.ID = id

			goals[i] = goal
			json.NewEncoder(w).Encode(goal)
		}
	}

}

func removeGoal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for i, goal := range goals {
		if goal.ID == id {
			goals = append(goals[:i], goals[i+1:]...)
			json.NewEncoder(w).Encode(goals)
		}
	}
}
