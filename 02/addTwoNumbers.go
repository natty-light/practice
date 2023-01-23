package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var arr1, arr2 = []int{2, 4, 3}, []int{5, 6, 4}
	var l1, l2 = constructLinkedList(arr1), constructLinkedList(arr2)

	printLinkedList(l1)
	printLinkedList(l2)
	var sum *ListNode = addTwoNumbers(l1, l2)
	printLinkedList(sum)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var numOnePtr, numTwoPtr, sumPtr *ListNode
	var sumHead *ListNode = &ListNode{Val: 0}
	var value, x, y, carry int

	numOnePtr, numTwoPtr, sumPtr = l1, l2, sumHead
	carry = 0

	for numOnePtr != nil || numTwoPtr != nil {
		if numOnePtr == nil {
			x = 0
		} else {
			x = numOnePtr.Val
			numOnePtr = numOnePtr.Next
		}
		if numTwoPtr == nil {
			y = 0
		} else {
			y = numTwoPtr.Val
			numTwoPtr = numTwoPtr.Next
		}
		value = x + y + carry
		carry = value / 10
		sumPtr.Next = &ListNode{Val: value % 10}
		sumPtr = sumPtr.Next
	}
	if carry > 0 {
		sumPtr.Next = &ListNode{Val: carry}
	}
	return sumHead.Next
}

func constructLinkedList(num []int) *ListNode {
	var head *ListNode = &ListNode{}
	var ptr *ListNode = head
	for i, elem := range num {
		ptr.Val = elem
		if i < len(num) {
			ptr.Next = &ListNode{}
			ptr = ptr.Next
		}
	}
	return head
}

func printLinkedList(head *ListNode) {
	var ptr *ListNode = head
	for ptr.Next != nil {
		fmt.Printf("%d ", ptr.Val)
		ptr = ptr.Next
	}
	fmt.Println()
}
