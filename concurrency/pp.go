package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //Creamos un canal para monitorear las Goroutines
	go doSomething(c)   //Llamamos a la función con go para generar un Goroutine
	pp := <-c
	fmt.Println(pp) //main esperara a que este canal reciba el mensaje

}

func doSomething(c chan int) { //Recibimos un canal de tipo int
	time.Sleep(3 * time.Second) //Poner a dormir
	fmt.Println("Done")
	c <- 1 // le envia el valor de 1
}
