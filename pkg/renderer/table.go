package renderer

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	"github.com/hashfunc/karpenter-simulator/pkg/simulation/budget"
)

func PrintBudgetTables(result budget.SimulationResult) {
	emptyValues := result[karpenterv1.DisruptionReasonEmpty]
	driftedValues := result[karpenterv1.DisruptionReasonDrifted]
	underutilizedValues := result[karpenterv1.DisruptionReasonUnderutilized]

	var rows [][]string
	for i := 0; i < 48; i++ {
		hour := i / 2
		minute := (i % 2) * 30
		timeStr := fmt.Sprintf("%02d:%02d", hour, minute)

		row := []string{
			timeStr,
			strconv.Itoa(emptyValues[i]),
			strconv.Itoa(driftedValues[i]),
			strconv.Itoa(underutilizedValues[i]),
		}
		rows = append(rows, row)
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		Headers("Time", "Empty", "Drifted", "Underutilized").
		Rows(rows...)

	fmt.Println(t)
}
