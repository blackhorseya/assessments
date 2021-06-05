// +build integration

package repo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/blackhorseya/amazingtalker/pb"
	"github.com/stretchr/testify/suite"
)

var (
	slug1 = "at-1"

	tutor1 = &pb.Tutor{
		ID:                "1",
		Slug:              "at-1",
		Name:              "Amazing Teacher 1",
		Headline:          "Hi I'm a English Teacher",
		Introduction:      ".........",
		PriceInfo:         &pb.PriceInfo{Trial: 5, Normal: 10},
		TeachingLanguages: []int32{2, 3},
	}

	tutor2 = &pb.Tutor{
		ID:                "2",
		Slug:              "at-2",
		Name:              "Amazing Teacher 2",
		Headline:          "Hi I'm a English Teacher",
		Introduction:      ".........",
		PriceInfo:         &pb.PriceInfo{Trial: 1, Normal: 5},
		TeachingLanguages: []int32{3},
	}
)

type repoSuite struct {
	suite.Suite
	repo IRepo
}

func (s *repoSuite) SetupTest() {
	if repo, err := CreateRepo("../../../../configs/app.yaml"); err != nil {
		panic(err)
	} else {
		s.repo = repo
	}
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_GetInfoBySlug() {
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Tutor
		wantErr bool
	}{
		{
			name:    "at-1 then tutor nil",
			args:    args{slug: slug1},
			want:    tutor1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.GetInfoBySlug(contextx.Background(), tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfoBySlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfoBySlug() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *repoSuite) Test_impl_ListByLangSlug() {
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		args    args
		want    []*pb.Tutor
		wantErr bool
	}{
		{
			name:    "english then tutors nil",
			args:    args{slug: "english"},
			want:    []*pb.Tutor{tutor1, tutor2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.ListByLangSlug(contextx.Background(), tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListByLangSlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListByLangSlug() got = %v, want %v", got, tt.want)
			}
		})
	}
}
