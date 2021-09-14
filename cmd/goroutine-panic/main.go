package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	group, _ := errgroup.WithContext(context.Background())
	for i := 0; i < 10; i++ {
		group.Go(allGood)
	}
	group.Go(panicMuch)

	if err := group.Wait(); err != nil {
		fmt.Println("finished")
	}
}

func panicMuch() error {
	defer recov()
	select {
	case <-time.After(time.Second):
	}

	panic("derp")
}

func allGood() error {
	select {
	case <-time.After(time.Second):
	}

	return nil
}

func recov() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
