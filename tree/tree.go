package tree

import (
    "github.com/ShiinaOrez/AlgorithmGo/typedefs"
    "github.com/kataras/iris/core/errors"
)

type Type typedefs.TreeType

type Node struct {
    Value      Type
    LinkNodes  []*Node
    Belong     *Tree
}

type Tree struct {
    Root       *Node
    Size       int
}


func BuildNode() *Node {
    node := new(Node)
    return node
}

func CreateTree() *Tree {
    tree := new(Tree)
    return tree
}

// SetRootNode
// @Description: Set a node to tree's root node, please make sure node belong to this tree
func (tree *Tree) SetRootNode(node *Node) error {
    if node.Belong == tree{
        tree.Root = node
    } else {
        return nil
    }
    return errors.New("[Error]: Node must in this tree!")
}

// JoinTree
// @Description: Join a node to a tree, if node used belong another tree, it will delete it.
func (node *Node) JoinTree(tree *Tree) error {
    if node.Belong == tree {
        return errors.New("[Error]: Node already in this tree!")
    } else {
        if node.Belong != nil {
            err := node.Belong.DelNode(node)
            if err != nil {
                return errors.New("[Error]: Delete node from tree failed!")
            }
        }
        node.Belong = tree
        tree.Size += 1
    }
    return nil
}

func (tree *Tree) DelNode(node *Node) error {
    if node.Belong != tree {
        return errors.New("[Error]: Before you delnode from a tree, please make sure node belong it.")
    } else {
        tree.Size -= 1
        err := node.CleanLinks()
        if err != nil {
            return errors.New("[Error]: Clean node's link failed")
        }
    }
    return nil
}

func (node *Node) CleanLinks() error {
    node.LinkNodes = nil
    return nil
}

func (node *Node) LinkTo(to *Node) error {
    if node.Belong == to.Belong && node.Belong != nil {
        if to.In(node.LinkNodes) {
            return errors.New("[Error]: Nodes already linked.")
        } else {
            node.LinkNodes = append(node.LinkNodes, to)
            return nil
        }
    } else if node.Belong == nil || to.Belong == nil {
        return errors.New("[Error]: Nodes must belong to a tree.")
    } else if node.Belong != to.Belong {
        return errors.New("[Error]: Nodes must belong to a same tree!")
    }
    return nil
}

func (node *Node) In(sli []*Node) bool{
    for _, per := range sli {
        if per == node {
            return true
        }
    }
    return false
}