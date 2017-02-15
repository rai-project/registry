package etcd

import (
	"time"

	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type etcdConfig struct {
	Provider  string        `json:"provider" config:"registry.provider" default:"etcd"`
	Endpoints []string      `json:"endpoints" config:"registry.endpoints"`
	Username  string        `json:"username" config:"registry.username"`
	Password  string        `json:"-" config:"registry.password"`
	Timeout   time.Duration `json:"timeout" config:"registry.timeout"`
}

var (
	Config = &etcdConfig{}
)

func (etcdConfig) ConfigName() string {
	return "Etcd"
}

func (etcdConfig) setDefaults() {
}

func (a *etcdConfig) Read() {
	vipertags.Fill(a)
}

func (c etcdConfig) String() string {
	return pp.Sprintln(c)
}

func (c etcdConfig) Debug() {
	log.Debug("Etcd Config = ", c)
}

func init() {
	config.Register(Config)
}
