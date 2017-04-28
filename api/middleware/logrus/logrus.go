package logrus

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"time"
)

// Logger returns a middleware handler that logs the request as it goes in and the response as it goes out.
func Logger() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
		start := time.Now()

		addr := req.Header.Get("X-Real-IP")
		if addr == "" {
			addr = req.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = req.RemoteAddr
			}
		}

		log.Printf("Started %s %s for %s", req.Method, req.URL.Path, addr)

		rw := res.(martini.ResponseWriter)
		c.Next()

		log.Printf("hahahhahahah %v %s in %v\n", rw.Status(), http.StatusText(rw.Status()), time.Since(start))
	}
}
