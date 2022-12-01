package main

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Engineer Name: " + e.Name
}

type Manager struct {
	Name string
}

func (m *Manager) GetName() string {
	return "Manager Name: " + m.Name
}

func PrintDevName(e Employee) {
	println(e.GetName())
}

func main() {
	e := &Engineer{Name: "John"}
	m := &Manager{Name: "Jose"}

	PrintDevName(e)
	PrintDevName(m)
}
