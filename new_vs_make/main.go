package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {

	fmt.Println("Let's take some anotations about new and make")

	//La funcion new(T) retorna un *T
	intPtr := new(int)
	fmt.Println(intPtr)
	fmt.Println(*intPtr) //vease como el valor al que apunte el puntero es el zero value del tipo

	pointPtr := new(Point)
	fmt.Println(pointPtr.x)//Lo mismo pasa con los valores de las structs
	
	//La funcion new es poco usada porque hay maneras mas idiomaticas de lograr lo mismo de forma mas idiomatica
	p2Ptr := &Point{}//aqui queda mas claro que el resultado es un puntero a un Point "vacio"
	fmt.Println(p2Ptr.x, p2Ptr.x)//Vease como se consigue el mosmo resultado.

	//Problemas:
	//m := *new(map[string]int)
	//m["ADN"]=10 con esta linea the codigo el programa entra en panico, porque el mapa es nil,
	//new creo una direccion en memoria pero el no inicialoizo el valor.
	//esto es porque el zeroed value de un map (o slice o chanel) es nil  

	//La funcion make resulve esto
	//make solo funciona para slice, map y chanel
	//make inicializa, es decir, le asigna un valor a lo que esta en el puntero, no lo deja como nil
	//ademas retorna directamente ese valor

	s := make([]int, 10, 100)//Inicializa un slice de 10 elementos que apunta a un array de 100 elementos
	s[5] = 150//vease como esto no rompe nanda, porque s no es nil, si hubieramos usado new, the program will crash in this line.
	fmt.Println(s)

	//Lo mismo pero con el map
	m := make(map[string]int)
	//ojo, esto es equivalente a hacer m := map[string]int{}
	m["ADN"] = 32//No problem here, the map is already inicialized.
	fmt.Println(m)

	//En resumen:
	//new crea un puntero a un valor del tipo que se le da como argumento,
	//este valor esta inizialodado como su zero value correspondiente, int => 0, strig => "", etc ...
	//esto puede genera problemas con valores que apuntan a estrucutes subyacentes
	//como los slice, los maps y los maps
	//Ejemplo
	s2Ptr := new([]int)
	s2 := *s2Ptr
	//Vease como aqui la capacity de s2 es 0 lo que quiere decir que en este punto su underlaying array es un array vacio
	//por eso esta linea rompera el programa s2[10] = 10
	fmt.Println(cap(s2))





}
