package graph

type Type int

const (
	Directed Type = iota
	Undirected
)

type Graph struct {
	// Type identifies whether the graph is directed or undirected.
	Type  Type
	hash  map[int]Node
	edges map[int]map[int]interface{}
}

func New(graphType Type) *Graph {
	graph := new(Graph)
	graph.Type = graphType
	graph.hash = make(map[int]Node)
	graph.edges = make(map[int]map[int]interface{})
	return graph
}

type Node interface {
	ID() int
}

func (g *Graph) AddNode(node Node) {
	g.hash[node.ID()] = node
}

func (g *Graph) AddEdge(u, v int, edgeData interface{}) {
	if _, exists := g.hash[u]; !exists {
		panic("Starting node doesn't exist")
	}
	if _, exists := g.hash[v]; !exists {
		panic("Ending node doesn't exist")
	}

	if g.edges[u] == nil {
		g.edges[u] = make(map[int]interface{})
	}
	g.edges[u][v] = edgeData

	if g.Type == Undirected {
		if g.edges[v] == nil {
			g.edges[v] = make(map[int]interface{})
		}
		g.edges[v][u] = edgeData
	}
}

func (g *Graph) Node(id int) (Node, bool) {
	if node, exists := g.hash[id]; exists {
		return node, true
	} else {
		return nil, false
	}
}

func (g *Graph) Edge(u, v int) (interface{}, bool) {
	if _, exists := g.hash[u]; !exists {
		return nil, false
	}
	if _, exists := g.hash[v]; !exists {
		return nil, false
	}
	if edge, exists := g.edges[u][v]; exists {
		return edge, true
	}
	return nil, false
}

func (g *Graph) Nodes(f func(Node)) {
	for id := range g.hash {
		f(g.hash[id])
	}
}

func (g *Graph) Edges(f func(u, v Node, edgeData interface{})) {
	if g.Type == Directed {
		for i := range g.edges {
			for j := range g.edges[i] {
				f(g.hash[i], g.hash[j], g.edges[i][j])
			}
		}
	} else {
		visited := make(map[int]bool)
		for i := range g.edges {
			for j := range g.edges[i] {
				if !visited[j] {
					f(g.hash[i], g.hash[j], g.edges[i][j])
					visited[i] = true
				}
			}
		}
	}
}

func (g *Graph) Neighbours(u int, f func(v Node, edgeData interface{})) {
	if _, exists := g.hash[u]; !exists {
		panic("Node doesn't exist")
	}
	for v := range g.edges[u] {
		f(g.hash[v], g.edges[u][v])
	}
}
