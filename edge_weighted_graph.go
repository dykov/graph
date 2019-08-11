package graph

type edge struct {
	from *Vertex
	to   *Vertex
}

// EWGraph represents an edge-weighted graph in graph theory.
// It contains a map shows existence of directed edges from the map key vertex
// to map value vertices and edges' weights.
// The field 'weight' contains a value of edge between vertices and
// it can be explained as map[ edge{ VERTEX_FROM , VERTEX_TO } ]EDGE_WEIGHT
type EWGraph struct {
	*Graph
	weight map[edge]int
}

// NewEWGraph returns new object type *EWGraph.
func NewEWGraph() *EWGraph {
	return &EWGraph{NewGraph(), make(map[edge]int)}
}

// AddEWRelation adds directed relations between vertex 'from'
// and vertices 'to' in the graph 'g' with the edge's weight.
func (g *EWGraph) AddEWRelation(from *Vertex, to *Vertex, weight int) {

	g.relations[from] = append(g.relations[from], to)
	if _, exist := g.relations[to]; !exist {
		g.relations[to] = nil
	}

	g.weight[edge{from, to}] = weight

}

// AddRelation adds directed relations between vertex 'from'
// and vertices 'to' in the graph 'g'.
// All edges' weights are 0.
func (g *EWGraph) AddRelation(from *Vertex, to ...*Vertex) {

	for _, v := range to {
		g.AddEWRelation(from, v, 0)
	}

}

// AddBidirectedRelation adds bidirected relations between vertex 'vert'
// and vertices 'verts' in the graph 'g'.
// All edges' weights are 0.
func (g *EWGraph) AddBidirectedRelation(vert *Vertex, verts ...*Vertex) {

	g.AddRelation(vert, verts...)

	for _, v := range verts {
		g.AddRelation(v, vert)
	}

}

// GetRelations returns slice of *Vertex - vertices which have directed edges from vertex 'vert'.
func (g *EWGraph) GetRelations(vert *Vertex) []*Vertex {
	return g.Graph.GetRelations(vert)
}
