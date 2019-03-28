package segmenttree

import (
	"errors"
	"github.com/ShiinaOrez/AlgorithmGo/typedefs"
)

type Type typedefs.SegmentTreeType

type Node struct {
	Range     [2]int
	Value     Type
	Left      *Node
	Right     *Node
	Father    *Node
}

type SegmentTree struct {
	Root      *Node
}

func BuildNode() *Node {
	return new(Node)
}

func CreateSegmentTree() *SegmentTree {
	return new(SegmentTree)
}

func (node *Node) Set(from *Node, sli []Type, r [2]int) (Type, error) {
	if sli == nil {
		return 0, errors.New("[Error]: Can not use nil to set a segment tree's node.")
	} else {
		node.Father = from
		node.Range = r
		if node.Range[0] == node.Range[1] {
			if len(sli) == 1 {
				node.Value = sli[0]
			} else {
				return 0, errors.New("[Error]: Range error!")
			}
		} else {
			var mid int
			mid = (node.Range[0] + node.Range[1]) / 2
			node.Left = BuildNode()
			node.Right = BuildNode()
			leftSum, err := node.Left.Set(node, sli[:mid+1], [2]int{node.Range[0], node.Range[0] + mid})
			if err != nil {
				return 0, err
			}
			rightSum, err := node.Right.Set(node, sli[mid+1:], [2]int{node.Range[0] + mid + 1, node.Range[1]})
			if err != nil {
				return 0, err
			}
			node.Value = leftSum + rightSum
		}
	}
	return node.Value, nil
}

func (tree *SegmentTree) Init(sli []Type) error {
	if sli == nil {
		return errors.New("[Error]: Can not use nil to init segment tree!")
	}
	if tree.Root == nil {
		tree.Root = BuildNode()
	}
	_, err := tree.Root.Set(nil, sli, [2]int{0, len(sli)})
	if err != nil {
		return err
	}
	return nil
}

func (node *Node) GetRangeSum(r [2]int) (Type, error) {
	if node.Range[0] == node.Range[1] {
		if r[0] == r[1] {
			if r[0] == node.Range[0] && r[1] == node.Range[1] {
				return node.Value, nil
			} else {
				return 0, errors.New("[Error]: Internal error, smallest range not equal to smallest node.")
			}
		} else {
			return 0, errors.New("[Error]: Internal error, range error.")
		}
	} else {
		if r[0] < node.Range[0] || r[1] > node.Range[1] {
			return 0, errors.New("[Error]: Out of range!")
		} else {
			if r[0] == node.Range[0] && r[1] == node.Range[1] {
				return node.Value, nil
			} else {

			}
		}
	}
}

func (tree *SegmentTree) Query(r [2]int) (Type, error) {
	if tree.Root == nil {
		return 0, errors.New("[Error]: Tree haven't init!")
	} else {
		if tree.Root.Range[0] > r[0] || tree.Root.Range[1] < r[1] {
			return 0, errors.New("[Error]: Out of range!")
		} else {
			res, err := tree.Root.GetRangeSum(r)
		}
	}
}