package service

import (
	"context"
	"errors"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/repository"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type FindAccountSuite struct {
	suite.Suite
}

func TestFindAccountSuite(t *testing.T) {
	suite.Run(t, new(FindAccountSuite))
}

func (s *FindAccountSuite) TestNewAccountFinder() {
	type args struct {
		account repository.Account
	}
	tests := []struct {
		name string
		args args
		want *FinderAccount
	}{
		{
			name: "should successfully build NewAccountFinder",
			args: args{
				account: new(mocks.Account),
			},
			want: &FinderAccount{
				repository: new(mocks.Account),
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewAccountFinder(tt.args.account)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewAccountFinder() = %v, want %v", got, tt.want)
		})
	}
}

func (s *FindAccountSuite) TestFinderAccount_Finder() {
	type fields struct {
		repository *mocks.Account
	}
	type args struct {
		ctx context.Context
		ID  int64
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
			name: "should successfully find account when repository return account",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: &model.Account{
				ID:             1,
				DocumentNumber: mock.Anything,
			},
			wantErr: false,
			mock: func(account *mocks.Account) {
				account.On("FindByID", mock.Anything, mock.Anything).Return(&model.Account{
					ID:             1,
					DocumentNumber: mock.Anything,
				}, nil).Once()
			},
		},
		{
			name: "should return error when repository.FindByID return error",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want:    nil,
			wantErr: true,
			mock: func(account *mocks.Account) {
				account.On("FindByID", mock.Anything, mock.Anything).Return(nil, errors.New("error to found account"))
			},
		},
		{
			name: "should return error when repository.FindByID return account as nil",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want:    nil,
			wantErr: true,
			mock: func(account *mocks.Account) {
				account.On("FindByID", mock.Anything, mock.Anything).Return(nil, nil)
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock(tt.fields.repository)
			f := &FinderAccount{
				repository: tt.fields.repository,
			}
			got, err := f.Finder(tt.args.ctx, tt.args.ID)

			s.Assert().True(tt.wantErr == (err != nil), "Finder() error = %v, wantErr %v", err, tt.wantErr)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Finder() got = %v, want %v", got, tt.want)

			tt.fields.repository.AssertExpectations(s.T())
		})
	}
}
