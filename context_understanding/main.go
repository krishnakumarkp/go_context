package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type key string

const keyname key = "name"

func main() {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		wg     sync.WaitGroup
	)

	//ctx, cancel = context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, keyname, "krishna")
	d := 5 * time.Second
	wg.Add(1)
	go func() {
		defer wg.Done()
		sleepAndTalkContextAwareValue(ctx, d, "hellow world!")
	}()

	//select {
	//case <-time.After(6 * time.Second):
	//cancel()
	//}
	wg.Wait()
}

func sleepAndTalk(ctx context.Context, duration time.Duration, msg string) {
	time.Sleep(duration)
	fmt.Println(msg)

}

func sleepAndTalkChannel(ctx context.Context, duration time.Duration, msg string) {
	select {
	case <-time.After(duration):
		fmt.Println(msg)
	}

}

func sleepAndTalkContextAware(ctx context.Context, duration time.Duration, msg string) {
	select {
	case <-time.After(duration):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
	return
}

func sleepAndTalkContextAwareValue(ctx context.Context, duration time.Duration, msg string) {
	select {
	case <-time.After(duration):
		fmt.Println(msg)
		value := ctx.Value(keyname)
		fmt.Println(value)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
	return
}
