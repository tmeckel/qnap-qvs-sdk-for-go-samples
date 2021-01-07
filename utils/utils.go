package utils

import "net/http"

// SetHeader set a header for a HTTP request
func SetHeader(r *http.Request, key, value string) {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(key, value)
}
