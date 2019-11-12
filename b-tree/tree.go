package btree

import (
	_ "fmt"
)

func NewNode(leaf bool) *Node {
	return &Node{
		Leaf: leaf,
		Nodes: make([]*Node, T+1),
	}
}

func NewBTree() *Tree {
	return &Tree{
		Root: NewNode(true),
		Count: 0,
	}
}

func (tree *Tree) Insert(key string, data interface{}) {
	root := tree.Root
	split, newEle, newL, newR := root.Insert(key, data)
	if split {
		newRoot := NewNode(false)
		newRoot.Elements = append(newRoot.Elements, newEle)
		newRoot.Nodes = make([]*Node, 2)
		newRoot.Nodes[0], newRoot.Nodes[1] = newL, newR
		tree.Root = newRoot
	}
	tree.Count += 1
}

func (node *Node) Insert(key string, data interface{}) (bool, Element, *Node, *Node) {
	if !node.Leaf {
		if len(node.Nodes) < node.N+1 {
			newNodes := make([]*Node, node.N+1)
			copy(newNodes, node.Nodes)
			node.Nodes = newNodes
		}
		for i, n := range node.Nodes {
			if n == nil {
				node.Nodes[i] = NewNode(true)
			}
		}
		if len(node.Elements) == 0 {
			node.Elements = []Element{Element{
				Key: key,
				Data: data,
			}}
			node.N += 1
		} else if key < node.Elements[0].Key {
			if split, newEle, newLeft, newRight := node.Nodes[0].Insert(key, data); split {
				node.Elements = append([]Element{newEle}, node.Elements...)
				node.Nodes[0] = newRight
				node.Nodes = append([]*Node{newLeft}, node.Nodes...)
				node.N += 1
			}
		} else {
			for index, ele := range node.Elements {
				if key > ele.Key {
					if node.Nodes[index+1] == nil {
						// fmt.Println(node.Elements, node.N, index, node.Nodes, "!!!")
					}
					if s, nE, nL, nR := node.Nodes[index+1].Insert(key, data); s {
						node.Elements = append(append(node.Elements[:index+1], nE), node.Elements[index+1:]...)
						node.Nodes[index+1] = nR
						node.Nodes = append(append(node.Nodes[:index+1], nL), node.Nodes[index+1])
						node.N += 1
					}
					break
				}
			}
		}
		if node.N >= T {
			nE, nL, nR := node.Split()
			return true, nE, nL, nR
		}
		return false, Element{}, nil, nil
	}
	// node is leaf
	if node.N == 0 {
		node.Elements = append(node.Elements, Element{
			Key:  key,
			Data: data,
		})
		node.N += 1
		// node.Nodes = make([]*Node, 2)
		return false, Element{}, nil, nil
	} else if node.N < T {
		for index, ele := range node.Elements {
			if ele.Key > key {
				// fmt.Println("	before eles:", node.Elements)
				leftPart := make([]Element, index+1)
				copy(leftPart, node.Elements[:index])
				leftPart[index] = Element{
					Key: key,
					Data: data,
				}
				node.Elements = append(leftPart, node.Elements[index:]...)
				// fmt.Println("	after eles: ", node.Elements)
				// node.Nodes = append(node.Nodes, &Node{})
				node.N += 1
				break
			}
		}
		if key > node.Elements[len(node.Elements)-1].Key {
			node.Elements = append(node.Elements, Element{
				Key: key,
				Data: data,
			})
			node.N += 1
		}
	}
	if node.N == T {
		nE, nL, nR := node.Split()
		return true, nE, nL, nR
	}
	return false, Element{}, nil, nil
}

func (node *Node) Split() (Element, *Node, *Node) {
	// fmt.Println("Split!")
	newLeftNode := &Node{
		N:        node.N/2,
		Leaf:     node.Leaf,
		Elements: make([]Element, node.N/2),
		Nodes:    make([]*Node, node.N/2),
	}
	newRightNode := &Node{
		N:        node.N-node.N/2-1,
		Leaf:     node.Leaf,
		Elements: make([]Element, node.N-node.N/2-1),
		Nodes:    make([]*Node, node.N-node.N/2+1),
	}
	// fmt.Println("Before:", node.Elements)
	// fmt.Println("New Father:", node.Elements[node.N/2])
	copy(newLeftNode.Elements, node.Elements[:node.N/2])
	// fmt.Println("New Left:", newLeftNode.Elements)
	copy(newRightNode.Elements, node.Elements[node.N/2+1:])
	// fmt.Println("New Right:", newRightNode.Elements)
	if !node.Leaf {
		copy(newLeftNode.Nodes, node.Nodes[:node.N/2+1])
		copy(newRightNode.Nodes, node.Nodes[node.N/2+1:])
	}
	return node.Elements[node.N/2], newLeftNode, newRightNode
}

func (tree *Tree) Search(key string) string {
	res := tree.Root.Search(key)
	if res != nil {
		return res.(string)
	}
	return "[Warn]: Value Not Found!"
}

func (node *Node) Search(key string) interface{} {
	// fmt.Println(node.Elements, "search", key)
	if node.N == 0 {
		return nil
	}
	if node.N == 1 {
		if node.Elements[0].Key == key {
			return node.Elements[0].Data
		}
		if !node.Leaf {
			if node.Elements[0].Key > key {
				if len(node.Nodes) == 0 {
					return nil
				}
				if node.Nodes[0] == nil {
					node.Nodes[0] = NewNode(true)
				}
				return node.Nodes[0].Search(key)
			} else {
				if len(node.Nodes) < 2 {
					return nil
				}
				if node.Nodes[1] == nil {
					node.Nodes[1] = NewNode(true)
				}
				return node.Nodes[1].Search(key)
			}
		}
		return nil
	}
	left := 0
	right := node.N-1
	if node.Elements[left].Key > key {
		if !node.Leaf {
			return node.Nodes[0].Search(key)
		}
	}
	if node.Elements[right].Key < key {
		if !node.Leaf {
			return node.Nodes[node.N].Search(key)
		}
	}
	for left < right {
		if left == right-1 {
			if key == node.Elements[left].Key {
				return node.Elements[left].Data
			}
			if key == node.Elements[right].Key {
				return node.Elements[right].Data
			}
			if !node.Leaf {
				return node.Nodes[right].Search(key)
			} else {
				return nil
			}
		} else {
			mid := (left+right)/2
			if node.Elements[mid].Key == key {
				return node.Elements[mid].Data
			} else if node.Elements[mid].Key < key {
				left = mid 
			} else {
				right = mid
			}
		}
	}
	return nil
}