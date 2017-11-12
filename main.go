package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use: "hello",
	}

	sayCmd := &cobra.Command{
		Use: "say NAME",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("NAME is required.")
			}

			fmt.Printf("Hello %s!\n", args[0])

			return nil
		},
	}

	completionCmd := &cobra.Command{
		Use: "completion",
		RunE: func(cmd *cobra.Command, args []string) error {
			return root.GenBashCompletion(os.Stdout)
		},
	}

	root.AddCommand(sayCmd, completionCmd)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
