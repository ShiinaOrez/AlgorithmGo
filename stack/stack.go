package stack

import (
	_"fmt"
	"github.com/ShiinaOrez/AlgorithmGo/typedefs"
)

type Type typedefs.StackType

type StackNode struct {
	Next   *StackNode
	Value  Type
}

type Stack struct {
	Top   *StackNode
	Size  int
}

func (stack *Stack) Push (value Type) int {
	newNode := new(StackNode)
	newNode.Value = value
	if stack.Top == nil {
		stack.Top = newNode
		stack.Size = 1
	}else{
		newNode.Next = stack.Top
		stack.Top = newNode
		stack.Size += 1
	}
	return stack.Size
}

func (stack *Stack) Pop () {
	if (stack.Top == nil) || (stack.Size == 0) {
		return 
	}else{
		if (stack.Top.Next == nil) || (stack.Size == 1) {
			stack.Top = nil
			stack.Size = 0
			return
		}else{
			stack.Top = stack.Top.Next
			stack.Size -= 1
		}
	}
	return
}

func (stack *Stack) Peek () Type{
	if (stack.Size > 0) && (stack.Top != nil) {
		return stack.Top.Value
	}else{
		return new(StackNode).Value
	}
}

func (stack *Stack) Empty () bool {
	if (stack.Size == 0) || (stack.Top == nil) {
		return true
	}else{
		return false
	}
}
