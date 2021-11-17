package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)


type Servidor struct { //Estructura de datos del servidor
	Tema       string
	Protocolo  string
	Puerto     string
	Conexiones int
}

type InfoServer struct { //Estructura para enviar respuesta al cliente
	Cadena string //Texto de los datos del servidor
	Conex  int    //Numero de conexiones en el servidor
	Puerto string //Puerto tpc para el servidor
	Tema   string //Tema del servidor
}

type Server struct{}



func clientMiddle() { //Recibir datos de los servidores
	var servidorAnime Servidor
	var servidorSeries Servidor
	var servidorPeliculas Servidor
    
	fmt.Println("Informacion servidores")
	
	c, err := rpc.Dial("tcp", "127.0.0.1:0111")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Call("Server.ServidorAnime", "cliente", &servidorAnime) 
	if err != nil {
		fmt.Println(err)
	}
	servidorMap[servidorAnime.Tema] = servidorAnime 

	c, err = rpc.Dial("tcp", "127.0.0.1:0222")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Call("ServerJ.ServidorPeliculas", "cliente", &servidorPeliculas) 
	if err != nil {
		fmt.Println(err)
	}

	servidorMap[servidorPeliculas.Tema] = servidorPeliculas

	c, err = rpc.Dial("tcp", "127.0.0.1:0333")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Call("Server.ServidorSeries", "cliente", &servidorSeries)
	if err != nil {
		fmt.Println(err)
	}
	servidorMap[servidorSeries.Tema] = servidorSeries

}

func (this *Server) ObtenerSalas(servidor string, reply *[3]InfoServer) error { //Función para obtener los servidores disponibles
	var servidorA [3]InfoServer //Cadena final a imprimir

	cont := 0 
	fmt.Println("Comunicando con el cliente\n")
	for _, c := range servidorMap { 
		servidorA[cont].Cadena = strconv.Itoa(cont+1) + ") " + c.Tema + " conectados: " + strconv.Itoa(c.Conexiones) 
		servidorA[cont].Conex = c.Conexiones                                                                        
		servidorA[cont].Puerto = c.Puerto
		servidorA[cont].Tema = c.Tema
		cont++
	}
	*reply = servidorA
	return nil
}

func server() {

	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		clientMiddle()
		go rpc.ServeConn(c)
	}
}

var servidorMap = make(map[string]Servidor)
func main() {

	go server()
	var input string
	fmt.Scanln(&input)

}
