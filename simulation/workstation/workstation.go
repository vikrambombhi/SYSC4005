package workstation

import (
	"fmt"
	"sync"
	"time"

	"github.com/vikrambombhi/SYSC4005/simulation/component"
)

type ws struct {
	buffers []chan component.SendVal
	timing  []float64
	name    string
	wg      *sync.WaitGroup
	vals    map[string][]time.Duration
	mutex   *sync.Mutex
}

// Workstation creates a workstation entity and has it start assembeling products
func Workstation(wg *sync.WaitGroup, buffers []chan component.SendVal, timing []float64, name string) {
	ws := ws{
		buffers,
		timing,
		name,
		wg,
		make(map[string][]time.Duration),
		&sync.Mutex{},
	}
	go ws.start()
}

func (ws ws) start() {
	ws.startAssembleing()
	ws.wg.Done()
}

func (ws ws) startAssembleing() {
	last := time.Now()
	for i, assembeTime := range ws.timing {
		var wg sync.WaitGroup
		for _, buffer := range ws.buffers {
			wg.Add(1)
			go func(buffer chan component.SendVal, wg *sync.WaitGroup) {
				val := <-buffer
				wg.Done()
				timeSince := time.Since(val.T)
				ws.mutex.Lock()
				ws.vals[val.Name] = append(ws.vals[val.Name], timeSince)
				ws.mutex.Unlock()
			}(buffer, &wg)
		}
		wg.Wait()
		// Simulates assembly time
		time.Sleep(time.Duration(assembeTime * 1000000000))
		fmt.Printf("%s done making product #%d in %s\n", ws.name, i+1, time.Since(last))
		last = time.Now()
		if i == 100 {
			fmt.Printf("\n\n\n %v \n\n", ws.vals)
		}
	}
}
