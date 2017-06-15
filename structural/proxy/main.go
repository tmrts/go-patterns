package main

import "fmt"

// For example:
// we must a execute some command
// so before that we must to create new terminal session
// and provide our user name and command
func main() {
	var excResp string // response of executed command
	var excErr error   // error of executing command

	// Create new instance of Proxy terminal
	t := new(Terminal)

	// Try execute command with invalid user
	if excResp, excErr = t.ExecuteCommand("badGopher", "capture_land"); excErr != nil {
		fmt.Printf("%s\n", excErr.Error()) // Prints: You are not allowed to execute commands
	}

	// Try execute incorrect command
	if excResp, excErr = t.ExecuteCommand("gopher", "doJob"); excErr != nil {
		fmt.Printf("%s\n", excErr.Error()) // Prints: I know only how to say hello (command: say_hi)
	}

	// Show correct result of execution
	fmt.Println(excResp) // Prints: Hi gopher
}

/*
	From that it's can be different terminals realizations with different methods, propertys, yda yda...
*/

// GopherTerminal for example:
// Its a "huge" structure with different public methods
type GopherTerminal struct {
	// user is a current authorized user
	user string
}

// Execute just runs known commands for current authorized user
func (gt *GopherTerminal) Execute(cmd string) (resp string, err error) {

	if cmd != "say_hi" {
		err = fmt.Errorf("Unknown command")
	}

	resp = fmt.Sprintf("Hi %s", gt.user)

	return
}

/*
	And now we will create owr proxy to deliver user and commands to specific objects
*/

// Terminal is a implementation of Proxy, it's  validates and sends data to GopherTerminal
// As example before send commands, user must be authorized
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

	// if user allowed to execute send commands then, for example we can decide which terminal can be used, remote or local etc..
	// but for example we just creating new instance of terminal,
	// set current user and send user command to execution in terminal
	t.gopherTerminal = new(GopherTerminal)
	t.gopherTerminal.user = user
	if resp, err = t.gopherTerminal.Execute(command); err != nil {
		err = fmt.Errorf("I know only how to say hello (command: say_hi)")
		return
	}

	return
}
