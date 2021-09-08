package convert

import "github.com/blackhorseya/csv2json/internal/entity"

// IBusiness declare convert service functions
type IBusiness interface {
	LoadFromCSV(content string) (users []*entity.Profile, err error)
}
