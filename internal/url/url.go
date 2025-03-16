package url

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

const (
	// ErrFlag is the error message when no flag is specified.
	ErrFlag = "please specify either --encode or --decode flag"

	// ErrDecode is the error message when decoding fails.
	ErrDecode = "failed to decode the input URL"

	// ErrNoArg is the error message when there's no arg provided
	ErrNoArg = "please provide url"
)

// NewURLCommand creates a new cobra command for URL encoding and decoding.
func NewURLCommand() *cobra.Command {
	var encodeFlag bool
	var decodeFlag bool

	cmd := &cobra.Command{
		Use:   "url",
		Short: "URL encoding and decoding",
		Long:  `Encode and decode URLs using the net/url package.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !encodeFlag && !decodeFlag || encodeFlag && decodeFlag {
				return fmt.Errorf(ErrFlag)
			}
			if encodeFlag {
				encoded := url.QueryEscape(args[0])
				cmd.Println(encoded)
				return nil
			}
			if decodeFlag {
				decoded, err := url.QueryUnescape(args[0])
				if err != nil {
					return fmt.Errorf("%s: %w", ErrDecode, err)
				}
				cmd.Println(decoded)
				return nil
			}
			return fmt.Errorf(ErrNoArg)
		},
	}

	cmd.Flags().BoolVarP(&encodeFlag, "encode", "e", false, "Encode the input URL")
	cmd.Flags().BoolVarP(&decodeFlag, "decode", "d", false, "Decode the input URL")

	return cmd
}
