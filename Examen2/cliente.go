package main

import (
	"bufio"
	"container/list"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"net/rpc"
	"os"
)

const CREAR_USUARIO = -1
const TERMINAR_SERVIDOR = -2

const MENSAJE = 1
const MOSTRAR_CHAT = 3
const SALIR = 0

const (
	SALA_NIÑOS    = 1
	SALA_ADULTOS  = 2
	SALA_VIEJITOS = 3
)

type Client struct {
	Operacion int8
	ID        string
	Mensaje   string
}
type InformServer struct { //Estructura para enviar respuesta al cliente
	Cadena string //Texto de los datos del servidor
	Conex  int    //Numero de conexiones en el servidor
	Puerto string //Puerto tpc para el servidor
	Tema   string //Tema del servidor
}

func conectarUsuario(c net.Conn) bool {
	var client Client

	conexionActiva = true
	fmt.Print("Id: ")
	fmt.Scanln(&id)
	fmt.Println("Conectando...")
	client.ID = id
	client.Operacion = CREAR_USUARIO
	err := gob.NewEncoder(c).Encode(client)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Conectado :)")
	return true
}

func clienteM(c net.Conn) {
	var err error

	for conexionActiva {

		var client Client
		err = gob.NewDecoder(c).Decode(&client)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch client.Operacion {
		case CREAR_USUARIO:
			fmt.Println(client.Mensaje)
			chat.PushBack(client.Mensaje)
		case MENSAJE:
			recibirMensaje(client)
		case SALIR:
			fmt.Println(client.Mensaje)
			chat.PushBack(client.Mensaje)
		case TERMINAR_SERVIDOR:
			fmt.Println("El servidor se ha desconectado!")
			conexionActiva = false
		}
	}
}

func recibirMensaje(client Client) {
	mensaje := client.ID + ": " + client.Mensaje
	fmt.Println(mensaje)
	chat.PushBack(mensaje)
}

func mostrarChat() {

	fmt.Println("Chat")
	for msj := chat.Front(); msj != nil; msj = msj.Next() {
		fmt.Println(msj.Value)
	}
}

var id string
var conexionActiva bool
var chat list.List
var scanner *bufio.Scanner

func obtenerOpcionMenu() (opcMenu int) {
	fmt.Println("1) Enviar Mensaje")
	fmt.Println("3) Mostrar Chat")
	fmt.Println("0) Salir")
	fmt.Print("> ")
	fmt.Scanln(&opcMenu)
	return opcMenu
}

func enviarMensaje(c net.Conn) {
	var client Client

	client.ID = id
	client.Operacion = MENSAJE
	fmt.Println("Escribe un mensaje")
	fmt.Print("> ")
	client.Mensaje = readLine()
	err := gob.NewEncoder(c).Encode(client)
	if err != nil {
		fmt.Println(err)
		return
	}
	chat.PushBack("Tú: " + client.Mensaje)
}

func cerrarConexion(c net.Conn) {
	var peticion Client

	peticion.ID = id
	peticion.Operacion = SALIR
	err := gob.NewEncoder(c).Encode(peticion)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Close()
	conexionActiva = false
}

func imprimirChat(c net.Conn) {
	for conexionActiva {
		switch obtenerOpcionMenu() {
		case MENSAJE:
			enviarMensaje(c)
		case MOSTRAR_CHAT:
			mostrarChat()
		case SALIR:
			conexionActiva = false
			cerrarConexion(c)
		}
		if conexionActiva {
			pausar()
		}
	}
}

func pausar() {
	fmt.Print("\nPresiona ENTER para continuar...")
	fmt.Scanln()
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Solicitando las opciones de salas...")
	var servA [3]InformServer

	err = c.Call("Server.ObtenerSalas", "cliente", &servA)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SERVIDORES")
		for _, s := range servA {
			fmt.Println(s.Cadena)
		}

		var opcion int64
		fmt.Scanln(&opcion)

		switch opcion {

		case SALA_NIÑOS:
			c, err := net.Dial("tcp", ":0011")
			if err != nil {
				fmt.Println(err)
				return
			}
			puerto := servA[0].Puerto
			tema := servA[0].Tema
			conectador := servA[0].Conex
			fmt.Println("Entraste a " ,tema ,"con el puerto", puerto, "Activos: ", conectador)
			conectarUsuario(c)
			pausar()
			go clienteM(c)
			imprimirChat(c)

		case SALA_ADULTOS:
			c, err := net.Dial("tcp", ":0022")
			if err != nil {
				fmt.Println(err)
				return
			}
			puerto := servA[1].Puerto
			tema := servA[1].Tema
			conectador := servA[1].Conex
			fmt.Println("Entraste a ", tema , "con el puerto" , puerto , "Activos: " , conectador)
			conectarUsuario(c)
			pausar()
			go clienteM(c)
			imprimirChat(c)

		case SALA_VIEJITOS:
			c, err := net.Dial("tcp", ":0033")
			if err != nil {
				fmt.Println(err)
				return
			}
			puerto := servA[2].Puerto
			tema := servA[2].Tema
			conectador := servA[2].Conex
			fmt.Println("Entraste a ", tema , "con el puerto" ,puerto ,"Activos: ", conectador)
			conectarUsuario(c)
			pausar()
			go clienteM(c)
			imprimirChat(c)
		}

	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	client()
}

func readLine() string {
	var stdin *bufio.Reader
	var line []rune
	var c rune
	var err error

	stdin = bufio.NewReader(os.Stdin)
	for {
		c, _, err = stdin.ReadRune()
		if err == io.EOF || c == '\n' {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading standard input\n")
			os.Exit(1)
		}
		line = append(line, c)
	}
	return string(line[:len(line)])
}
