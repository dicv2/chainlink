package client

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	commonassets "github.com/smartcontractkit/chainlink-common/pkg/assets"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	commonclient "github.com/smartcontractkit/chainlink/v2/common/client"
	evmconfig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/chaintype"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
)

const BALANCE_OF_ADDRESS_FUNCTION_SELECTOR = "0x70a08231"

var _ Client = (*chainClient)(nil)

// Client is the interface used to interact with an ethereum node.
type Client interface {
	Dial(ctx context.Context) error
	Close()
	// ChainID locally stored for quick access
	ConfiguredChainID() *big.Int

	// NodeStates returns a map of node Name->node state
	// It might be nil or empty, e.g. for mock clients etc
	NodeStates() map[string]commonclient.NodeState

	TokenBalance(ctx context.Context, address common.Address, contractAddress common.Address) (*big.Int, error)
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	LINKBalance(ctx context.Context, address common.Address, linkAddress common.Address) (*commonassets.Link, error)

	// Wrapped RPC methods
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	BatchCallContext(ctx context.Context, b []ethrpc.BatchElem) error
	// BatchCallContextAll calls BatchCallContext for every single node including
	// sendonlys.
	// CAUTION: This should only be used for mass re-transmitting transactions, it
	// might have unexpected effects to use it for anything else.
	BatchCallContextAll(ctx context.Context, b []ethrpc.BatchElem) error

	// HeadByNumber and HeadByHash is a reimplemented version due to a
	// difference in how block header hashes are calculated by Parity nodes
	// running on Kovan, Avalanche and potentially others. We have to return our own wrapper type to capture the
	// correct hash from the RPC response.
	HeadByNumber(ctx context.Context, n *big.Int) (*evmtypes.Head, error)
	HeadByHash(ctx context.Context, n common.Hash) (*evmtypes.Head, error)
	SubscribeNewHead(ctx context.Context) (<-chan *evmtypes.Head, ethereum.Subscription, error)
	// LatestFinalizedBlock - returns the latest finalized block as it's returned from an RPC.
	// CAUTION: Using this method might cause local finality violations. It's highly recommended
	// to use HeadTracker to get latest finalized block.
	LatestFinalizedBlock(ctx context.Context) (head *evmtypes.Head, err error)

	SendTransactionReturnCode(ctx context.Context, tx *types.Transaction, fromAddress common.Address) (commonclient.SendTxReturnCode, error)

	// Wrapped Geth client methods
	// blockNumber can be specified as `nil` to imply latest block
	// if blocks, transactions, or receipts are not found - a nil result and an error are returned
	// these methods may not be compatible with non Ethereum chains as return types may follow different formats
	// suggested options: use HeadByNumber/HeadByHash (above) or CallContext and parse with custom types
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	SequenceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (evmtypes.Nonce, error)
	TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	LatestBlockHeight(ctx context.Context) (*big.Int, error)

	HeaderByNumber(ctx context.Context, n *big.Int) (*types.Header, error)
	HeaderByHash(ctx context.Context, h common.Hash) (*types.Header, error)

	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error)

	IsL2() bool

	// Simulate the transaction prior to sending to catch zk out-of-counters errors ahead of time
	CheckTxValidity(ctx context.Context, from common.Address, to common.Address, data []byte) *SendError
}

func ContextWithDefaultTimeout() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), commonclient.QueryTimeout)
}

type chainClient struct {
	multiNode *commonclient.MultiNode[
		*big.Int,
		*RpcClient,
	]
	txSender     *commonclient.TransactionSender[*types.Transaction, *big.Int, *RpcClient]
	logger       logger.SugaredLogger
	chainType    chaintype.ChainType
	clientErrors evmconfig.ClientErrors
}

