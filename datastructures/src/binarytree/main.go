package main

// Keeps the node
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

//Keeps the root
type BinaryTree struct {
	root *BinaryNode
}

// Insert in Root
func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

// Imsert in Branches
func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	}

	if data < n.data {
		// Insert in left
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}

	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}

		} else {
			n.right.insert(data)
		}
	}

}

func print() {

}
func main() {
	tree := BinaryTree{} // Initialize interface
	tree.insert(100)     // Create Root Node
	tree.insert(25)
	tree.insert(55)
	tree.insert(15)
	tree.insert(5)
	//print(tree.root,)
}
