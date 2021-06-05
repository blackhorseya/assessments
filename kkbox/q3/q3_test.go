package q3

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "baaabbabbb then 7",
			args: args{S: "baaabbabbb"},
			want: 7,
		},
		{
			name: "babba then 5",
			args: args{S: "babba"},
			want: 5,
		},
		{
			name: "abaaaa then 4",
			args: args{S: "abaaaa"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
