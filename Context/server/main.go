package main
import (
	"fmt"
	"net/http"
	"time"
	"context"
_	"sync"
	"../log"
)

func main() {
	http.HandleFunc("/", log.Decorator(handler));
	fmt.Println(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int(42), int64(24))
	log.Println(ctx, "Handler Started")

	defer log.Println(ctx, "Handler Closed")

	select {
		case <-time.After(2 * time.Second):
			fmt.Fprintln(w, "Anupam")
		case <-ctx.Done():
			log.Println(ctx, "Context Cancelled")
	}
}

