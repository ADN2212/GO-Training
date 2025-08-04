package main

import (
	"fmt"
	"time"
	"sync"
)

//Ojo: La main function corre en su propia gorutine
func main() {
	//1-Uso basico de una gorutine.
	//go count("X")//Lanzar un gorutine permite que la gorutine principal pueda seguir su ejecucion antes de que esta termine.
	//count("Y")

	//2-Uso de los grupos de espera:
	var wg sync.WaitGroup
	wg.Add(1)//Aumenta el contador de gorutines en uno
	//Usar una IFE sirve para poder usar el Done justo despues de que termine la ejecucion del count
	go func ()  {
		count("Z", 5)
		wg.Done()//reduce el contador de gorutines en 1
	}()
	
	//Si se hicera go count("A", 100) sin usar los wg, la gorutine principal terminaria antes de que esta se completara.
	//go count("A", 100)
	wg.Add(1)
	go func ()  {
		count("A", 100)
		wg.Done()	
	}()
	wg.Wait()//Hace que la gorutine principal tenga que esperar a que las demas terminen antes de que se termine la ejecucion del programa.
}

func count(thing string, times int) {
	for i :=1; i<= times; i++ {
		fmt.Println(thing, "-", i)
		time.Sleep(time.Microsecond * 500)
	}
}
