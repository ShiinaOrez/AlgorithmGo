//Min-Root-Heap
package heap

import (
	"github.com/ShiinaOrez/AlgorithmGo/stack"
	"github.com/ShiinaOrez/AlgorithmGo/typedefs"
)

type Type typedefs.HeapType

type Heap struct {
	Root *HeapNode
	Size int
}

type HeapNode struct {
	Father  *HeapNode
	LeftNode *HeapNode
	RightNode *HeapNode
	Value Type
}

func BuildNode () *HeapNode {
    return new(HeapNode)
}

func Heaplify (now *HeapNode, value Type, comp func(Type, Type) bool) {
	var rv Type
    if comp(value, now.Value) {
		rv = now.Value
		now.Value = value
	}else{
		rv = value
	}
	if now.LeftNode == nil {
		now.LeftNode = BuildNode()
		now.LeftNode.Value = rv
		return
	}else{
		if now.RightNode == nil {
			now.RightNode = BuildNode()
			now.RightNode.Value = rv
			return
		}else{
			if comp(now.LeftNode.Value, now.RightNode.Value) {
				Heaplify(now.LeftNode, rv, comp)
			}else{
				Heaplify(now.RightNode, rv, comp)
			}
		}
	}
}

func HeapUp (now *HeapNode, comp func(Type, Type) bool) {
	if now.Father == nil {
		return
	}else{
		if comp(now.Father.Value, now.Value) {
			return
		}else{
			now.Father.Value, now.Value = now.Value, now.Father.Value
			HeapUp(now.Father, comp)
			return
		}
	}
}

func HeapDown (now *HeapNode, comp func(Type, Type) bool) {
	if (now.LeftNode == nil) && (now.RightNode == nil) {
		return
	}
	if (now.LeftNode != nil) && (now.RightNode != nil) {
		if comp(now.LeftNode.Value, now.RightNode.Value) {
			now.Value, now.LeftNode.Value = now.LeftNode.Value, now.Value
			HeapDown(now.LeftNode, comp)
		}else{
			now.Value, now.RightNode.Value = now.RightNode.Value, now.Value
			HeapDown(now.RightNode, comp)
		}
	}else{
		if (now.RightNode == nil) && comp(now.LeftNode.Value, now.Value) {
			now.Value, now.LeftNode.Value = now.LeftNode.Value, now.Value
			HeapDown(now.LeftNode, comp)
		}
		if (now.LeftNode == nil) && comp(now.RightNode.Value, now.Value){
			now.Value, now.RightNode.Value = now.RightNode.Value, now.Value
			HeapDown(now.RightNode, comp)
		}
	}
	return
}

func (heap *Heap) Pop (comp func(Type, Type) bool) {
	if heap.Size <= 0 {
		return
	}else{
		if heap.Size == 1{
			heap.Root = nil
			heap.Size = 0
			return
		}
		stk := new(stack.Stack)
		num := heap.Size
		for num > 0 {
			stk.Push(stack.Type(num&1))
			num /= 2
		}
		now := heap.Root
		stk.Pop()
		for !stk.Empty() {
			if stk.Size == 1 {
				if stk.Peek() == 1 {
					heap.Root.Value = now.RightNode.Value
					now.RightNode = nil
				}else{
					heap.Root.Value = now.LeftNode.Value
					now.LeftNode = nil
				}
			}
			if stk.Peek() == 1{
				now = now.RightNode
			}else{
				now = now.LeftNode
			}
			stk.Pop()
		}
		heap.Size -= 1
		HeapDown(heap.Root, comp)
	}
	return
}

func (heap *Heap) Peek () Type {
    if heap.Root == nil{
		return *new(Type)
	}else{
		return heap.Root.Value
	}
}

func (heap *Heap) Push (value Type, comp func(Type, Type) bool) int {
	heap.Size += 1
	if heap.Root == nil {
		heap.Root = BuildNode()
		heap.Size = 1
		heap.Root.Value = value
		return heap.Size
	}else{
		stk := new(stack.Stack)
		num := heap.Size
		for num > 0 {
			stk.Push(stack.Type(num&1))
			num /= 2
		}
		now := heap.Root
		stk.Pop()
		for !stk.Empty() {
			if stk.Size == 1 {
				if stk.Peek() == 1 {
					now.RightNode = new(HeapNode)
					now.RightNode.Father = now
				}else{
					now.LeftNode = new(HeapNode)
					now.LeftNode.Father = now
				}
			}
			if stk.Peek() == 1{
				now = now.RightNode
			}else{
				now = now.LeftNode
			}
			stk.Pop()
		}
		now.Value = value
		HeapUp(now, comp)
	}
	return heap.Size
}

func (heap *Heap) Sort (comp func(Type, Type) bool) []Type{
	size := heap.Size
	var res []Type
	for i := 1; i <= size; i++ {
		res = append(res, heap.Peek())
		heap.Pop(comp)
	}
	return res
}