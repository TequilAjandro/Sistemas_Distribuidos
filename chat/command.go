package main

type commandID int

// Definimos los comandos para los clientes
const (
	CMD_NICK commandID = iota
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     commandID
	client *client
	args   []string
}
