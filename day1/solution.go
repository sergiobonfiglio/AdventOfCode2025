package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {

	curr := 50
	zeroes := 0

	for _, line := range strings.Split(input, "\n") {
		dir := line[0]
		numR := line[1:]

		num, _ := strconv.Atoi(numR)
		num = num % 100

		switch dir {
		case 'L':
			curr = (curr - num)
			if curr < 0 {
				curr = 100 + curr
			}
		case 'R':
			curr = (curr + num) % 100
		}

		if curr == 0 {
			zeroes++
		}

	}

	return zeroes
}

func part2(input string) any {

	curr := 50
	zeroes := 0

	for _, line := range strings.Split(input, "\n") {
		dir := line[0]
		numR := line[1:]

		num, _ := strconv.Atoi(numR)
		turns := num / 100
		zeroes += turns

		num = num % 100


		prev := curr
		switch dir {
		case 'L':
			curr = (curr - num)
			if curr < 0 {
				if prev != 0 {
					zeroes++
				}
				curr = 100 + curr
			}
		case 'R':
			diff := (curr + num)
			curr = diff % 100

			if curr != 0 && prev != 0 && curr < prev {
				zeroes++
			}
		}

		if curr == 0 {
			zeroes++
		}
	}

	return zeroes
}
