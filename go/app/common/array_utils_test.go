package common

import (
	"reflect"
	"testing"
)

func TestFillArrayWithValue(t *testing.T) {
	type args struct {
		array []int
		value int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "simple", args: args{array: []int{1, 2, 3}, value: 1}, want: []int{1, 1, 1}},
		{name: "made array", args: args{array: make([]int, 3), value: 2}, want: []int{2, 2, 2}},
		{name: "negative value", args: args{array: []int{0}, value: -3}, want: []int{-3}},
		{name: "empty array", args: args{array: []int{}, value: 4}, want: []int{}},
		{name: "empty array", args: args{array: nil, value: 5}, want: nil},
	}
	for _, tt := range tests {
		FillArrayWithValue(tt.args.array, tt.args.value)
		if got := tt.args.array; !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FillArrayWithValue() = %v, want %v", got, tt.want)
		}
	}
}
