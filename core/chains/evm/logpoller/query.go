package logpoller

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"

	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type bytesProducer interface {
	Bytes() []byte
}

func concatBytes[T bytesProducer](byteSlice []T) pq.ByteaArray {
	var output [][]byte
	for _, b := range byteSlice {
		output = append(output, b.Bytes())
	}
	return output
}

// queryArgs is a helper for building the arguments to a postgres query created by DbORM
// Besides the convenience methods, it also keeps track of arguments validation and sanitization.
type queryArgs struct {
	args map[string]interface{}
	err  []error
}

func newQueryArgs(chainId *big.Int) *queryArgs {
	return &queryArgs{
		args: map[string]interface{}{
			"evm_chain_id": utils.NewBig(chainId),
		},
		err: []error{},
	}
}

func newQueryArgsForEvent(chainId *big.Int, address common.Address, eventSig common.Hash) *queryArgs {
	return newQueryArgs(chainId).
		withAddress(address).
		withEventSig(eventSig)
}

func (q *queryArgs) withEventSig(eventSig common.Hash) *queryArgs {
	return q.withCustomHashArg("event_sig", eventSig)
}

func (q *queryArgs) withEventSigArray(eventSigs []common.Hash) *queryArgs {
	q.args["event_sig_array"] = concatBytes(eventSigs)
	return q
}

func (q *queryArgs) withAddress(address common.Address) *queryArgs {
	q.args["address"] = address
	return q
}

func (q *queryArgs) withAddressArray(addresses []common.Address) *queryArgs {
	q.args["address_array"] = concatBytes(addresses)
	return q
}

func (q *queryArgs) withStartBlock(startBlock int64) *queryArgs {
	q.args["start_block"] = startBlock
	return q
}

func (q *queryArgs) withEndBlock(endBlock int64) *queryArgs {
	q.args["end_block"] = endBlock
	return q
}

func (q *queryArgs) withWordIndex(wordIndex int) *queryArgs {
	q.args["word_index"] = wordIndex
	return q
}

func (q *queryArgs) withWordValueMin(wordValueMin common.Hash) *queryArgs {
	return q.withCustomHashArg("word_value_min", wordValueMin)
}

func (q *queryArgs) withWordValueMax(wordValueMax common.Hash) *queryArgs {
	return q.withCustomHashArg("word_value_max", wordValueMax)
}

func (q *queryArgs) withConfs(confs Confirmations) *queryArgs {
	q.args["confs"] = confs
	return q
}

func (q *queryArgs) withTopicIndex(index int) *queryArgs {
	// Only topicIndex 1 through 3 is valid. 0 is the event sig and only 4 total topics are allowed
	if !(index == 1 || index == 2 || index == 3) {
		q.err = append(q.err, fmt.Errorf("invalid index for topic: %d", index))
	}
	// Add 1 since postgresql arrays are 1-indexed.
	q.args["topic_index"] = index + 1
	return q
}

func (q *queryArgs) withTopicValueMin(valueMin common.Hash) *queryArgs {
	return q.withCustomHashArg("topic_value_min", valueMin)
}

func (q *queryArgs) withTopicValueMax(valueMax common.Hash) *queryArgs {
	return q.withCustomHashArg("topic_value_max", valueMax)
}

func (q *queryArgs) withTopicValues(values []common.Hash) *queryArgs {
	q.args["topic_values"] = concatBytes(values)
	return q
}

func (q *queryArgs) withBlockTimestampAfter(after time.Time) *queryArgs {
	q.args["block_timestamp_after"] = after
	return q
}

func (q *queryArgs) withTxHash(hash common.Hash) *queryArgs {
	return q.withCustomHashArg("tx_hash", hash)
}

func (q *queryArgs) withCustomHashArg(name string, arg common.Hash) *queryArgs {
	q.args[name] = arg.Bytes()
	return q
}

func (q *queryArgs) withCustomArg(name string, arg any) *queryArgs {
	q.args[name] = arg
	return q
}

func (q *queryArgs) toArgs() (map[string]interface{}, error) {
	if len(q.err) > 0 {
		return nil, errors.Join(q.err...)
	}
	return q.args, nil
}
