package main

import (
	"context"
	"fmt"
	"time"
	"bufio"
	"os"
)

func sleepAndTalk(ctx context.Context, t time.Duration, msg string) {
	for {
		select {
			case <-time.After(t):
				fmt.Println("Hello")
				return
			case <-ctx.Done():
				fmt.Println("Context is cancelled")
				return
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()
	sleepAndTalk(ctx, 5 * time.Second, "Anupam")
}
