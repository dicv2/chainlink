package pipeline

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/mocks"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

var _ types.ContractStateReader = (*fakeContractStateReader)(nil)

type fakeContractStateReader struct {
	returnValue any
	returnError error
}

func (f *fakeContractStateReader) Start(ctx context.Context) error {
	return nil
}

func (f *fakeContractStateReader) Close() error {
	return nil
}

func (f *fakeContractStateReader) Ready() error {
	return nil
}

func (f *fakeContractStateReader) HealthReport() map[string]error {
	return nil
}

func (f *fakeContractStateReader) Name() string {
	return "FakeContractStateReader"
}

func (f *fakeContractStateReader) GetLatestValue(ctx context.Context, contractName, method string, confidenceLevel primitives.ConfidenceLevel, params, returnVal any) error {
	returnVal = f.returnValue
	return f.returnError
}

func (f *fakeContractStateReader) Bind(ctx context.Context, bindings []types.BoundContract) error {
	return nil
}

func TestOnChainReadTask(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		contractAddress       string
		contractName          string
		methodName            string
		params                string
		relayConfig           map[string]interface{}
		relay                 string
		vars                  Vars
		inputs                []Result
		expected              interface{}
		expectedErrorCause    error
		expectedErrorContains string
	}{
		{
			"test",
			"contractAddress1",
			"contractName1",
			"methodName1",
			"{}",
			map[string]interface{}{
				"chainID":     "chainID",
				"chainReader": "{\n    \"contracts\": {\n        \"median\": {\n            \"contractABI\": \"[{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"requester\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"configDigest\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint32\\\",\\\"name\\\":\\\"epoch\\\",\\\"type\\\":\\\"uint32\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint8\\\",\\\"name\\\":\\\"round\\\",\\\"type\\\":\\\"uint8\\\"}],\\\"name\\\":\\\"RoundRequested\\\",\\\"type\\\":\\\"event\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"latestTransmissionDetails\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"configDigest\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"internalType\\\":\\\"uint32\\\",\\\"name\\\":\\\"epoch\\\",\\\"type\\\":\\\"uint32\\\"},{\\\"internalType\\\":\\\"uint8\\\",\\\"name\\\":\\\"round\\\",\\\"type\\\":\\\"uint8\\\"},{\\\"internalType\\\":\\\"int192\\\",\\\"name\\\":\\\"latestAnswer_\\\",\\\"type\\\":\\\"int192\\\"},{\\\"internalType\\\":\\\"uint64\\\",\\\"name\\\":\\\"latestTimestamp_\\\",\\\"type\\\":\\\"uint64\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"}]\",\n            \"configs\": {\n                \"chainSpecificName\": \"latestTransmissionDetails\",\n                \"outputModifications\": [\n                    {\n                        \"Fields\": [\n                            \"LatestTimestamp_\"\n                        ],\n                        \"type\": \"epoch to time\"\n                    },\n                    {\n                        \"Fields\": {\n                            \"LatestAnswer_\": \"LatestAnswer\",\n                            \"LatestTimestamp_\": \"LatestTimestamp\"\n                        },\n                        \"type\": \"rename\"\n                    }\n                ]\n            }\n        }\n    }\n}",
			},
			"network",
			NewVarsFrom(map[string]interface{}{
				"foo": []byte("foo bar"),
			}),
			nil,
			[]byte("baz quux"),
			nil,
			"",
		},
	}

	r := mocks.NewRelayer(t)
	fcsr := &fakeContractStateReader{}
	r.On("NewContractStateReader", mock.Anything, mock.Anything).Return(fcsr, nil)
	relayers := map[types.RelayID]loop.Relayer{
		types.NewRelayID("network", "chainID"): r,
	}
	lggr := logger.TestLogger(t)

	csrm, err := newContractReaderManager(testutils.Context(t), relayers, lggr)
	require.NoError(t, err)
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			task := OnChainRead{
				BaseTask:        NewBaseTask(0, "onChainRead", nil, nil, 0),
				ContractAddress: test.contractAddress,
				ContractName:    test.contractName,
				MethodName:      test.methodName,
				Params:          test.params,
				RelayConfig:     test.relayConfig,
				Relay:           test.relay,
				csrm:            csrm,
			}

			result, runInfo := task.Run(testutils.Context(t), lggr, test.vars, test.inputs)
			require.False(t, runInfo.IsPending)
			require.False(t, runInfo.IsRetryable)

			if test.expectedErrorCause != nil || test.expectedErrorContains != "" {
				require.Nil(t, result.Value)
				if test.expectedErrorCause != nil {
					require.Equal(t, test.expectedErrorCause, errors.Cause(result.Error))
				}
				if test.expectedErrorContains != "" {
					require.Contains(t, result.Error.Error(), test.expectedErrorContains)
				}
			} else {
				require.NoError(t, result.Error)
				require.Equal(t, test.expected, result.Value)
			}
		})
	}
}