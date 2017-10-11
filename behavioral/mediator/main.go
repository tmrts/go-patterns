package main

import (
	"container/list"
	"fmt"
	"time"
)

/*
	Mediator
*/

// ChatRoomMediator main interface between colleague and mediator
type ChatRoomMediator interface {
	JoinUser(uc UserColleague)
	SendMessage(uc UserColleague)
}

// ChatRoom is a concrete mediator with all joined colleagues
type ChatRoom struct {
	users *list.List
}

// JoinUser join concrete colleagues to concrete mediator (ChatRoom)
func (cr *ChatRoom) JoinUser(uc UserColleague) {
	cr.users.PushBack(uc)
}

// SendMessage use mediator to communicate between colleagues
func (cr *ChatRoom) SendMessage(uc UserColleague) {
	for e := cr.users.Front(); e != nil; e = e.Next() {
		if e.Value == uc {
			ct := time.Now().Format(time.Kitchen)
			fmt.Printf("|%s| %s -> : %s \n", ct, uc.GetName(), uc.GetMessage())
		}
	}
}

// NewChatRoom creates concrete mediator
func NewChatRoom() *ChatRoom {
	return &ChatRoom{list.New()}
}

/*
	Colleague
*/

// UserColleague implements colleague to communicate with mediator
type UserColleague interface {
	GetName() string
	GetMessage() string
}

// User represents concrete colleague
type User struct {
	name     string
	message  string
	mediator ChatRoomMediator
}

// GetMessage returns last wrote message
func (u *User) GetName() string {
	return u.name
}

// GetMessage returns last wrote message
func (u *User) GetMessage() string {
	return u.message
}

// WriteMessage concrete colleague send's message via mediator
func (u *User) writeMessage(message string) {
	u.message = message

	// Ask mediator to send message
	u.mediator.SendMessage(u)
}

// AddUser will make relationship between concrete mediator and concrete colleague
func AddUser(name string, chatRoom ChatRoomMediator) *User {
	// Create new colleague(user) and join mediator(room)
	user := &User{name: name, message: "", mediator: chatRoom}

	// Join colleague(user) to mediator(room)
	chatRoom.JoinUser(user)

	return user
}

/*
	Showcase
*/
func main() {
	// Create mediator
	barrelHouse := NewChatRoom()

	// Create colleagues and join to mediator
	JohnInBarrelHouse := AddUser("John Doe", barrelHouse)
	AlmaInBarrelHouse := AddUser("Alma", barrelHouse)

	// Now colleagues will send messages via mediator
	JohnInBarrelHouse.writeMessage("Hello gophers")
	AlmaInBarrelHouse.writeMessage("Hey John!")
	JohnInBarrelHouse.writeMessage("Alma, let's discuss Mediator pattern?")

	// Example of output
	// |8:33PM| John Doe -> : Hello gophers
	// |8:33PM| Alma -> : Hey John!
	// |8:34PM| John Doe -> : Alma, let's discuss Mediator pattern?
}
