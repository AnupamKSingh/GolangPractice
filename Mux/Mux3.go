package main

import (
	"net/http"
	"time"
	"fmt"
	"io"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"

)

type Todo struct {
	ID string	`json:"id"`
	Name string		`json:"name"`
	Completed bool	`json:"completed"`
	Due time.Time	`json:"due"`
}

type Todos []Todo

type Route struct {
	Name string
	Method string
	Path string
	HttpHandler http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {
		"todo",
		"GET",
		"/todo/{todos}",
		todoShow,
	},
	Route {
		"toPut",
		"POST",
		"/todo",
		todoPost,
	},
}

func NewRouter () *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _,route := range routes	{
				router.
				Methods (route.Method).
				Path (route.Path).
				Name (route.Name).
				Handler (route.HttpHandler)
			}
	return router
}

func todoShow (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars (r)
	todo := vars ["todos"]
	fmt.Fprintln (w, "Hello ", todo)
}

func todoPost (w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll (io.LimitReader (r.Body,12121212))
	if err != nil {
		fmt.Println (err)
	}
	if err := r.Body.Close; err != nil {
		fmt.Println (err)
	}
	if err = json.Unmarshal (body, &todo);err != nil {
		w.Header ().Set ("Content-type", "application/json; charset=UTF-8")
		w.WriteHeader (432)
		if err := json.NewEncoder (w).Encode (err); err != nil {
			fmt.Println (err)
		}
	}
	w.Header().Set("Content-type","application/json; charset=UTF-8")
	w.WriteHeader (http.StatusCreated)
	if err := json.NewEncoder(w).Encode (todo);err != nil {
		fmt.Println (err)
	}
}


func main () {
	router := NewRouter ()
	log.Fatal (http.ListenAndServe (":8080", router))
}
