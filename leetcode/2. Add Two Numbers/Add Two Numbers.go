package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var remainder, sum, nextValue int
	dummy := new(ListNode)
	node := dummy

	for l1 != nil || l2 != nil || remainder > 0 {
		sum = 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += remainder
		nextValue = sum % 10
		remainder = sum / 10

		node.Next = &ListNode{nextValue, nil}
		node = node.Next
	}

	return dummy.Next
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)

	for result != nil {
		print(result.Val)
		result = result.Next
	}
}
