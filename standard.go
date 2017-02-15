package registry

import (
  "github.com/rai-project/config"
  "github.com/rai-project/registry/etcd"
)

var (
	Standard Registry
)


func	Register(s *Service, opts...RegisterOption) error {
  return std.Register(s, opts...)
}
func Deregister(s *Service) error {
  return std.Deregister(s)
}

func GetService(s string) ([]*Service, error) {
  return std.GetService(s)
}

func ListServices() ([]*Service, error) {
  return std.ListServices()
}

func String() string {
  return std.String()
}
