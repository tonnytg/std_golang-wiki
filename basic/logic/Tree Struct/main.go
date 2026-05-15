package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	v := TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2} }
	v.Left = &TreeNode{Val: 0, Right: &TreeNode{Val: 1}}
	v.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}

	fmt.Println(v.Left.Val)
	fmt.Println(v.Val)
	fmt.Println(v.Right.Val)
}