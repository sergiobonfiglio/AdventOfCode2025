package main

import (
	"slices"
	"strconv"
	"strings"
)

func part1(input string) any {

	parts := strings.Split(input, "\n\n")
	rangesInput := parts[0]
	ingredientsInput := parts[1]

	ranges := []Range{}
	for _, line := range strings.Split(rangesInput, "\n") {
		rParts := strings.Split(line, "-")
		start, _ := strconv.ParseInt(rParts[0], 10, 64)
		end, _ := strconv.ParseInt(rParts[1], 10, 64)
		ranges = append(ranges, Range{
			Start: start,
			End:   end,
		})
	}
	slices.SortFunc(ranges, rangeCompare)

	freshCount := 0
	for _, line := range strings.Split(ingredientsInput, "\n") {

		id, _ := strconv.ParseInt(line, 10, 64)
		isFresh := false
		for _, r := range ranges {
			if r.Contains(id) {
				isFresh = true
				break
			}
			if r.Start > id {
				break
			}
		}
		if isFresh {
			freshCount++
		}

	}

	return freshCount
}

func part2(input string) any {
	parts := strings.Split(input, "\n\n")
	rangesInput := parts[0]

	rangeSet := RangeSet{}
	for _, line := range strings.Split(rangesInput, "\n") {
		rParts := strings.Split(line, "-")
		start, _ := strconv.ParseInt(rParts[0], 10, 64)
		end, _ := strconv.ParseInt(rParts[1], 10, 64)

		rangeSet.Add(Range{
			Start: start,
			End:   end,
		})
	}
	count := int64(0)
	for r := range rangeSet {
		rLen := r.End - r.Start + 1
		count += rLen
	}

	return count

}

type RangeSet map[Range]struct{}

func (rs RangeSet) Add(r Range) {

	overlap := &Range{Start: r.Start, End: r.End}
	overlapping := []*Range{}
	for existing := range rs {
		if existing.Overlaps(r) {
			overlap.Start = min(overlap.Start, existing.Start)
			overlap.End = max(overlap.End, existing.End)
			overlapping = append(overlapping, &existing)		
		}
	}

	for _, o := range overlapping {
		delete(rs, *o)
	}

	rs[*overlap] = struct{}{}
}

type Range struct {
	Start int64
	End   int64
}

func (r Range) Contains(n int64) bool {
	return n >= r.Start && n <= r.End
}

func (r Range) Overlaps(other Range) bool {
	return r.Start <= other.End && other.Start <= r.End
}

func rangeCompare(a, b Range) int {
	if a.Start < b.Start {
		return -1
	} else if a.Start > b.Start {
		return 1
	}
	return 0
}
