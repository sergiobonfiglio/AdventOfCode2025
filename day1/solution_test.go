package main

import "testing"


func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name:  "example",
			input: `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
{
			name:  "example",
			input: `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
			want:  6,
		},
{
			name:  "example2",
			input: `R1000
R50
R1
L1
L1
R2
L1000`,
			want:  23,
		},
{
			name:  "example3",
			input: `L1000`,
			want:  10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
