package main

import (
	"testing"
)

func Test_findFirstVal(t *testing.T) {
	type args struct {
		s   []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not found",
			args: args{
				s:   []int{3, 4, 5},
				val: 2,
			},
			want: -1,
		},
		{
			name: "t1",
			args: args{
				s:   []int{1, 1, 1, 1, 1, 1, 1, 1},
				val: 1,
			},
			want: 0,
		},
		{
			name: "t2",
			args: args{
				s:   []int{1, 1, 2, 2, 2, 3, 3, 3},
				val: 2,
			},
			want: 2,
		},
		{
			name: "t3",
			args: args{
				s:   []int{1, 2, 3, 4, 5, 6, 7, 8},
				val: 5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstVal(tt.args.s, tt.args.val); got != tt.want {
				t.Errorf("findFirstVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLastVal(t *testing.T) {
	type args struct {
		s   []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not found",
			args: args{
				s:   []int{4, 5, 6},
				val: 7,
			},
			want: -1,
		},
		{
			name: "t1",
			args: args{
				s:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				val: 1,
			},
			want: 9,
		},
		{
			name: "t2",
			args: args{
				s:   []int{1, 1, 2, 2, 2, 3, 3, 3, 3, 3},
				val: 2,
			},
			want: 4,
		},
		{
			name: "t3",
			args: args{
				s:   []int{1, 1, 2, 2, 2, 5, 6, 7, 8, 9},
				val: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLastVal(tt.args.s, tt.args.val); got != tt.want {
				t.Errorf("findLastVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFirstGtVal(t *testing.T) {
	type args struct {
		s   []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not found",
			args: args{
				s:   []int{4, 5, 6},
				val: 7,
			},
			want: -1,
		},
		{
			name: "t1",
			args: args{
				s:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
				val: 5,
			},
			want: 5,
		},
		{
			name: "t2",
			args: args{
				s:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
				val: 1,
			},
			want: 1,
		},
		{
			name: "t3",
			args: args{
				s:   []int{10, 20, 30, 40, 50, 60, 70},
				val: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstGtVal(tt.args.s, tt.args.val); got != tt.want {
				t.Errorf("findFirstGtVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLastLtVal(t *testing.T) {
	type args struct {
		s   []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not found",
			args: args{
				s:   []int{4, 5, 6},
				val: 3,
			},
			want: -1,
		},
		{
			name: "t1",
			args: args{
				s:   []int{3, 3, 3},
				val: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLastLtVal(tt.args.s, tt.args.val); got != tt.want {
				t.Errorf("findLastLtVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
