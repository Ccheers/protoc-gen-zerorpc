package svc

import (
	"github.com/Ccheers/protoc-gen-zerorpc/example/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
