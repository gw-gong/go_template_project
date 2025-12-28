package types

import "github.com/gw-gong/gwkit-go/grpc/consul"

type RpcServer struct {
	Port            int                     `yaml:"port" mapstructure:"port"`
	RegisterEntries []*consul.RegisterEntry `yaml:"register_entries" mapstructure:"register_entries"`
}
