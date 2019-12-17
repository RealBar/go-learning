package main



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	p := root
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	for l1 != nil && l2 != nil{
		p.Next = new(ListNode)
		if l1.Val < l2.Val{
			p.Next.Val = l1.Val
			l1 = l1.Next
		}else{
			p.Next.Val = l2.Val
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil{
		p.Next = l1
	}else if l2 != nil{
		p.Next = l2
	}
	return root.Next
}