package base64_test

import (
	"testing"

	. "github.com/chaewonkong/devkit/internal/base64"
	"github.com/chaewonkong/devkit/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestURLCommand(t *testing.T) {
	for _, tt := range []struct {
		name    string
		decoded string
		encoded string
		encode  bool
	}{
		{
			name:    "Test Base64 encode command",
			encoded: "bGlvbmhlYXJ0QGxpb25oZWFydC5odA==",
			decoded: "lionheart@lionheart.ht",
			encode:  true,
		},
		{
			name:    "Test Base64 encode command",
			encoded: "bGlvbmhlYXJ0QGxpb25oZWFydC5odA==",
			decoded: "lionheart@lionheart.ht",
			encode:  false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewBase64Command()
			var expected string
			if tt.encode {
				cmd.SetArgs([]string{"--encode", tt.decoded})
				expected = tt.encoded
			} else {
				cmd.SetArgs([]string{"--decode", tt.encoded})
				expected = tt.decoded
			}

			mw := test.MockWriter{}
			cmd.SetOut(&mw)

			err := cmd.Execute()
			assert.NoError(t, err)

			assert.Equal(t, expected, mw.String())
		})
	}
}
