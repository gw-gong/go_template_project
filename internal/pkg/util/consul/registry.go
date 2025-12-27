package consul

import (
	"fmt"

	"github.com/gw-gong/go-template-project/internal/config/types"

	"github.com/gw-gong/gwkit-go/grpc/consul"
	"github.com/gw-gong/gwkit-go/log"
	"github.com/gw-gong/gwkit-go/util/str"
)

func RegisterServices(serviceConfigs []*types.ServiceConfig, port int) (func(), error) {
	serviceRegistryMap := make(map[string]consul.ConsulRegistry, len(serviceConfigs))
	for _, serviceConfig := range serviceConfigs {
		serviceRegistry, err := consul.NewConsulRegistry(serviceConfig.ServiceName)
		if err != nil {
			return nil, fmt.Errorf("new consul registry failed: %w, serviceName: %s", err, serviceConfig.ServiceName)
		}

		serviceID := str.GenerateUUID()
		err = serviceRegistry.Register(serviceID, port, serviceConfig.Tags)
		if err != nil {
			return nil, fmt.Errorf("register service failed: %w, serviceName: %s, serviceID: %s", err, serviceConfig.ServiceName, serviceID)
		}

		serviceRegistryMap[serviceID] = serviceRegistry
	}
	return func() {
		for serviceID, serviceRegistry := range serviceRegistryMap {
			err := serviceRegistry.Deregister(serviceID)
			if err != nil {
				log.Error("deregister service failed", log.Err(err), log.Str("serviceID", serviceID))
			}
		}
		log.Info("deregister services success")
	}, nil
}
