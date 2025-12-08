package main

import (
	"AdventOfCode2025/utils"
	"slices"
	"strings"
)

func part1(input string) any {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		nums := utils.ToIntArray(line, "")
		
		first, ix := maxDigitWithIndex(nums[:len(nums)-1])
		second := slices.Max(nums[ix+1:])

		sum += (first * 10) + second
	}

	return sum
}

func part2(input string) any {
	sum := int64(0)
	for _, line := range strings.Split(input, "\n") {
		nums := utils.ToInt64Array(line, "")
		
		maxDigits := 12

		left := 0
		right := len(nums)-maxDigits

		lineSum := int64(0)
		for i := 0; i < maxDigits; i++ {
			curr, ix := maxDigitWithIndex(nums[left:right])
			lineSum = lineSum*10 + curr
			left = left + ix + 1
			right = len(nums)-(maxDigits-(i+1))+1
		}
		sum += lineSum

	}

	return sum
}




func maxDigitWithIndex[T int | int64](nums []T) (T, int) {
	max := nums[0]
	index := 0
	for i, n := range nums {
		if n > max {
			max = n
			index = i
		}
		if max == 9 {
			break
		}
	}
	return max, index
}
