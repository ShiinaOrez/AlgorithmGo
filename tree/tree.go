package tree

import  (
    "fmt"
    "github.com/ShiinaOrez/AlgorithmGo/bstree"
    "github.com/ShiinaOrez/AlgorithmGo/graph"
    "github.com/ShiinaOrez/AlgorithmGo/typedefs"
)

type Type typedefs.GraphType

type Tree struct {
    graph.Graph
    RootNode  int
}

func (tree *Tree) Dfs(from, now int) {
    // If now is root, please transform from as -1
    if tree.Directed {
        fmt.Println("Error: Tree is not a Directed Graph!")
        return
    }
    if nowNode := tree.NodeTree.Find(now); nowNode != nil {
        rv := nowNode.Out.Range()
        for _, node := range(rv) {
            if node == from {
                continue
            } else {
                tree.Dfs(now, node)
            }
        }
    }

}