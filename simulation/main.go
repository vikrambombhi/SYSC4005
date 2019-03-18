package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/vikrambombhi/SYSC4005/simulation/component"
	"github.com/vikrambombhi/SYSC4005/simulation/inspector"
	"github.com/vikrambombhi/SYSC4005/simulation/workstation"
)

func main() {
	dataDir := flag.String("data", "./", "directory data files are in")
	alternativeDesign := flag.Bool("alt", false, "Flag to specify if alternative design should be used")
	flag.Parse()
	fmt.Println("alternative design flag set to:", *alternativeDesign)
	var wg sync.WaitGroup
	ws1 := readFile("../data/ws1.dat")
	ws2 := readFile("../data/ws2.dat")
	ws3 := readFile("../data/ws3.dat")

	servinsp1 := readFile(*dataDir + "/servinsp1.dat")
	servinsp22 := readFile(*dataDir + "/servinsp22.dat")
	servinsp23 := readFile(*dataDir + "/servinsp23.dat")

	ws1Component1 := make(chan string, 2)
	workstation.Workstation(&wg, []chan string{ws1Component1}, ws1, "ws1")
	wg.Add(1)

	ws2Component1 := make(chan string, 2)
	ws2Component2 := make(chan string, 2)
	workstation.Workstation(&wg, []chan string{ws2Component1, ws2Component2}, ws2, "ws2")
	wg.Add(1)

	ws3Component1 := make(chan string, 2)
	ws3Component3 := make(chan string, 2)
	workstation.Workstation(&wg, []chan string{ws3Component1, ws3Component3}, ws3, "ws3")
	wg.Add(1)

	var component1 component.Component
	if *alternativeDesign {
		component1 = component.CreateComponent([]chan string{ws3Component1, ws2Component1, ws1Component1}, servinsp1, "component1")
	} else {
		component1 = component.CreateComponent([]chan string{ws1Component1, ws2Component1, ws3Component1}, servinsp1, "component1")
	}
	inspector.Inspector([]*component.Component{&component1}, "inspector1")

	component2 := component.CreateComponent([]chan string{ws2Component2}, servinsp22, "component2")
	component3 := component.CreateComponent([]chan string{ws3Component3}, servinsp23, "component3")
	inspector.Inspector([]*component.Component{&component2, &component3}, "inspector2")

	wg.Wait()
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
