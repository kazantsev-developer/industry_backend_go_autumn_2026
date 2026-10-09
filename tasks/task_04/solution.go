package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}

	firstDiff := nums[1] - nums[0]

	count := 1
	sum := firstDiff
	minDiff := firstDiff
	maxDiff := firstDiff

	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		count++
		sum += diff
		if diff < minDiff {
			minDiff = diff
		}
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return Stats{
		Count: count,
		Sum:   sum,
		Min:   minDiff,
		Max:   maxDiff,
	}
}
