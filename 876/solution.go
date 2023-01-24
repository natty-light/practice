package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	var head = buildLinkedList(slice)
	printLinkedList(head)
	fmt.Println(middleNode(head).Val)

	slice = []int{1, 2, 3, 4, 5, 6}
	head = buildLinkedList(slice)
	printLinkedList(head)
	fmt.Println(middleNode(head).Val)
}

func middleNode(head *ListNode) *ListNode {
	var fastPtr, slowPtr = head, head
	for fastPtr != nil && fastPtr.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}
	return slowPtr
}

func buildLinkedList(slice []int) *ListNode {
	var head *ListNode = &ListNode{}
	var ptr *ListNode = head
	for i, elem := range slice {
		ptr.Val = elem
		if i < len(slice)-1 {
			ptr.Next = &ListNode{}
			ptr = ptr.Next
		}
	}
	return head
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf(" %d ", head.Val)
		head = head.Next
	}
}
