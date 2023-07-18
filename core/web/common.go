package web

import (
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm"
)

var (
	ErrMissingChainID = errors.New("evmChainID does not match any local chains")
	ErrInvalidChainID = errors.New("invalid evmChainID")
	ErrMultipleChains = errors.New("more than one chain available, you must specify evmChainID parameter")
)

func getChain(legacyChains *evm.Chains, chainIDstr string) (chain evm.Chain, err error) {
	if legacyChains.Len() > 1 {
		return nil, ErrMultipleChains
	}

	if chainIDstr != "" && chainIDstr != "<nil>" {
		chain, err = legacyChains.Get(chainIDstr)
		if err != nil {
			return nil, ErrMissingChainID
		}
		return chain, nil
	}

	chain, err = legacyChains.Default()
	if err != nil {
		return nil, err
	}
	return chain, nil
}
