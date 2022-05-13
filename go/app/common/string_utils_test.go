package common

import (
	"reflect"
	"testing"
)

func TestRoundStrings(t *testing.T) {
	type args struct {
		strings   []string
		maxLength []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "no rounding",
			args: args{strings: []string{"Lorem", "Lorem ipsum dolor"}, maxLength: []int{5, 17}},
			want: []string{"Lorem", "Lorem ipsum dolor"},
		},
		{
			name: "little rounding",
			args: args{strings: []string{"Lorem ipsum", "Lorem ipsum dolor"}, maxLength: []int{4, 13}},
			want: []string{"L...", "Lorem ipsu..."},
		},
		{
			name: "Max rounding",
			args: args{strings: []string{"Lorem", "Lorem", "Lorem", "Lorem", "Lorem"}, maxLength: []int{4, 3, 2, 1, 0}},
			want: []string{"L...", "...", "..", ".", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundStrings(tt.args.strings, tt.args.maxLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoundStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
