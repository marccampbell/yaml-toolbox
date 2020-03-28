package cli

import (
	"io/ioutil"

	"github.com/marccampbell/yaml-remarshaler/pkg/remarshaler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RemarshalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "remarshal",
		Short:         "",
		Long:          ``,
		SilenceUsage:  true,
		SilenceErrors: false,
		Args:          cobra.MinimumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// this does an inplace remarshal only right now
			input, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			remarshaled, err := remarshaler.FixUpYAML(input)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(args[0], remarshaled, 0644)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
