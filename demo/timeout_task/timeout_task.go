package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(1100))
	defer cancel()
	res := TestTimeout(ctx, time.Millisecond*time.Duration(1100))

	fmt.Println(res)
}

func TestTimeout(ctx context.Context, timeout time.Duration) (res string) {
	timer := time.NewTimer(timeout)

	go func(ctx context.Context) {
		time.Sleep(time.Second)
		res = "success"
	}(ctx)

	select {
	case <-ctx.Done():
		res = "ctx is canceled"
		return
	case <-timer.C:
		res = "timeout"
		return
	}
}
