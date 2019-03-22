package algorithmtest

import (
    "fmt"
    "github.com/ShiinaOrez/AlgorithmGo/tree"
    _ "github.com/ShiinaOrez/AlgorithmGo/typedefs"
)

func TreeTest() {
    tre := tree.CreateTree()
    rotNode := tree.BuildNode()
    rotNode.Value = 999
    err := rotNode.JoinTree(tre)
    CheckErr(err)
    err = tre.SetRootNode(rotNode)
    CheckErr(err)
    var sli []*tree.Node
    for i:=0; i<=10; i++ {
        newNode := tree.BuildNode()
        newNode.Value = (tree.Type)(i)
        newNode.JoinTree(tre)
        sli = append(sli, newNode)
        if i>= 1 {
            newNode.LinkTo(sli[newNode.Value/2])
            sli[newNode.Value/2].LinkTo(newNode)
        }
    }
    rotNode.LinkTo(sli[0])
    sli[0].LinkTo(rotNode)
    for _, i := range tre.Range() {
        fmt.Println(i.Value)
    }
}

func CheckErr (err error) {
    if err != nil {
        fmt.Println(err.Error())
    }
}
