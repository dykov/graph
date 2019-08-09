package bfs

import (
	"graph"
)

// SearchWithBFS looks for searched value in the graph gr.
// It returns true if the value is found.
func SearchWithBFS(gr *graph.Graph, startVertex *graph.Vertex, searchedValue interface{}) bool {
	var search = gr.GetRelations(startVertex)
	var searched = make(map[*graph.Vertex]bool)

	for len(search) > 0 {

		checkVertex := search[0]
		search = search[1:]

		if !searched[checkVertex] { // if checkVertex was not checked yet.

			if checkVertex.Value == interface{}(searchedValue) {
				return true
			}

			search = append(search, gr.GetRelations(checkVertex)...)
			searched[checkVertex] = true

		}

	}

	return false

}
