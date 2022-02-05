package selectpkg

import (
	"errors"
	"net/http"
	"time"
)

var ErrTimeout = errors.New("Request timed out")
var timeout = 10 * time.Second

func Racer(url1, url2 string) (string, error) {
	return ConfigurableRacer(url1, url2, timeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
