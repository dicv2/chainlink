// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"

	event "github.com/ethereum/go-ethereum/event"

	generated "github.com/smartcontractkit/chainlink/core/gethwrappers/generated"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"

	vrf_coordinator "github.com/smartcontractkit/chainlink/core/gethwrappers/ocr2vrf/generated/vrf_coordinator"
)

// VRFCoordinatorInterface is an autogenerated mock type for the VRFCoordinatorInterface type
type VRFCoordinatorInterface struct {
	mock.Mock
}

// AcceptOwnership provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Address provides a mock function with given fields:
func (_m *VRFCoordinatorInterface) Address() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// FilterOwnershipTransferRequested provides a mock function with given fields: opts, from, to
func (_m *VRFCoordinatorInterface) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*vrf_coordinator.VRFCoordinatorOwnershipTransferRequestedIterator, error) {
	ret := _m.Called(opts, from, to)

	var r0 *vrf_coordinator.VRFCoordinatorOwnershipTransferRequestedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *vrf_coordinator.VRFCoordinatorOwnershipTransferRequestedIterator); ok {
		r0 = rf(opts, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vrf_coordinator.VRFCoordinatorOwnershipTransferRequestedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOwnershipTransferred provides a mock function with given fields: opts, from, to
func (_m *VRFCoordinatorInterface) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*vrf_coordinator.VRFCoordinatorOwnershipTransferredIterator, error) {
	ret := _m.Called(opts, from, to)

	var r0 *vrf_coordinator.VRFCoordinatorOwnershipTransferredIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *vrf_coordinator.VRFCoordinatorOwnershipTransferredIterator); ok {
		r0 = rf(opts, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vrf_coordinator.VRFCoordinatorOwnershipTransferredIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterRandomWordsFulfilled provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*vrf_coordinator.VRFCoordinatorRandomWordsFulfilledIterator, error) {
	ret := _m.Called(opts)

	var r0 *vrf_coordinator.VRFCoordinatorRandomWordsFulfilledIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *vrf_coordinator.VRFCoordinatorRandomWordsFulfilledIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vrf_coordinator.VRFCoordinatorRandomWordsFulfilledIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterRandomnessFulfillmentRequested provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts) (*vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequestedIterator, error) {
	ret := _m.Called(opts)

	var r0 *vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequestedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequestedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequestedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterRandomnessRequested provides a mock function with given fields: opts, nextBeaconOutputHeight
func (_m *VRFCoordinatorInterface) FilterRandomnessRequested(opts *bind.FilterOpts, nextBeaconOutputHeight []uint64) (*vrf_coordinator.VRFCoordinatorRandomnessRequestedIterator, error) {
	ret := _m.Called(opts, nextBeaconOutputHeight)

	var r0 *vrf_coordinator.VRFCoordinatorRandomnessRequestedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []uint64) *vrf_coordinator.VRFCoordinatorRandomnessRequestedIterator); ok {
		r0 = rf(opts, nextBeaconOutputHeight)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vrf_coordinator.VRFCoordinatorRandomnessRequestedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []uint64) error); ok {
		r1 = rf(opts, nextBeaconOutputHeight)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForgetConsumerSubscriptionID provides a mock function with given fields: opts, consumers
func (_m *VRFCoordinatorInterface) ForgetConsumerSubscriptionID(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, consumers)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []common.Address) *types.Transaction); ok {
		r0 = rf(opts, consumers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []common.Address) error); ok {
		r1 = rf(opts, consumers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfirmationDelays provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) GetConfirmationDelays(opts *bind.CallOpts) ([8]*big.Int, error) {
	ret := _m.Called(opts)

	var r0 [8]*big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) [8]*big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([8]*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProducer provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) GetProducer(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalLinkBalance provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) GetTotalLinkBalance(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IBeaconPeriodBlocks provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) IBeaconPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IStartSlot provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) IStartSlot(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LINK provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) LINK(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaxNumWords provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) MaxNumWords(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MinDelay provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) MinDelay(opts *bind.CallOpts) (uint16, error) {
	ret := _m.Called(opts)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint16); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NUMCONFDELAYS provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	ret := _m.Called(opts)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint8); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Owner provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseLog provides a mock function with given fields: log
