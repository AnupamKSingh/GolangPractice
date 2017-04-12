package main
import (
	"net/http"
_	"log"
_	"strconv"
_	"fmt"
)

func main()	{
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/",fs)
	http.ListenAndServe(":8989", mux)
}
