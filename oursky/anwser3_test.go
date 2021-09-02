package oursky

import "testing"

func Test_recur(t *testing.T) {
	type args struct {
		n   float64
		cur float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "2 1 then 1.5",
			args: args{n: 2, cur: 1},
			want: 1.5,
		},
		{
			name: "5 3 then 3.8",
			args: args{n: 5, cur: 3},
			want: 3.8,
		},
		{
			name: "5 5 then 5.8",
			args: args{n: 5, cur: 5},
			want: 5.8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recur(tt.args.n, tt.args.cur); got != tt.want {
				t.Errorf("recur() = %v, want %v", got, tt.want)
			}
		})
	}
}
