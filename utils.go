package registry

import (
	"encoding/base64"
	"strings"

	store "github.com/rai-project/libkv/store"
	"github.com/rai-project/config"
	"github.com/rai-project/utils"
)

func decrypt(s string) string {
	q, err := base64.StdEncoding.DecodeString(s)
	if err == nil {
		s = string(q)
	}
	if utils.IsEncryptedString(s) {
		c, err := utils.DecryptStringBase64(config.App.Secret, s)
		if err == nil {
			return c
		}
	}
	return s
}

func cleanupEndpoints(endpoints []string) []string {
	es := make([]string, len(endpoints))
	for ii, e := range endpoints {
		es[ii] = strings.TrimLeft(strings.TrimLeft(e, "http://"), "https://")
	}
	return es
}

func getProvider(s string) store.Backend {
	switch strings.ToLower(s) {
	case "consul":
		return store.CONSUL
	case "zk", "zookeeper":
		return store.ZK
	case "bolt", "boltdb":
		return store.BOLTDB
	case "etcd":
		return store.ETCD
	case "mock":
		return store.Backend("mock")
	default:
		return store.ETCD
	}
	return store.ETCD
}

func cleanup(name string) string {
	name = strings.Replace(name, "p3srservices.", "p3sr/services/", -1)
	name = strings.Replace(name, "cermineproto.", "p3sr/services/", -1)
	name = strings.Replace(name, ".", "/", -1)
	name = strings.TrimSuffix(name, "Service")
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	return name
}
