package oursky

import "testing"

func Test_isSubset(t *testing.T) {
	type args struct {
		set1 []string
		set2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[A,B,C,D,E] [A,E,D] then true",
			args: args{set1: []string{"A", "B", "C", "D", "E"}, set2: []string{"A", "E", "D"}},
			want: true,
		},
		{
			name: "[A,B,C,D,E] [A,D,Z] then false",
			args: args{set1: []string{"A", "B", "C", "D", "E"}, set2: []string{"A", "D", "Z"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubset(tt.args.set1, tt.args.set2); got != tt.want {
				t.Errorf("isSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
