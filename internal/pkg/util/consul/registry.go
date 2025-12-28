package consul

import (
	"fmt"

	"github.com/gw-gong/gwkit-go/grpc/consul"
	"github.com/gw-gong/gwkit-go/log"
)

func RegisterServices(consulClient consul.ConsulClient, registerEntries []*consul.RegisterEntry, port int) (func(), error) {
	serviceIDs := make([]string, 0, len(registerEntries))
	for _, registerEntry := range registerEntries {
		err := consulClient.Register(registerEntry, port, false)
		if err != nil {
			return nil, fmt.Errorf("register service failed: %w, serviceName: %s", err, registerEntry.ServiceName)
		}
		serviceIDs = append(serviceIDs, registerEntry.ServiceID)
	}
	return func() {
		for _, serviceID := range serviceIDs {
			err := consulClient.Deregister(serviceID)
			if err != nil {
				log.Error("deregister service failed", log.Err(err), log.Str("serviceID", serviceID))
			}
		}
		log.Info("deregister services success")
	}, nil
}
