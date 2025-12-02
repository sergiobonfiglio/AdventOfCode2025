package main

import (
	"strconv"
	"strings"
)

type Range struct {
	Start int64
	End   int64
}

func newRange(s string) Range {
	parts := strings.Split(s, "-")
	var r Range
	var err error
	r.Start, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return r
	}
	r.End, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return r
	}
	return r
}

func part1(input string) any {

	rangesStr := strings.Split(input, ",")
	var ranges []Range
	for _, rStr := range rangesStr {
		ranges = append(ranges, newRange(rStr))
	}

	invalidSum := int64(0)
	for _, rr := range ranges {
		for i := rr.Start; i <= rr.End; i++ {

			s := strconv.FormatInt(i, 10)
			isEvenDigits := len(s)%2 == 0
			if !isEvenDigits {
				continue
			}

			if s[0:len(s)/2] == s[len(s)/2:] {
				invalidSum += i
			}

		}
	}

	return invalidSum
}

func part2(input string) any {
	rangesStr := strings.Split(input, ",")
	var ranges []Range
	for _, rStr := range rangesStr {
		ranges = append(ranges, newRange(rStr))
	}

	invalidSum := int64(0)
	for _, rr := range ranges {
		for i := rr.Start; i <= rr.End; i++ {
			s := strconv.FormatInt(i, 10)

			if hasRepeated(s) {
				invalidSum += i
			}

		}
	}

	return invalidSum
}

func hasRepeated(s string) bool {

	if len(s) == 2 {
		return s[0] == s[1]
	}

	for rep := 2; rep <= len(s); rep++ {

		chunkLen := len(s) / rep

		if chunkLen*rep != len(s) {
			//not divisible
			continue
		}

		prevChunk := s[0:chunkLen]
		isRepeated := true
		for chunk := 1; chunk < rep; chunk++ {
			currentChunk := s[chunk*chunkLen : (chunk+1)*chunkLen]
			if currentChunk != prevChunk {
				//not repeated
				isRepeated = false
				break
			}
		}
		if isRepeated {
			return true
		}

	}

	return false
}
