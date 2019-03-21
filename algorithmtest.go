package main

import (
    "fmt"
    "github.com/ShiinaOrez/AlgorithmGo/bstree"
    "github.com/ShiinaOrez/AlgorithmGo/graph"
    "github.com/ShiinaOrez/AlgorithmGo/heap"
    "github.com/ShiinaOrez/AlgorithmGo/queue"
    "github.com/ShiinaOrez/AlgorithmGo/stack"
    "github.com/ShiinaOrez/AlgorithmGo/stringmatch"
    _ "github.com/ShiinaOrez/AlgorithmGo/typedefs"
    "github.com/ShiinaOrez/AlgorithmGo/test"
    _ "reflect"
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

    /*Queue Test**/
    fmt.Printf("Queue Test:\n")
    quq := new(queue.Queue)
    var q queue.Type
    for q = 0; q<=8; q++ {
        quq.Push(q)
    }
    for !quq.Empty() {
        fmt.Printf("%d ", quq.Peek())
        quq.Pop()
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
    heapres := h.Sort(func (a heap.Type, b heap.Type) bool {return a < b;})
    for _, v := range(heapres) { fmt.Printf("%d ", v) }
    fmt.Printf("\n------\n")

    type I interface {
        Comp(*bstree.BSTree)  bool
    }

    /*Binary Sort Tree Test**/
    fmt.Printf("BinarySortTree Test:\n")
    bst := new(bstree.BSTree)
    var u bstree.Type
    for u = 0; u<=10; u++ {
        bst.Push(u, func (a bstree.Type, b bstree.Type) bool {return a < b;})
    }
    for u = 20; u>=11; u-- {
        bst.Push(u, func (a bstree.Type, b bstree.Type) bool {return a < b;})
    }
    res := bst.Sort()
    for _, v := range(res) { fmt.Printf("%d ", v) }
    fmt.Printf("\n------\n")

    /*Graph Test**/
    fmt.Printf("Graph Test:\n")
    g := new(graph.Graph)
    g.Insert(15, 3); g.Insert(2, 4); g.Insert(30, 1); g.Insert(22, 16)
    fmt.Println(g.Related(2, 30))
    g.Relate(2, 30)
    fmt.Println(g.Related(2, 30))
    fmt.Println(g.NodeTree.Find(2))

    master := "abc in abcd"
    temp := "n"
    exist, pos := stringmatch.Naive(master, temp)
    fmt.Println(exist, pos)
    exist, pos = stringmatch.Kmp(master, temp)
    fmt.Println(exist, pos)
    exist, pos = stringmatch.BoyerMoore(master, temp)
    fmt.Println(exist, pos)

    algorithmtest.RunTest()
}
