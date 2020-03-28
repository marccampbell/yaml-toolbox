package cli

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/marccampbell/yaml-toolbox/pkg/remarshaler"
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
			path := args[0]

			fi, err := os.Stat(path)
			if err != nil {
				return err
			}
			if fi.IsDir() {
				err := filepath.Walk(path,
					func(path string, info os.FileInfo, err error) error {
						if err != nil {
							return err
						}

						if info.IsDir() {
							return nil
						}

						if err := remarshalFileInPlace(path); err != nil {
							return err
						}

						return nil
					})
				if err != nil {
					return err
				}
			} else {
				if err := remarshalFileInPlace(path); err != nil {
					return err
				}
			}

			return nil
		},
	}

	return cmd
}

func remarshalFileInPlace(filename string) error {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	remarshaled, err := remarshaler.RemarshalYAML(input)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, remarshaled, 0644)
	if err != nil {
		return err
	}

	return nil
}
