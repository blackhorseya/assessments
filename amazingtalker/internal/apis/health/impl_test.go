package health

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/amazingtalker/internal/biz/health/mocks"
	"github.com/blackhorseya/amazingtalker/internal/pkg/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *mocks.IBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	s.r = gin.New()
	s.r.Use(middlewares.ContextMiddleware())
	s.r.Use(middlewares.LoggerMiddleware())

	s.mock = new(mocks.IBiz)
	handler, err := CreateIHandler(s.mock)
	if err != nil {
		panic(err)
	}

	s.handler = handler
}

func (s *handlerSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_Readiness() {
	s.r.GET("/api/readiness", s.handler.Readiness)

	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "readiness then 500 error",
			args: args{mock: func() {
				s.mock.On("Readiness", mock.Anything).Return(false, errors.New("error")).Once()
			}},
			wantCode: 500,
		},
		{
			name: "readiness then 503 false",
			args: args{mock: func() {
				s.mock.On("Readiness", mock.Anything).Return(false, nil).Once()
			}},
			wantCode: 503,
		},
		{
			name: "readiness then 200 true",
			args: args{mock: func() {
				s.mock.On("Readiness", mock.Anything).Return(true, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/readiness")
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Readiness() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Liveness() {
	s.r.GET("/api/liveness", s.handler.Liveness)

	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "liveness then 500 error",
			args: args{mock: func() {
				s.mock.On("Liveness", mock.Anything).Return(false, errors.New("error")).Once()
			}},
			wantCode: 500,
		},
		{
			name:     "liveness then 503 false",
			args:     args{mock: func() {
				s.mock.On("Liveness", mock.Anything).Return(false, nil).Once()
			}},
			wantCode: 503,
		},
		{
			name:     "liveness then 200 true",
			args:     args{mock: func() {
				s.mock.On("Liveness", mock.Anything).Return(true, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/liveness")
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Liveness() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}
