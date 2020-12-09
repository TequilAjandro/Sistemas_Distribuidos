package main

import (
	"fmt"
	"log"
	"net"
)

var s = newServer()

func start() {
	// s := newServer()
	go s.run()

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
	go start()
	var input string
	fmt.Scanln(&input)
	log.Printf("Servidor muerto")
	s.stop()
}
