package mocking

import (
	"bytes"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	sleep = "sleep"
	write = "write"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	Convey("Should test countdown", t, func() {
		buffer := &bytes.Buffer{}
		spyCountdownOperations := new(SpyCountdownOperations)
		Countdown(buffer, spyCountdownOperations)

		got := buffer.String()
		want := `3
2
1
Go!`
		So(got, ShouldEqual, want)
		So(len(spyCountdownOperations.Calls), ShouldEqual, 4)
	})

	Convey("Should call sleep after every print", t, func() {
		// buffer := &bytes.Buffer{}
		spyCountdownOperations := new(SpyCountdownOperations)
		Countdown(spyCountdownOperations, spyCountdownOperations)

		// 		got := buffer.String()
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		// So(got, ShouldEqual, want)
		So(spyCountdownOperations.Calls, ShouldResemble, want)
	})

	Convey("Should call sleep after every print", t, func() {
		// buffer := &bytes.Buffer{}
		spyCountdownOperations := new(SpyCountdownOperations)
		Countdown(spyCountdownOperations, spyCountdownOperations)

		// 		got := buffer.String()
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		// So(got, ShouldEqual, want)
		So(spyCountdownOperations.Calls, ShouldResemble, want)
	})
}

func TestConfigurableSleeper(t *testing.T) {
	Convey("Should init configured time and function", t, func() {
		sleepTime := 5 * time.Second
		sleeper := ConfigurableSleeper{}
		sleeper.Init(sleepTime, time.Sleep)
		So(sleeper.duration, ShouldEqual, sleepTime)
		So(sleeper.sleep, ShouldEqual, time.Sleep)
	})
	Convey("Should sleep for configured time", t, func() {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}

		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()
		So(spyTime.durationSlept, ShouldEqual, sleepTime)
	})
}
