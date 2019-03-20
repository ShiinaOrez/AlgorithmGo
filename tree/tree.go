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

func (tree *Tree) SetRootNode(node *Node) error {
    if node.Belong == tree{
        tree.Root = node
    } else {
        return nil
    }
    return errors.New("[Error]: Node must in this tree!")
}

func (node *Node) JoinTree(tree *Tree) error {
    if node.Belong == tree {
        return errors.New("[Error]: Node already in this tree!")
    } else {
        if node.Belong != nil {
            node.Belong.Size -= 1
        }
        node.Belong = tree
        tree.Size += 1
    }
    return nil
}