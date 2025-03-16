package uuid

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// NewUUIDCommand creates a new cobra command for generating UUIDs.
func NewUUIDCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uuid",
		Short: "Generate UUID",
		Long:  `Generate a new UUID with google/uuid package.`,
		Run: func(cmd *cobra.Command, args []string) {
			uuid := uuid.New()
			cmd.Println(uuid.String())
		},
	}

	return cmd
}
