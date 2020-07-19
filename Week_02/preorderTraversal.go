package Week_02

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
给定一个二叉树，返回它的 前序 遍历。

 示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 */
var preRes []int

//递归
func preorderTraversal(root *TreeNode) []int {
	preRes = make([]int,0)
	preDfs(root)
	return res
}

func preDfs(root *TreeNode){
	if root != nil{
		preRes = append(preRes,root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
}

//迭代方法
func preorderTraversalStack(root *TreeNode) []int{
	var res []int
	var s []*TreeNode
	for len(s) > 0 || root != nil{
		for root != nil{
			res = append(res,root.Val)
			s = append(s,root.Right)
			root = root.Left
		}
		idx := len(s) - 1
		root = s[idx]
		s = s[:idx]
	}
	return res
}