func NewChainClient(
	lggr logger.Logger,
	selectionMode string,
	leaseDuration time.Duration,
	nodes []commonclient.Node[*big.Int, *RpcClient],
	sendonlys []commonclient.SendOnlyNode[*big.Int, *RpcClient],
	chainID *big.Int,
	clientErrors evmconfig.ClientErrors,
	deathDeclarationDelay time.Duration,
) Client {
	chainFamily := "EVM"
	multiNode := commonclient.NewMultiNode[*big.Int, *RpcClient](
		lggr,
		selectionMode,
		leaseDuration,
		nodes,
		sendonlys,
		chainID,
		chainFamily,
		deathDeclarationDelay,
	)

	classifySendError := func(tx *types.Transaction, err error) commonclient.SendTxReturnCode {
		return ClassifySendError(err, clientErrors, logger.Sugared(logger.Nop()), tx, common.Address{}, false)
	}

	txSender := commonclient.NewTransactionSender[*types.Transaction, *big.Int, *RpcClient](
		lggr,
		chainID,
		chainFamily,
		multiNode,
		classifySendError,
		0, // use the default value provided by the implementation
	)

	return &chainClient{
		multiNode:    multiNode,
		txSender:     txSender,
		logger:       logger.Sugared(lggr),
		clientErrors: clientErrors,
	}
}

func (c *chainClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.BalanceAt(ctx, account, blockNumber)
}

// BatchCallContext - sends all given requests as a single batch.
// Request specific errors for batch calls are returned to the individual BatchElem.
// Ensure the same BatchElem slice provided by the caller is passed through the call stack
// to ensure the caller has access to the errors.
// Note: some chains (e.g Astar) have custom finality requests, so even when FinalityTagEnabled=true, finality tag
// might not be properly handled and returned results might have weaker finality guarantees. It's highly recommended
// to use HeadTracker to identify latest finalized block.
func (c *chainClient) BatchCallContext(ctx context.Context, b []ethrpc.BatchElem) error {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return err
	}
	return rpc.BatchCallContext(ctx, b)
}

// Similar to BatchCallContext, ensure the provided BatchElem slice is passed through
func (c *chainClient) BatchCallContextAll(ctx context.Context, b []ethrpc.BatchElem) error {
	var wg sync.WaitGroup
	defer wg.Wait()

	// Select main RPC to use for return value
	main, selectionErr := c.multiNode.SelectRPC()
	if selectionErr != nil {
		return selectionErr
	}

	doFunc := func(ctx context.Context, rpc *RpcClient, isSendOnly bool) {
		if rpc == main {
			return
		}
		// Parallel call made to all other nodes with ignored return value
		wg.Add(1)
		go func(rpc *RpcClient) {
			defer wg.Done()
			err := rpc.BatchCallContext(ctx, b)
			if err != nil {
				c.logger.Debugw("Secondary node BatchCallContext failed", "err", err)
			} else {
				c.logger.Debug("Secondary node BatchCallContext success")
			}
		}(rpc)
	}

	if err := c.multiNode.DoAll(ctx, doFunc); err != nil {
		return err
	}
	return main.BatchCallContext(ctx, b)
}

// TODO-1663: return custom Block type instead of geth's once client.go is deprecated.
func (c *chainClient) BlockByHash(ctx context.Context, hash common.Hash) (b *types.Block, err error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return b, err
	}
	return rpc.BlockByHashGeth(ctx, hash)
}

// TODO-1663: return custom Block type instead of geth's once client.go is deprecated.
func (c *chainClient) BlockByNumber(ctx context.Context, number *big.Int) (b *types.Block, err error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return b, err
	}
	return rpc.BlockByNumberGeth(ctx, number)
}

func (c *chainClient) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return err
	}
	return rpc.CallContext(ctx, result, method, args...)
}

func (c *chainClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.CallContract(ctx, msg, blockNumber)
}

func (c *chainClient) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.PendingCallContract(ctx, msg)
}

func (c *chainClient) Close() {
	_ = c.multiNode.Close()
	_ = c.txSender.Close()
}

func (c *chainClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.CodeAt(ctx, account, blockNumber)
}

func (c *chainClient) ConfiguredChainID() *big.Int {
	return c.multiNode.ChainID()
}

func (c *chainClient) Dial(ctx context.Context) error {
	err := c.multiNode.Start(ctx)
	if err != nil {
		return err
	}
	return c.txSender.Start(ctx)
}

func (c *chainClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return 0, err
	}
	return rpc.EstimateGas(ctx, call)
}
func (c *chainClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.FilterEvents(ctx, q)
}

