package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []struct {
		ShortDescription string `json:"shortDescription"`
		Price            string `json:"price"`
	} `json:"items"`
	Total string `json:"total"`
}

type IDJ struct {
	Id string `json:"id"`
}

type pointVal struct {
	Points int `json:"points"`
}

var scores = make(map[string]int)

// gets score from map using id and returns json with points
func points(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var score int = scores[id]

	pointJson := pointVal{Points: score}
	json.NewEncoder(w).Encode(pointJson)

}

// gets points value from receipt and stores it in map, returns json with id
func process(w http.ResponseWriter, r *http.Request) {
	var rec receipt

	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		fmt.Fprint(w, "Bad Request")
	} else {
		var idNum int = rand.Int()
		var score int = calcPoints(rec)
		var idTag string = strconv.Itoa(idNum)

		scores[idTag] = score

		idJson := IDJ{Id: idTag}
		json.NewEncoder(w).Encode(idJson)
	}
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/receipts/process", process).Methods("POST")
	myRouter.HandleFunc("/receipts/{id:[0-9]+}/points", points).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
