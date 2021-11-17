package main

import (
	"./composicion"
	"fmt"
)

func main() {
	var eleccion int
	var Con composicion.ContenidoWeb
	var titulo, formato, canales string
	var titulo1, formato1, duracion string
	var titulo2, formato2, frame string
	var ban bool = true
	Con = composicion.ContenidoWeb{
		Multi: []composicion.Multimedia{},
	}
	for ban {
	menu :=
		`
    Bienvenido a Funciones
    [ 1 ] Imagen
    [ 2 ] Audio
    [ 3 ] Video
    [ 4 ] Mostrar todo
	[ 5 ] Salir
    ¿Qué opcion quieres?
`
	fmt.Print(menu)
	
	fmt.Scanln(&eleccion)
	
		switch eleccion {
		case 1:
			fmt.Print("Titulo: ")
			fmt.Scanln(&titulo)
			fmt.Print("Formato: ")
			fmt.Scanln(&formato)
			fmt.Print("Canales: ")
			fmt.Scanln(&canales)
			Con.Multi = append(Con.Multi, &composicion.Imagen{titulo, formato, canales})
			
		case 2:
			fmt.Print("Titulo: ")
			fmt.Scanln(&titulo1)
			fmt.Print("Formato: ")
			fmt.Scanln(&formato1)
			fmt.Print("Duracion: ")
			fmt.Scanln(&duracion)
			Con.Multi = append(Con.Multi, &composicion.Audio{titulo1, formato1, duracion})
		case 3:
			fmt.Print("Titulo: ")
			fmt.Scanln(&titulo2)
			fmt.Print("Formato: ")
			fmt.Scanln(&formato2)
			fmt.Print("Frames: ")
			fmt.Scanln(&frame)
			Con.Multi = append(Con.Multi, &composicion.Video{titulo2, formato2, frame})
		case 4:
			fmt.Println(Con.Mostrar())
		case 5:
			ban=false
		default:
			fmt.Println("Opcion incorrecta")
		}
	}
}
