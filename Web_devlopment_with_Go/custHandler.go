package main
import (
	"fmt"
	"net/http"
	"log"
)

type messageHandler struct{
	message string
}

func (m *messageHandler) ServeHTTP (w http.ResponseWriter, r *http.Request)	{
	fmt.Fprintf (w, m.message)
}

func main()	{
	mux := http.NewServeMux()
	mh1 := &messageHandler{"Welcome to mh1"}
	mux.Handle("/mh1", mh1)
	mh2 := &messageHandler{"Welcome to mh2"}
	mux.Handle("/mh2", mh2)

	log.Println ("ListenAndServe is started")
	http.ListenAndServe(":8989", mux)
}
