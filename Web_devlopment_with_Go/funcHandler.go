package main
import (
	"fmt"
	"net/http"
	"log"
)

func messageHandler1(w http.ResponseWriter, r *http.Request)	{
	fmt.Fprintf(w, "You are in messageHandler1")
}

func messageHandler2(w http.ResponseWriter, r *http.Request)	{
	fmt.Fprintf(w, "Tou are in messageHandler2")
}

func main()	{
	mux := http.NewServeMux()
	mh1 := http.HandlerFunc(messageHandler1)
	mux.Handle("/mh1", mh1)

	mh2 := http.HandlerFunc(messageHandler2)
	mux.Handle("/mh2", mh2)

	log.Println ("ListenAndServe os started")
	http.ListenAndServe(":8989", mux)
}
