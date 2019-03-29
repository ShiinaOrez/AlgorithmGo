## AlgorithmGo.segmentree

-----

This package defines a simple **segment tree** data structure. Segment Tree in order to deal with the following problems: Query a range of data, CURD a range of data, and dynamic change the range value. Because of a range of linear like a segment(both have left endpoint and right endpoint), so it be called like segment tree. The idea of segment tree algorithm comes from the classic method of divide and conquer. So the time complexity of its single operation is O(log2 n).

In this version(2019.3.29v), my segment tree dosen't support the function of base a slice which have different data to update segment, so if you want to change it, please use **ItemAddValue** method.

In this package, I just provide you with some simple and basic method. If you want create more new method, please fork my code and edit it.

-----

## Package Information:

### Package: segmenttree

#### Type: Type
#### Type: Node
#### Type: SegmentTree
#### Mehtod: BuildNode() newNode *Node;
#### Mehtod: CreateSegmentTree() newTree *SegmentTree;
#### Method: SegmentTree.Init(sli []Type) error;
#### Method: Node.Set(from *Node, sli []Type, r [2]int) value Type, error;
#### Method: SegmentTree.Query(r [2]int) value Type, error;
#### Method: Node.GetRangeSum(r [2]int) value Type, error;
#### Method: SegmentTree.ItemAddValue(index int, value Type) error;
#### Method: Node.ItemAddValue(index int, value Type) bool;
#### Method: SegmentTree.RangeAddSameValue(r [2]int, value Type) error;
#### Method: Node.RangeAddSameValue(r [2]int, value Type) error;
#### Method: Node.Print()
#### Method: Node.ShowNode()
#### Method: SegmentTree.ShowTree()

-----

## Some Details:

This package just defined a **simple** segment tree, please add new method by your self.

Because of Go language dosen't support generic programming for now, so tree.Type needs you define it by yourself at typedefs package.

-----

## In the End:

### Contributor: ShiinaOrez
