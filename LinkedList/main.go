package main

import "fmt"

type Node struct {
	val      int
	nextNode *Node//Since we area passing &node when doing *&node = node so we area using a real node value, 
	// but we are using the value that is in the pointer.
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

//Para poder hacer cambios en la lista es necesario que los argumentos de la funcion sean
//punteros a los nodos, porque lo contrario se estarian editando copias de los nodos lo caul no generaria cambios en la lista original 
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
