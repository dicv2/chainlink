package vrf_test

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/solidity_vrf_coordinator_interface"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/vrfkey"
	"github.com/smartcontractkit/chainlink/core/services/pg"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	"github.com/smartcontractkit/chainlink/core/services/vrf"
	"github.com/smartcontractkit/chainlink/core/testdata/testspecs"
)

func TestIntegration_VRF_JPV2(t *testing.T) {
	tests := []struct {
		name    string
		eip1559 bool
	}{
		{"legacy mode", false},
		{"eip1559 mode", true},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			config, _ := heavyweight.FullTestDB(t, fmt.Sprintf("vrf_jpv2_%v", test.eip1559), true, true)
			config.Overrides.GlobalEvmEIP1559DynamicFees = null.BoolFrom(test.eip1559)
			key := cltest.MustGenerateRandomKey(t)
			cu := newVRFCoordinatorUniverse(t, key)
			incomingConfs := 2
			app := cltest.NewApplicationWithConfigAndKeyOnSimulatedBlockchain(t, config, cu.backend, key)
			require.NoError(t, app.Start(testutils.Context(t)))

			jb, vrfKey := createVRFJobRegisterKey(t, cu, app, incomingConfs)
			require.NoError(t, app.JobSpawner().CreateJob(&jb))

			_, err := cu.consumerContract.TestRequestRandomness(cu.carol,
				vrfKey.PublicKey.MustHash(), big.NewInt(100))
			require.NoError(t, err)
			cu.backend.Commit()
			t.Log("Sent test request")
			// Mine the required number of blocks
			// So our request gets confirmed.
			for i := 0; i < incomingConfs; i++ {
				cu.backend.Commit()
			}
			var runs []pipeline.Run
			gomega.NewWithT(t).Eventually(func() bool {
				runs, err = app.PipelineORM().GetAllRuns()
				require.NoError(t, err)
				// It possible that we send the test request
				// before the job spawner has started the vrf services, which is fine
				// the lb will backfill the logs. However we need to
				// keep blocks coming in for the lb to send the backfilled logs.
				cu.backend.Commit()
				return len(runs) == 1 && runs[0].State == pipeline.RunStatusCompleted
			}, cltest.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue())
			assert.Equal(t, pipeline.RunErrors([]null.String{{}}), runs[0].FatalErrors)
			assert.Equal(t, 4, len(runs[0].PipelineTaskRuns))
			assert.NotNil(t, 0, runs[0].Outputs.Val)

			// Ensure the eth transaction gets confirmed on chain.
			gomega.NewWithT(t).Eventually(func() bool {
				q := pg.NewQ(app.GetSqlxDB(), app.GetLogger(), app.GetConfig())
				uc, err2 := txmgr.CountUnconfirmedTransactions(q, key.Address.Address(), cltest.FixtureChainID)
				require.NoError(t, err2)
				return uc == 0
			}, cltest.WaitTimeout(t), 100*time.Millisecond).Should(gomega.BeTrue())

			// Assert the request was fulfilled on-chain.
			gomega.NewWithT(t).Eventually(func() bool {
				rfIterator, err := cu.rootContract.FilterRandomnessRequestFulfilled(nil)
				require.NoError(t, err, "failed to subscribe to RandomnessRequest logs")
				var rf []*solidity_vrf_coordinator_interface.VRFCoordinatorRandomnessRequestFulfilled
				for rfIterator.Next() {
					rf = append(rf, rfIterator.Event)
				}
				return len(rf) == 1
			}, cltest.WaitTimeout(t), 500*time.Millisecond).Should(gomega.BeTrue())
		})
	}
}

