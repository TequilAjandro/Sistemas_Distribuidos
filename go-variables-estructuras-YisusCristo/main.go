package main

import "fmt"

type Imagen struct {
	titulo, formato string
	canales         int
}

func (img Imagen) Mostrar() {
	fmt.Println("\nImagen",
		"\nTitulo: ", img.titulo,
		"\nFormato: ", img.formato,
		"\nCanales: ", img.canales)
	fmt.Println()
}

type Audio struct {
	titulo, formato string
	segundos        int
}

func (audio Audio) Mostrar() {
	fmt.Println("\nAudio",
		"\nTitulo: ", audio.titulo,
		"\nFormato: ", audio.formato,
		"\nSegundos: ", audio.segundos)
	fmt.Println()
}

type Video struct {
	titulo, formato string
	frames          int
}

func (video Video) Mostrar() {
	fmt.Println("\nVideo",
		"\nTitulo: ", video.titulo,
		"\nFormato: ", video.formato,
		"\nFrames: ", video.frames)
	fmt.Println()
}

type Multimedia interface {
	Mostrar()
}
type ContenidoWeb struct {
}

func Menu() int {
	var opcion int
	fmt.Println("1.- Crear una imagen",
		"\n2.- Crear un audio",
		"\n3.- Crear un video",
		"\n4.- Mostrar contenido",
		"\n\nElija una opcion")
	fmt.Scanln(&opcion)
	return opcion
}

func main() {
	var titulo string
	var formato string

	for {
		opcion := Menu()
		fmt.Println()
		switch opcion {
		case 1:
			var canales int
			fmt.Println("Ingrese el titulo, formato y numero de canales")
			fmt.Scanln(&titulo)
			fmt.Scanln(&formato)
			fmt.Scanln(&canales)
			img := Imagen{titulo, formato, canales}
			img.Mostrar()
		case 2:
			var segundos int
			fmt.Println("Ingrese el titulo, formato y numero de segundos")
			fmt.Scanln(&titulo)
			fmt.Scanln(&formato)
			fmt.Scanln(&segundos)
			audio := Imagen{titulo, formato, segundos}
			audio.Mostrar()
		case 3:
			var frames int
			fmt.Println("Ingrese el titulo, formato y numero de frames")
			fmt.Scanln(&titulo)
			fmt.Scanln(&formato)
			fmt.Scanln(&frames)
			video := Imagen{titulo, formato, frames}
			video.Mostrar()
		}
	}
}
