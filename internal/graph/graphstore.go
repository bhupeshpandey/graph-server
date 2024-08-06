package graph

import (
	"encoding/json"
	"fmt"
	"github.com/bhupeshpandey/graph-server/internal/model"
	"github.com/google/uuid"
	"net/http"
	"sync"
)

type GraphStore struct {
	sync.RWMutex
	Graphs map[string]model.Graph
}

func NewGraphStore() *GraphStore {
	return &GraphStore{
		RWMutex: sync.RWMutex{},
		Graphs:  make(map[string]model.Graph),
	}
}

func (graphStore *GraphStore) GetGraph(id string) (model.Graph, bool) {
	graphStore.RLock()
	graph, exists := graphStore.Graphs[id]
	graphStore.RUnlock()
	return graph, exists
}

func (graphStore *GraphStore) AddGraph(input []byte) (string, int, error) {
	var graph model.Graph
	err := json.Unmarshal(input, &graph)
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	uid, err := uuid.NewV7()
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	id := uid.String()
	graphStore.Lock()
	graphStore.Graphs[id] = graph
	graphStore.Unlock()
	return id, http.StatusOK, err
}

func (graphStore *GraphStore) DeleteGraph(id string) int {
	graphStore.Lock()
	_, exists := graphStore.Graphs[id]
	if !exists {
		return http.StatusNotFound
	}
	delete(graphStore.Graphs, id)
	graphStore.Unlock()
	return http.StatusOK
}

func (graphStore *GraphStore) FindShortestPath(id, startPosition, endPosition string) ([]string, error) {
	graphStore.RLock()
	graph, exists := graphStore.Graphs[id]
	graphStore.RUnlock()
	if !exists {
		return nil, fmt.Errorf("the graph with the id %s is not found", id)
	}

	if startPosition == endPosition {
		return []string{startPosition}, nil
	}

	visited := make(map[string]bool)
	queue := [][]string{{startPosition}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]

		if visited[node] {
			continue
		}

		for _, neighbor := range graph.Vertices[node] {
			newPath := append([]string{}, path...)
			newPath = append(newPath, neighbor)

			if neighbor == endPosition {
				return newPath, nil
			}

			queue = append(queue, newPath)
		}

		visited[node] = true
	}
	return nil, fmt.Errorf("no path found")
}
