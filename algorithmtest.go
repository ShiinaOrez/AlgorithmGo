package main

import (
    _"github.com/ShiinaOrez/AlgorithmGo/typedefs"
    "github.com/ShiinaOrez/AlgorithmGo/stack"
    "github.com/ShiinaOrez/AlgorithmGo/heap"
    "fmt"
)

func main() {
    /*Stack Test**/
    fmt.Printf("Stack Test:\n")
    stk := new(stack.Stack)
    var i stack.Type
    for i = 0; i<=8; i++ {
        stk.Push(i)
    }
    for !stk.Empty() {
        fmt.Printf("%d ", stk.Peek())
        stk.Pop()
    }
    fmt.Printf("\n------\n")

    /*Heap Test**/
    fmt.Printf("Heap Test:\n")
    h := new(heap.Heap)
    var j heap.Type
    for j = 0; j<=10; j++ {
        h.Push(j, func (a heap.Type, b heap.Type) bool {return a < b;})
    }
    for j = 20; j>=11; j-- {
        h.Push(j, func (a heap.Type, b heap.Type) bool {return a < b;})
    }
    h.Sort(func (a heap.Type, b heap.Type) bool {return a < b;})
}
