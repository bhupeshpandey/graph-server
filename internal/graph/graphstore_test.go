package graph

import (
	"encoding/json"
	"github.com/bhupeshpandey/graph-server/internal/model"
	"testing"
)

func TestGraphStore_AddGetGraph(t *testing.T) {
	gs := NewGraphStore()

	edges := map[string][]string{
		"A": {"B"},
		"B": {"A", "C"},
		"C": {"B"},
	}

	graph := model.Graph{Vertices: edges}
	marshal, err := json.Marshal(graph)
	if err != nil {
		return
	}

	graphID, _, err := gs.AddGraph(marshal)
	if err != nil {
		return
	}
	graph, exists := gs.GetGraph(graphID)
	if !exists {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(graph.Vertices) != len(edges) {
		t.Fatalf("expected %d edges, got %d", len(edges), len(graph.Vertices))
	}
}

func TestGraphStore_GetNonExistentGraph(t *testing.T) {
	gs := NewGraphStore()
	_, exists := gs.GetGraph("nonexistent")
	if exists {
		t.Fatalf("expected error, got nil")
	}
}

func TestGraphStore_DeleteGraph(t *testing.T) {
	gs := NewGraphStore()
	graphID := "graph1"
	edges := map[string][]string{
		"A": {"B"},
	}

	marshal, err := json.Marshal(model.Graph{Vertices: edges})
	if err != nil {
		return
	}

	_, _, err = gs.AddGraph(marshal)
	if err != nil {
		return
	}
	gs.DeleteGraph(graphID)
	_, exists := gs.GetGraph(graphID)
	if exists {
		t.Fatalf("expected error, got nil")
	}
}

func TestGraphStore_ShortestPath(t *testing.T) {
	gs := NewGraphStore()
	graphID := "graph1"
	edges := map[string][]string{
		"A": {"B"},
		"B": {"A", "C"},
		"C": {"B"},
	}

	marshal, err := json.Marshal(model.Graph{Vertices: edges})
	if err != nil {
		return
	}

	_, _, err = gs.AddGraph(marshal)
	if err != nil {
		return
	}

	path, err := gs.FindShortestPath(graphID, "A", "C")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expectedPath := []string{"A", "C"}
	if len(path) != len(expectedPath) {
		t.Fatalf("expected path length %d, got %d", len(expectedPath), len(path))
	}
	for i, v := range path {
		if v != expectedPath[i] {
			t.Fatalf("expected %s at position %d, got %s", expectedPath[i], i, v)
		}
	}
}
