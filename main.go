package main

import (
	"fmt"
	"os"
	"time"

	"github.com/imbdb/go-with-tests/mocking"
)

func GoWithTests() string {
	return "Go with Tests"
}

func main() {
	fmt.Println(GoWithTests())
	sleeper := &mocking.ConfigurableSleeper{}
	sleeper.Init(1*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
