package q2

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,-2,-3,5] then 1",
			args: args{A: []int{1, -2, -3, 5}},
			want: 1,
		},
		{
			name: "[1,2,3,-5] then -1",
			args: args{A: []int{1, 2, 3, -5}},
			want: -1,
		},
		{
			name: "[1,2,0,-5] then 0",
			args: args{A: []int{1, 2, 0, -5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
