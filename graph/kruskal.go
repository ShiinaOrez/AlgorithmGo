package graph

import (
	"github.com/ShiinaOrez/AlgorithmGo/unionfindset"
	"sort"
)

func (g *Graph) MST() *Graph {
	if g.Directed {
		return nil
	} else {
		mst := new(Graph)
		mst.Directed = false
		mst.Size = g.Size
		mst.NodeTree = g.NodeTree
		ufs := new(unionfindset.UnionFindSet)
		sort.Sort(g.ArcSli)
		for _, arc := range(g.ArcSli) {
			if len(mst.ArcSli) == mst.Size-1 {
				break
			}
			if ufs.Find(arc.From.Index) != ufs.Find(arc.To.Index) {
				ufs.Union(arc.From.Index, arc.To.Index)
				mst.Relate(arc.From.Index, arc.To.Index)
			}
		}
		return mst
	}
}