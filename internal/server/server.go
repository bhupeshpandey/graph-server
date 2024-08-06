package server

import (
	"github.com/bhupeshpandey/graph-server/internal/graph"
	"net/http"
)

type Server struct {
	graphStore *graph.GraphStore
	mux        *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()
	server := &Server{
		graphStore: graph.NewGraphStore(),
		mux:        mux,
	}
	server.registerHandlers(mux)

	return server
}

func (server *Server) Serve() error {
	err := http.ListenAndServe(":2007", server.mux)
	if err != nil {
		return err
	}

	return nil
}

func (server *Server) registerHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/graph", server.handlePostGraph)
	mux.HandleFunc("/graph/{id}/shortestpath/{start}/{end}", server.handleGetShortestPath)
	mux.HandleFunc("/graph/{id}", server.handleDeleteGraph)
}

func httpRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {

	}
}
