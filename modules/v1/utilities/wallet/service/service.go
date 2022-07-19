package service

import (
	api "ewallet-blockhain/app/contracts"
	"ewallet-blockhain/modules/v1/utilities/wallet/repository"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Service interface {
	GetBalance() (*big.Int, error)
	GetMyWallet() (common.Address, error)
	Deposite(amount int, key *bind.TransactOpts) (*types.Transaction, error)
	Withdraw(amount int, key *bind.TransactOpts) (*types.Transaction, error)
}

type service struct {
	repository repository.Repository
	blockhain  *api.Api
}

func NewService(repository repository.Repository, blockhain *api.Api) *service {
	return &service{repository, blockhain}
}

func (s *service) GetBalance() (*big.Int, error) {
	balance, err := s.blockhain.Balance(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (s *service) GetMyWallet() (common.Address, error) {
	mywallet, err := s.blockhain.Admin(&bind.CallOpts{})
	if err != nil {
		return mywallet, err
	}
	return mywallet, nil
}

func (s *service) Deposite(amount int, key *bind.TransactOpts) (*types.Transaction, error) {
	deposit, err := s.blockhain.Deposite(key, big.NewInt(int64(amount)))
	if err != nil {
		return deposit, err
	}
	return deposit, nil
}

func (s *service) Withdraw(amount int, key *bind.TransactOpts) (*types.Transaction, error) {
	withdraw, err := s.blockhain.Withdrawl(key, big.NewInt(int64(amount)))
	if err != nil {
		return withdraw, err
	}
	return withdraw, nil
}
