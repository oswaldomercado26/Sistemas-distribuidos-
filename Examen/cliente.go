package main

import (
	"bufio"
	"container/list"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)





type Peticion struct {
	Operacion int8
	ID  string
	Mensaje   string
	Archivo   File
}

type File struct {
	Nombre string
	Tamano int64
	Datos []byte
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
			fmt.Fprintf(os.Stderr, "Error \n")
			os.Exit(1)
		}
		line = append(line, c)
	}
	return string(line[:len(line)])
}
func conexionUsuario(c net.Conn) bool {
	var peticion Peticion

	conexionActiva = true
	fmt.Print("Nombre de Usuario: ")
	fmt.Scanln(&id)
	fmt.Println("Conectando...")
	peticion.ID = id
	peticion.Operacion = -1
	err := gob.NewEncoder(c).Encode(peticion)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func enviarArchivo(c net.Conn) {
	var peticion Peticion
	var ruta string

	peticion.ID = id
	peticion.Operacion = 2
	fmt.Println("Escribe la ruta de tu archivo")
	fmt.Scanln(&ruta)

	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No se pudo abrir el archivo:", err)
		return
	}
	defer archivo.Close()
	stats, err := archivo.Stat()
	if err != nil {
		fmt.Println("No se pudieron leer los stats:", err)
		return
	}
	_, file := filepath.Split(ruta)//cambia toda la ruta solo por el nombre del archivo
	fmt.Println(file)
	peticion.Archivo.Nombre = file
	peticion.Archivo.Tamano = stats.Size()
	peticion.Archivo.Datos = make([]byte, stats.Size())
	archivo.Read(peticion.Archivo.Datos)

	// Enviar archivo
	err = gob.NewEncoder(c).Encode(peticion)
	if err != nil {
		fmt.Println(err)
		return
	}
	chat.PushBack("Tu: " + peticion.Archivo.Nombre)
	readLine()
}


func cliente(c net.Conn) {
	var err error

	for conexionActiva {
		
		var peticion Peticion
		err = gob.NewDecoder(c).Decode(&peticion)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch peticion.Operacion {
		case -1:
			fmt.Println(peticion.Mensaje)
			chat.PushBack(peticion.Mensaje)
		case 1:
			recibir(peticion)
		case 2:
			recibirArchivo(peticion)
		case 0:
			fmt.Println(peticion.Mensaje)
			chat.PushBack(peticion.Mensaje)
		case -2:
			fmt.Println("El servidor se ha desconectado!")
			conexionActiva = false
		}
	}
}

func recibir(peticion Peticion) {
	mensaje := peticion.ID + ": " + peticion.Mensaje
	fmt.Println(mensaje)
	chat.PushBack(mensaje)
}

func recibirArchivo(peticion Peticion){
	archivo, err := os.Create(peticion.Archivo.Nombre)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer archivo.Close()
	archivo.Write(peticion.Archivo.Datos)
	archivo.Sync() 
	mensaje := peticion.ID + " [archivo]: " + peticion.Archivo.Nombre
	fmt.Println(mensaje)
	chat.PushBack(mensaje)
}

func mostrar() {
	fmt.Println("Chat")
	for msj := chat.Front(); msj != nil; msj = msj.Next() {
		fmt.Println(msj.Value)
	}
	readLine()
}
var id string
var conexionActiva bool
var chat list.List

func main(){
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	conexionUsuario(c) 
	fmt.Print("\nPresiona ENTER para continuar...")
	fmt.Scanln()
	go cliente(c)
	imprimirMenu(c)
	
}



func enviarMensaje(c net.Conn) {
	var peticion Peticion

	peticion.ID = id
	peticion.Operacion = 1
	fmt.Println("Escribe un mensaje")
	fmt.Print("> ")
	peticion.Mensaje = readLine()
	err := gob.NewEncoder(c).Encode(peticion)
	if err != nil {
		fmt.Println(err)
		return
	}
	chat.PushBack("Tú: " + peticion.Mensaje)
}



func cerrarConexion(c net.Conn) {
	var peticion Peticion

	peticion.ID = id
	peticion.Operacion = 0
	err := gob.NewEncoder(c).Encode(peticion)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Close()
	conexionActiva = false
}

func imprimirMenu(c net.Conn) {
	var eleccion int
	for conexionActiva {
		menu :=
		`
Bienvenido a Chatito
[ 1 ] Enviar mensaje
[ 2 ] Enviar archivo
[ 3 ] Mostrar Chat
[ 0 ] Terminar 
¿Qué opcion quieres?
`
	fmt.Print(menu)
	fmt.Scanln(&eleccion)
	readLine()
		switch eleccion {
		case 1:
			enviarMensaje(c)
		case 2:
			enviarArchivo(c)
		case 3:
			mostrar()
		case 0:
			conexionActiva = false
			cerrarConexion(c)
		}
		if conexionActiva {
			fmt.Scanln()
		}
	}
}


