package selectpkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {

	Convey("Should return faster url", t, func() {
		slowServer := makeDelayedServer(2 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)
		So(err, ShouldBeNil)
		So(got, ShouldEqual, want)
	})

	Convey("returns an error if a server doesn't respond within 10s", t, func() {
		serverA := makeDelayedServer(2 * time.Millisecond)
		serverB := makeDelayedServer(3 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 1*time.Millisecond)

		So(err, ShouldEqual, ErrTimeout)
	})
}
