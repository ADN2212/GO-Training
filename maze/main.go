package main

import "fmt"
import "bufio"
import "os"
import "strings"

type postion struct {
	row int
	col int
}

func main() {

	var maze [10][10]string

	maze[0] = [10]string{"#", "#", "$", "#", "#", "#", "#", "#", "#", "#"}
	maze[1] = [10]string{"#", "#", " ", " ", " ", " ", " ", " ", " ", "#"}
	maze[2] = [10]string{"#", "#", "#", " ", " ", "#", " ", "#", " ", "#"}
	maze[3] = [10]string{"#", "#", "#", " ", "#", "#", " ", "#", " ", "#"}
	maze[4] = [10]string{"#", " ", " ", " ", "#", "#", " ", "#", " ", "#"}
	maze[5] = [10]string{"#", "#", "#", " ", "#", " ", " ", "#", " ", "#"}
	maze[6] = [10]string{"#", "#", " ", " ", "#", " ", "#", "#", " ", "#"}
	maze[7] = [10]string{"#", "#", " ", "#", " ", " ", "#", "#", " ", "#"}
	maze[8] = [10]string{"#", " ", "#", "#", " ", "#", "#", "#", " ", "#"}
	maze[9] = [10]string{"#", "#", "#", "#", " ", "#", "#", "#", "#", "#"}

	fmt.Println("Andel diablo estamos atrapados en un laberinto, busquemos una forma de escapar ...")
	fmt.Println("pulsa 'i' para subir, 'k' para bajar y 'j', 'l' para ir a la izquerda y derecha respectivamente.")

	showMaze(&maze)

	winPos := postion{row: 9, col: 4}
	prevPos := postion{row: 0, col: 2}
	pos := postion{row: 0, col: 2}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Has tu movimiento: ")

		char, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Ha ocurrido un error raro ...")
			return
		}

		char = strings.Replace(char, "\n", "", -1)

		switch char {
		case "i":
			move(&pos, "up")
		case "k":
			move(&pos, "down")

		case "j":
			move(&pos, "left")

		case "l":
			move(&pos, "rigth")

		default:
			fmt.Println("Maldito animal, lee bien las instrucciones mmg")
			continue
		}

		if pos.col < 0 || pos.col > 9 || pos.row < 0 || pos.row > 9 {
			fmt.Println("Compai, ute eh vico? se salio del bloque.")
			break
		}

		if maze[pos.row][pos.col] == "#" {
			maze[prevPos.row][prevPos.col] = " "
			maze[pos.row][pos.col] = "X"
			showMaze(&maze)
			fmt.Println("Usted se la acaba de beber")
			break
		}

		if maze[pos.row][pos.col] == " " {
			maze[prevPos.row][prevPos.col] = " "
			maze[pos.row][pos.col] = "$"
			prevPos.row = pos.row
			prevPos.col = pos.col
			showMaze(&maze)
		}

		if pos.col == winPos.col && pos.row == winPos.row {
			maze[pos.row][pos.col] = "♡"
			showMaze(&maze)
			fmt.Println("Muy duro miop, ganaste")
			break
		}
	}
}

func showMaze(m *[10][10]string) {
	var currRowStr string
	for i := range m {
		for j := range m[i] {
			currRowStr += m[i][j]
		}
		fmt.Println(currRowStr)
		currRowStr = ""
	}
}

func move(pos *postion, direction string) {
	switch direction {
	case "up":
		pos.row -= 1
	case "down":
		pos.row += 1
	case "left":
		pos.col -= 1
	case "rigth":
		pos.col += 1
	}
}
