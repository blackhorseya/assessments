package trans

import (
	"glasnostic/internal/app/entities"
)

// IBiz declare a trans biz service function
type IBiz interface {
	LoadCSV(path string, num int) ([]*entities.Profile, error)
	LoadJSON(path string) ([]*entities.Profile, error)
	ToCSVByHeader(people []*entities.Profile, header string) (string, error)
}
