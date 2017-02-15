package registry

import (
	"context"
	"crypto/tls"
	"time"
)

type Options struct {
	Endpoints []string
	Timeout   time.Duration
	Secure    bool
	TLSConfig *tls.Config

	Context context.Context
}

type Option func(*Options)

type RegisterOptions struct {
	TTL     time.Duration
	Context context.Context
}

type RegisterOption func(*RegisterOptions)
