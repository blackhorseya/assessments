package entities

import (
	"reflect"
	"testing"
)

var (
	p1 = &Profile{
		Id:          "1",
		FirstName:   "Marc",
		LastName:    "Smith",
		Email:       "marc@csvjson.com",
		Description: "Writer of Java",
		Role:        "Dev",
		Phone:       "541-754-3010",
	}
)

func TestNewProfileFromLine(t *testing.T) {
	type args struct {
		line []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Profile
		wantErr bool
	}{
		{
			name: "1,Marc,Smith,marc@csvjson.com,Writer of Java,Dev,541-754-3010 then profile",
			args: args{line: []string{
				"1",
				"Marc",
				"Smith",
				"marc@csvjson.com",
				"Writer of Java",
				"Dev",
				"541-754-3010",
			}},
			want:    p1,
			wantErr: false,
		},
		{
			name: "2,John,Young,john@csvjson.com,Interested in MHW,HR,541-75-3010 then nil error",
			args: args{line: []string{
				"2",
				"John",
				"Young",
				"john@csvjson.com",
				"Interested in MHW",
				"HR",
				"541-75-3010",
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProfileFromLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProfileFromLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProfileFromLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfile_ToLineByHeader(t *testing.T) {
	type fields struct {
		Id          string
		FirstName   string
		LastName    string
		Email       string
		Description string
		Role        string
		Phone       string
	}
	type args struct {
		header string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "header then string",
			fields: fields{
				Id:          "1",
				FirstName:   "Marc",
				LastName:    "Smith",
				Email:       "marc@csvjson.com",
				Description: "Writer of Java",
				Role:        "Dev",
				Phone:       "541-754-3010",
			},
			args:    args{header: "FirstName,LastName,Email,Description,Role,Phone,Id"},
			want:    "Marc,Smith,marc@csvjson.com,Writer of Java,Dev,541-754-3010,1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				Id:          tt.fields.Id,
				FirstName:   tt.fields.FirstName,
				LastName:    tt.fields.LastName,
				Email:       tt.fields.Email,
				Description: tt.fields.Description,
				Role:        tt.fields.Role,
				Phone:       tt.fields.Phone,
			}
			got, err := p.ToLineByHeader(tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToLineByHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToLineByHeader() got = %v, want %v", got, tt.want)
			}
		})
	}
}
