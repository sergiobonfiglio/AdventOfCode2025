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
			input: `3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
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
		want  int64
	}{
		{
			name:  "example",
			input: `3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
			want:  14,
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
