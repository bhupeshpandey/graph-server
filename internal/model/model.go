package model

type Graph struct {
	Vertices map[string][]string `json:"vertices"`
}

type ShortestPathResponse struct {
	Path  []string `json:"path"`
	Error string   `json:"error,omitempty"`
}
