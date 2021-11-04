// Package rpc provides an app implementation for micro
package rpc

import (
	"context"
	mbroker "gitlab.ziroom.com/rent-web/micro/broker/memory"
	"gitlab.ziroom.com/rent-web/micro/client"
	rpcClient "gitlab.ziroom.com/rent-web/micro/client/rpc"
	"gitlab.ziroom.com/rent-web/micro/common/component"
	"gitlab.ziroom.com/rent-web/micro/common/constant"
	"gitlab.ziroom.com/rent-web/micro/registry/memory"
	"gitlab.ziroom.com/rent-web/micro/server"
	rpcServer "gitlab.ziroom.com/rent-web/micro/server/rpc"
	"gitlab.ziroom.com/rent-web/micro/service"
	tmem "gitlab.ziroom.com/rent-web/micro/transport/memory"
)

type rpcApp struct {
	opts service.Options
}

func (s *rpcApp) Name(name string) {
	s.opts.Server.Init(
		server.Name(name),
	)
}

// Init initialises options. Additionally it calls cmd.Init
// which parses command line flags. cmd.Init is only called
// on first Init.
func (s *rpcApp) Init(opts ...service.Option) {
	// process options
	for _, o := range opts {
		o(&s.opts)
	}
}

func (s *rpcApp) Options() service.Options {
	return s.opts
}

func (s *rpcApp) Call(name, ep string, req, rsp interface{}) error {
	r := s.Client().NewRequest(name, ep, req)
	return s.Client().Call(context.Background(), r, rsp)
}

func (s *rpcApp) Handle(v interface{}) error {
	h := s.Server().NewHandler(v)
	return s.Server().Handle(h)
}

func (s *rpcApp) Broadcast(topic string, msg interface{}) error {
	m := s.Client().NewMessage(topic, msg)
	return s.Client().Publish(context.Background(), m)
}

func (s *rpcApp) Subscribe(topic string, v interface{}) error {
	sub := s.Server().NewSubscriber(topic, v)
	return s.Server().Subscribe(sub)
}

func (s *rpcApp) Client() client.Client {
	return s.opts.Client
}

func (s *rpcApp) Server() server.Server {
	return s.opts.Server
}

func (s *rpcApp) String() string {
	return "rpc"
}

func (s *rpcApp) Start() error {
	for _, fn := range s.opts.BeforeStart {
		if err := fn(); err != nil {
			return err
		}
	}

	if err := s.opts.Server.Start(); err != nil {
		return err
	}

	for _, fn := range s.opts.AfterStart {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func (s *rpcApp) Stop() error {
	var gerr error

	for _, fn := range s.opts.BeforeStop {
		if err := fn(); err != nil {
			gerr = err
		}
	}

	if err := s.opts.Server.Stop(); err != nil {
		return err
	}

	for _, fn := range s.opts.AfterStop {
		if err := fn(); err != nil {
			gerr = err
		}
	}

	return gerr
}

func (s *rpcApp) Run() error {
	if err := s.Start(); err != nil {
		return err
	}

	// wait on context cancel
	<-s.opts.Context.Done()

	return s.Stop()
}

// NewApp returns a new micro app
func NewApp(opts ...service.Option) *rpcApp {
	//default init rpc client&server
	c := rpcClient.NewClient()
	s := rpcServer.NewServer()

	//default component
	b := mbroker.NewBroker()
	r := memory.NewRegistry()
	t := tmem.NewTransport()

	// set client options
	c.Init(
		client.Broker(b),
		client.Registry(r),
		client.Transport(t),
	)

	// set server options
	s.Init(
		server.Broker(b),
		server.Registry(r),
		server.Transport(t),
	)

	// define local opts
	options := service.Options{
		Broker:   b,
		Client:   c,
		Server:   s,
		Registry: r,
		Context:  context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	return &rpcApp{
		opts: options,
	}
}

func init() {
	component.SetServiceFactory(constant.DefaultProtocol, &rpcServiceFactory{})
}

type rpcServiceFactory struct{}

func (rpcServiceFactory) GetService(options service.Options) service.Service {
	return &rpcApp{
		opts: options,
	}
}
