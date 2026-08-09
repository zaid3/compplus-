// Copyright (c) 2026 Probo Inc <hello@probo.com>.
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

package connect_v1

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.probo.inc/probo/pkg/server/gqlutils"
)

type authRatePolicy struct {
	ipLimit    int
	emailLimit int
	window     time.Duration
}

type authRateBucket struct {
	count     int
	expiresAt time.Time
}

type publicAuthRateLimiter struct {
	mu        sync.Mutex
	buckets   map[string]authRateBucket
	lastPrune time.Time
}

const maxPublicAuthRateBuckets = 20000

var (
	publicAuthLimiter = &publicAuthRateLimiter{buckets: make(map[string]authRateBucket)}
	publicAuthPolicies = map[string]authRatePolicy{
		"signIn":                  {ipLimit: 20, emailLimit: 10, window: 5 * time.Minute},
		"signUp":                  {ipLimit: 8, emailLimit: 3, window: 15 * time.Minute},
		"forgotPassword":          {ipLimit: 10, emailLimit: 3, window: 15 * time.Minute},
		"resendVerificationEmail": {ipLimit: 10, emailLimit: 3, window: 15 * time.Minute},
		"resetPassword":           {ipLimit: 10, window: 15 * time.Minute},
		"verifyEmail":             {ipLimit: 20, window: 15 * time.Minute},
	}
)

func (l *publicAuthRateLimiter) allow(key string, limit int, window time.Duration) bool {
	if key == "" || limit <= 0 {
		return true
	}

	now := time.Now()

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.lastPrune.IsZero() || now.Sub(l.lastPrune) >= time.Minute {
		for bucketKey, bucket := range l.buckets {
			if !now.Before(bucket.expiresAt) {
				delete(l.buckets, bucketKey)
			}
		}
		l.lastPrune = now
	}

	bucket, exists := l.buckets[key]
	if !exists || !now.Before(bucket.expiresAt) {
		if !exists && len(l.buckets) >= maxPublicAuthRateBuckets {
			// Fail closed instead of allowing an attacker to grow memory without
			// bound by rotating arbitrary email addresses.
			return false
		}
		l.buckets[key] = authRateBucket{count: 1, expiresAt: now.Add(window)}
		return true
	}

	if bucket.count >= limit {
		return false
	}

	bucket.count++
	l.buckets[key] = bucket
	return true
}

func emailDeliveryConfigured() bool {
	return strings.TrimSpace(os.Getenv("PROBOD_SMTP_ADDR")) != "" &&
		strings.TrimSpace(os.Getenv("PROBOD_SMTP_USER")) != "" &&
		strings.TrimSpace(os.Getenv("PROBOD_SMTP_PASSWORD")) != "" &&
		strings.TrimSpace(os.Getenv("PROBOD_MAILER_SENDER_EMAIL")) != ""
}

func requestClientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		host = strings.TrimSpace(r.RemoteAddr)
	}

	peerIP := net.ParseIP(host)
	if peerIP == nil {
		return host
	}

	// Only honor proxy headers when the direct peer is a local/private reverse
	// proxy. Walk X-Forwarded-For from right to left so a client-supplied value
	// on the left cannot bypass the limiter when Traefik appends the real peer.
	if peerIP.IsLoopback() || peerIP.IsPrivate() {
		parts := strings.Split(r.Header.Get("X-Forwarded-For"), ",")
		for i := len(parts) - 1; i >= 0; i-- {
			candidate := net.ParseIP(strings.TrimSpace(parts[i]))
			if candidate == nil || candidate.IsLoopback() || candidate.IsPrivate() || candidate.IsUnspecified() {
				continue
			}
			return candidate.String()
		}

		if candidate := net.ParseIP(strings.TrimSpace(r.Header.Get("X-Real-IP"))); candidate != nil {
			return candidate.String()
		}
	}

	return peerIP.String()
}

func normalizedEmailFromVariables(value any) string {
	switch v := value.(type) {
	case map[string]any:
		if email, ok := v["email"].(string); ok {
			return strings.ToLower(strings.TrimSpace(email))
		}
		for _, nested := range v {
			if email := normalizedEmailFromVariables(nested); email != "" {
				return email
			}
		}
	case []any:
		for _, nested := range v {
			if email := normalizedEmailFromVariables(nested); email != "" {
				return email
			}
		}
	}
	return ""
}

func publicAuthRateLimitOperations(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	opCtx := graphql.GetOperationContext(ctx)
	if opCtx == nil || opCtx.Operation == nil {
		return next(ctx)
	}

	r := gqlutils.HTTPRequestFromContext(ctx)
	clientIP := requestClientIP(r)
	email := normalizedEmailFromVariables(opCtx.Variables)

	for _, selection := range opCtx.Operation.SelectionSet {
		field, ok := selection.(*ast.Field)
		if !ok {
			continue
		}

		policy, sensitive := publicAuthPolicies[field.Name]
		if !sensitive {
			continue
		}

		if (field.Name == "signUp" || field.Name == "forgotPassword" || field.Name == "resendVerificationEmail") && !emailDeliveryConfigured() {
			return publicAuthErrorResponse("Email authentication is temporarily unavailable.")
		}

		if !publicAuthLimiter.allow("gql:"+field.Name+":ip:"+clientIP, policy.ipLimit, policy.window) {
			gqlutils.HTTPResponseWriterFromContext(ctx).Header().Set("Retry-After", "60")
			return publicAuthErrorResponse("Too many authentication attempts. Please try again later.")
		}

		if email != "" && policy.emailLimit > 0 && !publicAuthLimiter.allow("gql:"+field.Name+":email:"+email, policy.emailLimit, policy.window) {
			gqlutils.HTTPResponseWriterFromContext(ctx).Header().Set("Retry-After", "60")
			return publicAuthErrorResponse("Too many authentication attempts. Please try again later.")
		}
	}

	return next(ctx)
}

func publicAuthErrorResponse(message string) graphql.ResponseHandler {
	return func(context.Context) *graphql.Response {
		return &graphql.Response{
			Errors: gqlerror.List{gqlerror.Errorf("%s", message)},
		}
	}
}

func magicLinkRateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !emailDeliveryConfigured() {
			http.Error(w, "Email authentication is temporarily unavailable.", http.StatusServiceUnavailable)
			return
		}

		const window = 15 * time.Minute
		clientIP := requestClientIP(r)
		if !publicAuthLimiter.allow("magic:ip:"+clientIP, 10, window) {
			w.Header().Set("Retry-After", "60")
			http.Error(w, "Too many authentication attempts. Please try again later.", http.StatusTooManyRequests)
			return
		}

		if err := r.ParseForm(); err == nil {
			email := strings.ToLower(strings.TrimSpace(r.FormValue("email")))
			if email != "" && !publicAuthLimiter.allow("magic:email:"+email, 3, window) {
				w.Header().Set("Retry-After", "60")
				http.Error(w, "Too many authentication attempts. Please try again later.", http.StatusTooManyRequests)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
