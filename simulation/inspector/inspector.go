package inspector

import (
	"math/rand"
	"time"

	"github.com/vikrambombhi/SYSC4005/simulation/component"
)

type inspector struct {
	components []*component.Component
	name       string
}

// Inspector creates a inspector entity and has it start inspecting components,
// the inspector places components onto buffer when done being inspected
func Inspector(components []*component.Component, name string) {
	inspector := inspector{components, name}
	go inspector.start()
}

func (inspector inspector) start() {
	// Guaranteed new seed every time
	rand.Seed(time.Now().UnixNano())
	for {
		rand := rand.Intn(len(inspector.components))
		component := inspector.components[rand]
		inspectComponent(component, inspector.name)
	}
}

// Inspects component and gives it to the first available buffer
func inspectComponent(component *component.Component, name string) {
	inspectTime := component.GetNextTime()
	// TODO: find least full buffer while sleeping
	time.Sleep(time.Duration(inspectTime * 1000000000))

	for {
		var leastFullBuffer chan string
		for _, buffer := range component.GetBuffers() {
			if leastFullBuffer == nil {
				leastFullBuffer = buffer
			} else if len(buffer) < len(leastFullBuffer) {
				leastFullBuffer = buffer
			}
		}
		select {
		case leastFullBuffer <- component.GetName():
			return
		default:
		}
	}
}
