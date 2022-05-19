package common

import (
	"reflect"
	"strconv"
	"testing"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

var maxIntStr = strconv.FormatUint(uint64(MaxInt), 10)

func TestParseLongDuration(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *LongDuration
		wantErr bool
	}{
		{name: "days", args: args{s: "3d"}, want: &LongDuration{Days: 3}, wantErr: false},
		{name: "weeks", args: args{s: "4w"}, want: &LongDuration{Days: 28}, wantErr: false},
		{name: "months", args: args{s: "5m"}, want: &LongDuration{Months: 5}, wantErr: false},
		{name: "years", args: args{s: "6y"}, want: &LongDuration{Years: 6}, wantErr: false},
		{name: "pointed value", args: args{s: "1.1y"}, want: nil, wantErr: true},
		{name: "negative value", args: args{s: "-1y"}, want: nil, wantErr: true},
		{name: "no number d", args: args{s: "d"}, want: nil, wantErr: true},
		{name: "no number unknown letter", args: args{s: "qw"}, want: nil, wantErr: true},
		{name: "unknown unit", args: args{s: "3qw"}, want: nil, wantErr: true},
		{name: "empty string", args: args{s: ""}, want: nil, wantErr: true},
		{name: "number only", args: args{s: "12"}, want: nil, wantErr: true},
		{name: "all units", args: args{s: "1d2w3m4y"}, want: &LongDuration{
			Days:   15,
			Months: 3,
			Years:  4,
		}, wantErr: false},
		{name: "unordered", args: args{s: "3m2w4y1d"}, want: &LongDuration{
			Days:   15,
			Months: 3,
			Years:  4,
		}, wantErr: false},
		{name: "unknown unit", args: args{s: "3m2w4y1d5qw"}, want: nil, wantErr: true},
		{name: "unknown unit", args: args{s: "3m2w4y5qw1d"}, want: nil, wantErr: true},
		{name: "unknown unit", args: args{s: "3m2w4yqw1d"}, want: nil, wantErr: true},
		{name: "zero case", args: args{s: "0"}, want: &LongDuration{}, wantErr: false},
		{name: "overflow", args: args{s: strconv.FormatUint(uint64(MaxInt)+1, 10) + "d"}, want: nil, wantErr: true},
		{name: "overflow", args: args{s: maxIntStr + "d"}, want: &LongDuration{Days: MaxInt}, wantErr: false},
		//{name: "overflow", args: args{s: maxIntStr + "w"}, want: &LongDuration{Days: MaxInt * 7}, wantErr: false}, // todo fix max weeks
		{name: "overflow", args: args{s: maxIntStr + "m"}, want: &LongDuration{Months: MaxInt}, wantErr: false},
		{name: "overflow", args: args{s: maxIntStr + "y"}, want: &LongDuration{Years: MaxInt}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLongDuration(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLongDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLongDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
