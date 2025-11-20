package main

import "fmt"

type Node struct {
	val      int
	nextNode *Node
}

func main() {

	node3 := Node{val: 3, nextNode: nil}
	node2 := Node{val: 2, nextNode: &node3}
	head1 := Node{val: 1, nextNode: &node2}

	//travers(head1)
	eraseNode(&head1, &head1)
	travers(head1)

	fmt.Println(sumLinkedListElement(head1))

}

// O(n)
func travers(head Node) {
	if head.nextNode != nil {
		fmt.Println(head.val)
		travers(*head.nextNode)
	} else {
		fmt.Println(head.val)
		fmt.Println("end")
	}
}

// O(n)
func sumLinkedListElement(head Node) int {
	if head.nextNode == nil {
		return head.val
	}
	return head.val + sumLinkedListElement(*head.nextNode)
}

func eraseNode(toErrase *Node, current *Node) {

	if current.val == toErrase.val {
		toErrase.nextNode = nil
		return
	}

	if current == nil {
		fmt.Println("Node not found")
		return
	}

	if current.nextNode.val == toErrase.val {
		current.nextNode = current.nextNode.nextNode
		return
	}

	eraseNode(toErrase, current.nextNode)

}
