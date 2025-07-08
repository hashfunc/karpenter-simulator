package command

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "karpenter-simulator",
		Short: "A Karpenter simulation tool",
		Long:  "A Karpenter simulation tool",
	}

	rootCmd.AddCommand(NewBudgetCommand())

	return rootCmd
}
