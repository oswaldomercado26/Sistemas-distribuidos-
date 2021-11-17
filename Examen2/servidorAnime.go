package main

import (
	"container/list"
	"encoding/gob"
	"fmt"
	"net"
	"net/rpc"
)



type Client struct {
	Operacion int8
	ID  string
	Mensaje   string
}
type Servidor struct {
	Tema       string
	Protocolo  string
	Puerto     string
	Conexiones int
}

func pausar() {
	fmt.Print("\nPresiona ENTER para continuar...")
	fmt.Scanln()
}

type Server struct{}

func server(){
	rpc.Register(new(Server))
	s,err := net.Listen("tcp",":0111")
	if err != nil{
		fmt.Println(err)
		return
	}
	for{
		c,err := s.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}

}

func conec(){
	s,err := net.Listen("tcp",":0011")
	if err != nil{
		fmt.Println(err)
		return
	}
	for{
		c,err := s.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		go handleCliente(c)
	}

}
var conectadosN int

func (this *Server) ServidorAnime(servidor string, reply *Servidor) error { //Función para obtener los servidores disponibles
	var info Servidor 
	info.Tema = "Servidor ANIME"
	info.Protocolo = "tcp"
	info.Puerto = "0011"
	info.Conexiones = conectadosN

	*reply = info
	return nil
}

func reciveMensaje(client Client, c net.Conn){
	chat.PushBack(client.ID+ ": " + client.Mensaje)
	reenviarPeticion(client, c)
}

func handleCliente(c net.Conn){
	var err error
	continuar:= true;
	
	for continuar{
		var client Client
		err = gob.NewDecoder(c).Decode(&client)
		if err != nil{
			fmt.Println(err)
			return
		}
		switch client.Operacion{

		case -1:
			clientes[client.ID]=c
			client.Mensaje = "Se conectó: "+ client.ID
			chat.PushBack(client.Mensaje)
			reenviarPeticion(client, c)
			conectadosN++

		case 1:
			go reciveMensaje(client,c)

		case 0:
			continuar=false
			delete(clientes, client.ID)
			client.Mensaje = "Se desconectó: " + client.ID
			chat.PushBack(client.Mensaje)
			enviarPeticionGeneral(client)

		}
		
	}

}

func reenviarPeticion(client Client, c net.Conn) {
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

func mostrarChat() {
	fmt.Println("Chat")
	for msj := chat.Front(); msj != nil; msj = msj.Next() {
		fmt.Println(msj.Value)
	}
}
func terminarServidor() {
	var client Client
	client.Operacion = -2
	enviarPeticionGeneral(client)
}

func enviarPeticionGeneral(peticion Client) {
	var err error
	for _, conexion := range clientes {
		err = gob.NewEncoder(conexion).Encode(peticion)
		if err != nil {
			fmt.Println(err)
		}
	}
}



var clientes = make(map[string]net.Conn)
var chat list.List

func main() {
	go server()
	go conec()
	
	var opcMenu int
	conexionActiva := true
	for conexionActiva {
		fmt.Println("1) Mostrar Mensajes")
		fmt.Println("0) Terminar Servidor")
		fmt.Print("> ")
		fmt.Scanln(&opcMenu)
		switch opcMenu {
		case 1:
			mostrarChat()
		case 0:
			conexionActiva = false
			terminarServidor()
		}
		if conexionActiva {
			pausar()
		}
	}
}
