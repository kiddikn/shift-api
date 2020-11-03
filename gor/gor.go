package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		mu   sync.Mutex
		done []string
	)

	eg, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		i := i
		eg.Go(func() error {
			ID, err := heavy(i)
			if err != nil {
				return err
			}
			mu.Lock()
			done = append(done, ID)
			mu.Unlock()
			return nil
		})
	}

	err := eg.Wait()
	for i := range done {
		fmt.Println(done[i])
	}
	if err != nil {
		return err
	}

	return nil
}

func heavy(i int) (string, error) {
	time.Sleep(time.Second * 1)
	fmt.Println(i)
	return uuid.New().String(), nil
}
