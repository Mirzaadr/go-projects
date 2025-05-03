package utils

import "net/http"

func ChainMiddleware(
	h http.Handler,
	m ...func(http.Handler) http.Handler) http.Handler {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}
	return wrapped
}
