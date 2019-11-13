package graph

type BSTreeNode struct {
	LeftNode  *BSTreeNode
	RightNode *BSTreeNode
	Value     *GraphNode
}

type BSTree struct {
	Root *BSTreeNode
	Size int
}

func BuildNode() *BSTreeNode {
	return new(BSTreeNode)
}

func (node *BSTreeNode) Comp(other *BSTreeNode) bool {
	return node.Value.Index < other.Value.Index
}

func (tree *BSTree) Push(value *GraphNode) int {
	if (tree.Size == 0) || (tree.Root == nil) {
		tree.Root = BuildNode()
		tree.Root.Value = value
		tree.Size = 1
		return tree.Size
	} else {
		now := tree.Root
		for now != nil {
			if value.Index < now.Value.Index {
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

func (tree *BSTree) Find(index int) *GraphNode {
	now := tree.Root
	for now != nil {
		if now.Value.Index == index {
			return now.Value
		} else {
			if index < now.Value.Index {
				now = now.LeftNode
			} else {
				now = now.RightNode
			}
		}
	}
	return nil
}
