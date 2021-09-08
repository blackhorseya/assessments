package convert

import (
	"strings"

	"github.com/blackhorseya/csv2json/internal/entity"
)

type impl struct {
}

// NewImpl serve caller to create an IBusiness
func NewImpl() IBusiness {
	return &impl{}
}

func (i *impl) LoadFromCSV(content string) (users []*entity.Profile, err error) {
	rows := strings.Split(content, "\n")

	var ret []*entity.Profile
	for i, row := range rows {
		if i == 0 || len(row) == 0 {
			continue
		}

		col := strings.Split(row, ",")
		ret = append(ret, &entity.Profile{
			ID:          col[0],
			FirstName:   col[1],
			LastName:    col[2],
			Email:       col[3],
			Description: col[4],
			Role:        col[5],
			Phone:       col[6],
		})
	}

	return ret, nil
}
