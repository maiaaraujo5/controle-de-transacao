package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/repository"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/repository/mocks"
	mocks2 "github.com/maiaaraujo5/controle-de-transacao/app/domain/service/mocks"
	"github.com/stretchr/testify/suite"
)

type TransactionCreatorSuite struct {
	suite.Suite
}

func TestTransactionCreatorSuite(t *testing.T) {
	suite.Run(t, new(TransactionCreatorSuite))
}

func (s *TransactionCreatorSuite) SetupSuite() {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2024, time.March, 28, 00, 31, 00, 0, time.UTC)
	})
}

func (s *TransactionCreatorSuite) TestNewTransactionCreator() {
	type args struct {
		transaction repository.Transaction
		finder      AccountFinder
	}
	tests := []struct {
		name string
		args args
		want *CreatorTransaction
	}{
		{
			name: "should successfully build NewTransactionCreator",
			args: args{
				transaction: new(mocks.Transaction),
				finder:      new(mocks2.AccountFinder),
			},
			want: &CreatorTransaction{
				repository:    new(mocks.Transaction),
				accountFinder: new(mocks2.AccountFinder),
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewTransactionCreator(tt.args.transaction, tt.args.finder)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewTransactionCreator() = %v, want %v", got, tt.want)
		})
	}
}

func (s *TransactionCreatorSuite) TestCreatorTransaction_Create() {
	type fields struct {
		repository    *mocks.Transaction
		accountFinder *mocks2.AccountFinder
	}
	type args struct {
		ctx         context.Context
		transaction *model.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Transaction
		wantErr bool
		mock    func(transaction *mocks.Transaction, account *mocks2.AccountFinder)
	}{
		{
			name: "should create a successfully transaction with Payment Operation",
			fields: fields{
				repository:    new(mocks.Transaction),
				accountFinder: new(mocks2.AccountFinder),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(4),
					Amount:    100,
				},
			},
			want: &model.Transaction{
				ID:        1,
				AccountID: 1,
				Operation: model.OperationType(4),
				Amount:    100,
				EventDate: time.Now(),
			},
			wantErr: false,
			mock: func(transaction *mocks.Transaction, account *mocks2.AccountFinder) {
				account.On("Finder", mock.Anything, mock.Anything).Return(&model.Account{
					ID:             1,
					DocumentNumber: mock.Anything,
				}, nil).Once()

				transaction.On("Save", mock.Anything, &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(4),
					Amount:    100,
					EventDate: time.Now(),
				}).Return(&model.Transaction{
					ID:        1,
					AccountID: 1,
					Operation: model.OperationType(4),
					Amount:    100,
					EventDate: time.Now(),
				}, nil)
			},
		},
		{
			name: "should create a successfully transaction with In Cash Operation",
			fields: fields{
				repository:    new(mocks.Transaction),
				accountFinder: new(mocks2.AccountFinder),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(1),
					Amount:    100,
				},
			},
			want: &model.Transaction{
				ID:        1,
				AccountID: 1,
				Operation: model.OperationType(1),
				Amount:    -100,
				EventDate: time.Now(),
			},
			wantErr: false,
			mock: func(transaction *mocks.Transaction, account *mocks2.AccountFinder) {
				account.On("Finder", mock.Anything, mock.Anything).Return(&model.Account{
					ID:             1,
					DocumentNumber: mock.Anything,
				}, nil).Once()

				transaction.On("Save", mock.Anything, &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(1),
					Amount:    -100,
					EventDate: time.Now(),
				}).Return(&model.Transaction{
					ID:        1,
					AccountID: 1,
					Operation: model.OperationType(1),
					Amount:    -100,
					EventDate: time.Now(),
				}, nil)
			},
		},
		{
			name: "should create a successfully transaction with Installment Payment operation",
			fields: fields{
				repository:    new(mocks.Transaction),
				accountFinder: new(mocks2.AccountFinder),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(2),
					Amount:    100,
				},
			},
			want: &model.Transaction{
				ID:        1,
				AccountID: 1,
				Operation: model.OperationType(2),
				Amount:    -100,
				EventDate: time.Now(),
			},
			wantErr: false,
			mock: func(transaction *mocks.Transaction, account *mocks2.AccountFinder) {
				account.On("Finder", mock.Anything, mock.Anything).Return(&model.Account{
					ID:             1,
					DocumentNumber: mock.Anything,
				}, nil).Once()

				transaction.On("Save", mock.Anything, &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(2),
					Amount:    -100,
					EventDate: time.Now(),
				}).Return(&model.Transaction{
					ID:        1,
					AccountID: 1,
					Operation: model.OperationType(2),
					Amount:    -100,
					EventDate: time.Now(),
				}, nil)
			},
		},
		{
			name: "should create a successfully transaction with withdraw operation",
			fields: fields{
				repository:    new(mocks.Transaction),
				accountFinder: new(mocks2.AccountFinder),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(3),
					Amount:    100,
				},
			},
			want: &model.Transaction{
				ID:        1,
				AccountID: 1,
				Operation: model.OperationType(3),
				Amount:    -100,
				EventDate: time.Now(),
			},
			wantErr: false,
			mock: func(transaction *mocks.Transaction, account *mocks2.AccountFinder) {
				account.On("Finder", mock.Anything, mock.Anything).Return(&model.Account{
					ID:             1,
					DocumentNumber: mock.Anything,
				}, nil).Once()

				transaction.On("Save", mock.Anything, &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(3),
					Amount:    -100,
					EventDate: time.Now(),
				}).Return(&model.Transaction{
					ID:        1,
					AccountID: 1,
					Operation: model.OperationType(3),
					Amount:    -100,
					EventDate: time.Now(),
				}, nil)
			},
		},
		{
			name: "should return error when transaction has type not valid",
			fields: fields{
				repository:    new(mocks.Transaction),
				accountFinder: new(mocks2.AccountFinder),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(5),
					Amount:    100,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(transaction *mocks.Transaction, account *mocks2.AccountFinder) {

			},
		},
		{
			name: "should return error when account finder returns error",
			fields: fields{
				repository:    new(mocks.Transaction),
				accountFinder: new(mocks2.AccountFinder),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(4),
					Amount:    100,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(transaction *mocks.Transaction, account *mocks2.AccountFinder) {
				account.On("Finder", mock.Anything, mock.Anything).Return(nil, errors.New("error to found account"))
			},
		},
		{
			name: "should return error when repository.Save returns error",
			fields: fields{
				repository:    new(mocks.Transaction),
				accountFinder: new(mocks2.AccountFinder),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID: 1,
					Operation: model.OperationType(4),
					Amount:    100,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(transaction *mocks.Transaction, account *mocks2.AccountFinder) {
				account.On("Finder", mock.Anything, mock.Anything).Return(&model.Account{
					ID:             1,
					DocumentNumber: mock.Anything,
				}, nil)

				transaction.On("Save", mock.Anything, mock.Anything).Return(nil, errors.New("error to save transaction"))
			},
		},
	}
	for _, tt := range tests {

		tt.mock(tt.fields.repository, tt.fields.accountFinder)

		s.Run(tt.name, func() {
			c := &CreatorTransaction{
				repository:    tt.fields.repository,
				accountFinder: tt.fields.accountFinder,
			}
			got, err := c.Create(tt.args.ctx, tt.args.transaction)

			s.Assert().True(tt.wantErr == (err != nil), "Create() error = %v, wantErr %v", err, tt.wantErr)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Create() got = %v, want %v", got, tt.want)
		})
	}
}
