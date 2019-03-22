## AlgorithmGo.tree

-----

This package defines a simple **tree** data structure. As we all know the tree is the most basic and very simple data structure. My tree could be a **directed** or **not directed**, its up to you. 

In this version(2019.3.21v), my tree dosen't support the function of determining loops, so when you create a new tree, please **take care of** it.

In this package, I just provide you with some simple and basic method. If you want create more new method, please fork my code and edit it.

-----

## Package Information:

### Package: tree

#### Type: Type
#### Type: Node
#### Type: Tree
#### Mehtod: BuildNode() newNode *Node;
#### Mehtod: CreateTree() newTree *Tree;
#### Method: Tree.SetRootNode(node *Node) error;
#### Method: Node.JoinTree(tree *Tree) error;
#### Method: Tree.DelNode(node *Node) error;
#### Method: Node.ClearLinks() error;
#### Method: Node.LinkTo(to *Node) error;
#### Method: Node.In(sli []*Node) bool;
#### Method: Node.Range(src []*Node, from *Node) sli []*Node;
#### Method: Tree.Range() sli []*Node;

-----

## Some Details:

This package just defined a simple tree, please add new method by your self.

Because of Go language dosen't support generic programming for now, so tree.Type needs you define it by yourself at typedefs package.

-----

## In the End:

### Contributor: ShiinaOrez
