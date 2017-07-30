package registry

import (
	"strings"
	"time"

	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/libkv/store"
	"github.com/rai-project/serializer"
	_ "github.com/rai-project/serializer/bson"
	_ "github.com/rai-project/serializer/json"
	_ "github.com/rai-project/serializer/jsonpb"
	"github.com/rai-project/vipertags"
)

type registryConfig struct {
	ProviderString          string        `json:"provider" config:"registry.provider"`
	Provider                store.Backend `json:"-" config:"-"`
	Endpoints               []string      `json:"endpoints" config:"registry.endpoints"`
	Username                string        `json:"username" config:"registry.username"`
	Password                string        `json:"-" config:"registry.password"`
	Timeout                 time.Duration `json:"timeout" config:"registry.timeout" default:"10s"`
	HeaderTimeoutPerRequest time.Duration `json:"header_timeout_per_request" config:"registry.header_timeout_per_request"`
	Bucket                  string        `json:"bucket" config:"registry.bucket"`
	AutoSync                bool          `json:"auto_Sync" config:"registry.auto_sync" default:"true"`
	Certificate             string        `json:"certificate" config:"registry.certificate"`

	Serializer     serializer.Serializer `json:"-" config:"-"`
	SerializerName string                `json:"serializer_name" config:"registry.serializer" default:"jsonpb"`

	done chan struct{} `json:"-" config:"-"`
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
	defer close(a.done)
	config.App.Wait()
	vipertags.Fill(a)
	if a.Certificate != "" {
		a.Certificate = decrypt(a.Certificate)
	}
	a.Provider = getProvider(a.ProviderString)
	if strings.TrimSpace(a.Bucket) == "" {
		a.Bucket = config.App.Name
	}
	a.Serializer, _ = serializer.FromName(a.SerializerName)
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
