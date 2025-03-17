package unix

import (
	"time"

	"github.com/spf13/cobra"
)

// NewUnixTimeCommand creates a new cobra command for unix timestamp generation.
func NewUnixTimeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unix",
		Short: "Generate Unix timestamp",
		Long:  "Epoch Time(Unix Time) generater in s, ms, μs, ns",
		Run: func(cmd *cobra.Command, args []string) {
			if nano, _ := cmd.Flags().GetBool("nano"); nano {
				cmd.Println(time.Now().UnixNano())
				return
			}

			if micro, _ := cmd.Flags().GetBool("micro"); micro {
				cmd.Println(time.Now().UnixMicro())
				return
			}

			if milli, _ := cmd.Flags().GetBool("milli"); milli {
				cmd.Println(time.Now().UnixMilli())

				return
			}

			cmd.Println(time.Now().Unix())
			return
		},
	}

	cmd.Flags().Bool("nano", false, "Print Unix timestamp in nanoseconds")
	cmd.Flags().Bool("milli", false, "Print Unix timestamp in milliseconds")
	cmd.Flags().Bool("micro", false, "Print Unix timestamp in microseconds")

	return cmd
}
