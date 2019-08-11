package graph

// SearchWithBFS changes an edge-weighted graph to a graph and calls SearchWithBFS() with it.
func (gr *EWGraph) SearchWithBFS(startVertex *Vertex, searchedValue interface{}) bool {
	return gr.Graph.SearchWithBFS(startVertex, searchedValue)
}

// SearchWithBFS looks for searched value in the graph 'gr'.
// It returns true if the value is found.
func (gr *Graph) SearchWithBFS(startVertex *Vertex, searchedValue interface{}) bool {
	var search = gr.GetRelations(startVertex)
	var searched = make(map[*Vertex]bool)

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
