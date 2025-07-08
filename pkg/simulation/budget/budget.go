package budget

import (
	"time"

	"k8s.io/utils/clock"
	clocktesting "k8s.io/utils/clock/testing"
	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"
)

type SimulationResult map[karpenterv1.DisruptionReason][]int

func Simulate(nodePool *karpenterv1.NodePool, numNodes int) (SimulationResult, error) {
	reasons := []karpenterv1.DisruptionReason{
		karpenterv1.DisruptionReasonEmpty,
		karpenterv1.DisruptionReasonDrifted,
		karpenterv1.DisruptionReasonUnderutilized,
	}

	var clocks []clock.Clock
	for hour := 0; hour < 24; hour++ {
		clocks = append(clocks, clocktesting.NewFakeClock(
			time.Date(2000, time.January, 1, hour, 0, 0, 0, time.UTC),
		))
		clocks = append(clocks, clocktesting.NewFakeClock(
			time.Date(2000, time.January, 1, hour, 30, 0, 0, time.UTC),
		))
	}

	simulationResult := SimulationResult{
		karpenterv1.DisruptionReasonEmpty: []int{},
		karpenterv1.DisruptionReasonDrifted: []int{},
		karpenterv1.DisruptionReasonUnderutilized: []int{},
	}
	for _, reason := range reasons {
		for _, c := range clocks {
			result, err := nodePool.GetAllowedDisruptionsByReason(c, numNodes, reason)
			if err != nil {
				return nil, err
			}
			simulationResult[reason] = append(simulationResult[reason], result)
		}
	}

	return simulationResult, nil
}
