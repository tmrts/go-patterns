package main

type memento struct{
	state string
}

func (m *memento) getSavedState() string{
	return m.state
}