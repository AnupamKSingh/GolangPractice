package main
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleWare(next http.Handler) http.Handler	{
	log.Println("Logging is called")
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("\nlog is invoked at start with %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("\nlog is invoked at end with %s %s", r.URL.Path, time.Since(start))
	})
}

func index(w http.ResponseWriter, r *http.Request)	{
	log.Println ("Index is called")
	fmt.Fprintf(w, "Index is invoked")
}

func about(w http.ResponseWriter, r *http.Request)	{
	log.Println("About is called")
	fmt.Fprintf(w, "About is invoked")
}

func main()	{
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	http.Handle("/", loggingMiddleWare(indexHandler))
	http.Handle("/about", loggingMiddleWare(aboutHandler))
	server := &http.Server{
		Addr : ":8989",
	}
	log.Println("Let's Begin")
	server.ListenAndServe()
}
