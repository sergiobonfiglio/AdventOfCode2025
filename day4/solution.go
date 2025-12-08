package main

import (
	"AdventOfCode2025/utils"
)

func part1(input string) any {

	matrix := utils.NewMatrixFromLines(input)

	accessible := 0
	for val, ok := matrix.Next(); ok; val, ok = matrix.Next() {

		if val != '@' {
			continue
		}

		curr := matrix.CurrCell()
		neighbors := curr.Neighbors()

		rolls := 0
		for _, ngb := range neighbors {
			if !matrix.IsIn(ngb.R, ngb.C) {
				continue
			}
			cellVal := matrix.GetAtCell(&ngb)
			if cellVal != nil && *cellVal == '@' {
				rolls++
			}
		}

		if rolls < 4 {
			accessible++
		}

	}

	return accessible
}

func part2(input string) any {
	matrix := utils.NewMatrixFromLines(input)

	hasChanges := true
	removed := 0

	for hasChanges {
		removable := []*utils.Cell{}
		for val, ok := matrix.Next(); ok; val, ok = matrix.Next() {

			if val != '@' {
				continue
			}

			curr := matrix.CurrCell()
			neighbors := curr.Neighbors()

			rolls := 0
			for _, ngb := range neighbors {
				if !matrix.IsIn(ngb.R, ngb.C) {
					continue
				}
				cellVal := matrix.GetAtCell(&ngb)
				if cellVal != nil && *cellVal == '@' {
					rolls++
				}
			}

			if rolls < 4 {
				removable = append(removable, curr)
			}
		}
		matrix.Reset()
		
		hasChanges = len(removable) > 0
		removed += len(removable)

		// remove rolls
		for _, cell := range removable {
			matrix.SetValAtCell(cell, '.')
		}

	}

	return removed
}
