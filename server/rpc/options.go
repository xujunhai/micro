package rpc

import (
	mbroker "gitlab.ziroom.com/rent-web/micro/broker/memory"
	"gitlab.ziroom.com/rent-web/micro/registry/memory"
	tmem "gitlab.ziroom.com/rent-web/micro/transport/memory"

	"gitlab.ziroom.com/rent-web/micro/codec"
	"gitlab.ziroom.com/rent-web/micro/server"
)

func newOptions(opt ...server.Option) server.Options {
	opts := server.Options{
		Codecs:           make(map[string]codec.NewCodec),
		Metadata:         map[string]string{},
		RegisterInterval: server.DefaultRegisterInterval,
		RegisterTTL:      server.DefaultRegisterTTL,
	}

	for _, o := range opt {
		o(&opts)
	}

	if opts.Broker == nil {
		opts.Broker = mbroker.NewBroker()
	}

	if opts.Registry == nil {
		opts.Registry = memory.NewRegistry()
	}

	if opts.Transport == nil {
		opts.Transport = tmem.NewTransport()
	}

	if opts.RegisterCheck == nil {
		opts.RegisterCheck = server.DefaultRegisterCheck
	}

	if len(opts.Address) == 0 {
		opts.Address = server.DefaultAddress
	}

	if len(opts.Name) == 0 {
		opts.Name = server.DefaultName
	}

	if len(opts.Id) == 0 {
		opts.Id = server.DefaultId
	}

	if len(opts.Version) == 0 {
		opts.Version = server.DefaultVersion
	}

	return opts
}
