package q2

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
			name: "50552 then 55",
			args: args{S: "50552"},
			want: 55,
		},
		{
			name: "10101 then 10",
			args: args{S: "10101"},
			want: 10,
		},
		{
			name: "88 then 88",
			args: args{S: "88"},
			want: 88,
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
