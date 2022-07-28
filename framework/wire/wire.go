//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeEvent() (Event, error) {
	// 实体是实体,关系是关系,两者之间专门表征
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
