package trans

import (
	"bufio"
	"encoding/json"
	"fmt"
	"glasnostic/internal/app/entities"
	"glasnostic/internal/pkg/worker"
	"io/ioutil"
	"os"
	"strings"
)

type impl struct {
}

// NewImpl serve caller to new a IBiz
func NewImpl() IBiz {
	return &impl{}
}

// LoadCSV serve user to load csv file to people
func (i *impl) LoadCSV(path string, num int) ([]*entities.Profile, error) {
	pool := worker.NewPool(num).Start()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	index := -1
	for scanner.Scan() {
		index++
		if index == 0 {
			continue
		}

		pool.In <-strings.Split(scanner.Text(), ",")
	}

	var people []*entities.Profile
	go func() {
		for person := range pool.Out {
			people = append(people, person)
		}
	}()

	<-pool.Stop

	return people, nil
}

// LoadJSON serve user to load json file to people
func (i *impl) LoadJSON(path string) ([]*entities.Profile, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var ret []*entities.Profile
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// ToCSVByHeader serve user to print people csv format by header
func (i *impl) ToCSVByHeader(people []*entities.Profile, header string) (string, error) {
	content := []string{header}

	for _, person := range people {
		line, err := person.ToLineByHeader(header)
		if err != nil {
			fmt.Println(err)
			continue
		}

		content = append(content, line)
	}

	return strings.Join(content, "\n"), nil
}
