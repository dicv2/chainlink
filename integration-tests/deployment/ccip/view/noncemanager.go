package view

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type NonceManager struct {
	Contract
	AuthorizedCallers []common.Address `json:"authorizedCallers"`
}

func (nm NonceManager) Address() common.Address {
	return common.HexToAddress(nm.Contract.Address)
}

func NonceManagerSnapshot(nm NonceManagerGetter) (NonceManager, error) {
	authorizedCallers, err := nm.GetAllAuthorizedCallers(nil)
	if err != nil {
		return NonceManager{}, err
	}
	tv, err := nm.TypeAndVersion(nil)
	if err != nil {
		return NonceManager{}, err
	}
	return NonceManager{
		Contract: Contract{
			TypeAndVersion: tv,
			Address:        nm.Address().Hex(),
		},
		// TODO: these can be resolved using an address book
		AuthorizedCallers: authorizedCallers,
	}, nil
}

type NonceManagerGetter interface {
	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)
	TypeAndVersion(opts *bind.CallOpts) (string, error)
	Address() common.Address
}
