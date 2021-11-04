// Package service encapsulates the client, server and other interfaces to provide a complete dapp
package service

import (
	"gitlab.ziroom.com/rent-web/micro/client"
	"gitlab.ziroom.com/rent-web/micro/server"
)

// service is an interface for distributed apps
type Service interface {
	// Set the current application name
	Name(string)
	// Call an application by name and endpoint
	Call(name, ep string, req, rsp interface{}) error
	// Register a handler e.g a public Go struct/method with signature func(context.Context, *Request, *Response) error
	Handle(v interface{}) error
	// Broadcast a message to all subscribers
	Broadcast(topic string, msg interface{}) error
	// Subscribe to broadcast messages. Signature is public Go func or struct with signature func(context.Context, *Message) error
	Subscribe(topic string, v interface{}) error
	// Run the application
	Run() error
	// Get the service's client
	Client() client.Client
	// Get the service's client
	Server() server.Server
}

//use configuration build service
type Factory interface {
	GetService(Options) Service
}
