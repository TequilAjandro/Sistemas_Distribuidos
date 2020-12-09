package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Clase cliente con atributos: conexión, nombre, sala, comandos
type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

// Se captura el comando y un mensaje del cliente y se envia  al servidor
func (c *client) readInput() {
	for {
		// Entrada de usuario
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		// Se parsea la entrada de usuario para obtener el comando y el texto
		msg = strings.Trim(msg, "\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		// Se envia el comando introducido por el cliente y el respectivo texto
		switch cmd {
		case "-nick":
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case "-msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case "-quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

// Función del cliente para recibir un error
func (c *client) err(err error) {
	c.conn.Write([]byte("err: " + err.Error() + "\n\r"))
}

// Función del cliente para recibir mensajes
func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n\r"))
}
