package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val:  1,
	}
	cur := head
	for i := 2 ; i <=5 ; i++ {
		cur.Next = &ListNode{
			Val:  i,
		}
		cur = cur.Next
	}
	printList(head)
	head = removeNthFromEnd(head, 6)
	printList(head)
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := removeNthRec(head, 0, n)
	if length == n {
		return head.Next
	}
	return head
}

func removeNthRec(node *ListNode, cur, n int) int {
	if node == nil {
		return cur
	}
	var length int
	length = removeNthRec(node.Next, cur + 1, n)
	fmt.Println(cur, length)
	if cur == length - n - 1 {
		node.Next = node.Next.Next
	}
	return length
}
