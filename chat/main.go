package main

import (
	"fmt"
	"log"
	"net"
)

// Creamos nuestro server
var s = newServer()

func start() {
	// Corremos el server
	go s.run()

	// Definimos un puerto
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("No se pudo inciar el servidor: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Servidor corriendo en el puerto :8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("No se pudo aceptar la conexión: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}

func main() {
	// Corremos todo
	go start()
	// Comando para salir
	var input string
	fmt.Scanln(&input)
	log.Printf("Servidor muerto")
	// Termina el servidor y guarda la bitacora en un archivo de texto
	s.stop()
}
