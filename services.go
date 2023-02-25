package ManageServer

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
	}
	return nil
}
