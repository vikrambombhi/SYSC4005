package workstation

import (
	"fmt"
	"sync"
	"time"
)

type ws struct {
	buffers []chan string
	timing  []float64
	name    string
	wg      *sync.WaitGroup
}

// Workstation creates a workstation entity and has it start assembeling products
func Workstation(wg *sync.WaitGroup, buffers []chan string, timing []float64, name string) {
	ws := ws{buffers, timing, name, wg}
	go ws.start()
}

func (ws ws) start() {
	ws.startAssembleing()
	ws.wg.Done()
}

func (ws ws) startAssembleing() {
	for i, assembeTime := range ws.timing {
		var wg sync.WaitGroup
		for _, buffer := range ws.buffers {
			wg.Add(1)
			go func(buffer chan string, wg *sync.WaitGroup) {
				getComponent(buffer, wg)
			}(buffer, &wg)
		}
		wg.Wait()
		// Simulates assembly time
		time.Sleep(time.Duration(assembeTime * 1000000000))
		fmt.Printf("%s done making product #%d\n", ws.name, i+1)
	}
}

func getComponent(buffer chan string, wg *sync.WaitGroup) {
	<-buffer
	wg.Done()
}
