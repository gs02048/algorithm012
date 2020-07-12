package Week_01

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preNode := &ListNode{}
	newNode := preNode
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			preNode.Next = l1
			l1 = l1.Next
		}else{
			preNode.Next = l2
			l2 = l2.Next
		}
		preNode = preNode.Next
	}
	if l1 != nil{
		preNode.Next = l1
	}
	if l2 != nil{
		preNode.Next = l2
	}
	return newNode.Next
}