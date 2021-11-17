package composicion

type ContenidoWeb struct {
	Multi []Multimedia
}

func (Con *ContenidoWeb) Mostrar() string {
	datos := ""
	for _, v := range Con.Multi {
		datos += v.Mostrar()
	}
	return datos
}

type Multimedia interface {
	Mostrar() string
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion string
}

type Video struct {
	Titulo  string
	Formato string
	Frames  string
}

func (i *Imagen) Mostrar() string {
	return "Imagen: "+ "\n" +
		"Titulo: " + i.Titulo + "\n" +
		"Formato: " + i.Formato + "\n" +
		"Canales: " + i.Canales + "\n"
}

func (a *Audio) Mostrar() string {
	return "Audio: "+ "\n" + 
		"Titulo: " + a.Titulo + "\n" +
		"Formato: " + a.Formato + "\n" +
		"Canales: " + a.Duracion + "\n"
}
func (v *Video) Mostrar() string {
	return "Video: "+ "\n" +
		"Titulo: " + v.Titulo + "\n" +
		"Formato: " + v.Formato + "\n" +
		"Canales: " + v.Frames + "\n"
}
