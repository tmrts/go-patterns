package main

import "fmt"

// Terminal represents a public methods
// whose allows to interact with "client"
type ITerminal interface {
	Execute(someCommand string) string
}

// GopherTerminal incapsulated terminal
type GopherTerminal struct {
	// User is a authorized user
	user string
}

// Implements Terminal interface for GopherTerminal
// Execute just run known command
func (gt *GopherTerminal) Execute(cmd string) (resp string) {
	resp = "Unknown command"

	if cmd == "say_hi" {
		resp = fmt.Sprintf("Hi %s", gt.user)
	}

	return
}

// Implements Proxy terminal to validate if user can use it
// As example before send command user, must be authorized to do it
type Terminal struct {
	gopherTerminal *GopherTerminal
}

// ExecuteCommand intercepts execution of command, implements authorizing user, validates it and
// poxing command to real terminal (gopherTerminal) method
func (t *Terminal) ExecuteCommand(user, command string) (resp string, err error) {
	// As we use terminal like proxy, then
	// we will intercept user name to validate if it's allowed to execute commands
	if user != "gopher" {
		err = fmt.Errorf("You are not allowed to execute commands")
		return
	}

	// if user allowed to execute send commands then,
	// create new instance of terminal, set current user and send user command to execution
	t.gopherTerminal = new(GopherTerminal)
	t.gopherTerminal.user = user
	if resp = t.gopherTerminal.Execute(command); resp == "Unknown command" {
		err = fmt.Errorf("I know only how to say hello (type: say_hi)")
		return
	}

	return
}

// For example:
// we must a execute some command
// so before that we must to create new terminal session
// and provide our user name and command
func main() {
	// Create new instance of terminal
	newTerm := new(Terminal)
	// Try execute command
	resp, err := newTerm.ExecuteCommand("gopher", "doJob")
	if err != nil {
		fmt.Println(err.Error()) // print Unknown command
	}

	// Handle result of execution
	fmt.Println(resp)
}
