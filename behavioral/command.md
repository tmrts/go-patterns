# Command pattern

The [command pattern](https://en.wikipedia.org/wiki/Command_pattern) encapsulates all information needed to perform an action or event trigger at later time.

```
                                              +-----------------+
                                              |    command      |
                                              +-----------------+
                                              | execute()       |
                                              +------^----------+
                                                     v
                                                     |
                                                     |
                                                     |
                   +------------------+              |
                   |     receiver     |              |
+--------+         +------------------+       +------+---------------+
| client +-------->| action()         |<------+    concreteCommand   |
|        |         +------------------+       +----------------------+
|        |                                    | state                |
|        |                                    +----------------------+      +---------------------+
|        +----------------------------------->| execute() 0-----------------+ receiver.action()   |
+--------+                                    +----------------------+      +---------------------+
```

## Implementation


### Receiver

Recevier works when **command's** execute method is called.

```go
type receiver interface {
	action(string)
}

type concreteReceiver struct{}

func (c *concreteReceiver) action(m string) {
	fmt.Println(m)
}
```

### Command

Command knows about receiver and invokes it's action method.

```go
type command interface {
	execute()
}

type concreteCommand struct {
	r receiver
}

func (c *concreteCommand) execute() {
	c.r.action("do action.")
}
```

### Invoker

Invoker knows and to execute given command which is pushed to history of commands.

```go
type invoker interface {
	execute(command)
}

type concreteInvoker struct {
	history []command
}

func (c *concreteInvoker) execute(cmd command) {
	c.history = append(c.history, cmd)
	cmd.execute()
}
```

### Client

Client desides which command it should execute at which points. To execute a command it passes one to invoker object.

```go
type client struct {
	inv invoker

	// commands available for client
	sampleCommand *concreteCommand
}

func newClient() *client {
	c := &client{}
	c.inv = &concreteInvoker{}
	c.sampleCommand = &concreteCommand{&concreteReceiver{}}

	return c
}

func (c *client) do() {
	// execute command by client. Could have multiple conditions to deside which command to execute.
	c.inv.execute(c.sampleCommand)
}
```
