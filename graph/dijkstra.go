package graph

func (index Type) getNode(g *Graph) *GraphNode {
	return g.NodeTree.Find(int(index))
}

func (node *GraphNode) getNode(g *Graph) *GraphNode {
	return node
}

type DijkstraArg interface {
	getNode( *Graph)   *GraphNode
}

//func (g *Graph) Dijkstra(from, to DijkstraArg) Type {
//	source := from.getNode(g)
//	target := to.getNode(g)
	
//}