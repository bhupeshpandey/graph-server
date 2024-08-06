package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer()
	if s.graphStore == nil {
		t.Errorf("expected graph to be initialized")
	}
	if s.mux == nil {
		t.Errorf("expected mux to be initialized")
	}
}

func TestServe(t *testing.T) {
	s := NewServer()
	err := s.Serve()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func Test_RegisterHandlers(t *testing.T) {
	s := NewServer()
	mux := http.NewServeMux()
	s.registerHandlers(mux)
	req, err := http.NewRequest("GET", "/graph", nil)
	if err != nil {
		t.Fatal(err)
	}
	h, _ := mux.Handler(req)
	if h == nil {
		t.Errorf("expected handler for /graph to be registered")
	}
	req, err = http.NewRequest("GET", "/graph/shortest_path/123", nil)
	if err != nil {
		t.Fatal(err)
	}
	h, _ = mux.Handler(req)
	if h == nil {
		t.Errorf("expected handler for /graph/shortest_path/{id} to be registered")
	}
	req, err = http.NewRequest("GET", "/graph/123", nil)
	if err != nil {
		t.Fatal(err)
	}
	h, _ = mux.Handler(req)
	if h == nil {
		t.Errorf("expected handler for /graph/{id} to be registered")
	}
}

func TestHandlePostGraph(t *testing.T) {
	req, err := http.NewRequest("POST", "/graph", bytes.NewBuffer([]byte(`{"nodes": ["A", "B", "C"], "edges": [{"from": "A", "to": "B"}, {"from": "B", "to": "C"}]}`)))
	if err != nil {
		t.Fatal(err)
	}
	s := NewServer()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handlePostGraph)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, status)
	}
}

func TestHandleGetShortestPath(t *testing.T) {
	req, err := http.NewRequest("GET", "/graph/shortest_path/A", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	s := NewServer()
	handler := http.HandlerFunc(s.handleGetShortestPath)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}
}

func TestHandleDeleteGraph(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/graph/A", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	s := NewServer()
	handler := http.HandlerFunc(s.handleDeleteGraph)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("expected status code %d, got %d", http.StatusNoContent, status)
	}
}

func BenchmarkServe(b *testing.B) {
	s := NewServer()
	for i := 0; i < b.N; i++ {
		err := s.Serve()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkHandlePostGraph(b *testing.B) {
	req, err := http.NewRequest("POST", "/graph", bytes.NewBuffer([]byte(`{"nodes": ["A", "B", "C"], "edges": [{"from": "A", "to": "B"}, {"from": "B", "to": "C"}]}`)))
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()
	s := NewServer()
	handler := http.HandlerFunc(s.handlePostGraph)
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
	}
}

func BenchmarkHandleGetShortestPath(b *testing.B) {
	req, err := http.NewRequest("GET", "/graph/shortest_path/A", nil)
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()
	s := NewServer()
	handler := http.HandlerFunc(s.handleGetShortestPath)
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
	}
}

func BenchmarkHandleDeleteGraph(b *testing.B) {
	req, err := http.NewRequest("DELETE", "/graph/A", nil)
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()
	s := NewServer()
	handler := http.HandlerFunc(s.handleDeleteGraph)
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
	}
}
