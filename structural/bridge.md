# Bridge Pattern
Decouples an interface from its implementation so that the two can vary independently

## Implementation

```go
type Info struct {
	Sender // who or which way to send information.        eg: QQ Wechat
	MsgTag // the tag of msg.  just like log level tag.    eg: COMMON IMPORTANCE
}

func (i Info) Sending(to, msg string) {
	msg = i.MsgTag(msg)
	i.Sender.Send(to, msg)
}

// interface

type Sender interface {
	Send(to, msg string)
}

type MsgTag func(string) string


// implementation

type QQSender struct {
}

func (QQSender) Send(to, msg string) {
	fmt.Println("[QQ] send to " + to + " with message: " + msg)
}

type WechatSender struct {
}

func (WechatSender) Send(to, msg string) {
	fmt.Println("[Wechat] send to " + to + " with message: " + msg)
}

func ImportanceTag(msg string) string {
	return "[IMPORTANCE] " + msg
}

func CommonTag(msg string) string {
	return "[COMMON] " + msg
}
```

## Usage
```go
info := &Info{MsgTag: ImportanceTag, Sender: QQSender{}}
info.Sending("cbping", "hello world")

info.Sender = WechatSender{}
info.Sending("cbping", "hello world")

info.MsgTag = CommonTag
info.Sending("cbping", "hello world")
```

