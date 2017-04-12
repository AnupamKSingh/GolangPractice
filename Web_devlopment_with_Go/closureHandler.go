package main
import (
_	"log"
	"fmt"
	"net/http"
)

func messageHandler(mess string) (http.Handler)	{
	switch mess {
		case "mh1":
			return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)	{
						fmt.Fprintf(w, "Message from mh1")
					})
		case "mh2":
			return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)	{
						fmt.Fprintf(w, "Message from mh2")
					})
	}
	return nil
}



func main()	{
	mux := http.NewServeMux()
	mux.Handle("/mh1", messageHandler("mh1"))
	mux.Handle("/mh2", messageHandler("mh2"))

	http.ListenAndServe(":8989", mux)
}
