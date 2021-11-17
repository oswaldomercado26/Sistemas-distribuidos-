package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

type Proceso struct {
	ID               int
	Valor            uint64
	ContinuarProceso bool
}

func (p *Proceso) IncrementarValor(){
	for p.ContinuarProceso {
		p.MostrarValor()
		p.Valor++
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Proceso) MostrarValor() {
	fmt.Printf(" %d:  %d\n", p.ID, p.Valor)
}

var proceso Proceso

func cliente(c net.Conn) {
	mensaje := "Nuevo"
	fmt.Println("", mensaje)
	err := gob.NewEncoder(c).Encode(mensaje)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewDecoder(c).Decode(&proceso)
	if err != nil {
		fmt.Println(err)
		return
	}
	proceso.ContinuarProceso = true
	go proceso.IncrementarValor()
}

func main() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	go cliente(c)

	var input string
	fmt.Scanln(&input)
	proceso.ContinuarProceso = false
	err = gob.NewEncoder(c).Encode(&proceso)
	if err != nil {
		fmt.Println(err)
		return
	}
}
