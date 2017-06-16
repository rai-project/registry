package registry

import (
	"context"

	"github.com/rai-project/libkv"
	store "github.com/rai-project/libkv/store"
	"github.com/rai-project/libkv/store/boltdb"
	"github.com/rai-project/libkv/store/consul"
	"github.com/rai-project/libkv/store/etcd"
	"github.com/rai-project/libkv/store/mock"
	"github.com/rai-project/libkv/store/zookeeper"
)

type Store interface {
	store.Store
}

func New(opts ...Option) (store.Store, error) {
	options := Options{
		Provider:          Config.Provider,
		Endpoints:         cleanupEndpoints(Config.Endpoints),
		Username:          Config.Username,
		Password:          Config.Password,
		Timeout:           Config.Timeout,
		TLSConfig:         nil,
		PersistConnection: true,
		Context:           context.Background(),
	}
	if Config.HeaderTimeoutPerRequest != 0 {
		HeaderTimeoutPerRequest(Config.HeaderTimeoutPerRequest)(&options)
	}
	if Config.Certificate != "" {
		TLSCertificate(Config.Certificate)(&options)
	}
	AutoSync(Config.AutoSync)(&options)
	for _, o := range opts {
		o(&options)
	}
	storeConfig := &store.Config{
		ClientTLS:         &store.ClientTLSConfig{},
		ConnectionTimeout: options.Timeout,
		Username:          options.Username,
		Password:          options.Password,
		TLS:               options.TLSConfig,
		Bucket:            options.Bucket,
		PersistConnection: options.PersistConnection,
		Context:           options.Context,
	}
	if options.Provider == store.Backend("mock") {
		return mock.New(options.Endpoints, storeConfig)
	}
	return libkv.NewStore(options.Provider, options.Endpoints, storeConfig)
}

func init() {
	consul.Register()
	boltdb.Register()
	etcd.Register()
	zookeeper.Register()
}
