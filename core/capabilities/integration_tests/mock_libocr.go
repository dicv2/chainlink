package integration_tests

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
)

type node struct {
	ocr3types.ReportingPlugin[[]byte]
	*ocr3.ContractTransmitter
	key ocr2key.KeyBundle
}

// mockLibOCR is a mock libocr implementation for testing purposes that simulates libocr protocol rounds without having
// to setup the libocr network
type mockLibOCR struct {
	nodes []*node
	f     uint8
}

func newMockLibOCR(f uint8) *mockLibOCR {
	return &mockLibOCR{f: f}
}

func (m *mockLibOCR) Start(ctx context.Context, t *testing.T, protocolRoundInterval time.Duration) {
	go func() {
		ticker := time.NewTicker(protocolRoundInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := m.simulateProtocolRound(ctx)
				if err != nil {
					require.FailNow(t, err.Error())
				}
			}
		}
	}()
}

func (m *mockLibOCR) AddNode(plugin ocr3types.ReportingPlugin[[]byte], transmitter *ocr3.ContractTransmitter, key ocr2key.KeyBundle) {
	m.nodes = append(m.nodes, &node{plugin, transmitter, key})
}

func (m *mockLibOCR) simulateProtocolRound(ctx context.Context) error {
	var seqNr uint64

	// randomly select a leader
	leader := m.nodes[rand.Intn(len(m.nodes))]

	outcomeCtx := ocr3types.OutcomeContext{
		SeqNr:           seqNr,
		PreviousOutcome: nil,
		Epoch:           0,
		Round:           0,
	}

	// get the query
	query, err := leader.Query(ctx, outcomeCtx)
	if err != nil {
		return fmt.Errorf("failed to get query: %w", err)
	}

	var observations []types.AttributedObservation
	for oracleID, node := range m.nodes {
		obs, err2 := node.Observation(ctx, outcomeCtx, query)
		if err2 != nil {
			return fmt.Errorf("failed to get observation: %w", err)
		}

		observations = append(observations, types.AttributedObservation{
			Observation: obs,
			Observer:    commontypes.OracleID(oracleID),
		})
	}

	var outcomes []ocr3types.Outcome
	for _, node := range m.nodes {
		outcome, err2 := node.Outcome(outcomeCtx, query, observations)
		if err2 != nil {
			return fmt.Errorf("failed to get outcome: %w", err)
		}

		if len(outcome) == 0 {
			return nil // wait until all nodes have an outcome for testing purposes
		}

		outcomes = append(outcomes, outcome)
	}

	// if all outcomes are equal proceed to reports
	for _, outcome := range outcomes {
		if !bytes.Equal(outcome, outcomes[0]) {
			return nil
		}
	}

	reports, err := leader.Reports(0, outcomes[0])
	if err != nil {
		return fmt.Errorf("failed to get reports: %w", err)
	}
	for _, report := range reports {
		// create signatures
		var signatures []types.AttributedOnchainSignature
		for i, node := range m.nodes {
			sig, err := node.key.Sign(types.ReportContext{}, report.Report)
			if err != nil {
				return fmt.Errorf("failed to sign report: %w", err)
			}

			signatures = append(signatures, types.AttributedOnchainSignature{
				Signer:    commontypes.OracleID(i),
				Signature: sig,
			})

			if uint8(len(signatures)) == m.f+1 {
				break
			}
		}

		for _, node := range m.nodes {
			accept, err := node.ShouldAcceptAttestedReport(ctx, seqNr, report)
			if err != nil {
				return fmt.Errorf("failed to check if report should be accepted: %w", err)
			}
			if !accept {
				continue
			}

			transmit, err := node.ShouldTransmitAcceptedReport(ctx, seqNr, report)
			if err != nil {
				return fmt.Errorf("failed to check if report should be transmitted: %w", err)
			}

			if !transmit {
				continue
			}

			err = node.Transmit(ctx, types.ConfigDigest{}, 0, report, signatures)
			if err != nil {
				return fmt.Errorf("failed to transmit report: %w", err)
			}
		}
	}

	return nil
}
