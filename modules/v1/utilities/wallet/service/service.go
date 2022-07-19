package service

import (
	api "ewallet-blockhain/app/contracts"
	"ewallet-blockhain/modules/v1/utilities/wallet/repository"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Service interface {
	GetBalance() (*big.Int, error)
	GetMyWallet() (common.Address, error)
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
