package service

import (
	"context"
	"errors"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"

	"github.com/maiaaraujo5/controle-de-transacao/app/domain/repository"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type AccountCreatorSuite struct {
	suite.Suite
}

func TestAccountCreatorSuite(t *testing.T) {
	suite.Run(t, new(AccountCreatorSuite))
}

func (s *AccountCreatorSuite) TestNewAccountCreator() {
	type args struct {
		account repository.Account
	}
	tests := []struct {
		name string
		args args
		want *CreatorAccount
	}{
		{
			name: "should successfully build NewAccountCreator",
			args: args{
				account: new(mocks.Account),
			},
			want: &CreatorAccount{
				repository: new(mocks.Account),
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewAccountCreator(tt.args.account)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewAccountCreator() = %v, want %v", got, tt.want)
		})
	}
}

func (s *AccountCreatorSuite) TestCreatorAccount_Create() {
	type fields struct {
		repository *mocks.Account
	}
	type args struct {
		ctx     context.Context
		account *model.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
		mock    func(account *mocks.Account)
	}{
		{
			name: "should successfully create a new account when no account with the same document number exists",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				account: &model.Account{
					DocumentNumber: mock.Anything,
				},
			},
			want: &model.Account{
				ID:             1,
				DocumentNumber: mock.Anything,
			},
			wantErr: false,
			mock: func(account *mocks.Account) {
				account.On("FindByDocumentNumber", mock.Anything, mock.Anything).Return(nil, nil).Once()
				account.On("Save", mock.Anything, mock.Anything).Return(&model.Account{
					ID:             1,
					DocumentNumber: mock.Anything,
				}, nil)
			},
		},
		{
			name: "should return error when repository.FindByDocumentNumber return error",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				account: &model.Account{
					DocumentNumber: mock.Anything,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(account *mocks.Account) {
				account.On("FindByDocumentNumber", mock.Anything, mock.Anything).Return(nil, errors.New("error to find account")).Once()
			},
		},
		{
			name: "should return error when one account with the same document number already exists",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				account: &model.Account{
					DocumentNumber: mock.Anything,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(account *mocks.Account) {
				account.On("FindByDocumentNumber", mock.Anything, mock.Anything).Return(&model.Account{
					ID:             1,
					DocumentNumber: mock.Anything,
				}, nil).Once()
			},
		},
		{
			name: "should return error when repository.Save returns error",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				account: &model.Account{
					DocumentNumber: mock.Anything,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(account *mocks.Account) {
				account.On("FindByDocumentNumber", mock.Anything, mock.Anything).Return(nil, nil).Once()
				account.On("Save", mock.Anything, mock.Anything).Return(nil, errors.New("error to save account"))
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock(tt.fields.repository)

			a := &CreatorAccount{
				repository: tt.fields.repository,
			}
			got, err := a.Create(tt.args.ctx, tt.args.account)

			s.Assert().True(tt.wantErr == (err != nil), "Create() error = %v, wantErr %v", err, tt.wantErr)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Create() = %v, want %v", got, tt.want)

			tt.fields.repository.AssertExpectations(s.T())
		})

	}
}
