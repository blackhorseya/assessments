package trans

import (
	"glasnostic/internal/app/entities"
	"reflect"
	"testing"
)

var (
	p1, err1 = entities.NewProfileFromLine([]string{"1", "Marc", "Smith", "marc@csvjson.com", "Writer of Java", "Dev", "541-754-3010"})
	p2, err2 = entities.NewProfileFromLine([]string{"2", "John", "Young", "john@csvjson.com", "Interested in MHW", "HR", "541-75-3010"})
	p3, err3 = entities.NewProfileFromLine([]string{"3", "Peter", "Scott", "peter@csvjson.com", "amateur boxer", "Dev", "541-754-3010"})
)

func Test_impl_LoadCSV(t *testing.T) {
	type args struct {
		path string
		num  int
	}
	tests := []struct {
		name    string
		args    args
		want    []*entities.Profile
		wantErr bool
	}{
		{
			name:    "path then people nil",
			args:    args{path: "../../../../test/raw.csv", num: 3},
			want:    []*entities.Profile{p1, p3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &impl{}
			got, err := i.LoadCSV(tt.args.path, tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadCSV() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_LoadJson(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []*entities.Profile
		wantErr bool
	}{
		{
			name:    "path then csv",
			args:    args{path: "../../../../test/export.json"},
			want:    []*entities.Profile{p1, p3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &impl{}
			got, err := i.LoadJson(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_ToCSVByHeader(t *testing.T) {
	type args struct {
		people []*entities.Profile
		header string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "print then string nil",
			args: args{people: []*entities.Profile{p1, p3}, header: "ID,FirstName,LastName,Email,Description,Role,Phone"},
			want: `ID,FirstName,LastName,Email,Description,Role,Phone
1,Marc,Smith,marc@csvjson.com,Writer of Java,Dev,541-754-3010
3,Peter,Scott,peter@csvjson.com,amateur boxer,Dev,541-754-3010`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &impl{}
			got, err := i.ToCSVByHeader(tt.args.people, tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCSVByHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToCSVByHeader() got = %v, want %v", got, tt.want)
			}
		})
	}
}
