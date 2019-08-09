package graph

// Vertex represents a vertex in graph theory.
// It contains a value of the vertex.
type Vertex struct {
	Value interface{}
}

// NewVertex returns new object type *Vertex with value val.
func NewVertex(val interface{}) *Vertex {
	return &Vertex{Value: val}
}

// Graph represents a graph in graph theory.
// It contains a map shows existence of directed edges from the map key vertex
// to map value vertices.
type Graph struct {
	relations map[*Vertex][]*Vertex
}

// NewGraph returns new object type *Graph.
func NewGraph() *Graph {
	return &Graph{relations: make(map[*Vertex][]*Vertex)}
}

// AddRelation adds directed relations between vertex v
// and vertices verts in the graph g.
func (g *Graph) AddRelation(v *Vertex, verts ...*Vertex) {
	g.relations[v] = append(g.relations[v], verts...)
}

// AddUndirectedRelation adds undirected relations between vertex v
// and vertices verts in the graph g.
func (g *Graph) AddUndirectedRelation(v *Vertex, verts ...*Vertex) {
	g.relations[v] = append(g.relations[v], verts...)
	for _, vert := range verts {
		g.relations[vert] = append(g.relations[vert], v)
	}
}

// GetRelations returns slice of *Vertex - verties which have directed edges from vertex v.
func (g *Graph) GetRelations(v *Vertex) []*Vertex {
	return g.relations[v]
}
