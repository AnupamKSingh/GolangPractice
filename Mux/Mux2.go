package main
import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
_	"io/ioutil"
	"github.com/gorilla/mux"
	"time"

)

type Todo struct {
	Name string		`json:"name"`
	Completed bool	`json:"completed"`
	Due time.Time	`json:"due"`
}

type Route struct {
	Name string
	Method string
	Path string
	HttpHandler http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {
		"Index",
		"GET",
		"/",
		index,
	},
	Route {
		"TodoIndex",
		"GET",
		"/todo",
		todoIndex,
		},
	Route {
		"TodoShow",
		"GET",
		"/todo/{todos}",
		todoShow,
	},
}

type Todos []Todo

func index (w http.ResponseWriter, r *http.Request)	{
	fmt.Fprintln (w, "Welcome!")
}

func todoIndex (w http.ResponseWriter, r *http.Request)	{
	todos := Todos{
				Todo {
					Name :"First Person",
					Completed: true,
				},
				Todo {
					Name :"Second Person",
					Completed :false,
				},
			}
			
	if err := json.NewEncoder(w).Encode (todos); err != nil {
		fmt.Println (err)
	}
}

func todoShow (w http.ResponseWriter, r *http.Request)	{
	vars := mux.Vars (r)
	todos := vars ["todos"]
	fmt.Fprintln (w, "This is ", todos)
}

func Newrouter () *mux.Router {
	router := mux.NewRouter ().StrictSlash (true)
	for _, route := range routes {
		router.
		Methods (route.Method).
		Path (route.Path).
		Name (route.Name).
		Handler (route.HttpHandler)
	}
	return router
}

func main()	{
	router := Newrouter ()
	log.Fatal (http.ListenAndServe (":8080", router))
}
			
