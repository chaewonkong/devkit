package url

import (
	"net/url"

	"github.com/spf13/cobra"
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
		Run: func(cmd *cobra.Command, args []string) {
			if !encodeFlag && !decodeFlag {
				cmd.PrintErr("Please specify either --encode or --decode flag")
				return
			}
			if encodeFlag {
				encoded := url.QueryEscape(args[0])
				cmd.Println(encoded)
				return
			}
			if decodeFlag {
				decoded, err := url.QueryUnescape(args[0])
				if err != nil {
					cmd.PrintErr(err)
					return
				}
				cmd.Println(decoded)
				return
			}
		},
	}

	cmd.Flags().BoolVarP(&encodeFlag, "encode", "e", false, "Encode the input URL")
	cmd.Flags().BoolVarP(&decodeFlag, "decode", "d", false, "Decode the input URL")

	return cmd
}
