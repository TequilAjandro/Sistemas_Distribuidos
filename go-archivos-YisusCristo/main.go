package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func WriteFile(title string, s []string) {
	file, err := os.OpenFile(title, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Error al crear el archivo: %s", err)
	}
	datawriter := bufio.NewWriter(file)

	for _, data := range s {
		_, _ = datawriter.WriteString(data + "\n")
	}
	datawriter.Flush()
	file.Close()
}

func main() {
	var line string
	var n int
	fmt.Scanln(&n)
	// Creación del Slice
	s := make([]string, n)
	// Llenado del Slice
	for i := 0; i < n; i++ {
		fmt.Scanln(&line)
		s[i] = line
	}
	// Ascendente
	sort.Strings(s)
	WriteFile("ascendente.txt", s)
	fmt.Println(s)
	// Descendente
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	WriteFile("descendente.txt", s)
	fmt.Println(s)
}
