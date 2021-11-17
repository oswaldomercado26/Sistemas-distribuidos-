package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var alumnos = make(map[string]map[string]float32)
var materias = make(map[string]map[string]float32)

func agregarCalificacion(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		var calificacion float32
		materia := strings.ToLower(req.FormValue("materia"))
		alumno := strings.ToLower(req.FormValue("alumno"))
		value, _ := strconv.ParseFloat(req.FormValue("calificacion"), 32)
		calificacion = float32(value)

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
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHTML("respuesta.html"),
			"Se agregó la calificación correctamente!",
		)
		fmt.Println(alumnos)
	}
}

func cargarHTML(a string) string {
	html, _ := ioutil.ReadFile(a)
	return string(html)
}

func formRegistro(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHTML("registro.html"),
	)
}

func formAlumno(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHTML("alumno.html"),
	)
}


func formMateria(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHTML("materia.html"),
	)
}

func promGeneral(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	totalCalificaciones := 0.0
	numCalificaciones := 0
	for _, mats := range alumnos {
		for _, calificacion := range mats {     
			totalCalificaciones += float64(calificacion)
		}
		numCalificaciones += len(mats)
	}
	promedio := float32(totalCalificaciones / float64(numCalificaciones))
	fmt.Fprintf(
		res,
		cargarHTML("respuesta.html"),
		fmt.Sprintf("El promedio general es: %.2f", promedio),

	)
}


func promedioMateria(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() error %v", err)
		return
	}
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	materia := strings.ToLower(req.FormValue("materia"))
	alumns, ok := materias[materia]
	if ok {
		totalCalificaciones := 0.0
		for _, calificacion := range alumns {
			totalCalificaciones += float64(calificacion)
		}
		promedio := float32(totalCalificaciones / float64(len(alumns)))
		fmt.Fprintf(
			res,
			cargarHTML("respuesta.html"),
			fmt.Sprintf("El promedio de la materia %s es: %.2f", materia, promedio),
		)
	} else {
		fmt.Fprintf(
			res,
			cargarHTML("respuesta.html"),
			"La materia "+materia+" no existe!",
		)
	}

}

func promedioAlumno(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() error %v", err)
		return
	}
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	nombre := strings.ToLower(req.FormValue("alumno"))
	mats, ok := alumnos[nombre]
	if ok {
		totalCalificaciones := 0.0
		for _, calificacion := range mats {
			totalCalificaciones += float64(calificacion)
		}
		promedio := float32(totalCalificaciones / float64(len(mats)))
		fmt.Fprintf(
			res,
			cargarHTML("respuesta.html"),
			fmt.Sprintf("El promedio del alumno %s es: %.2f", nombre, promedio),
		)
		} else {
		fmt.Fprintf(
			res,
			cargarHTML("respuesta.html"),
			"El alumno"+ nombre +" no existe!",
		)
		
	}

}

func main() {
	http.HandleFunc("/formregistro", formRegistro)
	http.HandleFunc("/registro", agregarCalificacion)
	http.HandleFunc("/formalum",formAlumno)
	http.HandleFunc("/alumno",promedioAlumno)
	http.HandleFunc("/general",promGeneral)
	http.HandleFunc("/materia",formMateria)
	http.HandleFunc("/prommateria",promedioMateria)

	fmt.Println("Arrancando el servidor...")
	http.ListenAndServe(":9000", nil)


}
