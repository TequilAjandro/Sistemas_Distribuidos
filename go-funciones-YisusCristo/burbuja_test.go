package main

import (
	"testing"
)

func TestBurbuja01(t *testing.T) {
	s := []int64{1, -10, 90, 14, -100, -2}
	s2 := []int64{-100, -10, -2, 1, 14, 90}

	Burbuja(s)

	for i, _ := range s {
		if s[i] != s2[i] {
			t.Errorf("No está ordenado")
		}
	}
}


func TestBurbuja02(t *testing.T) {
	s := []int64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2}
	s2 := []int64{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	Burbuja(s)

	for i, _ := range s {
		if s[i] != s2[i] {
			t.Errorf("No está ordenado")
		}
	}
}


func TestBurbuja03(t *testing.T) {
	s := []int64{}
	s2 := []int64{}

	Burbuja(s)

	for i, _ := range s {
		if s[i] != s2[i] {
			t.Errorf("No está ordenado")
		}
	}
}
