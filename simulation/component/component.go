package component

import "time"

// Component is used to model a component. It maintains a internal list of inspection times and buffers
type Component struct {
	buffers  []chan SendVal
	timing   []float64
	position int
	name     string
}

type SendVal struct {
	Name string
	T    time.Time
}

// CreateComponent creates a component entity
func CreateComponent(buffers []chan SendVal, timings []float64, name string) Component {
	return Component{buffers, timings, 0, name}
}

func (component Component) GetName() string {
	return component.name
}

func (component Component) GetNextTime() float64 {
	inspectTime := component.timing[component.position%len(component.timing)]
	component.position = component.position + 1
	return inspectTime
}

func (component Component) GetBuffers() []chan SendVal {
	return component.buffers
}
