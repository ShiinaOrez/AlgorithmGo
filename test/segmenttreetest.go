package algorithmtest

import (
	"fmt"
	"github.com/ShiinaOrez/AlgorithmGo/segmenttree"
	"log"
)

func SegmentTreeTest() {
	fmt.Println("------")
	fmt.Println("SegmentTree Test:")

	tree := segmenttree.CreateSegmentTree()
	sli := []segmenttree.Type{4, 3, 62, 8, 2, 6, 3, 8, 1, 3, 64, 23, 88, 23}
	tree.Init(sli)
	fmt.Print(sli, " [3->7]Sum:")
	res, err := tree.Query([2]int{3, 7})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(res)
	}
	tree.ItemAddValue(5, 10)
	fmt.Print(" [1->5]Sum:")
	res, err = tree.Query([2]int{1, 5})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(res)
	}
	err = tree.RangeAddSameValue([2]int{5, 7}, 3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(" [1->5]Sum:")
	res, err = tree.Query([2]int{1, 5})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(res)
	}
}