func (c *chainClient) HeaderByHash(ctx context.Context, h common.Hash) (head *types.Header, err error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return head, err
	}
	return rpc.HeaderByHash(ctx, h)
}

func (c *chainClient) HeaderByNumber(ctx context.Context, n *big.Int) (head *types.Header, err error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return head, err
	}
	return rpc.HeaderByNumber(ctx, n)
}

func (c *chainClient) HeadByHash(ctx context.Context, h common.Hash) (*evmtypes.Head, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.BlockByHash(ctx, h)
}

func (c *chainClient) HeadByNumber(ctx context.Context, n *big.Int) (*evmtypes.Head, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.BlockByNumber(ctx, n)
}

func (c *chainClient) IsL2() bool {
	return c.chainType.IsL2()
}

func (c *chainClient) LINKBalance(ctx context.Context, address common.Address, linkAddress common.Address) (*commonassets.Link, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.LINKBalance(ctx, address, linkAddress)
}

func (c *chainClient) LatestBlockHeight(ctx context.Context) (*big.Int, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.LatestBlockHeight(ctx)
}

func (c *chainClient) NodeStates() map[string]commonclient.NodeState {
	return c.multiNode.NodeStates()
}

func (c *chainClient) PendingCodeAt(ctx context.Context, account common.Address) (b []byte, err error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return b, err
	}
	return rpc.PendingCodeAt(ctx, account)
}

// TODO-1663: change this to evmtypes.Nonce(int64) once client.go is deprecated.
func (c *chainClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return 0, err
	}
	n, err := rpc.PendingSequenceAt(ctx, account)
	return uint64(n), err
}

func (c *chainClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	_, err := c.txSender.SendTransaction(ctx, tx)
	return err
}

func (c *chainClient) SendTransactionReturnCode(ctx context.Context, tx *types.Transaction, fromAddress common.Address) (commonclient.SendTxReturnCode, error) {
	err := c.SendTransaction(ctx, tx)
	returnCode := ClassifySendError(err, c.clientErrors, c.logger, tx, fromAddress, c.IsL2())
	return returnCode, err
}

func (c *chainClient) SequenceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (evmtypes.Nonce, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return 0, err
	}
	return rpc.SequenceAt(ctx, account, blockNumber)
}

func (c *chainClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (s ethereum.Subscription, err error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return s, err
	}
	return rpc.SubscribeFilterLogs(ctx, q, ch)
}

func (c *chainClient) SubscribeNewHead(ctx context.Context) (<-chan *evmtypes.Head, ethereum.Subscription, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, nil, err
	}

	ch, sub, err := rpc.SubscribeToHeads(ctx)
	if err != nil {
		return nil, nil, err
	}

	return ch, sub, nil
}

func (c *chainClient) SuggestGasPrice(ctx context.Context) (p *big.Int, err error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return p, err
	}
	return rpc.SuggestGasPrice(ctx)
}

func (c *chainClient) SuggestGasTipCap(ctx context.Context) (t *big.Int, err error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return t, err
	}
	return rpc.SuggestGasTipCap(ctx)
}

func (c *chainClient) TokenBalance(ctx context.Context, address common.Address, contractAddress common.Address) (*big.Int, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.TokenBalance(ctx, address, contractAddress)
}

func (c *chainClient) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.TransactionByHash(ctx, txHash)
}

// TODO-1663: return custom Receipt type instead of geth's once client.go is deprecated.
func (c *chainClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (r *types.Receipt, err error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return r, err
	}
	//return rpc.TransactionReceipt(ctx, txHash)
	return rpc.TransactionReceiptGeth(ctx, txHash)
}

func (c *chainClient) LatestFinalizedBlock(ctx context.Context) (*evmtypes.Head, error) {
	rpc, err := c.multiNode.SelectRPC()
	if err != nil {
		return nil, err
	}
	return rpc.LatestFinalizedBlock(ctx)
}

func (c *chainClient) CheckTxValidity(ctx context.Context, from common.Address, to common.Address, data []byte) *SendError {
	msg := ethereum.CallMsg{
		From: from,
		To:   &to,
		Data: data,
	}
	return SimulateTransaction(ctx, c, c.logger, c.chainType, msg)
}
