/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package webkms

import "net/http"

// addHeaders function supports adding custom http headers.
type addHeaders func(req *http.Request) (*http.Header, error)

// HeadersOpts represents http header option.
type HeadersOpts struct {
	HeadersFunc addHeaders
}

// NewOpt creates a new empty additional http headers option.
// Not to be used directly. It's intended for implementations of remoteKMS.
// Use WithHeaders() option function below instead.
func NewOpt() *HeadersOpts {
	return &HeadersOpts{}
}

// HeadersOpt are the remoteKMS headers option.
type HeadersOpt func(opts *HeadersOpts)

// WithHeaders option is for setting additional http request headers (since it's a function, it can call a remote
// authorization server to fetch the necessary info needed in these headers).
func WithHeaders(addHeadersFunc addHeaders) HeadersOpt {
	return func(opts *HeadersOpts) {
		opts.HeadersFunc = addHeadersFunc
	}
}
