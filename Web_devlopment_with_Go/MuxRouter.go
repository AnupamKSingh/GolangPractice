package main
import (
_	"fmt"
	"log"
	"encoding/json"
	"strconv"
	"time"
	"net/http"

	"github.com/gorilla/mux"
)

type Note struct {
	Title string         `json:"title"`
	Description string   `json:"Description"`
	CreatedOn time.Time  `json:"CreatedOn"`
}

var Id int = 0
var noteStore = make(map[string]Note)

func GetNoteHandler(w http.ResponseWriter, r *http.Request)	{
	var notes []Note
	for _, note := range noteStore {
		notes = append(notes, note)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(&notes)
	if err != nil {
		log.Printf("Error happened while doing json.Marshal")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func PostNoteHandler(w http.ResponseWriter, r *http.Request)	{
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Printf("Error happened while doing json.NewDecoder(r.Body).decode(&notes)")
	}
	Id ++
	k := strconv.Itoa(Id)
	note.CreatedOn = time.Now()
	noteStore[k] = note
	j, err := json.Marshal(note)
	if err != nil {
		log.Printf("json.Marshal(note)")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func PutNoteHandler(w http.ResponseWriter, r *http.Request)	{
	var noteNew Note
	err := json.NewDecoder(r.Body).Decode(&noteNew)
	if err != nil {
		log.Printf ("Error happened in json.NewDecoder(r.Body).Decode(&note)")
	}
	vars := mux.Vars(r)
	k := vars["id"]
	if _, ok := noteStore[k];ok {
		noteNew.CreatedOn = time.Now()
		delete(noteStore, k)
		noteStore[k] = noteNew
	} else {
		log.Printf("Not present in noteStore")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request)	{
	vars := mux.Vars(r)
	k := vars["id"]
	log.Printf(k)
	if note, ok := noteStore[k]; ok {
		log.Printf(note.Title)
		log.Printf(note.Description)
		delete(noteStore, k)
	} else {
		log.Printf("No entry found with this index")
	}
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}


func main()	{
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := http.Server{
		Addr : ":8989",
		Handler: r,
	}

	server.ListenAndServe()
}
