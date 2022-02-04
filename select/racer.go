package selectpkg

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) string {
	duration1 := measureResponseTime(url1)
	duration2 := measureResponseTime(url1)

	if duration1 < duration2 {
		return url1
	} else {
		return url2
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
