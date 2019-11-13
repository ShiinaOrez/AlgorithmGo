package bstree

import (
	"github.com/ShiinaOrez/AlgorithmGo/typedefs"
)

type Type typedefs.BSTreeType

type BSTreeNode struct {
	LeftNode  *BSTreeNode
	RightNode *BSTreeNode
	Value     Type
}

type BSTree struct {
	Root *BSTreeNode
	Size int
}

func BuildNode() *BSTreeNode {
	return new(BSTreeNode)
}

func (node *BSTreeNode) Comp(other *BSTreeNode) bool {
	return node.Value < other.Value
}

func (tree *BSTree) Push(value Type, comp func(Type, Type) bool) int {
	if (tree.Size == 0) || (tree.Root == nil) {
		tree.Root = BuildNode()
		tree.Root.Value = value
		tree.Size = 1
		return tree.Size
	} else {
		now := tree.Root
		for now != nil {
			if comp(value, now.Value) {
				if now.LeftNode != nil {
					now = now.LeftNode
				} else {
					now.LeftNode = BuildNode()
					now.LeftNode.Value = value
					break
				}
			} else {
				if now.RightNode != nil {
					now = now.RightNode
				} else {
					now.RightNode = BuildNode()
					now.RightNode.Value = value
					break
				}
			}
		}
		tree.Size += 1
		return tree.Size
	}
}

func (now *BSTreeNode) Traversal(res []BSTreeNode) []BSTreeNode {
	if now.LeftNode != nil {
		res = now.LeftNode.Traversal(res)
	}
	res = append(res, *now)
	if now.RightNode != nil {
		res = now.RightNode.Traversal(res)
	}
	return res
}

func (tree *BSTree) Find(target Type, comp func(Type, Type) bool) bool {
	now := tree.Root
	for now != nil {
		if now.Value == target {
			return true
		} else {
			if comp(target, now.Value) {
				now = now.LeftNode
			} else {
				now = now.RightNode
			}
		}
	}
	return false
}

func (tree *BSTree) Sort() []Type {
	var res []Type
	var rv []BSTreeNode
	rv = tree.Root.Traversal(rv)
	for _, node := range rv {
		res = append(res, node.Value)
	}
	return res
}
