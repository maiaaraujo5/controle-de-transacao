package handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/service"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/service/mocks"
	"github.com/maiaaraujo5/controle-de-transacao/app/server/rest/middlewares"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type FindAccountSuite struct {
	suite.Suite
	echo *echo.Echo
}

func TestFindAccountSuite(t *testing.T) {
	suite.Run(t, new(FindAccountSuite))
}

func (s *FindAccountSuite) SetupSuite() {
	e := echo.New()
	e.Use(middlewares.ErrorMiddleware)
	s.echo = e
}

func (s *FindAccountSuite) TestNewFindAccount() {
	validate := validator.New()

	type args struct {
		finder   service.AccountFinder
		validate *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want *FindAccount
	}{
		{
			name: "should successfully build NewFindAccount",
			args: args{
				finder:   new(mocks.AccountFinder),
				validate: validate,
			},
			want: &FindAccount{
				service:   new(mocks.AccountFinder),
				validator: validate,
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewFindAccount(tt.args.finder, tt.args.validate)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewFindAccount() = %v, want %v", got, tt.want)
		})
	}
}

func (s *FindAccountSuite) TestFindAccount_Handle() {
	type fields struct {
		service   *mocks.AccountFinder
		validator *validator.Validate
	}
	type args struct {
		id string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            bool
		wantHttpStatusCode int
		responseBody       string
		mock               func(service *mocks.AccountFinder)
	}{
		{
			name: "should return HTTP 200 ok when account was found",
			fields: fields{
				service:   new(mocks.AccountFinder),
				validator: validator.New(),
			},
			args: args{
				id: "1",
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusOK,
			responseBody:       `{"account_id":1,"document_number":"123"}` + "\n",
			mock: func(service *mocks.AccountFinder) {
				service.On("Finder", mock.Anything, mock.Anything).Return(&model.Account{
					ID:             1,
					DocumentNumber: "123",
				}, nil).Once()
			},
		},
		{
			name: "should return HTTP 500 Internal Server Error when service return error",
			fields: fields{
				service:   new(mocks.AccountFinder),
				validator: validator.New(),
			},
			args: args{
				id: "1",
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusInternalServerError,
			responseBody:       `{"code":500,"description":"error to find account"}` + "\n",
			mock: func(service *mocks.AccountFinder) {
				service.On("Finder", mock.Anything, mock.Anything).Return(nil, errors.New("error to find account")).Once()
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			tt.mock(tt.fields.service)

			h := &FindAccount{
				service:   tt.fields.service,
				validator: tt.fields.validator,
			}

			s.echo.GET("/v1/accounts/:accountId", h.Handle)

			req := httptest.NewRequest(http.MethodGet, "/v1/accounts/"+tt.args.id, nil)
			rec := httptest.NewRecorder()

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			s.echo.ServeHTTP(rec, req)

			s.Assert().Equal(tt.wantHttpStatusCode, rec.Code)
			s.Assert().Equal(tt.responseBody, rec.Body.String())
			tt.fields.service.AssertExpectations(s.T())

		})
	}
}
