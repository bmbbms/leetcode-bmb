package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumber(t *testing.T) {
	type parms struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args parms
		want []int
	}{
		{
			name: "test1",
			args: parms{
				[]int{1, 3, 5, 7},
				4,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumber(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
