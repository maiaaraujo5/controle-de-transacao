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
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

type CreateTransactionSuite struct {
	suite.Suite
	echo *echo.Echo
}

func TestCreateTransactionSuite(t *testing.T) {
	suite.Run(t, new(CreateTransactionSuite))
}

func (s *CreateTransactionSuite) SetupSuite() {
	e := echo.New()
	e.Use(middlewares.ErrorMiddleware)
	s.echo = e
}

func (s *CreateTransactionSuite) TestNewCreateTransaction() {
	validate := validator.New()

	type args struct {
		service  service.TransactionCreator
		validate *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want *CreateTransaction
	}{
		{
			name: "should successfully build NewCreateTransaction",
			args: args{
				service:  new(mocks.TransactionCreator),
				validate: validate,
			},
			want: &CreateTransaction{
				service:   new(mocks.TransactionCreator),
				validator: validate,
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewCreateTransaction(tt.args.service, tt.args.validate)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewCreateTransaction() = %v, want %v", got, tt.want)
		})
	}
}

func (s *CreateTransactionSuite) TestCreateTransaction_Handle() {
	type fields struct {
		service   *mocks.TransactionCreator
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
		mock               func(service *mocks.TransactionCreator)
	}{
		{
			name: "should return HTTP 201 Created when create transaction is successfully",
			fields: fields{
				service:   new(mocks.TransactionCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1,"operation_type_id":1,"amount":100.0}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusCreated,
			responseBody:       `{"id":1,"account_id":1,"operation_type_id":1,"amount":100,"event_date":"2024-03-30T15:29:39.000719683Z"}` + "\n",
			mock: func(service *mocks.TransactionCreator) {
				service.On("Create", mock.Anything, mock.Anything).Return(&model.Transaction{
					ID:        1,
					AccountID: 1,
					Operation: 1,
					Amount:    100.0,
					EventDate: time.Date(2024, 03, 30, 15, 29, 39, 719683, time.UTC),
				}, nil)
			},
		},
		{
			name: "should return HTTP 422 when account id was not sent",
			fields: fields{
				service:   new(mocks.TransactionCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"operation_type_id":1,"amount":100.0}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusUnprocessableEntity,
			responseBody:       `{"code":422,"description":"The server understands the content type of the request entity but was unable to process the contained instructions.","validationError":[{"path":"CreateTransaction.AccountID","field":"AccountID","value":0,"message":"{AccountID} is a required field with type int64"}]}` + "\n",
			mock: func(service *mocks.TransactionCreator) {
			},
		},
		{
			name: "should return HTTP 422 when operation_type_id was not sent",
			fields: fields{
				service:   new(mocks.TransactionCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1,"amount":100.0}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusUnprocessableEntity,
			responseBody:       `{"code":422,"description":"The server understands the content type of the request entity but was unable to process the contained instructions.","validationError":[{"path":"CreateTransaction.OperationTypeID","field":"OperationTypeID","value":0,"message":"{OperationTypeID} is a required field with type int64"}]}` + "\n",
			mock: func(service *mocks.TransactionCreator) {
			},
		},
		{
			name: "should return HTTP 422 when amount was not sent",
			fields: fields{
				service:   new(mocks.TransactionCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1,"operation_type_id":1}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusUnprocessableEntity,
			responseBody:       `{"code":422,"description":"The server understands the content type of the request entity but was unable to process the contained instructions.","validationError":[{"path":"CreateTransaction.Amount","field":"Amount","value":0,"message":"{Amount} is a required field with type float64"}]}` + "\n",
			mock: func(service *mocks.TransactionCreator) {
			},
		},
		{
			name: "should return 400 Bad Request when some type was sent in different type",
			fields: fields{
				service:   new(mocks.TransactionCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":"1","operation_type_id":1,"amount":100.0}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusBadRequest,
			responseBody:       `{"code":400,"description":"Unmarshal type error: expected=int64, got=string, field=account_id, offset=17"}` + "\n",
			mock: func(service *mocks.TransactionCreator) {
			},
		},
		{
			name: "should return HTTP 500 Internal Server Error when create transaction return error",
			fields: fields{
				service:   new(mocks.TransactionCreator),
				validator: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1,"operation_type_id":1,"amount":100.0}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusInternalServerError,
			responseBody:       `{"code":500,"description":"error to create transaction"}` + "\n",
			mock: func(service *mocks.TransactionCreator) {
				service.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("error to create transaction"))
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {

			tt.mock(tt.fields.service)

			h := &CreateTransaction{
				service:   tt.fields.service,
				validator: tt.fields.validator,
			}

			s.echo.POST("/v1/transactions", h.Handle)

			req := httptest.NewRequest(http.MethodPost, "/v1/transactions", tt.args.body)
			rec := httptest.NewRecorder()

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			s.echo.ServeHTTP(rec, req)

			s.Assert().Equal(tt.wantHttpStatusCode, rec.Code)
			s.Assert().Equal(tt.responseBody, rec.Body.String())
			tt.fields.service.AssertExpectations(s.T())

		})
	}
}
