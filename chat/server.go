package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// Estructura del server
type server struct {
	rooms    room
	commands chan command
}

// Constructor para el servidor
func newServer() *server {
	return &server{
		rooms:    room{Name: "general", Members: nil, Messages: nil},
		commands: make(chan command),
	}
}

// Función para correr el servidor
func (s *server) run() {
	for cmd := range s.commands {
		// Según el comando recibido de un cliente es la función que ejecutara
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args[1])
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

// Detiene el server y llama a que se guarden los mensajes en un archivo de texto
func (s *server) stop() {
	s.rooms.saveMessages()
}

// Se añade un cliente al servidor
func (s *server) newClient(conn net.Conn) {
	log.Printf("\nUn nuevo cliente se ha conectatdo al servidor: %s\n\r", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "",
		commands: s.commands,
	}

	c.msg(fmt.Sprintf("Te haz conectado al servidor\r"))
	c.readInput()
}

// Definimos el nombre de nuestro cliente
func (s *server) nick(c *client, nick string) {
	c.nick = nick
	msg := "Bienvenido al chat general " + nick
	c.msg(string(msg))
	s.join(c, "general")
}

// Unimos al cliente a nuestra sala
func (s *server) join(c *client, roomName string) {
	r := &s.rooms
	r.Members = append(r.Members, c)
	c.room = r
	msg := "Se ha unido a la sala " + c.nick
	log.Printf(msg)
	r.broadcast(c, string(msg))
}

// Se envia el mensaje de un cliente al resto, además de mostrarse en el servidor
func (s *server) msg(c *client, args []string) {
	msg := strings.Join(args[1:len(args)], " ")
	log.Printf("Mensaje -> %s: %s", strings.TrimSpace(c.nick), msg)
	c.room.broadcast(c, string(strings.TrimSpace(c.nick)+": "+msg))
}

// Función para que un cliente se desconecte
func (s *server) quit(c *client) {
	log.Printf("%s abandono el chat", strings.TrimSpace(c.nick))

	c.msg("Hasta luego\r")
	c.conn.Close()
}
