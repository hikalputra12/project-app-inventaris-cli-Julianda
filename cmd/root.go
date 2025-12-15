package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "InventorySystem-CLI",
		Short: "A CLI for manage an inventory items",
		Long:  `this CLI application serves as a robust terminal-based tool for comprehensive asset and inventory life cycle management, linking operational stock control with necessary financial reporting.`,
	}
)

// Execute excutes the root command
func Execute() error {
	return rootCmd.Execute()
}
