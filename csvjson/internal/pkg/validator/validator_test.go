package validator

import "testing"

func TestValidPhone(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "541-754-3010 then true",
			args: args{"541-754-3010"},
			want: true,
		},
		{
			name: "123-456-00 then false",
			args: args{"123-456-00"},
			want: false,
		},
		{
			name: "1111-1111 then false",
			args: args{"1111-1111"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidPhone(tt.args.number); got != tt.want {
				t.Errorf("ValidPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}
