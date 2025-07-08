package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
	"sigs.k8s.io/yaml"

	"github.com/hashfunc/karpenter-simulator/pkg/renderer"
	"github.com/hashfunc/karpenter-simulator/pkg/simulation/budget"
)

var (
	nodeCount int
)

func NewBudgetCommand() *cobra.Command {
	budgetCmd := &cobra.Command{
		Use:   "budget [yaml-file-path]",
		Short: "Simulate budget disruptions for a NodePool",
		Long:  "Simulate budget disruptions for a NodePool configuration from a YAML file",
		Args:  cobra.ExactArgs(1),
		RunE:  runBudgetSimulation,
	}

	budgetCmd.Flags().IntVarP(&nodeCount, "nodes", "n", 100, "Number of nodes to simulate")

	return budgetCmd
}

func runBudgetSimulation(cmd *cobra.Command, args []string) error {
	yamlFilePath := args[0]

	nodePool, err := loadNodePoolFromYAML(yamlFilePath)
	if err != nil {
		return fmt.Errorf("failed to load YAML file: %w", err)
	}

	result, err := budget.Simulate(nodePool, nodeCount)
	if err != nil {
		return fmt.Errorf("failed to run simulation: %w", err)
	}

	fmt.Printf("\nSimulation Result for %d nodes:\n", nodeCount)
	renderer.PrintBudgetTables(result)

	return nil
}

func loadNodePoolFromYAML(filePath string) (*karpenterv1.NodePool, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var nodePool karpenterv1.NodePool
	if err := yaml.Unmarshal(data, &nodePool); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &nodePool, nil
}
