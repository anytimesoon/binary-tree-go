package main

type Number interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Node[T Number, K any] struct {
	key    T
	data   K
	left   *Node[T, K]
	right  *Node[T, K]
	parent *Node[T, K]
	level  int
}

func (node *Node[T, K]) GetRoot() *Node[T, K] {
	if node.parent == nil {
		return node
	} else {
		return node.parent.GetRoot()
	}
}

func newNode[T Number, K any](key T, data K, parent *Node[T, K]) *Node[T, K] {
	return &Node[T, K]{key, data, nil, nil, parent, 0}
}

func (node *Node[T, K]) traverseFullyLeft() *Node[T, K] {
	if node.left == nil {
		return node
	} else {
		return node.left.traverseFullyLeft()
	}
}

func (node *Node[T, K]) traverseFullyRight() *Node[T, K] {
	if node.right == nil {
		return node
	} else {
		return node.right.traverseFullyRight()
	}
}

func (node *Node[T, K]) ascOrderVals(vals *[]T) {
	if node != nil {
		node.left.ascOrderVals(vals)
		*vals = append(*vals, node.key)
		node.right.ascOrderVals(vals)
	}
}

func (node *Node[T, K]) descOrderVals(vals *[]T) {
	if node != nil {
		node.right.descOrderVals(vals)
		*vals = append(*vals, node.key)
		node.left.descOrderVals(vals)
	}
}

func (node *Node[T, K]) insertNode(key T, data K) *Node[T, K] {
	if node.key == key {
		return node
	}

	switch {
	case key < node.key && node.left != nil:
		node.left.insertNode(key, data)
	case key < node.key && node.left == nil:
		node.left = newNode[T, K](key, data, node)
	case key > node.key && node.right != nil:
		node.right.insertNode(key, data)
	case key > node.key && node.right == nil:
		node.right = newNode[T, K](key, data, node)
	}

	node.level = max(node.left.height(), node.right.height()) + 1

	return node.rebalance()
}

func (node *Node[T, K]) find(x T) bool {
	var next *Node[T, K]

	if node.key == x {
		return true
	}

	if x < node.key && node.left != nil {
		next = node.left
	} else if x > node.key && node.right != nil {
		next = node.right
	} else {
		return false
	}

	return next.find(x)
}

func (node *Node[T, K]) rebalance() *Node[T, K] {
	switch {
	case node.balance() < -1 && node.left.balance() == -1:
		return node.rotateRight()
	case node.balance() > 1 && node.right.balance() == 1:
		return node.rotateLeft()
	case node.balance() < -1 && node.left.balance() == 1:
		return node.rotateLeftRight()
	case node.balance() > 1 && node.right.balance() == -1:
		return node.rotateRightLeft()
	}
	return node
}

func max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func (node *Node[T, K]) balance() int {
	return node.right.height() - node.left.height()
}

func (node *Node[T, K]) rotateRight() *Node[T, K] {
	l := node.left
	node.left = l.right
	l.right = node
	l.parent = node.parent
	node.parent = l
	node.level = max(node.left.height(), node.right.height()) + 1
	l.level = max(l.left.height(), l.right.height()) + 1
	return l
}

func (node *Node[T, K]) rotateLeft() *Node[T, K] {
	r := node.right
	node.right = r.left
	r.left = node
	r.parent = node.parent
	node.parent = r
	node.level = max(node.left.height(), node.right.height()) + 1
	r.level = max(r.left.height(), r.right.height()) + 1
	return r
}

func (node *Node[T, K]) rotateLeftRight() *Node[T, K] {
	node.left = node.left.rotateLeft()
	n := node.rotateRight()
	n.level = max(node.left.height(), node.right.height()) + 1
	return n
}

func (node *Node[T, K]) rotateRightLeft() *Node[T, K] {
	node.right = node.right.rotateRight()
	n := node.rotateLeft()
	n.level = max(node.left.height(), node.right.height()) + 1
	return n
}

func (node *Node[T, K]) height() int {
	if node == nil {
		return 0
	}
	return node.level
}
