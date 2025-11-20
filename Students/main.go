package main

import "fmt"

type Student struct {
	fullName string
	mat string
	average float32
	next *Student 
}

func main() {

	st2 := Student{fullName: "Nina", mat: "0202020", average: 85.5, next: nil}
	st1 := Student{fullName: "Juan", mat: "11010101", average: 75.5, next: nil}
	st3 := Student{fullName: "Eric", mat: "09090909", average: 65, next: nil}

	st1.next = &st2
	st2.next = &st3

	showStudents(st1)

	fmt.Println("------------------------------")

	fmt.Println(howManyPassed(st1))
}

func showStudents(stu Student) {
	if stu.next == nil {
		fmt.Printf("name: %v, mat: %v, average: %v \n", stu.fullName, stu.mat, stu.average)
		return
	}

	fmt.Printf("name: %v, mat: %v, average: %v \n", stu.fullName, stu.mat, stu.average)
	showStudents(*stu.next)

}

func howManyPassed(stu Student) int {
	if stu.next == nil {
		if stu.average >= 70 {
			fmt.Printf("name: %v, mat: %v, average: %v \n", stu.fullName, stu.mat, stu.average)
			return 1
		}
		return 0
	}

	if stu.average >= 70 {
		fmt.Printf("name: %v, mat: %v, average: %v \n", stu.fullName, stu.mat, stu.average)
		return 1 + howManyPassed(*stu.next)
	} else {
		return 0 + howManyPassed(*stu.next)
	}

}
