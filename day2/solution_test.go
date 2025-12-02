package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "example",
			input: `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`,
			want:  1227775554,
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
			name:  "example0",
			input: `11-12`,
			want:  11,
		},
		{
			name:  "example1",
			input: `1010-1010`,
			want:  1010,
		},
		{
			name:  "example2",
			input: `111-111`,
			want:  111,
		},
		{
			name:  "example2.1",
			input: `110-110`,
			want:  0,
		},
		{
			name:  "example3",
			input: `565656-565656`,
			want:  565656,
		},
		{
			name:  "example4",
			input: `1188511885-1188511885`,
			want:  1188511885,
		},
		{
			name:  "example",
			input: `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`,
			want:  4174379265,
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
