//Source code adapted from https://en.wikipedia.org/wiki/Abstract_factory_pattern
//Added "otherOS" in case this code is executed on other os than windows, mac os x
package main

import (
	"fmt"
	"runtime"
)

type iButton interface {
	paint()
}

type iGUIFactory interface {
	createButton() iButton
}

type winFactory struct {
}

func (WF *winFactory) createButton() iButton {
	return newWinButton()
}

type osxFactory struct {
}

func (WF *osxFactory) createButton() iButton {
	return newOSXButton()
}

type otherOSFactory struct {
}

func (WF *otherOSFactory) createButton() iButton {
	return newOtherOSButton()
}

type winButton struct {
}

func (wb *winButton) paint() {
	fmt.Println("WinButton")
}

func newWinButton() *winButton {
	return &winButton{}
}

type osxButton struct {
}

func (ob *osxButton) paint() {
	fmt.Println("OSXButton")
}

func newOSXButton() *osxButton {
	return &osxButton{}
}

type otherOSButton struct {
}

func (ob *otherOSButton) paint() {
	fmt.Println("OtherOSButton")
}

func newOtherOSButton() *otherOSButton {
	return &otherOSButton{}
}

func main() {
	var factory iGUIFactory

	switch runtime.GOOS {
	case "windows":
		factory = &winFactory{}
		break
	case "darwin":
		factory = &osxFactory{}
		break
	default:
		factory = &otherOSFactory{}
	}

	button := factory.createButton()
	button.paint()
}
