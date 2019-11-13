package btree

import (
	"fmt"
)

type Element struct {
	Key  string
	Data interface{}
}

type Output interface {
	String() string
}

type Input interface {
	Update(string, interface{})
}

type Node struct {
	N        int
	Leaf     bool
	Elements []Element // len(Keys) == N
	Nodes    []*Node
}

type Tree struct {
	Root  *Node
	Count int
}

func (ele *Element) String() string {
	return fmt.Sprintf("[%s]: %s", ele.Key, ele.Data.(string))
}

func (ele *Element) Update(key string, data interface{}) {
	ele.Key = key
	ele.Data = data
}
