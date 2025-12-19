package main

import "fmt"
import "slices"
import "cmp"

func main() {
	fmt.Println("Let sort some shit")

	//Basic sort
	strs := []string{"c", "a", "b"}
	slices.Sort(strs) //This is a inplace sort
	fmt.Println("Sorted strs => ", strs)

	//this is alos valid for numbers:
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Sorted ints => ", ints)

	//Is a slice sorted ?
	fmt.Println("Ints are sorted ? ", slices.IsSorted(ints))
	fmt.Println("Strs are sorted ? ", slices.IsSorted(strs))

	//Tambien se pueden ordenar slices en base a otros critarios usando "Funciones comparadoras"
	compareByLen := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	names := []string{"adonis", "juan", "nina", "manetha", "habilidad"}
	slices.SortFunc(names, compareByLen) //Esto tambien es un sort in place...
	fmt.Println(names)

	//Tambien se pueden ordenar objetos mas complejos en base a estas compare functions:
	type Person struct {
		name string
		age  uint
	}

	people := []Person{
		{name: "Adonis", age: 31},
		{name: "Manetha", age: 36},
		{name: "Habilidad", age: 29},
		{name: "Nina", age: 55},
		{name: "Juan", age: 65},
	}

	//Vease como se pueden usar las funciones anonimas como callbacks flow JS
	slices.SortFunc(people, func(p1, p2 Person) int {
		return cmp.Compare(p1.age, p2.age)
	})
	fmt.Println(people)

	//Tambien se pueden usar Sorted functions para no mutar el slice originar
	people2 := []Person{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 5},
		{"Vera", 24},
		{"Zac", 20},
	}

	//Este aproach es preferible si eres un FP lover.
	sortedPoeple := slices.SortedStableFunc(slices.Values(people2), func(p1, p2 Person) int {
		return cmp.Compare(p1.age, p2.age)
	})

	fmt.Println("Unsorted people => ", people2)
	fmt.Println("Sorted poeple => ", sortedPoeple)

}
