package main

import "fmt"

type BinarySearchTree[T Number] interface {
	GetSmallestVal() T
	GetBiggestVal() T
	GetSmallestNode() *Node[T]
	GetBiggestNode() *Node[T]
	SortAscVals() []T
	SortDescVals() []T
	IsPresent(T) bool
	Add(T)
}

type Tree[T Number] struct {
	Root       *Node[T]
	TotalNodes int
}

func NewTree[T Number](data T) BinarySearchTree[T] {
	return Tree[T]{
		Root:       createNode[T](data, nil),
		TotalNodes: 1,
	}
}

func (t Tree[T]) GetSmallestVal() T {
	lowest := t.Root.traverseFullyLeft()
	return lowest.Data
}

func (t Tree[T]) GetSmallestNode() *Node[T] {
	return t.Root.traverseFullyLeft()
}

func (t Tree[T]) GetBiggestVal() T {
	biggest := t.Root.traverseFullyRight()
	return biggest.Data
}

func (t Tree[T]) GetBiggestNode() *Node[T] {
	return t.Root.traverseFullyRight()
}

func (t Tree[T]) SortAscVals() []T {
	vals := make([]T, 0)
	t.Root.ascOrderVals(&vals)
	return vals
}

func (t Tree[T]) SortDescVals() []T {
	vals := make([]T, 0)
	t.Root.descOrderVals(&vals)
	return vals
}

func (t Tree[T]) IsPresent(x T) bool {
	return t.Root.find(x)
}

func (t Tree[T]) Add(data T) {
	t.Root.insertNode(data)
}

func main() {
	root := NewTree(100)
	root.Add(2)
	root.Add(150)

	fmt.Println(root.IsPresent(150))
	fmt.Println(root.IsPresent(100))
	fmt.Println(root.IsPresent(2))
	fmt.Println(root.IsPresent(6))
	//fmt.Printf("%+v", root.NodesInAscOrder())
}
