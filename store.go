package registry

import (
	"strings"

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
	endpoints := []string{}
	for _, e := range Config.Endpoints {
		endpoints = append(endpoints, strings.TrimLeft(strings.TrimLeft(e, "http://"), "https://"))
	}
	libkv.NewStore(Config.Provider, endpoints, &store.Config{
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
