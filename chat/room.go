package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type room struct {
	Name     string
	Members  []*client
	Messages []string
}

func (r *room) broadcast(sender *client, msg string) {
	for _, m := range r.Members {
		if sender.nick != m.nick {
			m.msg(msg)
		}
	}
	r.Messages = append(r.Messages, (sender.conn.RemoteAddr().String() + " " + strings.TrimSpace(sender.nick) + ": " + msg))
}

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
