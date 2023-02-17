package upstream

import (
	"fmt"
	"net/http"

	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/middleware"
	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/logger"
)

const defaultStaticResponseCode = 200

// newStaticResponseHandler creates a new staticResponseHandler that serves a
// a static response code.
func newStaticResponseHandler(upstream string, code *int) http.Handler {
	return &staticResponseHandler{
		code:     derefStaticCode(code),
		upstream: upstream,
	}
}

// staticResponseHandler responds with a static response with the given response code.
type staticResponseHandler struct {
	code     int
	upstream string
}

// ServeHTTP serves a static response.
func (s *staticResponseHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	scope := middleware.GetRequestScope(req)
	// If scope is nil, this will panic.
	// A scope should always be injected before this handler is called.
	scope.Upstream = s.upstream

	rw.WriteHeader(s.code)
	_, err := fmt.Fprintf(rw, "Authenticated")
	if err != nil {
		logger.Errorf("Error writing static response: %v", err)
	}
}

// derefStaticCode returns the derefenced value, or the default if the value is nil
func derefStaticCode(code *int) int {
	if code != nil {
		return *code
	}
	return defaultStaticResponseCode
}
