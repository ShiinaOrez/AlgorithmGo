package graph

import (
	_ "github.com/ShiinaOrez/AlgorithmGo/queue"
	"github.com/ShiinaOrez/AlgorithmGo/typedefs"
)

type Type typedefs.GraphType

type GraphNode struct {
	Index int
	Value Type
	In    BSTree
	Out   BSTree
}

type Arc struct {
	From  *GraphNode
	To    *GraphNode
	Value int
}
type ArcSlice []Arc

type Graph struct {
	NodeTree BSTree
	ArcSli   ArcSlice
	Size     int
	Directed bool //Directed Graph
}

//sort interface:
func (s ArcSlice) Len() int           { return len(s) }
func (s ArcSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ArcSlice) Less(i, j int) bool { return s[i].Value < s[j].Value }

func (g *Graph) HasNode(index int) bool {
	if g.NodeTree.Find(index) != nil {
		return true
	}
	return false
}

func (g *Graph) Related(from int, to int) bool {
	if node := g.NodeTree.Find(from); node != nil {
		if node.Out.Find(to) != nil {
			if g.Directed {
				return true
			} else {
				if t := g.NodeTree.Find(to); t != nil {
					if t.Out.Find(from) != nil {
						return true
					}
				}
			}

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

func (g *Graph) Relate(from, to int) bool {
	if !g.HasNode(from) || !g.HasNode(to) {
		return false
	}
	if g.Related(from, to) {
		return false
	} else {
		FROM := g.NodeTree.Find(from)
		TO := g.NodeTree.Find(to)
		if g.Directed {
			g.ArcSli = append(g.ArcSli, Arc{FROM, TO, 0})
		} else {
			g.ArcSli = append(g.ArcSli, Arc{FROM, TO, 0})
		}
		FROM.Out.Push(TO)
		FROM.In.Push(TO)
		TO.In.Push(FROM)
		TO.Out.Push(FROM)
	}
	return true
}
