// Package http provides a http transport
package http

import (
	"github.com/divisionone/go-micro/transport"
	"github.com/divisionone/go-micro/transport/http"
)

/*
	HTTP transport is the default synchronous communication mechanism for go-micro.
	Implementation here https://godoc.org/github.com/divisionone/go-micro/transport/http
	We add a link here for completeness
*/

func NewTransport(opts ...transport.Option) transport.Transport {
	return http.NewTransport(opts...)
}
