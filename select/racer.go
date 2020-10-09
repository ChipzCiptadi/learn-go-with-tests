package racer

import (
	"fmt"
	"net/http"
	"time"
)

var serverTimeout = 3 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, serverTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Server timeout waiting for %q and %q", a, b)
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
