package cli

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/marccampbell/yaml-toolbox/pkg/splitter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func SplitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "split",
		Short:         "",
		Long:          ``,
		SilenceUsage:  true,
		SilenceErrors: false,
		Args:          cobra.MinimumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			inputFile := args[0]

			v := viper.GetViper()
			outputDir := v.GetString("out")
			if _, err := os.Stat(outputDir); os.IsNotExist(err) {
				err := os.MkdirAll(outputDir, 0755)
				if err != nil {
					return err
				}
			}

			fi, err := os.Stat(inputFile)
			if err != nil {
				return err
			}
			if fi.IsDir() {
				err := filepath.Walk(inputFile,
					func(path string, info os.FileInfo, err error) error {
						if err != nil {
							return err
						}

						if info.IsDir() {
							return nil
						}

						if err := splitFile(path, outputDir); err != nil {
							return err
						}

						return nil
					})
				if err != nil {
					return err
				}
			} else {
				if err := splitFile(inputFile, outputDir); err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().String("out", "", "the output directory to store the split files in")

	return cmd
}

func splitFile(filename string, outputDir string) error {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	splitDocs, err := splitter.SplitYAML(input)
	if err != nil {
		return err
	}

	for filename, content := range splitDocs {
		err = ioutil.WriteFile(filepath.Join(outputDir, filename), content, 0644)
		if err != nil {
			return err
		}

	}

	return nil
}
