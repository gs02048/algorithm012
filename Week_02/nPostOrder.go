package Week_02
/*
590. N叉树的后序遍历

给定一个 N 叉树，返回其节点值的后序遍历。

例如，给定一个 3叉树 :

返回其后序遍历: [5,6,3,2,4,1].
说明: 递归法很简单，你可以使用迭代法完成此题吗?

 */
type Node struct {
    Val int
    Children []*Node
}

var nres []int
//递归
func postorder(root *Node) []int {
	ndfs(root)
	return nres
}

func ndfs(root *Node){
	if root != nil{
		for _,child := range root.Children{
			ndfs(child)
		}
		nres = append(nres,root.Val)
	}
}
