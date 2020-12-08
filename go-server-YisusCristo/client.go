package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	process "./packages"
)

var flag = false
var p process.Process

func client() {
	c, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(flag)
	if err != nil {
		fmt.Println(err)
	}

	if !flag {
		err = gob.NewDecoder(c).Decode(&p)
		if err != nil {
			fmt.Println(err)
			return
		}
		flag = true
	}

	c.Close()

}

func main() {
	go client()

	go func() {
		for {
			if p.Value != 0 && p.Id != 0 {
				fmt.Println("-----------------------")
				fmt.Println(p.Id, " : ", p.Value)
				p.Value++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)

	c, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(flag)
	if err != nil {
		fmt.Println(err)
	}

	err = gob.NewEncoder(c).Encode(p)
	if err != nil {
		fmt.Println(err)
	}
	flag = false

	fmt.Println("Ended Client\n")
}
