package core_http_response

import "net/http"

var (
	StatusCodeUnitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUnitialized,
	}
}

func (rw *ResponseWriter) WriteHeader(stutusCode int) {
	rw.ResponseWriter.WriteHeader(stutusCode)
	rw.statusCode = stutusCode
}

func (rw *ResponseWriter) GetStatusCode() int {
	if rw.statusCode == StatusCodeUnitialized {
		return http.StatusOK
	}

	return rw.statusCode
}
