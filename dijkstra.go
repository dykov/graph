package graph

import (
	"errors"
	"fmt"
)

// Dijkstra finds shortest path between vertices in the edge-weighted graph 'gr'.
// It returns map of costs for each vertex, map of parent of vertices and error
// if some edge has a negative cost value.
func (gr *EWGraph) Dijkstra(start *Vertex) (map[*Vertex]int, map[*Vertex]*Vertex, error) {

	var unchecked []*Vertex
	var costs = make(map[*Vertex]int, len(unchecked))
	var parents = make(map[*Vertex]*Vertex, len(unchecked))
	for v := range gr.relations {
		// fmt.Println(v)
		unchecked = append(unchecked, v)
		costs[v] = -1
		parents[v] = nil
	}

	costs[start] = 0

	for len(unchecked) > 0 {

		currentVert, index := lowestCostUnchecked(unchecked, costs)

		vertCost := costs[currentVert]
		vertNeighbors := gr.GetRelations(currentVert)

		// fmt.Print("!currentVert:", currentVert.Value, " neighbors:")
		// for _, n := range vertNeighbors {
		// 	fmt.Print(n, ",")
		// }
		// fmt.Println()

		for _, n := range vertNeighbors {
			if gr.weight[edge{currentVert, n}] < 0 {
				return nil, nil,
					errors.New("Negative weight edge between " + fmt.Sprint(currentVert) + " and " + fmt.Sprint(n))
			}
			newcost := vertCost + gr.weight[edge{currentVert, n}]
			if costs[n] == -1 || costs[n] > newcost {
				costs[n] = newcost
				parents[n] = currentVert
			}
		}

		unchecked = append(unchecked[:index], unchecked[index+1:]...) // drop from slice

	}

	return costs, parents, nil

}

func lowestCostUnchecked(unchecked []*Vertex, costs map[*Vertex]int) (*Vertex, int) {

	var cheapvert = unchecked[0] // default: the firts unchecked vertex is the cheapest (has min cost)
	var index = 0                // index in the slice

	for i, vert := range unchecked {

		// fmt.Println("v:", vert.Value, " c:", costs[vert])

		if costs[cheapvert] == -1 ||
			costs[vert] != -1 && costs[vert] < costs[cheapvert] {
			cheapvert = vert
			index = i
		}
	}

	// fmt.Println("CHEAP v:", cheapvert.Value, " c:", costs[cheapvert])
	return cheapvert, index

}
