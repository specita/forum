// The requestid package is a partini middleware that will parse
// the X-Request-Id header from a request and make it injectible.
package requestid

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/pborman/uuid"
)

const header = "X-Request-Id"

// Our custom type to map the context to.
type RequestId string

// Options for the middleware.
type Options struct {
	Generate bool
}

// A middleware to extract the request id.
func Middleware(options *Options) martini.Handler {
	return func(c martini.Context, req *http.Request, res http.ResponseWriter) {
		requestId := req.Header.Get(header)

		if requestId == "" && options.Generate {
			uuid := generateRequestId()
			req.Header.Set(header, uuid)
			requestId = uuid
		}

		c.Map(RequestId(requestId))
		c.Next()
		res.Header().Set(header, requestId)
	}
}

func generateRequestId() string {
	return uuid.NewUUID().String()
}
