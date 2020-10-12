package server

import (
	"crypto/tls"
	"net/http"

	"github.com/micro/micro/v3/internal/api/resolver"
	"github.com/micro/micro/v3/internal/api/server/acme"
)

// Server serves api requests
type Server interface {
	Address() string
	Init(opts ...Option) error
	Handle(path string, handler http.Handler)
	Start() error
	Stop() error
}

type Option func(o *Options)

type Options struct {
	EnableACME   bool
	EnableCORS   bool
	ACMEProvider acme.Provider
	EnableTLS    bool
	ACMEHosts    []string
	TLSConfig    *tls.Config
	Resolver     resolver.Resolver
	Wrappers     []Wrapper
}

type Wrapper func(h http.Handler) http.Handler

func WrapHandler(w ...Wrapper) Option {
	return func(o *Options) {
		o.Wrappers = append(o.Wrappers, w...)
	}
}

func EnableCORS(b bool) Option {
	return func(o *Options) {
		o.EnableCORS = b
	}
}

func EnableACME(b bool) Option {
	return func(o *Options) {
		o.EnableACME = b
	}
}

func ACMEHosts(hosts ...string) Option {
	return func(o *Options) {
		o.ACMEHosts = hosts
	}
}

func ACMEProvider(p acme.Provider) Option {
	return func(o *Options) {
		o.ACMEProvider = p
	}
}

func EnableTLS(b bool) Option {
	return func(o *Options) {
		o.EnableTLS = b
	}
}

func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

func Resolver(r resolver.Resolver) Option {
	return func(o *Options) {
		o.Resolver = r
	}
}
