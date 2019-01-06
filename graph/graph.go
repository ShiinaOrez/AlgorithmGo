package graph

import (
	"github.com/ShiinaOrez/AlgorithmGo/typedefs"
)

type Type typedefs.GraphType

type GraphNode struct {
	Index  int
	Value  Type
	In     [] *GraphNode
	Out    [] *GraphNode
}

type Graph struct {
	Nodes  [] *GraphNode
	Size   int
}

func (g *Graph) HasNode(index int) bool {
	for _, node := range(g.Nodes) {
		if node.Index == index {
			return true
		}
	}
	return false
}

func (g *Graph) Related(from int, to int) bool {
	for _, node := range(g.Nodes) {
		if node.Index == from {
			for _, t := range(node.Out) {
				if t.Index == to {
					return true
				}
			}
			break
		}
	}
	return false
}

func (g *Graph) Insert(index int, value Type) bool {
	if g.HasNode(index) {
		return false
	}
	node := new(GraphNode)
	node.Value = value
	node.Index = index
	g.Nodes = append(g.Nodes, node)
	g.Size += 1
	return true
}

func (g *Graph) Relate(from int, to int) bool {
	if !g.HasNode(from) || !g.HasNode(to) {
		return false
	}
	if g.Related(from, to) {
		return false
	}else{
		var FROM *GraphNode
		var TO *GraphNode
		for _, node := range(g.Nodes) {
			if node.Index == from {
				FROM = node
			}
			if node.Index == to {
				TO = node
			}
			if FROM != nil && TO != nil {
				break
			}
		}
		FROM.Out = append(FROM.Out, TO)
		TO.In = append(TO.In, FROM)
		return true
	}
}