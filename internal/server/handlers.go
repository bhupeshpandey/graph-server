package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (server *Server) handlePostGraph(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	readInput, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": err})
		return
	}
	id, statusCode, err := server.graphStore.AddGraph(readInput)
	if err != nil {
		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": err})
		return
	}

	indent, err := json.MarshalIndent(map[string]interface{}{"id": id}, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(indent)
}

func (server *Server) handleGetShortestPath(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	parsedURL, err := url.Parse(r.URL.RequestURI())
	if err != nil {
		// Handle error
	}

	paths := parsedURL.Path
	parts := strings.Split(paths, "/")

	id := parts[2]
	start := parts[4]
	end := parts[5]

	path, err := server.graphStore.FindShortestPath(id, start, end)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"path": path})
}

func (server *Server) handleDeleteGraph(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	parsedURL, err := url.Parse(r.URL.RequestURI())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": err})
	}

	paths := parsedURL.Path
	parts := strings.Split(paths, "/")

	id := parts[2]

	statusCode := server.graphStore.DeleteGraph(id)
	w.WriteHeader(statusCode)
}
