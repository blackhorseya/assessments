package entities

import (
	"errors"
	"fmt"
	"glasnostic/internal/pkg/validator"
	"reflect"
	"strings"
)

// NewProfileFromLine new a profile from csv line, separate with comma
func NewProfileFromLine(line []string) (*Profile, error) {
	if ok := validator.ValidPhone(line[6]); !ok {
		return nil, errors.New("format of phone number is error")
	}

	return &Profile{
		ID:          line[0],
		FirstName:   line[1],
		LastName:    line[2],
		Email:       line[3],
		Description: line[4],
		Role:        line[5],
		Phone:       line[6],
	}, nil
}

// ToLineByHeader print profile by header
func (x *Profile) ToLineByHeader(header string) (string, error) {
	val := reflect.ValueOf(x).Elem()
	cols := strings.Split(header, ",")
	var ret string
	for _, col := range cols {
		field := val.FieldByName(col)
		ret += fmt.Sprintf("%v,", field.Interface())
	}

	return strings.TrimRight(ret, ","), nil
}
