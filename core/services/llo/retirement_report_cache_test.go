package llo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func Test_RetirementReportCache(t *testing.T) {
	lggr := logger.TestLogger(t)
	// db := pgtest.NewSqlxDB(t)
	// orm := &retirementReportCacheORM{ds: db}
	rrc := newRetirementReportCache(lggr, nil)
	exampleAttestedRetirementReport := []byte{1, 2, 3, 4, 5}
	exampleDigest := ocrtypes.ConfigDigest{1}

	t.Run("AttestedRetirementReport", func(t *testing.T) {
		attestedRetirementReport, err := rrc.AttestedRetirementReport(exampleDigest)
		assert.NoError(t, err)
		assert.Nil(t, attestedRetirementReport)

		rrc.c[exampleDigest] = exampleAttestedRetirementReport

		attestedRetirementReport, err = rrc.AttestedRetirementReport(exampleDigest)
		assert.NoError(t, err)
		assert.Equal(t, exampleAttestedRetirementReport, attestedRetirementReport)
	})
	t.Run("CheckAttestedRetirementReport", func(t *testing.T) {
		t.Run("invalid", func(t *testing.T) {
			_, err := rrc.CheckAttestedRetirementReport(exampleDigest, []byte("not valid"))
			assert.EqualError(t, err, "foo")
		})
		t.Run("valid", func(t *testing.T) {
			rr, err := rrc.CheckAttestedRetirementReport(exampleDigest, exampleAttestedRetirementReport)
			assert.NoError(t, err)
			assert.Equal(t, "foo", rr)
		})
	})
	t.Run("Store", func(t *testing.T) {
	})
	t.Fatal("TODO")
}

func Test_RetirementReportCache_ORM(t *testing.T) {
	t.Fatal("TODO")
}
