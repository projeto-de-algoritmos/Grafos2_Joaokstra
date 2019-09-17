package main

import(
    "fmt"
    "sync"
    "strconv"
)
type Node struct{
    Key string
    Value int
}

type Edge struct{
    to *Node
    weight int
}

func (n *Node) String() string {
    return fmt.Sprintf("%v", n.Key)
}

type Graph struct{
	nodes []*Node
    edges map[Node][]Edge
    mutex  sync.RWMutex
}

func (g *Graph) AddNode(n *Node){
    g.mutex.Lock()
    g.nodes = append(g.nodes, n)
    g.mutex.Unlock()
}

func (g *Graph) AddEdge(node1, node2 *Node, weight int){
    g.mutex.Lock()
	if g.edges == nil{
		g.edges = make(map[Node][]Edge)
	}
	// non directional graph implementation, so edges point both ways
	g.edges[*node1] = append(g.edges[*node1], Edge{node2, weight})
    g.edges[*node2] = append(g.edges[*node2], Edge{node1, weight})
    g.mutex.Unlock()
}

func (g *Graph) GetNode(key string)*Node{
    for _, n := range g.nodes {
        if(n.Key == key){
            return n
        }
    }
    return nil
}

func (g *Graph) String() {
    g.mutex.RLock()
    s := ""
    for i := 0; i < len(g.nodes); i++ {
        s += g.nodes[i].String() + " -> "
        near := g.edges[*g.nodes[i]]
        for j := 0; j < len(near); j++ {
            s += near[j].to.String() + "(" + strconv.Itoa(near[j].weight)+") "
        }
        s += "\n"
    }
    fmt.Println(s)
    g.mutex.RUnlock()
}

func (g *Graph) Dijkstra(origin, destiny *Node) (int, []*Node) {
    h := newHeap()
    h.push(path{value: 0, nodes: []*Node{origin}})
    visited := make(map[*Node]bool)

    for len(*h.values) > 0 {
        p := h.pop()
        node := p.nodes[len(p.nodes)-1]

        if visited[node] {
            continue
        }

        if node == destiny {
            return p.value, p.nodes
        }

        for _, e := range g.edges[*node] {
            if !visited[e.to] {
                h.push(path{value: p.value + e.weight, nodes: append([]*Node{}, append(p.nodes, e.to)...)})
            }
        }

        visited[node] = true
    }

    return 0, nil
}


