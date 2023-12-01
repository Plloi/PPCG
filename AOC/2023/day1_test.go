package main

import (
	"strings"
	"testing"
)

func Test_day1Part1(t *testing.T) {
	type args struct {
		calibrate []string
	}

	var inputDay1 []string
	inputDay1, _ = readLines("day1.txt")

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				calibrate: strings.Split("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet", "\n"),
			},
			want: 142,
		},
		{
			name: "RealDeal",
			args: args{
				calibrate: inputDay1,
			},
			want: 55029,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1Part1(tt.args.calibrate); got != tt.want {
				t.Errorf("day1Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day1Part2(t *testing.T) {
	type args struct {
		calibrate []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				calibrate: strings.Split("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen", "\n"),
			},
			want: 281,
		},
		{
			name: "Test 2",
			args: args{
				calibrate: strings.Split("3oneeighttwo", "\n"),
			},
			want: 32,
		},
		{
			name: "Test 3",
			args: args{
				calibrate: strings.Split("twofourfive485", "\n"),
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1Part2(tt.args.calibrate); got != tt.want {
				t.Errorf("day1Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
