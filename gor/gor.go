package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
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

		// ここでbreakしてもほぼ無視される
		// select {
		// case <-ctx.Done():
		// 	break
		// default:
		// }

		i := i
		eg.Go(func() error {
			// これが無いと他のgoroutineでerrorが起きても
			// heavyが実行されてしまう
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			ID, err := heavy(ctx, i)
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
	fmt.Println("処理されたgoroutine数：" + strconv.Itoa(len(done)))
	if err != nil {
		return err
	}

	return nil
}

func heavy(ctx context.Context, i int) (string, error) {
	// 適当な数の時にエラーを発生させる
	// goroutineの発生順序は保証されていないので、
	// 実行する毎にエラーの発生するタイミングが変わり
	// トータルで処理されるgoroutine数が変わる=>適切にキャンセル処理がされていると考えられる
	if i == 3 {
		return "", errors.New("spuriously error")
	}

	time.Sleep(500 * time.Millisecond)
	return uuid.New().String(), nil
}
