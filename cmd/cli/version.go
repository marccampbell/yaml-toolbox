package cli

import (
	"fmt"

	"github.com/marccampbell/yaml-toolbox/pkg/version"
	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the current version and exit",
		Long:  `Print the current version and exit`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%s\n", version.Version())

			return nil
		},
	}
	return cmd
}
