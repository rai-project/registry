package registry

import (
	"github.com/docker/libkv"
	store "github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
	"github.com/docker/libkv/store/consul"
	"github.com/docker/libkv/store/etcd"
	"github.com/docker/libkv/store/zookeeper"
)

type Store interface {
	store.Store
}

func NewStore() {
	libkv.NewStore(Config.Provider, Config.Endpoints, &store.Config{
		ConnectionTimeout: Config.Timeout,
		Username:          Config.Username,
		Password:          Config.Password,
		PersistConnection: true,
	})
}

func init() {
	consul.Register()
	boltdb.Register()
	etcd.Register()
	zookeeper.Register()
}
