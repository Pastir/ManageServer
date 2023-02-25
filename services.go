package ManageServer

type (
	// Service - service  control
	Service interface {
		// Init - init service initialization => configuration, connections...
		Init() error

		// Close - close connections, release resources
		Close() error

		// Check - service availability check
		Check() error
	}

	ServiceWatcher struct {
		Services []Service
	}
)
