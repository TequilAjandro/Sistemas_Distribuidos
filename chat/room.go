package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Definimos la sala a la que se uniran todos los clientes
type room struct {
	// Nombre de la sala
	Name string
	// Todos los miembros en la sala
	Members []*client
	// Todos los mensajes que se han enviado en la sala
	Messages []string
}

// Función para enviar un mensaje a todos los demas clientes
func (r *room) broadcast(sender *client, msg string) {
	for _, m := range r.Members {
		if sender.nick != m.nick {
			m.msg(msg)
		}
	}
	// Añade el mensaje enviado por un usuario a la bitacora
	r.Messages = append(r.Messages, (sender.conn.RemoteAddr().String() + " " + strings.TrimSpace(sender.nick) + ": " + msg))
}

// Guarda todos los mensajes de la sala en un archivo de texto
func (r *room) saveMessages() {
	fmt.Println("Dentro de la funcion jajajaj")
	file, err := os.OpenFile("registro.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(file)
	fmt.Println("Justo antes, len ", len(r.Messages))
	for _, data := range r.Messages {
		_, _ = datawriter.WriteString(data + "\n")
	}
	datawriter.Flush()
	file.Close()
	for _, data := range r.Messages {
		fmt.Println(data)
	}
}
