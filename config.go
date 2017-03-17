package registry

import (
	"strings"
	"time"

	"github.com/docker/libkv/store"
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type registryConfig struct {
	ProviderString string        `json:"provider" config:"registry.provider"`
	Provider       store.Backend `json:"-" config:"-"`
	Endpoints      []string      `json:"endpoints" config:"registry.endpoints"`
	Username       string        `json:"username" config:"registry.username"`
	Password       string        `json:"-" config:"registry.password"`
	Timeout        time.Duration `json:"timeout" config:"registry.timeout"`
	done           chan struct{} `json:"-" config:"-"`
}

var (
	Config = &registryConfig{
		done: make(chan struct{}),
	}
)

func (registryConfig) ConfigName() string {
	return "Registry"
}

func (a *registryConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *registryConfig) Read() {
	vipertags.Fill(a)
	switch strings.ToLower(a.ProviderString) {
	case "consul":
		a.Provider = store.CONSUL
	case "zk", "zookeeper":
		a.Provider = store.ZK
	case "bolt", "boltdb":
		a.Provider = store.BOLTDB
	case "etcd":
		a.Provider = store.ETCD
	default:
		a.Provider = store.ETCD
	}
}

func (c registryConfig) Wait() {
	<-c.done
}

func (c registryConfig) String() string {
	return pp.Sprintln(c)
}

func (c registryConfig) Debug() {
	log.Debug("Registry Config = ", c)
}

func init() {
	config.Register(Config)
}
