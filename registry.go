package registry

type Registry struct {
	Register(*Service, ...RegisterOption) error
	Deregister(*Service) error
	GetService(string) ([]*Service, error)
	ListServices() ([]*Service, error)
	String() string
}
