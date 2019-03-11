package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/vikrambombhi/SYSC4005/simulation/inspector"
	"github.com/vikrambombhi/SYSC4005/simulation/workstation"
)

func main() {
	ws1 := readFile("../data/ws1.dat")
	ws2 := readFile("../data/ws2.dat")
	ws3 := readFile("../data/ws3.dat")

	ws1Component1 := make(chan bool, 2)
	workstation.Workstation([]chan bool{ws1Component1}, ws1, "ws1")

	ws2Component1 := make(chan bool, 2)
	ws2Component2 := make(chan bool, 2)
	workstation.Workstation([]chan bool{ws2Component1, ws2Component2}, ws2, "ws2")

	ws3Component1 := make(chan bool, 2)
	ws3Component3 := make(chan bool, 2)
	workstation.Workstation([]chan bool{ws3Component1, ws3Component3}, ws3, "ws3")

	inspector.Inspector([]chan bool{ws1Component1, ws2Component1, ws3Component1}, "inspector1")
	inspector.Inspector([]chan bool{ws2Component2, ws3Component3}, "inspector2")

	// todo: Replace this with a sync wait
	for {
	}
}

func readFile(filename string) []float64 {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)

	data := make([]float64, 0)
	s, e := Readln(r)
	for e == nil {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		data = append(data, f)
		s, e = Readln(r)
	}

	return data
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
