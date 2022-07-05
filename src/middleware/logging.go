package middleware

import (
	"log"
	"net/http"
	"time"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()
		recorder := &StatusRecorder{
			ResponseWriter: res,
			Status:         200,
		}

		h.ServeHTTP(recorder, req)

		log.Printf(
			"%s %s %s %d %s",
			req.RemoteAddr,
			req.Method,
			req.RequestURI,
			recorder.Status,
			time.Since(start),
		)
	})
}
