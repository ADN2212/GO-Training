package main

import (
	"context"
	"fmt"
	//"golang.org/x/text/cases"
	"time"
)

func main() {

	//Este es una especia de contexto padre,
	ctx := context.Background()

	//Ejemplo - 1: Pasando valores usando contextos.

	//Este segundo es un contexto hijo, del contexto principal
	//ctx2 := context.WithValue(ctx, "bar", 10)//porque no se pueden user strs como claves.
	//performTask(ctx2)

	//Ejemplo - 2: Cancelando un contexto para evitar que se siga pasando informacion a traves de el
	// ctx3, cancel := context.WithCancel(ctx)
	// //Ojo, si el contexto padre de un contexto es cancelado esta tambien sera cancelado!
	// ctx4 := context.WithValue(ctx3, "foo", 10)
	// go performTask2(ctx4)
	// time.Sleep(2 * time.Second)
	// cancel()//Esto detendra la ejecucion de la go rutine ...
	// time.Sleep(1 * time.Second)

	//Ejemplo - 3: Contexto con un timpo de vida definido
	ctx4, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()
	go performTask3(ctx4)
	time.Sleep(6 * time.Second)

}

// func performTask(ctx context.Context) {
// 	bar := ctx.Value("bar").(int)//Al parecer el tipo no puede ser inferido, por eso el type assertion
// 	//El programa se rompe si el tipo no es el esperado.
// 	fmt.Printf("This value was retrived from the context => %v \n", bar)
// }

// func performTask2(ctx context.Context) {
// 	for {
// 		select {
// 		case <- ctx.Done():
// 			fmt.Println("Tarea canselada X")
// 			return
// 		default:
// 			//En caso deque el contexto no haya sido cancelado se ejecuta una tarea
// 			fmt.Println(ctx.Value("foo"))
// 			//Esto simula el tiempo de ejecucion de la task
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}
// }

func performTask3(ctx context.Context) {
	//I asume that this patter is usual
	//A while true like loop that gets data for a ctx unitl it is done.
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Tarea completada o tiempo limite completado => ", ctx.Err().Error())
			return
		default:
			fmt.Println("Ejecutando tarea lalala lalala ...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
