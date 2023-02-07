package main

type Number interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Node[T Number] struct {
	Data   T
	Left   *Node[T]
	Right  *Node[T]
	Parent *Node[T]
}

func (node *Node[T]) GetRoot() *Node[T] {
	if node.Parent == nil {
		return node
	} else {
		return node.Parent.GetRoot()
	}
}

func createNode[T Number](data T, parent *Node[T]) *Node[T] {
	return &Node[T]{data, nil, nil, parent}
}

func (node *Node[T]) traverseFullyLeft() *Node[T] {
	if node.Left == nil {
		return node
	} else {
		return node.Left.traverseFullyLeft()
	}
}

func (node *Node[T]) traverseFullyRight() *Node[T] {
	if node.Right == nil {
		return node
	} else {
		return node.Right.traverseFullyRight()
	}
}

func (node *Node[T]) ascOrderVals(vals *[]T) {
	if node != nil {
		node.Left.ascOrderVals(vals)
		*vals = append(*vals, node.Data)
		node.Right.ascOrderVals(vals)
	}
}

func (node *Node[T]) descOrderVals(vals *[]T) {
	if node != nil {
		node.Right.descOrderVals(vals)
		*vals = append(*vals, node.Data)
		node.Left.descOrderVals(vals)
	}
}

func (node *Node[T]) insertNode(data T) *Node[T] {
	var newNode *Node[T]
	if newNode != nil {
		return newNode
	} else {
		if data > node.Data {
			if node.Right == nil {
				newNode = createNode[T](data, node)
				node.Right = newNode
				return newNode
			}
			node.Right.insertNode(data)
		} else {
			if node.Left == nil {
				newNode = createNode[T](data, node)
				node.Left = newNode
				return newNode
			}
			node.Left.insertNode(data)
		}
	}

	return nil
}
