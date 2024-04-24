package wrapper

import "net/http"

type HandlerWrapper func(http.ResponseWriter, *http.Request)

// ServeHTTP calls f(w, r).
func (f HandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