func TestIntegration_VRF_WithBHS(t *testing.T) {
	config, _ := heavyweight.FullTestDB(t, "vrf_with_bhs", true, true)
	config.Overrides.GlobalEvmEIP1559DynamicFees = null.BoolFrom(true)
	key := cltest.MustGenerateRandomKey(t)
	cu := newVRFCoordinatorUniverse(t, key)
	incomingConfs := 2
	config.Overrides.BlockBackfillDepth = null.IntFrom(500)
	app := cltest.NewApplicationWithConfigAndKeyOnSimulatedBlockchain(t, config, cu.backend, key)
	require.NoError(t, app.Start(testutils.Context(t)))

	// Create VRF job but do not start it yet
	jb, vrfKey := createVRFJobRegisterKey(t, cu, app, incomingConfs)

	// Create BHS job and start it
	_ = createAndStartBHSJob(t, key.Address.String(), app, cu.bhsContractAddress.String(),
		cu.rootContractAddress.String(), "")

	// Create a VRF request
	_, err := cu.consumerContract.TestRequestRandomness(cu.carol,
		vrfKey.PublicKey.MustHash(), big.NewInt(100))
	require.NoError(t, err)

	cu.backend.Commit()
	requestBlock := cu.backend.Blockchain().CurrentHeader().Number

	// Wait 101 blocks.
	for i := 0; i < 100; i++ {
		cu.backend.Commit()
	}

	// Wait for the blockhash to be stored
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		cu.backend.Commit()
		_, err := cu.bhsContract.GetBlockhash(&bind.CallOpts{
			Pending:     false,
			From:        common.Address{},
			BlockNumber: nil,
			Context:     nil,
		}, requestBlock)
		if err == nil {
			return true
		} else if strings.Contains(err.Error(), "execution reverted") {
			return false
		} else {
			t.Fatal(err)
			return false
		}
	}, cltest.WaitTimeout(t), time.Second).Should(gomega.BeTrue())

	// Wait another 160 blocks so that the request is outside the 256 block window
	for i := 0; i < 160; i++ {
		cu.backend.Commit()
	}

	// Start the VRF job and wait until it's processed
	require.NoError(t, app.JobSpawner().CreateJob(&jb))

	var runs []pipeline.Run
	gomega.NewWithT(t).Eventually(func() bool {
		runs, err = app.PipelineORM().GetAllRuns()
		require.NoError(t, err)
		cu.backend.Commit()
		return len(runs) == 1 && runs[0].State == pipeline.RunStatusCompleted
	}, 10*time.Second, 1*time.Second).Should(gomega.BeTrue())
	assert.Equal(t, pipeline.RunErrors([]null.String{{}}), runs[0].FatalErrors)
	assert.Equal(t, 4, len(runs[0].PipelineTaskRuns))
	assert.NotNil(t, 0, runs[0].Outputs.Val)

	// Ensure the eth transaction gets confirmed on chain.
	gomega.NewWithT(t).Eventually(func() bool {
		q := pg.NewQ(app.GetSqlxDB(), app.GetLogger(), app.GetConfig())
		uc, err2 := txmgr.CountUnconfirmedTransactions(q, key.Address.Address(), cltest.FixtureChainID)
		require.NoError(t, err2)
		return uc == 0
	}, 5*time.Second, 100*time.Millisecond).Should(gomega.BeTrue())

	// Assert the request was fulfilled on-chain.
	gomega.NewWithT(t).Eventually(func() bool {
		rfIterator, err := cu.rootContract.FilterRandomnessRequestFulfilled(nil)
		require.NoError(t, err, "failed to subscribe to RandomnessRequest logs")
		var rf []*solidity_vrf_coordinator_interface.VRFCoordinatorRandomnessRequestFulfilled
		for rfIterator.Next() {
			rf = append(rf, rfIterator.Event)
		}
		return len(rf) == 1
	}, 5*time.Second, 500*time.Millisecond).Should(gomega.BeTrue())
}

func createVRFJobRegisterKey(t *testing.T, u coordinatorUniverse, app *cltest.TestApplication, incomingConfs int) (job.Job, vrfkey.KeyV2) {
	vrfKey, err := app.KeyStore.VRF().Create()
	require.NoError(t, err)

	jid := uuid.FromStringOrNil("96a8a26f-d426-4784-8d8f-fb387d4d8345")
	expectedOnChainJobID, err := hex.DecodeString("3936613861323666643432363437383438643866666233383764346438333435")
	require.NoError(t, err)
	s := testspecs.GenerateVRFSpec(testspecs.VRFSpecParams{
		JobID:                    jid.String(),
		Name:                     "vrf-primary",
		CoordinatorAddress:       u.rootContractAddress.String(),
		MinIncomingConfirmations: incomingConfs,
		PublicKey:                vrfKey.PublicKey.String()}).Toml()
	jb, err := vrf.ValidatedVRFSpec(s)
	require.NoError(t, err)
	assert.Equal(t, expectedOnChainJobID, jb.ExternalIDEncodeStringToTopic().Bytes())

	p, err := vrfKey.PublicKey.Point()
	require.NoError(t, err)
	_, err = u.rootContract.RegisterProvingKey(
		u.neil, big.NewInt(7), u.neil.From, pair(secp256k1.Coordinates(p)), jb.ExternalIDEncodeStringToTopic())
	require.NoError(t, err)
	u.backend.Commit()
	return jb, vrfKey
}
