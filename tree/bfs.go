package tree

import (
	"errors"
)

type BFSNode struct {
	node    *Node
	from    int
}

func (tree *Tree) BFStep(now *Node, queue []BFSNode, index int) ([]BFSNode, error) {
	if queue[index].node != now {
		return nil, errors.New("[Error]: Something error internal!")
	}
	if now.Belong != tree {
		return nil, errors.New("[Error]: BFS now node must belong to this tree!")
	} else {
		for _, nod := range now.LinkNodes {
			if queue[queue[index].from].node == nod {
				continue
			} else {
				queue = append(queue, BFSNode{nod, index})
			}
		}
	}
	return queue, nil
}

func (tree *Tree) BFS() ([]*Node, error) {
	if tree.Root == nil {
		return nil, errors.New("[Error]: If you want use method BFS(), please make sure this tree has root node!")
	}
	var queue []BFSNode
	var index = 0
	var before = 0
	var err error
	queue = append(queue, BFSNode{tree.Root, before})
	for index<len(queue) {
		queue, err = tree.BFStep(queue[index].node, queue, index)
		if err != nil {
			return nil, err
		}
		index += 1
	}
	var res []*Node
	for _, q := range queue {
		res = append(res, q.node)
	}
	return res, nil
}