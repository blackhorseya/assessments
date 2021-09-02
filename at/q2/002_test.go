package q2

import (
	"reflect"
	"testing"
)

func Test_q2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1, 2, 3, 4] 4 then [1, 2]",
			args: args{nums: []int{1, 2, 3, 4}, target: 4},
			want: []int{1, 2},
		},
		{
			name: "[1,2,3] 3 then []",
			args: args{nums: []int{1, 2, 3}, target: 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := q2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("q2() = %v, want %v", got, tt.want)
			}
		})
	}
}
