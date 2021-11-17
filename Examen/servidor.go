package main

import (
	"container/list"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)


type Cliente struct {
	Operacion int8
	ID        string
	Mensaje   string
	Archivo   File
}

type File struct {
	Nombre string
	Tamaño int64
	Datos []byte
}


func server() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleCliente(c)
	}

}

func recibeMsg(cliente Cliente, c net.Conn) {
	lc.PushBack(cliente.ID + ": " + cliente.Mensaje)

	reenviarPeticion(cliente, c)
}

func handleCliente(c net.Conn) {
	var err error
	continuar := true
	for continuar {
		var cliente Cliente
		err = gob.NewDecoder(c).Decode(&cliente)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch cliente.Operacion {
		case -1:
			clientes[cliente.ID] = c
			cliente.Mensaje = "Se conectó: " + cliente.ID
			lc.PushBack(cliente.Mensaje)
			reenviarPeticion(cliente, c)

		case 1:
			go recibeMsg(cliente, c)

		case 2:
			go recibirArchivo(cliente, c)
		case 0:
			continuar = false
			delete(clientes, cliente.ID)
			cliente.Mensaje = "Se desconectó: " + cliente.ID
			lc.PushBack(cliente.Mensaje)
			enviarPeticionGeneral(cliente)

		}

	}

}
func recibirArchivo(cliente Cliente, c net.Conn) {
	lc.PushBack(cliente.ID + "[archivo]: " + cliente.Archivo.Nombre)
	reenviarPeticion(cliente, c)
}
func reenviarPeticion(client Cliente, c net.Conn) {
	var err error

	for _, conexion := range clientes {

		if conexion != c {
			err = gob.NewEncoder(conexion).Encode(client)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func mostrar() {
	fmt.Println("Chat")
	for msj := lc.Front(); msj != nil; msj = msj.Next() {
		fmt.Println(msj.Value)
	}
}
func terminar() {
	var cliente Cliente

	cliente.Operacion = -2
	enviarPeticionGeneral(cliente)
}

func enviarPeticionGeneral(cliente Cliente) {
	var err error
	for _, conexion := range clientes {
		err = gob.NewEncoder(conexion).Encode(cliente)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func respaldar() {
	nombreArchivo := "respaldo.txt"
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer archivo.Close()
	for msj := lc.Front(); msj != nil; msj = msj.Next() {
		archivo.WriteString(msj.Value.(string) + "\n")
	}
	archivo.Sync()
	fmt.Println("Respaldo exitoso!")
}

var clientes = make(map[string]net.Conn)
var lc list.List

func main() {
	go server()
	var eleccion int
	conexion := true
	for conexion {
		menu :=
		`
Bienvenido a Chatito
[ 1 ] Mostrar chat
[ 2 ] Respaldar chat
[ 0 ] Terminar 
¿Qué opcion quieres?
`
	fmt.Print(menu)
	fmt.Scanln(&eleccion)
		switch eleccion {
		case 1:
			mostrar()
		case 2:
			respaldar()
		case 0:
			conexion = false
			terminar()
		}
		if conexion {
			fmt.Scanln()
		}
	}
}
