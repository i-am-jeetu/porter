package main

import (
	"github.com/deislabs/porter/pkg/mixin/exec"
	"github.com/spf13/cobra"
)

func buildInstallCommand(e *exec.Exec) *cobra.Command {
	var opts struct {
		file string
	}
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Execute the install functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return e.Install(opts.file)
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&opts.file, "file", "f", "", "Path to the script to execute")
	return cmd
}