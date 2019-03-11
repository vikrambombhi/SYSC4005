package inspector

import (
	"fmt"
)

type inspector struct {
	buffers []chan bool
	name    string
}

// Create inspector and have it start inspecting components,
// place components onto buffer when done being inspected
func Inspector(buffers []chan bool, name string) {
	inspector := inspector{buffers, name}
	go inspector.start()
}

func (inspector inspector) start() {
	for {
		for i, buffer := range inspector.buffers {
			select {
			case buffer <- true:
				fmt.Printf("%s done inspecting component %d\n", inspector.name, i+1)
			}
		}
	}
}
