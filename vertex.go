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
