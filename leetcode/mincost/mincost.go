package mincost

func minCostClimbingStairs(cost []int) int {
	fromOneStepBack := 0
	fromTwoStepBack := 0

	for i := 2; i <= len(cost); i++ {
		curr := min(fromOneStepBack+cost[i-1], fromTwoStepBack+cost[i-2])
		fromTwoStepBack, fromOneStepBack = fromOneStepBack, curr
	}
	return fromOneStepBack
}
