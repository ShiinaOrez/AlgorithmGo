package graph

import (
	"github.com/ShiinaOrez/AlgorithmGo/typedefs"
	_"github.com/ShiinaOrez/AlgorithmGo/queue"
)

type Type typedefs.GraphType

type GraphNode struct {
	Index  int
	Value  Type
	In     BSTree
	Out    BSTree
}

type Graph struct {
	NodeTree BSTree
	Size   int
}

func (g *Graph) HasNode(index int) bool {
	if g.NodeTree.Find(index) != nil {
		return true
	}
	return false
}

func (g *Graph) Related(from int, to int) bool {
	if node := g.NodeTree.Find(from); node != nil {
		if node.Out.Find(to) != nil {
			return true
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
	g.NodeTree.Push(node)
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
		FROM := g.NodeTree.Find(from)
		TO := g.NodeTree.Find(to)
		FROM.Out.Push(TO)
		TO.In.Push(FROM)
		return true
	}
}