package url_test

import (
	"testing"

	"github.com/chaewonkong/devkit/internal/test"
	. "github.com/chaewonkong/devkit/internal/url"
	"github.com/stretchr/testify/assert"
)

func TestURLCommand(t *testing.T) {
	for _, tt := range []struct {
		name     string
		given    string
		expected string
		encode   bool
	}{
		{
			name:     "Test URL encode command",
			given:    "example.com?name=devkt&age=25",
			expected: "example.com%3Fname%3Ddevkt%26age%3D25",
			encode:   true,
		},
		{
			name:     "Test URL decode command",
			expected: "example.com?name=devkt&age=25",
			given:    "example.com%3Fname%3Ddevkt%26age%3D25",
			encode:   false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewURLCommand()
			if tt.encode {
				cmd.SetArgs([]string{"--encode", tt.given})
			} else {
				cmd.SetArgs([]string{"--decode", tt.given})
			}

			mw := test.MockWriter{}
			cmd.SetOut(&mw)

			err := cmd.Execute()
			assert.NoError(t, err)

			assert.Equal(t, tt.expected, mw.String())
		})
	}
}
