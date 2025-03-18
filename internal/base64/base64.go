package base64

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	errNoFlag = "Please specify either --encode or --decode flag"
)

// NewBase64Command creates a new cobra command for Base64 encoding and decoding.
func NewBase64Command() *cobra.Command {
	var encodeFlag bool
	var decodeFlag bool

	cmd := &cobra.Command{
		Use:   "base64",
		Short: "Base64 encoding and decoding",
		Long:  `Encode and decode strings using the base64 package.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !encodeFlag && !decodeFlag {
				return fmt.Errorf(errNoFlag)
			}
			if encodeFlag {
				base64Encoded := base64.StdEncoding.EncodeToString([]byte(args[0]))
				cmd.Println(base64Encoded)

				return nil
			}

			if decodeFlag {
				base64Decoded, err := base64.StdEncoding.DecodeString(args[0])
				if err != nil {
					return err
				}
				cmd.Println(string(base64Decoded))
				return nil
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&encodeFlag, "encode", "e", false, "Encode the input string")
	cmd.Flags().BoolVarP(&decodeFlag, "decode", "d", false, "Decode the input string")

	return cmd
}
