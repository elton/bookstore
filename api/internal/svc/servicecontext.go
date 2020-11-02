package svc

import (
	"github.com/elton/bookstore/api/internal/config"
	"github.com/elton/bookstore/rpc/add/adder"
	"github.com/elton/bookstore/rpc/check/checker"
	"github.com/tal-tech/go-zero/zrpc"
)

// ServiceContext 把相关的RPC服务引入到API的service层中
type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder
	Checker checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
