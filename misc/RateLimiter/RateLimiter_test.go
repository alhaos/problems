package RateLimiter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

// makeServer создаёт тестовый HTTP сервер с задержкой
func makeServer(t *testing.T, delay time.Duration) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		fmt.Fprintln(w, "ok")
	}))
}

func TestFetch_Basic(t *testing.T) {
	srv := makeServer(t, 0)
	defer srv.Close()

	urls := []string{srv.URL, srv.URL, srv.URL}
	results := Fetch(urls, 3)

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
		if r.Body == "" {
			t.Errorf("expected non-empty body for %s", r.URL)
		}
	}
}

func TestFetch_URLsPreserved(t *testing.T) {
	srv := makeServer(t, 0)
	defer srv.Close()

	urls := []string{srv.URL + "/a", srv.URL + "/b", srv.URL + "/c"}
	results := Fetch(urls, 3)

	for i, r := range results {
		if r.URL != urls[i] {
			t.Errorf("index %d: got URL %q, want %q", i, r.URL, urls[i])
		}
	}
}

func TestFetch_ErrorHandling(t *testing.T) {
	urls := []string{"http://localhost:1", "http://localhost:2"}
	results := Fetch(urls, 2)

	for _, r := range results {
		if r.Err == nil {
			t.Errorf("expected error for unreachable URL %s", r.URL)
		}
	}
}

func TestFetch_MaxConcurrent(t *testing.T) {
	var active atomic.Int32
	var maxSeen atomic.Int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cur := active.Add(1)
		defer active.Add(-1)

		// фиксируем максимум
		for {
			old := maxSeen.Load()
			if cur <= old || maxSeen.CompareAndSwap(old, cur) {
				break
			}
		}

		time.Sleep(20 * time.Millisecond)
		fmt.Fprintln(w, "ok")
	}))
	defer srv.Close()

	urls := make([]string, 10)
	for i := range urls {
		urls[i] = srv.URL
	}

	maxConcurrent := 3
	Fetch(urls, maxConcurrent)

	if got := int(maxSeen.Load()); got > maxConcurrent {
		t.Errorf("max concurrent requests = %d, want <= %d", got, maxConcurrent)
	}
}

func TestFetch_EmptyURLs(t *testing.T) {
	results := Fetch([]string{}, 3)
	if len(results) != 0 {
		t.Errorf("expected empty result, got %v", results)
	}
}
