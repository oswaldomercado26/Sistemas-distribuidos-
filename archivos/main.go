package main

import (
	
	"fmt"
	"os"
	"sort"
	"strings"
)

type Proceso struct {
	id uint64
	Prioridad int64
	Tiempo uint64
	Estatus string
}

type process []Proceso

func (a process) Len() int           { return len(a) }
func (a process) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a process) Less(i, j int) bool { return uint64(a[i].Prioridad) < uint64(a[j].Prioridad) }

func main() {

    var ban bool = true
	var eleccion int
	for ban {
	menu :=
		`
    Actividad 7
[ 1 ] Ordenar ascendente
[ 2 ] Ordenar descendente
[ 3 ] Proceso ascendente
[ 4 ] Proceso descendente
[ 5 ] Salir
    ¿Qué opcion quieres?
`
	fmt.Print(menu)
	
	fmt.Scanln(&eleccion)
	
		switch eleccion {
		case 1:
			var entrada string
			var c int=0
			file, err:= os.Create("ascendente.txt")
			if err != nil{
				fmt.Println(err)
				return
			}
			defer file.Close()	
			fmt.Print("Ingrese el numero de valores del arreglo: ")
			fmt.Scanln(&c)
			s:=make([]string,0,c)
			
			for i := 0; i < c; i++ {
				fmt.Printf("Ingrese la cadena: ")
				fmt.Scanln( &entrada)
				s = append(s, entrada)
		
			}
			//menor a mayor
			sort.Sort(sort.StringSlice(s))
			file.WriteString(strings.Join(s[:],"\n"))
			fmt.Println(s)
		                 
		case 2:
			
			var entrada string
			var c int=0
			file, err:= os.Create("descendente.txt")
			if err != nil{
				fmt.Println(err)
				return
			}
			defer file.Close()	
			fmt.Print("Ingrese el numero de valores del arreglo: ")
			fmt.Scanln(&c)
			s:=make([]string,0,c)
			
			for i := 0; i < c; i++ {
				fmt.Printf("Ingrese la cadena: ")
				fmt.Scanln( &entrada)
				s = append(s, entrada)
		
			}
			//mayor a menor
			sort.Sort(sort.Reverse(sort.StringSlice(s)))
		
			file.WriteString(strings.Join(s[:],"\n"))
			fmt.Println(s)
		case 3:
			ps:=[]Proceso{
				Proceso{id: 1 ,Prioridad: 1 ,Tiempo: 4 ,Estatus: "Entregado" },
				Proceso{id: 5 ,Prioridad: 4 ,Tiempo: 10 ,Estatus: "Saliendo" },
				Proceso{id: 10 ,Prioridad: 2 ,Tiempo: 5 ,Estatus: "Observacion" },
				Proceso{id: 2 ,Prioridad: 3 ,Tiempo: 9 ,Estatus: "Error" },
				Proceso{id: 8 ,Prioridad: 5 ,Tiempo: 1 ,Estatus: "Saliendo" },
			}
			fmt.Println("Ascendente")
			fmt.Println(ps)
			sort.Sort(process(ps))
			fmt.Println(ps)
		case 4:
			ps:=[]Proceso{
				Proceso{id: 1 ,Prioridad: 1 ,Tiempo: 4 ,Estatus: "Entregado" },
				Proceso{id: 5 ,Prioridad: 4 ,Tiempo: 10 ,Estatus: "Saliendo" },
				Proceso{id: 10 ,Prioridad: 2 ,Tiempo: 5 ,Estatus: "Observacion" },
				Proceso{id: 2 ,Prioridad: 3 ,Tiempo: 9 ,Estatus: "Error" },
				Proceso{id: 8 ,Prioridad: 5 ,Tiempo: 1 ,Estatus: "Saliendo" },
			}
			fmt.Println("Descendente")
			fmt.Println(ps)
			sort.Sort(sort.Reverse(process(ps)))
			fmt.Println(ps)
		case 5:
			ban=false
		default:
			fmt.Println("Opcion incorrecta")
		}
	}
}