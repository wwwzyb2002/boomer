package main

import (
	"fmt"
	"log"
	"time"

	"github.com/wwwzyb2002/boomer"
)

func foo(user *boomer.User) {
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	elapsed := time.Since(start)

	// Report your test result as a success, if you write it in python, it will looks like this
	// events.request_success.fire(request_type="http", name="foo", response_time=100, response_length=10)
	globalBoomer.RecordSuccess("http", "foo", elapsed.Nanoseconds()/int64(time.Millisecond), int64(10))
}

func bar(user *boomer.User) {
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	elapsed := time.Since(start)

	// Report your test result as a failure, if you write it in python, it will looks like this
	// events.request_failure.fire(request_type="udp", name="bar", response_time=100, exception=Exception("udp error"))
	globalBoomer.RecordFailure("udp", "bar", elapsed.Nanoseconds()/int64(time.Millisecond), "udp error")
}

var globalBoomer = boomer.NewStandaloneBoomer(10, 1)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	userConfig := &boomer.UserConfig{
		Tasks: []*boomer.Task{
			{
				Name:   "foo",
				Weight: 1000,
				Fn:     foo,
			},
			{
				Name:   "bar",
				Weight: 9000,
				Fn:     bar,
			},
		},
		StartFunc: func(user *boomer.User) error {
			fmt.Println("user start")
			return nil
		},
		StopFunc: func(user *boomer.User) {
			fmt.Println("user stop")
		},
		WaitTime: func() time.Duration {
			return 100 * time.Millisecond
		},
	}

	globalBoomer.AddOutput(boomer.NewConsoleOutput())
	globalBoomer.Run(userConfig)
}
