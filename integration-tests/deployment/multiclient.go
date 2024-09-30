package deployment

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-retry"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

const (
	RPC_DEFAULT_RETRY_ATTEMPTS = 10
	RPC_DEFAULT_RETRY_DELAY    = 1000 * time.Millisecond
)

type RetryConfig struct {
	Attempts uint64
	Delay    time.Duration
}

func defaultRetryConfig() RetryConfig {
	return RetryConfig{
		Attempts: RPC_DEFAULT_RETRY_ATTEMPTS,
		Delay:    RPC_DEFAULT_RETRY_DELAY,
	}
}

type RPC struct {
	WSURL string
	// TODO: http fallback needed for some networks?
}

// MultiClient should comply with the OnchainClient interface
var _ OnchainClient = &MultiClient{}

type MultiClient struct {
	logger logger.Logger
	*ethclient.Client
	Backups     []*ethclient.Client
	RetryConfig RetryConfig
}

func WithRetryConfig(attempts uint64, delay time.Duration) func(client *MultiClient) {
	return func(client *MultiClient) {
		client.RetryConfig = RetryConfig{
			Attempts: attempts,
			Delay:    delay,
		}
	}
}

func NewMultiClient(lggr logger.Logger, rpcs []RPC, opts ...func(client *MultiClient)) (*MultiClient, error) {
	if len(rpcs) == 0 {
		return nil, fmt.Errorf("No RPCs provided, need at least one")
	}
	mc := &MultiClient{
		logger: lggr,
	}
	clients := make([]*ethclient.Client, 0, len(rpcs))
	for i, rpc := range rpcs {
		client, err := ethclient.Dial(rpc.WSURL)
		if err != nil {
			lggr.Warnf("failed to dial rpc %d ending in %s, moving to next one", i+1, rpc.WSURL[len(rpc.WSURL)-4:])
			continue
		}
		clients = append(clients, client)
	}
	if len(clients) == 0 {
		return nil, fmt.Errorf("all RPCs failed, try again with different URLs")
	}
	mc.Client = clients[0]
	if len(clients) > 1 {
		mc.Backups = clients[1:]
	} else {
		lggr.Warn("Only one RPC provided, no backups available")
	}
	mc.RetryConfig = defaultRetryConfig()
	mc.logger = lggr

	for _, opt := range opts {
		opt(mc)
	}
	return mc, nil
}

func (mc *MultiClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	var header *types.Header
	err := mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		var err error
		header, err = client.HeaderByNumber(ctx, number)
		return err
	})
	return header, err
}

func (mc *MultiClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var result []byte
	err := mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		var err error
		result, err = client.CallContract(ctx, call, blockNumber)
		return err
	})
	return result, err
}

func (mc *MultiClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	var code []byte
	err := mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		var err error
		code, err = client.PendingCodeAt(ctx, account)
		return err
	})
	return code, err
}

func (mc *MultiClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	var count uint64
	err := mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		var err error
		count, err = client.PendingNonceAt(ctx, account)
		return err
	})
	return count, err
}

func (mc *MultiClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var price *big.Int
	err := mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		var err error
		price, err = client.SuggestGasPrice(ctx)
		return err
	})
	return price, err
}

func (mc *MultiClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	var gas uint64
	err := mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		var err error
		gas, err = client.EstimateGas(ctx, call)
		return err
	})
	return gas, err
}

func (mc *MultiClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	var tipCap *big.Int
	err := mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		var err error
		tipCap, err = client.SuggestGasTipCap(ctx)
		return err
	})
	return tipCap, err
}

func (mc *MultiClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt
	// TransactionReceipt might return ethereum.NotFound error if the transaction is not yet mined
	err := mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		var err error
		receipt, err = client.TransactionReceipt(ctx, txHash)
		return err
	}, ethereum.NotFound)
	return receipt, err
}

func (mc *MultiClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		return client.SendTransaction(ctx, tx)
	})
}

func (mc *MultiClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	var code []byte
	err := mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		var err error
		code, err = client.CodeAt(ctx, account, blockNumber)
		return err
	})
	return code, err
}

func (mc *MultiClient) NonceAt(ctx context.Context, account common.Address) (uint64, error) {
	var count uint64
	err := mc.retryWithBackups(ctx, func(client *ethclient.Client) error {
		var err error
		count, err = client.NonceAt(ctx, account, nil)
		return err
	})
	return count, err
}

func (mc *MultiClient) retryWithBackups(ctx context.Context, op func(*ethclient.Client) error, acceptedErrors ...error) error {
	var err2 error
	funcName := runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name()
	for i, client := range append([]*ethclient.Client{mc.Client}, mc.Backups...) {
		err2 = retry.Do(ctx, retry.WithMaxRetries(mc.RetryConfig.Attempts, retry.NewConstant(mc.RetryConfig.Delay)),
			func(ctx context.Context) error {
				err := op(client)
				if err != nil {
					// Check if the error is one of the accepted errors
					// If it is, log it and return nil
					for _, acceptedError := range acceptedErrors {
						if errors.Is(err, acceptedError) {
							return err
						}
					}
					mc.logger.Warnf("error %+v with client %d for op %s", err, i+1, funcName)
					return retry.RetryableError(err)
				}
				return nil
			})
		if err2 == nil {
			return nil
		}
		for _, acceptedError := range acceptedErrors {
			if errors.Is(err2, acceptedError) {
				return err2
			}
		}
		mc.logger.Infof("Client %d failed, trying next client\n", i+1)
	}
	return errors.Wrapf(err2, "All backup clients %v failed", mc.Backups)
}
