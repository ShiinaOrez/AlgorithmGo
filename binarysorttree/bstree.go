package bstree

import (
	_"fmt"
	"github.com/ShiinaOrez/AlgorithmGo/typedefs"
)

type Type typedefs.BSTreeType

type BSTreeNode struct {
	LeftNode  *BSTreeNode
	RightNode *BSTreeNode
	Value     Type
}

type BSTree struct {
	Root      *BSTreeNode
	Size      int
}