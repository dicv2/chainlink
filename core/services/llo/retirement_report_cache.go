package llo

import (
	"context"
	"fmt"
	"sync"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	llotypes "github.com/smartcontractkit/chainlink-common/pkg/types/llo"

	datastreamsllo "github.com/smartcontractkit/chainlink-data-streams/llo"
)

var _ RetirementReportCache = &retirementReportCache{}

type WriteOnlyRetirementReportCache interface {
	Store(ctx context.Context, cd ocr2types.ConfigDigest, retirementReport []byte, sigs []types.AttributedOnchainSignature) error
}

type RetirementReportCache interface {
	datastreamsllo.PredecessorRetirementReportCache
	WriteOnlyRetirementReportCache
}

type RetirementReportCacheService interface {
	services.Service
	RetirementReportCache
}

type RetirementReportCacheORM interface {
	StoreAttestedRetirementReport(ctx context.Context, cd ocr2types.ConfigDigest, attestedRetirementReport []byte) error
	LoadAttestedRetirementReports(ctx context.Context) (map[ocr2types.ConfigDigest][]byte, error)
}

type retirementReportCacheORM struct {
	ds sqlutil.DataSource
}

// TODO: Test ORM
// TODO: Test whole thing
func (o *retirementReportCacheORM) StoreAttestedRetirementReport(ctx context.Context, cd ocr2types.ConfigDigest, attestedRetirementReport []byte) error {
	_, err := o.ds.ExecContext(ctx, `
INSERT INTO retirement_report_cache (config_digest, attested_retirement_report, updated_at)
VALUES ($1, $2, NOW())
ON CONFLICT (config_digest) DO UPDATE
SET attested_retirement_report = $2, updated_at = NOW()
`, cd, attestedRetirementReport)
	if err != nil {
		return fmt.Errorf("StoreAttestedRetirementReport failed: %w", err)
	}
	return nil
}

func (o *retirementReportCacheORM) LoadAttestedRetirementReports(ctx context.Context) (map[ocr2types.ConfigDigest][]byte, error) {
	rows, err := o.ds.QueryContext(ctx, "SELECT config_digest, attested_retirement_report FROM retirement_report_cache")
	if err != nil {
		return nil, fmt.Errorf("LoadAttestedRetirementReports failed: %w", err)
	}
	defer rows.Close()

	reports := make(map[ocr2types.ConfigDigest][]byte)
	for rows.Next() {
		var cd ocr2types.ConfigDigest
		var arr []byte
		if err := rows.Scan(&cd, &arr); err != nil {
			return nil, fmt.Errorf("LoadAttestedRetirementReports failed: %w", err)
		}
		reports[cd] = arr
	}

	return reports, nil
}

// TODO
type RetirementReportVerifier interface {
	Verify(key types.OnchainPublicKey, digest types.ConfigDigest, seqNr uint64, r ocr3types.ReportWithInfo[llotypes.ReportInfo], signature []byte) bool
}

type retirementReportCache struct {
	services.Service
	eng      *services.Engine
	mu       sync.RWMutex
	c        map[ocr2types.ConfigDigest][]byte
	orm      RetirementReportCacheORM
	verifier RetirementReportVerifier
	codec    datastreamsllo.RetirementReportCodec
}

func NewRetirementReportCache(lggr logger.Logger, ds sqlutil.DataSource) RetirementReportCacheService {
	orm := &retirementReportCacheORM{ds: ds}
	return newRetirementReportCache(lggr, orm)
}

func newRetirementReportCache(lggr logger.Logger, orm RetirementReportCacheORM) *retirementReportCache {
	r := &retirementReportCache{
		c:   make(map[ocr2types.ConfigDigest][]byte),
		orm: orm,
	}
	r.Service, r.eng = services.Config{
		Name:  "RetirementReportCache",
		Start: r.start,
	}.NewServiceEngine(lggr)
	return r
}

func (r *retirementReportCache) AttestedRetirementReport(predecessorConfigDigest ocr2types.ConfigDigest) ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.c[predecessorConfigDigest], nil
}

func (r *retirementReportCache) CheckAttestedRetirementReport(predecessorConfigDigest ocr2types.ConfigDigest, attestedRetirementReport []byte) (datastreamsllo.RetirementReport, error) {
	// TODO: verify signatures
	fmt.Println("TRASH", attestedRetirementReport)
	panic("implement me")
}

func (r *retirementReportCache) verify(digest ocr2types.ConfigDigest, serializedAttestedRetirementReport []byte) (report []byte, err error) {
	var attestedRetirementReport AttestedRetirementReport
	if err := proto.Unmarshal(serializedAttestedRetirementReport, &attestedRetirementReport); err != nil {
		return nil, fmt.Errorf("Verify failed; failed to unmarshal protobuf: %w", err)
	}
	// TODO: The config poller needs to persist all configs so that it's accessible here later on
	// authorizedOnchainPublicKeys, f, err := r.getAuthorizedOnchainPublicKeys(digest)
	// if err != nil {
	//     return nil, fmt.Errorf("Verify failed; failed to get authorized keys: %w", err)
	// }
	// for _, sig := range attestedRetirementReport.Sigs {
	//     if !r.verifier.Verify(authorizedKey, digest, 0, ocr3types.ReportWithInfo{}, sig.Signature) {
	//         return nil, fmt.Errorf("Verify failed; signature verification failed")
	//     }
	// }
	panic("TODO")
}

func (r *retirementReportCache) Store(ctx context.Context, cd ocr2types.ConfigDigest, retirementReport []byte, sigs []types.AttributedOnchainSignature) error {
	// TODO: proto-encode?
	var pbSigs []*AttributedOnchainSignature
	for _, s := range sigs {
		pbSigs = append(pbSigs, &AttributedOnchainSignature{
			Signer:    uint32(s.Signer),
			Signature: s.Signature,
		})
	}
	attestedRetirementReport := AttestedRetirementReport{
		RetirementReport: retirementReport,
		Sigs:             pbSigs,
	}

	serialized, err := proto.Marshal(&attestedRetirementReport)
	if err != nil {
		return fmt.Errorf("Store failed; failed to marshal protobuf: %w", err)
	}

	r.mu.Lock()
	r.c[cd] = serialized
	r.mu.Unlock()

	return r.orm.StoreAttestedRetirementReport(ctx, cd, serialized)
}

// NOTE: Could do this lazily instead if we wanted to avoid a performance hit
// on application startup
func (r *retirementReportCache) start(ctx context.Context) error {
	// Load all attested retirement reports from the ORM
	// and store them in the cache
	attestedRetirementReports, err := r.orm.LoadAttestedRetirementReports(ctx)
	if err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.c = attestedRetirementReports
	return nil
}
