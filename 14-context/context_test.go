package _4_context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Spystore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *Spystore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *Spystore) Cancel() {
	s.cancelled = true
}

func (s *Spystore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store was not told to cancel")
	}
}

func (s *Spystore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("store was told to cancel")
	}
}

func TestHandler(t *testing.T) {
	data := "hello world !"
	t.Run("returns data from store", func(t *testing.T) {
		store := &Spystore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &Spystore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		store.assertWasCancelled()
	})
}
