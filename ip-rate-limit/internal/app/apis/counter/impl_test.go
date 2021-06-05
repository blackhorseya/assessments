package counter

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/ip-rate-limit/internal/app/biz/counter/mocks"
	"github.com/blackhorseya/ip-rate-limit/pkg/transports/http/middlewares"
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
	s.r.Use(middlewares.IPRateLimitMiddleware(1, 60))

	s.mock = new(mocks.IBiz)
	if handler, err := CreateHandler(s.mock); err != nil {
		panic(err)
	} else {
		s.handler = handler
	}
}

func (s *handlerSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_Count() {
	s.r.GET("", s.handler.Count)

	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody string
	}{
		{
			name: "count then 500 error",
			args: args{mock: func() {
				s.mock.On("Count", mock.Anything, mock.Anything, mock.Anything).Return(
					0, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: "Error",
		},
		{
			name: "count then 200 2",
			args: args{mock: func() {
				s.mock.On("Count", mock.Anything, mock.Anything, mock.Anything).Return(
					2, nil).Once()
			}},
			wantCode: 200,
			wantBody: "2",
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			gotBody, err := ioutil.ReadAll(got.Body)
			if err != nil {
				panic(err)
				return
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Count() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			s.EqualValuesf(tt.wantBody, gotBody, "Count() body = %v, wantBody = %v", gotBody, tt.wantBody)

			s.TearDownTest()
		})
	}
}
