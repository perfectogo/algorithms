package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(tree2str(&TreeNode{1, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}, &TreeNode{7, nil, nil}}))
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	if root.Left == nil && root.Right == nil {
		return fmt.Sprint(root.Val, "")
	}
	if root.Right == nil {
		return fmt.Sprint(root.Val, "("+tree2str(root.Left)+")")
	}
	return fmt.Sprint(root.Val, "("+tree2str(root.Left)+")("+tree2str(root.Right)+")")
}
