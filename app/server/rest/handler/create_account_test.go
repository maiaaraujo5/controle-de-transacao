package handler

import (
	errors2 "errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/errors"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/service"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/service/mocks"
	"github.com/maiaaraujo5/controle-de-transacao/app/server/rest/middlewares"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type CreateAccountSuite struct {
	suite.Suite
	echo *echo.Echo
}

func TestCreateAccountSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountSuite))
}

func (s *CreateAccountSuite) SetupSuite() {
	e := echo.New()
	e.Use(middlewares.ErrorMiddleware)
	s.echo = e

}

func (s *CreateAccountSuite) TestNewCreateAccount() {

	validate := validator.New()

	type args struct {
		service  service.AccountCreator
		validate *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want *CreateAccount
	}{
		{
			name: "should successfully build NewCreateAccount",
			args: args{
				service:  new(mocks.AccountCreator),
				validate: validate,
			},
			want: &CreateAccount{
				service:   new(mocks.AccountCreator),
				validator: validate,
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewCreateAccount(tt.args.service, tt.args.validate)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewCreateAccount() = %v, want %v", got, tt.want)
		})
	}
}

func (s *CreateAccountSuite) TestCreateAccount_Handle() {
	type fields struct {
		service   *mocks.AccountCreator
		validator *validator.Validate
	}
	type args struct {
		body io.Reader
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            bool
		wantHttpStatusCode int
		responseBody       string
		mock               func(service *mocks.AccountCreator)
	}{
		{
			name: "should return HTTP 201 Created when create account is successfully",
			fields: fields{
				service:   new(mocks.AccountCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"document_number":"123"}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusCreated,
			responseBody:       `{"account_id":1,"document_number":"123"}` + "\n",
			mock: func(service *mocks.AccountCreator) {
				service.On("Create", mock.Anything, mock.Anything).Return(&model.Account{
					ID:             1,
					DocumentNumber: "123",
				}, nil).Once()
			},
		},
		{
			name: "should return HTTP 409 Conflict when create account return accountAlreadyExists error",
			fields: fields{
				service:   new(mocks.AccountCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"document_number":"123"}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusConflict,
			responseBody:       `{"code":409,"description":"mock.Anything"}` + "\n",
			mock: func(service *mocks.AccountCreator) {
				service.On("Create", mock.Anything, mock.Anything).Return(nil, errors.AccountAlreadyExists(mock.Anything)).Once()
			},
		},
		{
			name: "should return HTTP 500 Internal Server Error when create account return error",
			fields: fields{
				service:   new(mocks.AccountCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"document_number":"123"}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusInternalServerError,
			responseBody:       `{"code":500,"description":"mock.Anything"}` + "\n",
			mock: func(service *mocks.AccountCreator) {
				service.On("Create", mock.Anything, mock.Anything).Return(nil, errors2.New(mock.Anything)).Once()
			},
		},
		{
			name: "should return HTTP 400 BadRequest when document_number was sent in another type",
			fields: fields{
				service:   new(mocks.AccountCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"document_number":123}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusBadRequest,
			responseBody:       `{"code":400,"description":"Unmarshal type error: expected=string, got=number, field=document_number, offset=22"}` + "\n",
			mock: func(service *mocks.AccountCreator) {
			},
		},
		{
			name: "should return HTTP 422 StatusUnprocessableEntity when dont send document_number",
			fields: fields{
				service:   new(mocks.AccountCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusUnprocessableEntity,
			responseBody:       `{"code":422,"description":"The server understands the content type of the request entity but was unable to process the contained instructions.","validationError":[{"path":"CreateAccount.DocumentNumber","field":"DocumentNumber","value":"","message":"{DocumentNumber} is a required field with type string"}]}` + "\n",
			mock: func(service *mocks.AccountCreator) {
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			tt.mock(tt.fields.service)

			h := &CreateAccount{
				service:   tt.fields.service,
				validator: tt.fields.validator,
			}

			s.echo.POST("/v1/accounts", h.Handle)

			req := httptest.NewRequest(http.MethodPost, "/v1/accounts", tt.args.body)
			rec := httptest.NewRecorder()

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			s.echo.ServeHTTP(rec, req)

			s.Assert().Equal(tt.wantHttpStatusCode, rec.Code)
			s.Assert().Equal(tt.responseBody, rec.Body.String())
			tt.fields.service.AssertExpectations(s.T())

		})
	}
}
