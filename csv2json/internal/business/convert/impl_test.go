package convert

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/csv2json/internal/entity"
)

func Test_impl_LoadFromCSV(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name      string
		args      args
		wantUsers []*entity.Profile
		wantErr   bool
	}{
		{
			name:      "load from csv",
			args:      args{content: `ID,FirstName,LastName,Email,Description,Role,Phone
1,Marc,Smith,marc@glasnostic.com,Writer of Java,Dev,541-754-3010
`},
			wantUsers: []*entity.Profile{
				{
					ID:          "1",
					FirstName:   "Marc",
					LastName:    "Smith",
					Email:       "marc@glasnostic.com",
					Description: "Writer of Java",
					Role:        "Dev",
					Phone:       "541-754-3010",
				},
			},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &impl{}
			gotUsers, err := i.LoadFromCSV(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFromCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUsers, tt.wantUsers) {
				t.Errorf("LoadFromCSV() gotUsers = %v, want %v", gotUsers, tt.wantUsers)
			}
		})
	}
}
