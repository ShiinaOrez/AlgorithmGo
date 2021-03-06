package segmenttree

import (
	"errors"
	"fmt"
	"github.com/ShiinaOrez/AlgorithmGo/typedefs"
)

type Type typedefs.SegmentTreeType

type Node struct {
	Range  [2]int
	Value  Type
	Left   *Node
	Right  *Node
	Father *Node
}

type SegmentTree struct {
	Root *Node
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
			mid = (node.Range[1] - node.Range[0]) / 2
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
	_, err := tree.Root.Set(nil, sli, [2]int{0, len(sli) - 1})
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
				mid := node.Left.Range[1]
				if r[0] > mid {
					return node.Right.GetRangeSum(r)
				} else if r[1] < mid+1 {
					return node.Left.GetRangeSum(r)
				} else {
					leftSum, err := node.Left.GetRangeSum([2]int{r[0], mid})
					if err != nil {
						return 0, err
					}
					rightSum, err := node.Right.GetRangeSum([2]int{mid + 1, r[1]})
					if err != nil {
						return 0, err
					}
					return leftSum + rightSum, nil
				}
			}
		}
	}
	return 0, errors.New("[Error]: Something error internal.")
}

func (tree *SegmentTree) Query(r [2]int) (Type, error) {
	var res Type
	if tree.Root == nil {
		return 0, errors.New("[Error]: Tree haven't init!")
	} else {
		if tree.Root.Range[0] > r[0] || tree.Root.Range[1] < r[1] {
			return 0, errors.New("[Error]: Out of range!")
		} else {
			var err error
			res, err = tree.Root.GetRangeSum(r)
			if err != nil {
				return 0, err
			}
		}
	}
	return res, nil
}

func (node *Node) ItemAddValue(index int, value Type) error {
	if index < node.Range[0] || index > node.Range[1] {
		return errors.New("[Error]: Out of range!")
	} else {
		node.Value += value
		if node.Range[0] == node.Range[1] {
			if node.Range[0] == index {
				return nil
			}
		} else {
			if index <= node.Left.Range[1] {
				err := node.Left.ItemAddValue(index, value)
				if err != nil {
					return err
				}
			} else {
				err := node.Right.ItemAddValue(index, value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (tree *SegmentTree) ItemAddValue(index int, value Type) error {
	if tree.Root == nil {
		return errors.New("[Error]: Tree haven't init!")
	} else {
		if index < tree.Root.Range[0] || index > tree.Root.Range[1] {
			return errors.New("[Error]: Out of range!")
		} else {
			err := tree.Root.ItemAddValue(index, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (node *Node) RangeAddSameValue(r [2]int, value Type) error {
	if node.Range[0] > r[0] || node.Range[1] < r[1] {
		return errors.New("[Error]: Out of range!")
	} else {
		node.Value += Type((r[1] - r[0] + 1) * int(value))
		if node.Range[0] == node.Range[1] {
			return nil
		}
		if node.Left.Range[1] >= r[1] {
			err := node.Left.RangeAddSameValue(r, value)
			if err != nil {
				return err
			}
		} else if node.Right.Range[0] <= r[0] {
			err := node.Right.RangeAddSameValue(r, value)
			if err != nil {
				return err
			}
		} else {
			err := node.Left.RangeAddSameValue([2]int{r[0], node.Left.Range[1]}, value)
			if err != nil {
				return err
			}
			err = node.Right.RangeAddSameValue([2]int{node.Right.Range[0], r[1]}, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (tree *SegmentTree) RangeAddSameValue(r [2]int, value Type) error {
	if tree.Root == nil {
		return errors.New("[Error]: Tree haven't init!")
	} else {
		if r[0] < tree.Root.Range[0] || r[1] > tree.Root.Range[1] {
			return errors.New("[Error]: Out of range.")
		} else {
			err := tree.Root.RangeAddSameValue(r, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (node *Node) Print() {
	fmt.Printf("[%d - %d]: %d\n", node.Range[0], node.Range[1], node.Value)
	return
}

func (node *Node) ShowNode() {
	node.Print()
	if node.Left != nil {
		node.Left.ShowNode()
	}
	if node.Right != nil {
		node.Right.ShowNode()
	}
	return
}

func (tree *SegmentTree) ShowTree() {
	tree.Root.ShowNode()
}
