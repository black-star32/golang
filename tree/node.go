package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

func (node treeNode)  print(){
	fmt.Print(node.value, " ")
}

func (node *treeNode) setValue(value int){
	if node == nil{
		fmt.Println("settting value to nil node, ignore")
		return
	}
	node.value = value
}

func createNode(value int) *treeNode{
	return &treeNode{value:value}
}

func (node *treeNode) traverse(){
	if node == nil{
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode
	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.print()
	fmt.Println()

	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()

	pRoot := &root
	pRoot.print()
	println()
	pRoot.setValue(200)
	pRoot.print()
	println()

	var ppRoot *treeNode
	ppRoot.setValue(200)
	ppRoot = &root
	ppRoot.setValue(300)
	ppRoot.print()

	root.traverse()
}
