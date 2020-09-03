package main

import "fmt"

type task interface {
	execute()
	setNext(task)
}

type turnOnLights struct {
	next task
}

func (turnOnLights *turnOnLights) execute() {
	fmt.Println("Turning the lights on...")
	if turnOnLights.next != nil {
		turnOnLights.next.execute()
	}
}

func (turnOnLights *turnOnLights) setNext(next task) {
	turnOnLights.next = next
}

type turnOnComputer struct {
	next task
}

func (turnOnComputer *turnOnComputer) execute() {
	fmt.Println("Turning the computer on...")
	if turnOnComputer.next != nil {
		turnOnComputer.next.execute()
	}
}

func (turnOnComputer *turnOnComputer) setNext(next task) {
	turnOnComputer.next = next
}

type openCodeEditor struct {
	next task
}

func (openCodeEditor *openCodeEditor) execute() {
	fmt.Println("Opening the code editor...")
	if openCodeEditor.next != nil {
		openCodeEditor.next.execute()
	}
}

func (openCodeEditor *openCodeEditor) setNext(next task) {
	openCodeEditor.next = next
}

type code struct {
	next task
}

func (code *code) execute() {
	fmt.Println("Start coding in go...")
	if code.next != nil {
		code.next.execute()
	}
}

func (code *code) setNext(next task) {
	code.next = next
}

func main() {
	turnOnLights := &turnOnLights{}
	turnOnComputer := &turnOnComputer{}
	openCodeEditor := &openCodeEditor{}
	code := &code{}

	turnOnLights.setNext(turnOnComputer)
	turnOnComputer.setNext(openCodeEditor)
	openCodeEditor.setNext(code)

	turnOnLights.execute()
	// Out:
	// Turning the lights on...
	// Turning the computer on...
	// Opening the code editor...
	// Start coding in go...

	openCodeEditor.execute()
	// Out:
	// Opening the code editor...
	// Start coding in go...
}
