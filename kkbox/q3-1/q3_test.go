package q3_1

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		riddle string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "?b?a then abca",
			args: args{riddle: "?b?a"},
			want: "abca",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.riddle); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
