package main

import (
	"fmt"
)

type Number interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type TreeNode[K Number] struct {
	Data   K
	Left   *TreeNode[K]
	Right  *TreeNode[K]
	Parent *TreeNode[K]
}

func createNode[K Number](data K, parent *TreeNode[K]) *TreeNode[K] {
	return &TreeNode[K]{data, nil, nil, parent}
}

func NewTree[K Number](data K) *TreeNode[K] {
	return createNode[K](data, nil)
}

func (node *TreeNode[K]) GetLowestValue() K {
	root := node.GetRoot()
	lowest := root.traverseFullyLeft()
	return lowest.Data
}

func (node *TreeNode[K]) traverseFullyLeft() *TreeNode[K] {
	if node.Left == nil {
		return node
	} else {
		return node.Left.traverseFullyLeft()
	}
}

func (node *TreeNode[K]) Add(data K) {
	root := node.GetRoot()
	root.insertNode(data)
}

func (node *TreeNode[K]) GetRoot() *TreeNode[K] {
	if node.Parent == nil {
		return node
	} else {
		return node.Parent.GetRoot()
	}
}

func (node *TreeNode[K]) insertNode(data K) *TreeNode[K] {
	var newNode *TreeNode[K]
	if newNode != nil {
		return newNode
	} else {
		if data > node.Data {
			if node.Right == nil {
				newNode = createNode[K](data, node)
				node.Right = newNode
				return newNode
			}
			node.Right.insertNode(data)
		} else {
			if node.Left == nil {
				newNode = createNode[K](data, node)
				node.Left = newNode
				return newNode
			}
			node.Left.insertNode(data)
		}
	}

	return nil
}

func main() {
	root := NewTree(100)
	root.Add(2)
	root.Add(150)

	fmt.Println(root.GetLowestValue())
}
