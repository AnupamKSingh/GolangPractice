package main
import (
	"fmt"
	"time"
_	"sync"
	"net/http"
	"io"
	"os"
	"context"
)

func main() {
	ctx := context.Background()


	ctx, cancel := context.WithTimeout(ctx,  6 * time.Second)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error status Code is ",res.StatusCode)
	    return
	}
	io.Copy(os.Stdout, res.Body)
}
