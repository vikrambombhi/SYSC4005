package inspector

import (
	"fmt"
	"math/rand"
	"time"
)

type inspector struct {
	components []Component
	name       string
}

// Component is used to model a component. It maintains a internal list of inspection times and buffers
type Component struct {
	buffers  []chan bool
	timing   []float64
	position int
	name     string
}

// CreateComponent creates a component entity
func CreateComponent(buffers []chan bool, timings []float64, name string) Component {
	return Component{buffers, timings, 0, name}
}

// Inspector creates a inspector entity and has it start inspecting components,
// the inspector places components onto buffer when done being inspected
func Inspector(components []Component, name string) {
	inspector := inspector{components, name}
	go inspector.start()
}

func (inspector inspector) start() {
	// Guaranteed new seed every time
	rand.Seed(time.Now().UnixNano())
	for {
		rand := rand.Intn(len(inspector.components))
		component := inspector.components[rand]
		inspectComponent(component)
		fmt.Printf("Done inspecting %s\n", component.name)
	}
}

// Inspects component and gives it to the first available buffer
func inspectComponent(component Component) {
	inspectTime := component.timing[component.position]
	// TODO: find least full buffer while sleeping
	time.Sleep(time.Duration(inspectTime * 1000000000))
	component.position++

	for {
		var leastFullBuffer chan bool
		for _, buffer := range component.buffers {
			if leastFullBuffer == nil {
				leastFullBuffer = buffer
			}
			if len(buffer) < len(leastFullBuffer) {
				leastFullBuffer = buffer
			}
		}
		select {
		case leastFullBuffer <- true:
			return
		}
	}
}
