package types

import "github.com/gw-gong/gwkit-go/grpc/consul"

type RpcServer struct {
	Port     int              `yaml:"port" mapstructure:"port"`
	Services []*ServiceConfig `yaml:"services" mapstructure:"services"`
}

type ServiceConfig struct {
	ServiceName consul.ServiceName `yaml:"service_name" mapstructure:"service_name"`
	Tags        []string           `yaml:"tags" mapstructure:"tags"`
}
