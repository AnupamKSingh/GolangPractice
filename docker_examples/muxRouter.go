package main

import (
    "fmt"
    "html"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc ("/deepoo", Deepoo)
	router.HandleFunc ("/anupam", Anupam)
	router.HandleFunc ("/anupam/{json}", JsonFunc)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Println ("Anupam")
}

func Deepoo (w http.ResponseWriter, r *http.Request) {
	fmt.Println ("I am Deepoo")
}

func Anupam (w http.ResponseWriter, r *http.Request) {
	fmt.Println ("I am Anupam")
}

func JsonFunc (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars (r)
	fmt.Println ("Variable Arguments")
	fmt.Println (vars["json"])
}
