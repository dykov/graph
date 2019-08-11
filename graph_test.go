package graph

import (
	"testing"
)

func TestDijkstra(t *testing.T) {

	graph, start, expCosts, expParents, expError := initTest1()

	var costs, par, err = graph.Dijkstra(start)

	if err != expError {
		t.Error("Error. Expected:", expError, " Got:", err, "\n")
	}

	for v := range costs {
		if costs[v] != expCosts[v] {
			t.Error("Cost error. Expected:", expCosts[v], " Got:", costs[v], "\n")
		}
	}

	for v := range par {
		if par[v] != expParents[v] {
			t.Error("Parent error. Expected:", expParents[v], " Got:", par[v], "\n")
		}
	}

}

// Example from "Grokking Algorithms" by Aditya Y. Bhargava.
func initTest1() (*EWGraph, *Vertex, map[*Vertex]int, map[*Vertex]*Vertex, error) {

	book := NewVertex("book")
	lp := NewVertex("LP")
	poster := NewVertex("poster")
	guitar := NewVertex("bass guitar")
	drums := NewVertex("drums")
	piano := NewVertex("piano")

	egr := NewEWGraph()
	egr.AddEWRelation(book, lp, 5)
	egr.AddEWRelation(book, poster, 0)
	egr.AddEWRelation(lp, guitar, 15)
	egr.AddEWRelation(lp, drums, 20)
	egr.AddEWRelation(poster, drums, 35)
	egr.AddEWRelation(poster, guitar, 30)
	egr.AddEWRelation(guitar, piano, 20)
	egr.AddEWRelation(drums, piano, 10)

	expectedCosts := map[*Vertex]int{
		book:   0,
		lp:     5,
		poster: 0,
		guitar: 20,
		drums:  25,
		piano:  35,
	}

	expectedParents := map[*Vertex]*Vertex{
		book:   nil,
		lp:     book,
		poster: book,
		guitar: lp,
		drums:  lp,
		piano:  drums,
	}

	return egr, book, expectedCosts, expectedParents, nil

}
