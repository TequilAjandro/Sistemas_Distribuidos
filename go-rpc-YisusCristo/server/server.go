package main

import (
	"fmt"
	"net"
	"net/rpc"
)

//Subject -> struct para la materia
type Subject struct {
	grade float64
}

//Student -> struct estudiante
type Student struct {
	name         string
	subject      []Subject
	averageGrade float64
}

//Average calcula el promedio
func (s *Student) Average() {
	var average float64 = 0.0
	for _, g := range s.subject {
		average += g.grade
	}
	s.averageGrade = average / float64(len(s.subject))
	fmt.Println("Como me gustan los culos")
}

//Server struct
type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

//Anos suculentos
func (this *Server) Anos(i int64, reply *string) error {
	*reply = "ANOOOOOOOOOOOOS"
	return nil
}

//Server
func server() {
	var jaja float64 = 55.56
	s := Student{}
	s.subject = append(s.subject, Subject{jaja})

	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	// var jaja float64 = 55.56
	// var jaja2 float64 = 666.10
	// var jaja3 float64 = 1.20

	// s := Student{}
	// sub1 := Subject{jaja}
	// sub2 := Subject{jaja2}
	// sub3 := Subject{jaja3}

	// s.subject = append(s.subject, sub1)
	// s.subject = append(s.subject, sub2)
	// s.subject = append(s.subject, sub3)

	// s.Average()
	// fmt.Printf("averageGrade %f", s.averageGrade)
	fmt.Println("Corriendose en tu cara")
	server()
}
