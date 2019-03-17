package workstation

import (
	"fmt"
	"sync"
	"time"
)

type ws struct {
	buffers []chan bool
	timing  []float64
	name    string
	wg      *sync.WaitGroup
}

// Workstation creates a workstation entity and has it start assembeling products
func Workstation(wg *sync.WaitGroup, buffers []chan bool, timing []float64, name string) {
	ws := ws{buffers, timing, name, wg}
	go ws.start()
}

func (ws ws) start() {
	var wg sync.WaitGroup

	for i, assembeTime := range ws.timing {
		for _, buffer := range ws.buffers {
			go func(buffer chan bool) {
				getComponent(buffer, &wg)
			}(buffer)
			wg.Add(1)
		}
		wg.Wait()
		// Simulates assembly time
		time.Sleep(time.Duration(assembeTime * 1000000000))
		fmt.Printf("%s done making product #%d\n\n\n", ws.name, i+1)
	}

	ws.wg.Done()
}

func getComponent(buffer chan bool, wg *sync.WaitGroup) {
	<-buffer
	wg.Done()
}
