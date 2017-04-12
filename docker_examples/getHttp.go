package main

import (
    "fmt"
  _  "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%q","Container name" )
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
