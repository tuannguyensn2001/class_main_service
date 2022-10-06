package cmd

import "github.com/spf13/cobra"

func GetRoot() *cobra.Command {
	root := &cobra.Command{}

	root.AddCommand(server())
	root.AddCommand(migrateUp())
	root.AddCommand(migrateDown())

	return root
}
