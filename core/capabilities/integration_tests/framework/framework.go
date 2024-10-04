package framework

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"

	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

func Context(tb testing.TB) context.Context {
	return testutils.Context(tb)
}

func MustNewSimTransactor(t testing.TB) *bind.TransactOpts {
	return testutils.MustNewSimTransactor(t)
}

func NewLegacyTransaction(nonce uint64, to common.Address, value *big.Int, gasLimit uint32, gasPrice *big.Int, data []byte) *types.Transaction {
	return cltest.NewLegacyTransaction(nonce, to, value, gasLimit, gasPrice, data)
}

func GetSimulatedChainID() *big.Int {
	return testutils.SimulatedChainID
}

func NewSimulatedBackend(t *testing.T, alloc core.GenesisAlloc, gasLimit uint32) *backends.SimulatedBackend {
	return cltest.NewSimulatedBackend(t, alloc, gasLimit)
}

func NewApplicationWithConfigV2AndKeyOnSimulatedBlockchain(
	t testing.TB,
	cfg chainlink.GeneralConfig,
	backend *backends.SimulatedBackend,
	flagsAndDeps ...interface{},
) TestApplication {
	return cltest.NewApplicationWithConfigV2AndKeyOnSimulatedBlockchain(t, cfg, backend, flagsAndDeps...)
}

type TestApplication interface {
	AddJobV2(ctx context.Context, j *job.Job) error
	Start(ctx context.Context) error
}
