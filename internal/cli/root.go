package cli

import (
	"github.com/ProvisionLabs/provisionhub/internal/application"
	"github.com/spf13/cobra"
)

// NewRootCommand constructs the command tree without starting an interactive session.
func NewRootCommand() *cobra.Command {
	info := application.Info()
	root := &cobra.Command{
		Use:           "provisionhub",
		Short:         info.Tagline,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.AddCommand(newVersionCommand())
	return root
}
