package tutors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/amazingtalker/internal/biz/tutors/mocks"
	"github.com/blackhorseya/amazingtalker/internal/pkg/responses"
	"github.com/blackhorseya/amazingtalker/internal/pkg/transports/http/middlewares"
	"github.com/blackhorseya/amazingtalker/pb"
	"github.com/gin-gonic/gin"
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

func (s *handlerSuite) Test_impl_GetInfoBySlug() {
	s.r.GET("/api/tutor/:slug", s.handler.GetInfoBySlug)

	type args struct {
		slug string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *responses.DataResp
	}{
		{
			name:     "missing slug then 404 error",
			args:     args{slug: ""},
			wantCode: 404,
			wantBody: nil,
		},
		{
			name: "at-1 then 500 error",
			args: args{slug: slug1, mock: func() {
				s.mock.On("GetInfoBySlug", mock.Anything, slug1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "at-1 then 404 nil",
			args: args{slug: slug1, mock: func() {
				s.mock.On("GetInfoBySlug", mock.Anything, slug1).Return(nil, nil).Once()
			}},
			wantCode: 404,
			wantBody: nil,
		},
		{
			name: "at-1 then 200 tutor",
			args: args{slug: slug1, mock: func() {
				s.mock.On("GetInfoBySlug", mock.Anything, slug1).Return(tutor1, nil).Once()
			}},
			wantCode: 200,
			wantBody: &responses.DataResp{Data: tutor1},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/tutor/%v", tt.args.slug)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *responses.DataResp
			if err := json.Unmarshal(body, &gotBody); err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetInfoBySlug() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_ListByLangSlug() {
	s.r.GET("/api/tutors/:slug", s.handler.ListByLangSlug)

	type args struct {
		slug string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *responses.DataResp
	}{
		{
			name:     "missing slug then 404 error",
			args:     args{slug: ""},
			wantCode: 404,
			wantBody: nil,
		},
		{
			name: "english then 500 error",
			args: args{slug: english, mock: func() {
				s.mock.On("ListByLangSlug", mock.Anything, english).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "english then 404 nil",
			args: args{slug: english, mock: func() {
				s.mock.On("ListByLangSlug", mock.Anything, english).Return(nil, nil).Once()
			}},
			wantCode: 404,
			wantBody: nil,
		},
		{
			name: "english then 200 tutors",
			args: args{slug: english, mock: func() {
				s.mock.On("ListByLangSlug", mock.Anything, english).Return([]*pb.Tutor{tutor1}, nil).Once()
			}},
			wantCode: 200,
			wantBody: &responses.DataResp{Data: []*pb.Tutor{tutor1}},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/tutors/%v", tt.args.slug)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *responses.DataResp
			if err := json.Unmarshal(body, &gotBody); err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "ListByLangSlug() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}
