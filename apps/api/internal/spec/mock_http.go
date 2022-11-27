package spec

import "net/http"

type MockHttpResponseWriter struct {
	StatusCode int
	Text       string
}

func NewMockHttpResponseWriter() *MockHttpResponseWriter {
	return &MockHttpResponseWriter{}
}

func (m *MockHttpResponseWriter) Header() http.Header {
	return make(map[string][]string)
}

func (m *MockHttpResponseWriter) Write(bytes []byte) (int, error) {
	m.Text = string(bytes)

	return len(bytes), nil
}

func (m *MockHttpResponseWriter) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}
