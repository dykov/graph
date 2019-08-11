package graph

// Graph represents a graph in graph theory.
// It contains a map shows existence of directed edges from the map key vertex
// to map value vertices.
// The field 'relations' can be explained as map[ *VERTEX_FROM ] []*VERTICES_TO
type Graph struct {
	relations map[*Vertex][]*Vertex
}

// NewGraph returns a new object of type *Graph.
func NewGraph() *Graph {
	return &Graph{relations: make(map[*Vertex][]*Vertex)}
}

// AddRelation adds directed edges between vertex 'from'
// and vertices 'to' in the graph 'g'.
func (g *Graph) AddRelation(from *Vertex, to ...*Vertex) {

	g.relations[from] = append(g.relations[from], to...)

	for _, v := range to {
		if _, exist := g.relations[v]; !exist {
			g.relations[v] = nil
		}
	}

}

// AddBidirectedRelation adds undirected edges between vertex 'vert'
// and vertices 'verts' in the graph 'g'.
func (g *Graph) AddBidirectedRelation(vert *Vertex, verts ...*Vertex) {

	g.AddRelation(vert, verts...)

	for _, v := range verts {
		g.AddRelation(v, vert)
	}

}

// GetRelations returns slice of *Vertex - vertices which have directed edges from vertex 'v'.
func (g *Graph) GetRelations(v *Vertex) []*Vertex {
	return g.relations[v]
}
