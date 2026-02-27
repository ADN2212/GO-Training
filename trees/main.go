package main

import "slices"
import "fmt"

type Color int

const (
	red Color = iota
	blue
)

type Node struct {
	id       uint
	name     string
	color    Color
	children []*Node
}

func innerSearch(node *Node, idsArray *[]uint) {
	if len(node.children) == 0 {
		return
	}
	for _, node := range node.children {
		if !slices.Contains(*idsArray, node.id) {
			*idsArray = append(*idsArray, node.id)
		}
	}
	for _, n := range node.children {
		innerSearch(n, idsArray)
	}
}

func (node *Node) descendants() []uint {
	var res = []uint{}
	innerSearch(node, &res)
	return res
}

func main() {

	fmt.Println("Trees!!!")

	bn1 := Node{id: 12, name: "Ismael", color: red, children: []*Node{}}
	bn2 := Node{id: 13, name: "Leo", color: red, children: []*Node{}}
	bn3 := Node{id: 14, name: "Martin", color: red, children: []*Node{}}
	bn4 := Node{id: 15, name: "Jhan", color: red, children: []*Node{}}

	nieto1 := Node{id: 7, name: "Erica", color: blue, children: []*Node{&bn1, &bn2}}
	nieto2 := Node{id: 8, name: "Eric", color: red, children: []*Node{&bn1, &bn2}}
	nieto3 := Node{id: 9, name: "Juan", color: red, children: []*Node{&bn3}}
	nieto4 := Node{id: 10, name: "Ana", color: blue, children: []*Node{&bn3, &bn4}}
	nieto5 := Node{id: 11, name: "Marcos", color: red, children: []*Node{&bn4}}

	hija1 := Node{id: 3, name: "Maria", color: blue, children: []*Node{&nieto1, &nieto2}}
	hijo1 := Node{id: 4, name: "Pedro", color: red, children: []*Node{&nieto1, &nieto2, &nieto3}}
	hija2 := Node{id: 5, name: "Juana", color: blue, children: []*Node{&nieto3, &nieto4, &nieto5}}
	hijo2 := Node{id: 6, name: "Paco", color: red, children: []*Node{&nieto4, &nieto5}}

	//append(padre.children, &hija1)

	padre := Node{id: 1, name: "Padrote", color: red, children: []*Node{&hija1, &hija2, &hijo1, &hijo2}}
	madre := Node{id: 2, name: "Madrota", color: blue, children: []*Node{&hija1, &hija2, &hijo1, &hijo2}}

	fmt.Println(padre)
	fmt.Println(madre)

	//Whole tree:
	fmt.Println(padre.descendants())

	//Tercer nivel:
	fmt.Println(nieto1.descendants())
	fmt.Println(nieto2.descendants())
	fmt.Println(nieto3.descendants())
	fmt.Println(nieto4.descendants())

	//Segundo nivel:
	fmt.Println(hija1.descendants())
	fmt.Println(hijo2.descendants())

}
