package test

import (
	"io"
	"strings"
)

var _ io.Writer = &MockWriter{}

// MockWriter is a mock implementation of io.Writer.
type MockWriter struct {
	output string
}

// Write implements io.Writer.
func (m *MockWriter) Write(p []byte) (n int, err error) {
	m.output = string(p)
	return len(p), nil
}

// String returns the output written to the mock writer.
func (m *MockWriter) String() string {
	return strings.TrimSpace(m.output)
}
