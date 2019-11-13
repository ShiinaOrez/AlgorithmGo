package btree

import (
	"fmt"
	"math/rand"
	"testing"
)

func mock(n int) []string {
	res := []string{}
	for i := 0; i < 100; i++ {
		b := make([]byte, n)
		for i := range b {
			b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
		}
		res = append(res, string(b))
	}
	return res
}

func Test(t *testing.T) {
	tree := NewBTree()
	mockData := mock(4)
	resData := []string{}
	for _, str := range mockData {
		tree.Insert(str, str)
	}
	for _, str := range mockData {
		resData = append(resData, tree.Search(str))
	}
	fmt.Println(mockData)
	fmt.Println(resData)
}
