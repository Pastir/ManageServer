package ManageServer

import (
	"fmt"
	"time"
)

type (
	// Service - service  control
	Service interface {
		// Init service initialization => configuration, connections...
		Init() error

		// Close connections, release resources
		Close() error

		// Check service availability check
		Check() error

		Name() string
	}

	ServiceWatcher struct {
		Services []Service
	}
)

func (sw *ServiceWatcher) InitServices() error {

	for _, serv := range sw.Services {
		err := serv.Init()
		if err != nil {
			return err
		}
		fmt.Println(time.Now().String() + " Initialization " + serv.Name() + " successful")
	}
	return nil
}

// AddService - add type Service
func (sw *ServiceWatcher) AddService(s Service) {
	sw.Services = append(sw.Services, s)
}
