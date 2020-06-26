package main

import (
	"fmt"
	//"golang.org/x/tools/container/intsets"
	"golang/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder(){
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

//func testSparse(){
//	s := intsets.Sparse{}
//
//	s.Intesrt(1)
//	s.Intesrt(1000)
//	s.Intesrt(1000000)
//	fmt.Println(s.Has(1000))
//	fmt.Println(s.Has(100000))
//
//}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Print()
	fmt.Println()

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	pRoot := &root
	pRoot.Print()
	println()
	pRoot.SetValue(200)
	pRoot.Print()
	println()

	var ppRoot *tree.Node
	ppRoot.SetValue(200)
	ppRoot = &root
	ppRoot.SetValue(300)
	ppRoot.Print()

	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	//testSparse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount ++
	})
	fmt.Println("Node count:", nodeCount)
}

