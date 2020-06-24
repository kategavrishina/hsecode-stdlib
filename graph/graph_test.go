package graph_test

import (
	"fmt"
	"hsecode.com/stdlib/graph"
	"sort"
	"testing"
)

type Int int

func (v Int) ID() int {
	return int(v)
}

type Person struct {
	Id   int
	Name string
}

func (d Person) ID() int { return d.Id }

func TestGraph_AddNode(t *testing.T) {
	G := graph.New(graph.Undirected)

	G.AddNode(Person{Id: 1, Name: "Fred Weasley"})
	node, _ := G.Node(1)

	fmt.Println(node.(Person).Name)

	// data will be overwritten due to the same ID
	G.AddNode(Person{Id: 1, Name: "George Weasley"})

	node, _ = G.Node(1)
	fmt.Println(node.(Person).Name)

	// Output:
	// Fred Weasley
	// George Weasley
}

func TestGraph_Edges(t *testing.T) {
	g := graph.New(graph.Directed)
	g.AddNode(Int(2))
	g.AddNode(Int(3))

	g.AddEdge(2, 3, "forward edge")
	g.AddEdge(3, 2, "backward edge")

	g.Edges(func(u, v graph.Node, e interface{}) {
		fmt.Println(u, v, e)
	})

	// Output:
	// 2 3 forward edge
	// 3 2 backward edge

	ug := graph.New(graph.Undirected)
	ug.AddNode(Int(2))
	ug.AddNode(Int(3))

	ug.AddEdge(2, 3, "single edge")
	ug.Edges(func(u, v graph.Node, e interface{}) {
		fmt.Println(u, v, e)
	})

	// Output: 2 3 single edge
}

func TestGraph_Node(t *testing.T) {
	g := graph.New(graph.Directed)
	g.AddNode(Int(2))

	fmt.Println(g.Node(2))
	fmt.Println(g.Node(8))

	// Output:
	// 2 true
	// <nil> false
}

func TestGraph_Nodes(t *testing.T) {
	g := graph.New(graph.Directed)
	g.AddNode(Int(2))
	g.AddNode(Int(3))

	g.Nodes(func(node graph.Node) {
		fmt.Println(node)
	})

	// Output: 2 3
}

func TestGraph_Neighbours(t *testing.T) {
	G := graph.New(graph.Undirected)

	// Create the following graph with no edge data:
	//
	//  1---2 \
	//  | / |  3
	//  5---4 /

	for i := 1; i <= 5; i++ {
		G.AddNode(Int(i))
	}
	G.AddEdge(1, 2, nil)
	G.AddEdge(2, 3, nil)
	G.AddEdge(3, 4, nil)
	G.AddEdge(4, 5, nil)
	G.AddEdge(5, 1, nil)
	G.AddEdge(2, 4, nil)
	G.AddEdge(2, 5, nil)

	// Compute neighbour nodes for 2
	nodes := make([]int, 0)
	G.Neighbours(2, func(v graph.Node, e interface{}) {
		nodes = append(nodes, v.ID())
	})

	sort.Ints(nodes)
	fmt.Println(nodes)

	// Output: [1 3 4 5]
}
