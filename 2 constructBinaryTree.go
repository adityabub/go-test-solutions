package main

import "fmt"

//Binary Tree Node
type Node struct {
	data        int
	left, right *Node
}

//map for storing inorder traversal for easy search
var mp map[int]int

// crete new node
func newNode(data int) *Node {
	node := Node{data: data}
	return &node
}

// Recursive function to construct binary tree
// Initial: inStrt=0 inEnd=n-1
// No error checking
func buildUtil(in []int, post []int, inStrt int, inEnd int, pIndex *int) *Node {

	// Base case
	if inStrt > inEnd {
		return nil
	}

	// take current node from Postorder traversal using postIndex
	// decrement postIndex

	curr := post[*pIndex]

	node := newNode(curr)
	*pIndex--

	// If this node has no children then return
	if inStrt == inEnd {
		return node
	}

	// find index of node in Inorder traversal
	iIndex := mp[curr]

	//construct left and right subtress
	node.right = buildUtil(in[:], post[:], iIndex+1, inEnd, pIndex)
	node.left = buildUtil(in[:], post[:], inStrt, iIndex-1, pIndex)

	return node
}

// create map for easy serach of indexes of inorder traversal and call buildUtil
func buildTree(in []int, post []int, len int) *Node {
	mp = make(map[int]int)
	for i := range in {
		k := in[i]
		v := i
		mp[k] = v
	}

	index := len - 1 //index in postorder
	return buildUtil(in[:], post[:], 0, len-1, &index)
}

// only testing
func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	preOrder(node.left)
	preOrder(node.right)
}

func main() {
	in := [5]int{9, 3, 15, 20, 7}
	post := [5]int{9, 15, 7, 20, 3}

	root := buildTree(in[:], post[:], 5)

	fmt.Println("Preorder of the constructed tree")
	preOrder(root)
}
