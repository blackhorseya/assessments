package q4

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		T *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "then 3",
			args: args{
				T: &Tree{
					X: 1,
					L: &Tree{
						X: 2,
						L: &Tree{
							X: 3,
							L: &Tree{X: 2},
						},
						R: &Tree{X: 6},
					},
					R: &Tree{
						X: 3,
						L: &Tree{X: 3},
						R: &Tree{
							X: 1,
							L: &Tree{X: 5},
							R: &Tree{X: 6},
						},
					},
				},
			},
			want: 3,
		},
		{
			name: "then 2",
			args: args{
				T: &Tree{
					X: 1,
					R: &Tree{
						X: 2,
						L: &Tree{
							X: 1,
						},
						R: &Tree{
							X: 1,
							L: &Tree{
								X: 4,
							},
						},
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.T); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
