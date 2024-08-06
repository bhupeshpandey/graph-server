package graph

import (
	"encoding/json"
	"github.com/bhupeshpandey/graph-server/internal/model"
	"testing"
)

func BenchmarkGraphStore_AddGraph(b *testing.B) {
	gs := NewGraphStore()
	edges := map[string][]string{
		"A": {"B"},
		"B": {"A", "C"},
		"C": {"B"},
	}
	for i := 0; i < b.N; i++ {
		marshal, err := json.Marshal(model.Graph{Vertices: edges})
		if err != nil {
			return
		}

		_, _, err = gs.AddGraph(marshal)
		if err != nil {
			return
		}
	}
}

func BenchmarkGraphStore_GetGraph(b *testing.B) {
	gs := NewGraphStore()
	edges := map[string][]string{
		"A": {"B"},
		"B": {"A", "C"},
		"C": {"B"},
	}
	marshal, err := json.Marshal(model.Graph{Vertices: edges})
	if err != nil {
		return
	}

	id, _, err := gs.AddGraph(marshal)
	if err != nil {
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, exists := gs.GetGraph(id)
		if !exists {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}
