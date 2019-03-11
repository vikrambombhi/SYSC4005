package workstation

import (
	"fmt"
	"sync"
	"time"
)

type ws struct {
	buffers []chan bool
	name    string
	timing  []float64
}

// Create Workstation and have it start assembeling products
func Workstation(buffers []chan bool, timing []float64, name string) {
	ws := ws{buffers, name, timing}
	go ws.start()
}

func (ws ws) start() {
	var wg sync.WaitGroup

	for _, assembeTime := range ws.timing {
		for _, buffer := range ws.buffers {
			go func(buffer chan bool) {
				getComponent(buffer, &wg)
			}(buffer)
			wg.Add(1)
		}
		wg.Wait()
		time.Sleep(time.Duration(assembeTime * 1000000000))
		fmt.Printf("%s making product\n", ws.name)
	}
}

func getComponent(buffer chan bool, wg *sync.WaitGroup) {
	<-buffer
	wg.Done()
}
