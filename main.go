package main

import "fmt"

type BinarySearchTree[T Number, K any] interface {
	GetSmallestVal() T
	GetBiggestVal() T
	GetSmallestNode() *Node[T, K]
	GetBiggestNode() *Node[T, K]
	SortAscVals() []T
	SortDescVals() []T
	IsPresent(T) bool
	Add(T, K)
}

type Tree[T Number, K any] struct {
	Root       *Node[T, K]
	TotalNodes int
}

func NewTree[T Number, K any](key T, data K) BinarySearchTree[T, K] {
	return &Tree[T, K]{
		Root:       newNode[T](key, data, nil),
		TotalNodes: 1,
	}
}

func (t *Tree[T, K]) GetSmallestVal() T {
	lowest := t.Root.traverseFullyLeft()
	return lowest.key
}

func (t *Tree[T, K]) GetSmallestNode() *Node[T, K] {
	return t.Root.traverseFullyLeft()
}

func (t *Tree[T, K]) GetBiggestVal() T {
	biggest := t.Root.traverseFullyRight()
	return biggest.key
}

func (t *Tree[T, K]) GetBiggestNode() *Node[T, K] {
	return t.Root.traverseFullyRight()
}

func (t *Tree[T, K]) SortAscVals() []T {
	vals := make([]T, 0)
	t.Root.ascOrderVals(&vals)
	return vals
}

func (t *Tree[T, K]) SortDescVals() []T {
	vals := make([]T, 0)
	t.Root.descOrderVals(&vals)
	return vals
}

func (t *Tree[T, K]) IsPresent(x T) bool {
	return t.Root.find(x)
}

func (t *Tree[T, K]) Add(key T, data K) {
	root := t.Root.insertNode(key, data)
	if t.Root != root {
		//updateRoot[T, K](&t, root)
		t.Root = root
	}
}

//func updateRoot[T Number, K any](t *Tree[T, K], root *Node[T, K]) {
//	t.Root = root
//}

func main() {
	tree := NewTree(100, "something")
	tree.Add(2, "something else")
	tree.Add(6, "something else")
	tree.Add(1, "something else")
	tree.Add(5, "something else")
	//tree.Add(3, "something else")
	tree.Add(150, "")
	tree.Add(151, "")
	tree.Add(152, "")
	tree.Add(149, "")

	arr := tree.SortAscVals()
	for _, n := range arr {
		fmt.Println(n)
	}
	//fmt.Printf("%+v", tree.NodesInAscOrder())
}
