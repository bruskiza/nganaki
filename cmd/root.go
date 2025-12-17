package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "nganaki",
		Short: "nganaki is a CLI tool",
		Long: `nganaki is a CLI tool for managing your application.
It provides various commands to interact with the system.`,
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)


	return rootCmd, nil

}


func Execute() int {
	c, err := NewRootCmd()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if err := c.Execute(); err != nil {
		_, _ = fmt.Fprintln(c.ErrOrStderr(), err)
		return 1
	}
	return 0
	
}