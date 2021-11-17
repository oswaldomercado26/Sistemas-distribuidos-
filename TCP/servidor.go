package main

import (
	"container/list"
	"encoding/gob"
	"fmt"
	"math"
	"net"
	"time"
)

const MAX_PROCESOS = 5

type Proceso struct {
	ID               int
	Valor            uint64
	ContinuarProceso bool
}

func (p *Proceso) IncrementarValor(){
	for p.Valor <= math.MaxUint64 && p.ContinuarProceso {
		fmt.Printf(" %d : %d\n", p.ID, p.Valor)
		p.Valor++
		time.Sleep(time.Millisecond * 500)
	}
}

var procesos list.List

func agregarProceso(proceso *Proceso) {
	proceso.ContinuarProceso = true
	go proceso.IncrementarValor()
	procesos.PushBack(proceso)

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

func handleCliente(c net.Conn) {
	var mensaje string
	err := gob.NewDecoder(c).Decode(&mensaje)
	if err != nil {
		fmt.Println(err)
		return
	}
	if mensaje == "Nuevo" {
		proceso := procesos.Remove(procesos.Front()).(*Proceso)
		proceso.ContinuarProceso = false
		err = gob.NewEncoder(c).Encode(*proceso)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = gob.NewDecoder(c).Decode(proceso)
		if err != nil {
			fmt.Println(err)
			return
		}
		agregarProceso(proceso)
		c.Close()
	}
}

func main() {
	for i := 0; i < MAX_PROCESOS; i++ {
		proceso := new(Proceso)
		proceso.ID = procesos.Len()
		agregarProceso(proceso)
	}
	go server()

	var input string
	fmt.Scanln(&input)
}
