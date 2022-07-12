package main

import (
	"errors"
	"fmt"
	"time"
)

type Message string

type Greeter struct {
	Grumpy  bool
	Message Message // <- adding a Message field
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

// provider
func NewMessage() Message {
	return Message("Hi there!")
}

// provider
func NewGreeterOld(m Message) Greeter {
	return Greeter{Message: m}
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

// provider
func NewEventOld(g Greeter) Event {
	return Event{Greeter: g}
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (g Greeter) Greet() Message {
	return g.Message
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	// 无依赖注入框架情况-通过线性关系完成依赖关系的逐步构建
	// 推导
	// 全局变量
	// 缺点: 会导致你的变量过于暴露
	// 初始化方法
	// 缺点: 不方便共享
	message := NewMessage()
	greeter := NewGreeter(message)
	event, _ := NewEvent(greeter)
	event.Start()

	// 有依赖注入框架情况
	e, _ := InitializeEvent()
	e.Start()
}