func (_m *VRFCoordinatorInterface) ParseLog(log types.Log) (generated.AbigenLog, error) {
	ret := _m.Called(log)

	var r0 generated.AbigenLog
	if rf, ok := ret.Get(0).(func(types.Log) generated.AbigenLog); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.AbigenLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferRequested provides a mock function with given fields: log
func (_m *VRFCoordinatorInterface) ParseOwnershipTransferRequested(log types.Log) (*vrf_coordinator.VRFCoordinatorOwnershipTransferRequested, error) {
	ret := _m.Called(log)

	var r0 *vrf_coordinator.VRFCoordinatorOwnershipTransferRequested
	if rf, ok := ret.Get(0).(func(types.Log) *vrf_coordinator.VRFCoordinatorOwnershipTransferRequested); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vrf_coordinator.VRFCoordinatorOwnershipTransferRequested)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferred provides a mock function with given fields: log
func (_m *VRFCoordinatorInterface) ParseOwnershipTransferred(log types.Log) (*vrf_coordinator.VRFCoordinatorOwnershipTransferred, error) {
	ret := _m.Called(log)

	var r0 *vrf_coordinator.VRFCoordinatorOwnershipTransferred
	if rf, ok := ret.Get(0).(func(types.Log) *vrf_coordinator.VRFCoordinatorOwnershipTransferred); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vrf_coordinator.VRFCoordinatorOwnershipTransferred)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseRandomWordsFulfilled provides a mock function with given fields: log
func (_m *VRFCoordinatorInterface) ParseRandomWordsFulfilled(log types.Log) (*vrf_coordinator.VRFCoordinatorRandomWordsFulfilled, error) {
	ret := _m.Called(log)

	var r0 *vrf_coordinator.VRFCoordinatorRandomWordsFulfilled
	if rf, ok := ret.Get(0).(func(types.Log) *vrf_coordinator.VRFCoordinatorRandomWordsFulfilled); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vrf_coordinator.VRFCoordinatorRandomWordsFulfilled)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseRandomnessFulfillmentRequested provides a mock function with given fields: log
func (_m *VRFCoordinatorInterface) ParseRandomnessFulfillmentRequested(log types.Log) (*vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequested, error) {
	ret := _m.Called(log)

	var r0 *vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequested
	if rf, ok := ret.Get(0).(func(types.Log) *vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequested); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequested)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseRandomnessRequested provides a mock function with given fields: log
func (_m *VRFCoordinatorInterface) ParseRandomnessRequested(log types.Log) (*vrf_coordinator.VRFCoordinatorRandomnessRequested, error) {
	ret := _m.Called(log)

	var r0 *vrf_coordinator.VRFCoordinatorRandomnessRequested
	if rf, ok := ret.Get(0).(func(types.Log) *vrf_coordinator.VRFCoordinatorRandomnessRequested); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vrf_coordinator.VRFCoordinatorRandomnessRequested)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessVRFOutputs provides a mock function with given fields: opts, vrfOutputs, juelsPerFeeCoin, blockHeight, blockHash
func (_m *VRFCoordinatorInterface) ProcessVRFOutputs(opts *bind.TransactOpts, vrfOutputs []vrf_coordinator.VRFBeaconTypesVRFOutput, juelsPerFeeCoin *big.Int, blockHeight uint64, blockHash [32]byte) (*types.Transaction, error) {
	ret := _m.Called(opts, vrfOutputs, juelsPerFeeCoin, blockHeight, blockHash)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []vrf_coordinator.VRFBeaconTypesVRFOutput, *big.Int, uint64, [32]byte) *types.Transaction); ok {
		r0 = rf(opts, vrfOutputs, juelsPerFeeCoin, blockHeight, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []vrf_coordinator.VRFBeaconTypesVRFOutput, *big.Int, uint64, [32]byte) error); ok {
		r1 = rf(opts, vrfOutputs, juelsPerFeeCoin, blockHeight, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Producer provides a mock function with given fields: opts
func (_m *VRFCoordinatorInterface) Producer(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RedeemRandomness provides a mock function with given fields: opts, requestID
func (_m *VRFCoordinatorInterface) RedeemRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	ret := _m.Called(opts, requestID)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int) *types.Transaction); ok {
		r0 = rf(opts, requestID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, *big.Int) error); ok {
		r1 = rf(opts, requestID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestRandomness provides a mock function with given fields: opts, numWords, subID, confirmationDelayArg
func (_m *VRFCoordinatorInterface) RequestRandomness(opts *bind.TransactOpts, numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	ret := _m.Called(opts, numWords, subID, confirmationDelayArg)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint16, uint64, *big.Int) *types.Transaction); ok {
		r0 = rf(opts, numWords, subID, confirmationDelayArg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint16, uint64, *big.Int) error); ok {
		r1 = rf(opts, numWords, subID, confirmationDelayArg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestRandomnessFulfillment provides a mock function with given fields: opts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments
func (_m *VRFCoordinatorInterface) RequestRandomnessFulfillment(opts *bind.TransactOpts, subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint64, uint16, *big.Int, uint32, []byte) *types.Transaction); ok {
		r0 = rf(opts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint64, uint16, *big.Int, uint32, []byte) error); ok {
		r1 = rf(opts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetConfirmationDelays provides a mock function with given fields: opts, confDelays
func (_m *VRFCoordinatorInterface) SetConfirmationDelays(opts *bind.TransactOpts, confDelays [8]*big.Int) (*types.Transaction, error) {
	ret := _m.Called(opts, confDelays)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, [8]*big.Int) *types.Transaction); ok {
		r0 = rf(opts, confDelays)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, [8]*big.Int) error); ok {
		r1 = rf(opts, confDelays)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetProducer provides a mock function with given fields: opts, addr
func (_m *VRFCoordinatorInterface) SetProducer(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, addr)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferLink provides a mock function with given fields: opts, recipient, juelsAmount
func (_m *VRFCoordinatorInterface) TransferLink(opts *bind.TransactOpts, recipient common.Address, juelsAmount *big.Int) (*types.Transaction, error) {
	ret := _m.Called(opts, recipient, juelsAmount)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, *big.Int) *types.Transaction); ok {
		r0 = rf(opts, recipient, juelsAmount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address, *big.Int) error); ok {
		r1 = rf(opts, recipient, juelsAmount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferOwnership provides a mock function with given fields: opts, to
func (_m *VRFCoordinatorInterface) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, to)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferRequested provides a mock function with given fields: opts, sink, from, to
func (_m *VRFCoordinatorInterface) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *vrf_coordinator.VRFCoordinatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, from, to)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *vrf_coordinator.VRFCoordinatorOwnershipTransferRequested, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *vrf_coordinator.VRFCoordinatorOwnershipTransferRequested, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferred provides a mock function with given fields: opts, sink, from, to
func (_m *VRFCoordinatorInterface) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *vrf_coordinator.VRFCoordinatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, from, to)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *vrf_coordinator.VRFCoordinatorOwnershipTransferred, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *vrf_coordinator.VRFCoordinatorOwnershipTransferred, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchRandomWordsFulfilled provides a mock function with given fields: opts, sink
func (_m *VRFCoordinatorInterface) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *vrf_coordinator.VRFCoordinatorRandomWordsFulfilled) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *vrf_coordinator.VRFCoordinatorRandomWordsFulfilled) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *vrf_coordinator.VRFCoordinatorRandomWordsFulfilled) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchRandomnessFulfillmentRequested provides a mock function with given fields: opts, sink
func (_m *VRFCoordinatorInterface) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequested) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequested) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *vrf_coordinator.VRFCoordinatorRandomnessFulfillmentRequested) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchRandomnessRequested provides a mock function with given fields: opts, sink, nextBeaconOutputHeight
func (_m *VRFCoordinatorInterface) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *vrf_coordinator.VRFCoordinatorRandomnessRequested, nextBeaconOutputHeight []uint64) (event.Subscription, error) {
	ret := _m.Called(opts, sink, nextBeaconOutputHeight)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *vrf_coordinator.VRFCoordinatorRandomnessRequested, []uint64) event.Subscription); ok {
		r0 = rf(opts, sink, nextBeaconOutputHeight)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *vrf_coordinator.VRFCoordinatorRandomnessRequested, []uint64) error); ok {
		r1 = rf(opts, sink, nextBeaconOutputHeight)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewVRFCoordinatorInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewVRFCoordinatorInterface creates a new instance of VRFCoordinatorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVRFCoordinatorInterface(t mockConstructorTestingTNewVRFCoordinatorInterface) *VRFCoordinatorInterface {
	mock := &VRFCoordinatorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
