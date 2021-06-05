package health

import (
	"errors"
	"testing"

	"github.com/blackhorseya/amazingtalker/internal/biz/health/repo/mocks"
	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
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

func (s *bizSuite) Test_impl_Readiness() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "readiness then false error",
			args: args{mock: func() {
				s.mock.On("Ping", mock.Anything).Return(false, errors.New("error")).Once()
			}},
			want:    false,
			wantErr: true,
		},
		{
			name: "readiness then true nil",
			args: args{mock: func() {
				s.mock.On("Ping", mock.Anything).Return(true, nil).Once()
			}},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.Readiness(contextx.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Readiness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Readiness() got = %v, want %v", got, tt.want)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_Liveness() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "liveness then false error",
			args: args{mock: func() {
				s.mock.On("Ping", mock.Anything).Return(false, errors.New("error")).Once()
			}},
			want:    false,
			wantErr: true,
		},
		{
			name: "liveness then true nil",
			args: args{mock: func() {
				s.mock.On("Ping", mock.Anything).Return(true, nil).Once()
			}},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.Liveness(contextx.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Liveness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Liveness() got = %v, want %v", got, tt.want)
			}

			s.TearDownTest()
		})
	}
}
