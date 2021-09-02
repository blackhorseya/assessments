package q1

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "20 then 26",
			args: args{n: 20},
			want: "26",
		},
		{
			name: "-7 then -10",
			args: args{n: -7},
			want: "-10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := q1(tt.args.n); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
