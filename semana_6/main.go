package main

import (
	"fmt"
	"time"
)


var b bool = false
//gorutine
func Proceso(aux uint, c chan uint) {
	id := aux
	i := 0
	for { 
		select{
		case nulo:= <-c:
			if nulo==id{
				return//llena de vacios el arreglo
			}else{
				c<-nulo
			
			}
		default:
			if b {
				fmt.Println(id, ":", i)
			}
			i++
			time.Sleep(time.Millisecond * 500)
			
	}
}	
}

func main() {
	id := uint(0)
	c := make(chan uint)
	var eleccion int
	var ban bool = true
	for ban {
		menu :=
			`
Bienvenido a Gorutines
[ 1 ] Añadir
[ 2 ] Mostrar
[ 3 ] Terminar 
[ 4 ] Salir
¿Qué opcion quieres?
`
		fmt.Print(menu)
		fmt.Scanln(&eleccion)
		switch eleccion {
		
			case 1:
				//llamada gorutine
				go Proceso(id, c)
				id++
	
			case 2:
				b = !b
	
			case 3:
				
				var eliminar uint
				fmt.Print("ID a eliminar: ")
				fmt.Scanln(&eliminar)
				c <- eliminar
	
			case 4:
				ban = false
		}
	}
}
