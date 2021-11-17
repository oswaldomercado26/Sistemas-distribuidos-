package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

type Registro struct {
	Alumno       string
	Materia      string
	Calificacion float32
}


func agregar(c *rpc.Client) {
	var registro Registro
	var respuesta string

	fmt.Println("Agregar Calificación")
	fmt.Print("Alumno: ")
	scanner.Scan()
	registro.Alumno = scanner.Text()
	fmt.Print("Materia: ")
	scanner.Scan()
	registro.Materia = scanner.Text()
	fmt.Print("Calificación: ")
	fmt.Scanln(&registro.Calificacion)
	err := c.Call("Server.AgregarCalificacion", registro, &respuesta)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Se agregó la calificación correctamente!")
	}
}

func promedioAlumno(c *rpc.Client) {
	var promedio float32
	var alumno string

	fmt.Println("Obtener Promedio de Alumno")
	fmt.Print("Alumno: ")
	scanner.Scan()
	alumno = scanner.Text()
	err := c.Call("Server.ObtenerPromedioAlumno", alumno, &promedio)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Promedio de", alumno, "=", promedio)
	}
}

func promedioGeneral(c *rpc.Client){

	var promedio float32
	var args string

	fmt.Println("Obtener Promedio General")
	err := c.Call("Server.ObtenerPromedioGeneral", args, &promedio)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Promedio General:", promedio)
	}
}


func promedioMateria(c *rpc.Client) {
	var promedio float32
	var materia string

	fmt.Println("Promedio de Materia")
	fmt.Print("Materia: ")
	scanner.Scan()
	materia = scanner.Text()
	err := c.Call("Server.ObtenerPromedioMateria", materia, &promedio)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Promedio de", materia, "=", promedio)
	}
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var eleccion int64
	for {

	menu :=	`
Bienvenido a RCP
[ 1 ] Agregar calificacion
[ 2 ] Promedio Alumno
[ 3 ] Promedio General
[ 4 ] Promedio Materia
[ 5 ] Terminar 
¿Qué opcion quieres?
`
	fmt.Print(menu)
	fmt.Scanln(&eleccion)

		switch eleccion {
		case 1:
			agregar(c)

		case 2:
			promedioAlumno(c)

		case 3:
			promedioGeneral(c)

		case 4:
			promedioMateria(c)

		case 5:
			return
		}

	}
}

var scanner *bufio.Scanner //liberar los espacios 

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	client()
}
