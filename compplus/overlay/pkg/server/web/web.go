// Copyright (c) 2025-2026 Probo Inc <hello@probo.com>.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package web

import (
	"net/http"

	"go.probo.inc/probo/apps/console"
	"go.probo.inc/probo/pkg/server/statichandler"
)

type Server struct {
	*statichandler.Server
}

func NewServer() (*Server, error) {
	gzipOptions := statichandler.GzipOptions{
		EnableFileTypeCheck: false,
	}

	spaServer, err := statichandler.NewServer(console.StaticFiles, "dist", gzipOptions)
	if err != nil {
		return nil, err
	}

	return &Server{
		Server: spaServer,
	}, nil
}

func setSecurityHeaders(w http.ResponseWriter) {
	// TLS terminates at the production reverse proxy, but browsers still receive
	// these application-level protections on the final HTTPS response.
	w.Header().Set("Strict-Transport-Security", "max-age=31536000")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "no-referrer")
	w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w)
	s.Server.ServeHTTP(w, r)
}

func (s *Server) ServeSPA(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w)
	s.Server.ServeSPA(w, r)
}
