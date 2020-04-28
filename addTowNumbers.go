package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	r := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: nil}}}
	print(addTwoNumbers(l, r))
}

//https://leetcode-cn.com/problems/add-two-numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}
	l := l1
	r := l2
	incr := false
	next := &result
	for {
		lval := 0
		rval := 0
		if l != nil {
			lval = l.Val
		}
		if r != nil {
			rval = r.Val
		}
		if incr {
			next.Val = lval + rval + 1
		} else {
			next.Val = lval + rval
		}
		incr = false
		if next.Val >= 10 {
			next.Val -= 10
			incr = true
		}
		if l != nil && l.Next != nil {
			l = l.Next
		} else {
			l = nil
		}
		if r != nil && r.Next != nil {
			r = r.Next
		} else {
			r = nil
		}
		if l == nil && r == nil {
			if incr {
				next.Next = &ListNode{Val: 1, Next: nil}
			} else {
				next.Next = nil
			}
			break
		}
		next.Next = &ListNode{}
		next = next.Next
	}
	return &result
}
