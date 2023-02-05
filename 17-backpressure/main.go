package main

import (
	"errors"
	"net/http"
	"time"
)

// struct that contains a buffered channel
type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	// initialize the channel and its "token" limited capacity
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

// execute Process for every time
func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		// run the function
		f()
		// release the token
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func workload() string {
	time.Sleep(3000 * time.Millisecond)
	return "done"
}

func main() {
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(workload()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(err.Error()))
		}
	})
	http.ListenAndServe(":8080", nil)
}
