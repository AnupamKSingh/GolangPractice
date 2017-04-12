package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/mux"	
)

type Person struct {
	ID string `json:"id,omitempty"`
	FirstName string	`json:"firstname,omitempty"`	
	LastName string	`json:"lastname,omitempty"`
	Age int `json:"age,omitempty"`
}

var people []Person

func CreatePerson (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person Person
	body, err := ioutil.ReadAll (io.LimitReader (r.Body,12121212))
    if err != nil {
        fmt.Println (err)
		fmt.Println ("Read")
    }
	fmt.Println (string(body))
	if err := json.NewDecoder(r.Body).Decode (person);err != nil {
//	if err := json.Unmarshal(body, &person);err != nil {
		fmt.Println ("Decode")
		fmt.Println (err)
		w.Header().Set("Content-type", "application/json; charset=UTF-8")
		w.WriteHeader (234)
	}
	fmt.Println (person)
	person.ID = id
	people = append (people ,person)
	if err := json.NewEncoder(w).Encode(people); err != nil {
		fmt.Println ("Encode")
		w.Header().Set("Content-type", "application/json; charset=UTF-8")
		w.WriteHeader (432)
	}
}

func main () {
	router := mux.NewRouter ()
	people = append(people, Person{ID: "2", FirstName: "Maria", LastName: "Raboy",Age:20})
	router.HandleFunc ("/people/{id}", CreatePerson).Methods ("POST")
	log.Fatal (http.ListenAndServe (":8080", router))
}
