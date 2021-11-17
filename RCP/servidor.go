package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"strings"
)


type Registro struct {
	Alumno       string
	Materia      string
	Calificacion float32
}


type Server struct{}

func (this *Server) AgregarCalificacion(registro Registro, respuesta *string) error{
	materia := strings.ToLower(registro.Materia)
	alumno := strings.ToLower(registro.Alumno)
	calificacion := registro.Calificacion
	_, ok := asignaturas[alumno]
	if !ok {
		asignaturas[alumno] = make(map[string]float32)
	}
	if _,err := materias[materia]; err{
		
		materias[materia][alumno] = calificacion
	}else{
		materias[materia] = make(map[string]float32)
		materias[materia][alumno] = calificacion
	}
	asignaturas[alumno][materia]= calificacion	
	
	fmt.Println(asignaturas)
	return nil
}

func (s *Server) ObtenerPromedioGeneral(args string, promedio *float32) error {
	totalCalificaciones := 0.0
	numCalificaciones := 0
	for _, mats := range asignaturas {
		for _, calificacion := range mats {     
			totalCalificaciones += float64(calificacion)
		}
		numCalificaciones += len(mats)
	}
	*promedio = float32(totalCalificaciones / float64(numCalificaciones))
	return nil
}

func (s *Server) ObtenerPromedioMateria(materia string, promedio *float32) error {
	materia = strings.ToLower(materia)
	alumns, ok := materias[materia]
	if ok {
		totalCalificaciones := 0.0
		for _, calificacion := range alumns {
			totalCalificaciones += float64(calificacion)
		}
		*promedio = float32(totalCalificaciones / float64(len(alumns)))
	} else {
		return errors.New("La materia " + materia + " no esta registrada!")
	}
	return nil
}


func (this *Server) ObtenerPromedioAlumno(nombre string, promedio *float32) error{
	nombre = strings.ToLower(nombre)
	mats, ok := asignaturas[nombre]
	if ok {
		totalCalificaciones := 0.0
		for _, calificacion := range mats {
			totalCalificaciones += float64(calificacion)
		}
		*promedio = float32(totalCalificaciones / float64(len(mats)))
	} else {
		return errors.New("El alumno " + nombre + " no tiene registro!")
	}
	return nil
}




func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	fmt.Println("servidor conectado")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

var asignaturas = make(map[string]map[string]float32)
var materias = make(map[string]map[string]float32)
func main(){
	go server()

	var input string
	fmt.Scanln(&input)

}