package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ls *ListNode = &ListNode{0, nil}
	var lsAux *ListNode = ls
	remainder := 0
	if l1 != nil || l2 != nil {
		for {
			if l1 != nil {
				lsAux.Val = l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				lsAux.Val = lsAux.Val + l2.Val
				l2 = l2.Next
			}
			lsAux.Val = lsAux.Val + remainder

			if lsAux.Val > 9 {
				lsAux.Val = lsAux.Val - 10
				remainder = 1
			} else {
				remainder = 0
			}

			if l1 != nil || l2 != nil {
				lsAux.Next = &ListNode{0, nil}
				lsAux = lsAux.Next
			} else {
				break
			}
		}
	}

	if remainder > 0 {
		lsAux.Next = &ListNode{remainder, nil}
	}

	return ls
}

func main() {
	//l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}}
	//l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}
	l1 := ListNode{5, nil}
	l2 := ListNode{5, nil}
	fmt.Println("l1:", l1)
	fmt.Println("l2:", l2)
	var ls = addTwoNumbers(&l1, &l2)

	for ls != nil {
		fmt.Println("result: ", ls.Val, ls.Next)
		ls = ls.Next
	}

}
