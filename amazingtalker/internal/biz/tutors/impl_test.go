package tutors

import (
	"errors"
	"reflect"
	"testing"

	"github.com/blackhorseya/amazingtalker/internal/biz/tutors/repo/mocks"
	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/blackhorseya/amazingtalker/pb"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	slug1 = "at-1"

	english = "english"

	tutor1 = &pb.Tutor{
		ID:           "xxx",
		Slug:         "at-1",
		Name:         "Amazing Teacher 1",
		Headline:     "Hi I'm a English Teacher",
		Introduction: ".........",
		PriceInfo: &pb.PriceInfo{
			Trial:  5,
			Normal: 10,
		},
		TeachingLanguages: []int32{123, 121},
	}
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	s.mock = new(mocks.IRepo)
	biz, err := CreateIBiz(s.mock)
	if err != nil {
		panic(err)
	}

	s.biz = biz
}

func (s *bizSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestBizSuite(t *testing.T) {
	suite.Run(t, new(bizSuite))
}

func (s *bizSuite) Test_impl_GetInfoBySlug() {
	type args struct {
		slug string
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Tutor
		wantErr bool
	}{
		{
			name: "at-1 then nil error",
			args: args{slug: slug1, mock: func() {
				s.mock.On("GetInfoBySlug", mock.Anything, slug1).Return(nil, errors.New("error")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "at-1 then tutor nil",
			args: args{slug: slug1, mock: func() {
				s.mock.On("GetInfoBySlug", mock.Anything, slug1).Return(tutor1, nil).Once()
			}},
			want:    tutor1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.GetInfoBySlug(contextx.Background(), tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfoBySlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfoBySlug() got = %v, want %v", got, tt.want)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_ListByLangSlug() {
	type args struct {
		slug string
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    []*pb.Tutor
		wantErr bool
	}{
		{
			name: "english then nil error",
			args: args{slug: english, mock: func() {
				s.mock.On("ListByLangSlug", mock.Anything, english).Return(nil, errors.New("error")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "english then tutors nil",
			args: args{slug: english, mock: func() {
				s.mock.On("ListByLangSlug", mock.Anything, english).Return([]*pb.Tutor{tutor1}, nil).Once()
			}},
			want:    []*pb.Tutor{tutor1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.ListByLangSlug(contextx.Background(), tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListByLangSlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListByLangSlug() got = %v, want %v", got, tt.want)
			}

			s.TearDownTest()
		})
	}
}
