package registry

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"time"

	"github.com/rai-project/libkv/store"
)

type Options struct {
	Provider          store.Backend
	Username          string
	Password          string
	Endpoints         []string
	Timeout           time.Duration
	Bucket            string
	TLSConfig         *tls.Config
	PersistConnection bool

	Context context.Context
}

type Option func(*Options)

type RegisterOptions struct {
	TTL     time.Duration
	Context context.Context
}

type RegisterOption func(*RegisterOptions)

func Bucket(s string) Option {
	return func(o *Options) {
		o.Bucket = s
	}
}

func Provider(s string) Option {
	return func(o *Options) {
		o.Provider = getProvider(s)
	}
}

func Username(s string) Option {
	return func(o *Options) {
		o.Username = s
	}
}

func Password(s string) Option {
	return func(o *Options) {
		o.Password = s
	}
}

func UsernamePassword(u string, p string) Option {
	return func(o *Options) {
		o.Username = u
		o.Password = p
	}
}

func Endpoint(addr string) Option {
	return Endpoints([]string{addr})
}

func Endpoints(addrs []string) Option {
	return func(o *Options) {
		o.Endpoints = cleanupEndpoints(addrs)
	}
}

func TLSCertificate(s string) Option {
	return func(o *Options) {
		var roots *x509.CertPool
		if o.TLSConfig != nil && o.TLSConfig.RootCAs != nil {
			roots = o.TLSConfig.RootCAs
		} else {
			roots = x509.NewCertPool()
		}
		cert := []byte(decrypt(s))
		roots.AppendCertsFromPEM(cert)

		o.TLSConfig = &tls.Config{
			RootCAs: roots,
		}
	}
}

func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

func PersistConnection(b bool) Option {
	return func(o *Options) {
		o.PersistConnection = b
	}
}

func HeaderTimeoutPerRequest(t time.Duration) Option {
	return func(o *Options) {
		o.Context = context.WithValue(o.Context, "HeaderTimeoutPerRequest", t)
	}
}

func AutoSync(v bool) Option {
	return func(o *Options) {
		o.Context = context.WithValue(o.Context, "AutoSync", v)
	}
}
