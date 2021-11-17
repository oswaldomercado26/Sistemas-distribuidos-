package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

)

type Alumno struct {
	Materia      string
	Calificacion float32
}

type Registro struct {
	Alumno       string
	Materia      string
	Calificacion float32
}

func Add(registro Registro) ([]byte) {
	jsonData := []byte(`{"code": "ok"}`) // Creamos la respuesta

	materia := strings.ToLower(registro.Materia)
	alumno := strings.ToLower(registro.Alumno)
	calificacion := registro.Calificacion

	_, ok := alumnos[alumno]
	if !ok {
		alumnos[alumno] = make(map[string]float32)
	}
	if _, err := materias[materia]; err {

		materias[materia][alumno] = calificacion
	} else {
		materias[materia] = make(map[string]float32)
		materias[materia][alumno] = calificacion
	}
	alumnos[alumno][materia] = calificacion

	fmt.Println(alumnos)
	
	return jsonData // retornamos la respuesta
}

func Put(nombre string, alumno Alumno) []byte {
	materia := strings.ToLower(alumno.Materia)
	_, ok := alumnos[nombre]
	if ok == false {
		return []byte(`{"code": "error", "message": "El nombre del alumno no existe"}`)
	}
	_, ok = alumnos[nombre][materia]
	if ok == false {
		return []byte(`{"code": "error", "message": "La materia no esta registrada para este alumno"}`)
	}
	alumnos[nombre][materia] = alumno.Calificacion
	materias[materia][nombre] = alumno.Calificacion
	return []byte(`{"code": "ok"}`)
}

func Delete(id string) ([]byte) {
	_, ok := alumnos[id] // buscamos la materia por nombre
	if ok == false { // si no se encuentra la tarea
		return []byte(`{"code": "noexiste"}`) // retornamos la respuesta noexiste
	}
	delete(alumnos, id) // eliminamos de admin la tarea con el id

	for _, alumns := range materias {
		_, ok := alumns[id]
		if ok {
			delete(alumns, id)
		}
	}
	return []byte(`{"code": "ok"}`)
}

func Get() ([]byte, error) {
	jsonData, err := json.MarshalIndent(alumnos, "", "    ") 
	if err != nil {
		return jsonData, nil
	}
	return jsonData, err
}	

func GetID(id string) ([]byte, error) {
	jsonData := []byte(`{}`) 
	nom, ok := alumnos[id]
	if ok == false { // si no existe la materia
		return jsonData, nil
	}	
	// si existe la tarea la convertimos a JSON
	jsonData, err := json.MarshalIndent(nom, "", "    ")
	if err != nil {
		return jsonData, err // retornamos el JSON
	}
	return jsonData, nil // error al convertir a JSON
}

func alumno_id(res http.ResponseWriter, req *http.Request) {
	
	nombre := strings.ToLower(strings.TrimPrefix(req.URL.Path, "/alumno/"))
	res_json := []byte(`{}`)
	var err error

	fmt.Println(req.Method) // imprimimos el método que se mando llamar
	switch req.Method {
	case "GET": 
		res_json, err = GetID(nombre) 
		if err != nil { 
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}

	case "DELETE": // si el método es DELETE
		res_json = Delete(nombre)	

	case "PUT": 
		var alumno Alumno
		err := json.NewDecoder(req.Body).Decode(&alumno)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res_json = Put(nombre, alumno)
	}
	res.Header().Set(
		"Content-Type",
		"application/json",
	)
	res.Write(res_json)
}
	

func alumno(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method) // imprimimos el método usando del Request 
	switch req.Method {
	case "POST": // si el método es POST
		var alumno Registro// creamos una estacia de Tarea
		// convertidos el JSON enviado a una Tarea
		err := json.NewDecoder(req.Body).Decode(&alumno) 
		if err != nil { // si no fue posible convertir el JSON a una tarea
			 // retornamos un error al cliente
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		// si se convirtio el JSON a Tarea
		res_json := Add(alumno) // agregamos la tarea al administrador de tareas
		// armamos el *header de* la respuesta al cliente que será un JSON
		res.Header().Set(
			"Content-Type",
			"application/json",
		)		
		res.Write(res_json) 

	case "GET":
		res_json, err := Get() 
		if err != nil { 
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return		
		}
		res.Header().Set(
			"Content-Type",
			"application/json",
		)
		res.Write(res_json) 
	}
}


var alumnos = make(map[string]map[string]float32)
var materias = make(map[string]map[string]float32)

func main() {
	
	http.HandleFunc("/alumno", alumno) // creamos el endpoint /tarea
	http.HandleFunc("/alumno/", alumno_id) // creamos el endpoint /tarea/
	fmt.Println("Corriendo RESTful API") // mensaje para empezar a probar
	http.ListenAndServe(":9000", nil) // arrancamos el servidor en el puerto :9000
}