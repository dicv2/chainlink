package models

import (
	"log"
	"math/big"
	"path"
	"reflect"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-go/utils"
)

type ORM struct {
	*storm.DB
}

func NewORM(dir string) *ORM {
	path := path.Join(dir, "db.bolt")
	orm := &ORM{initializeDatabase(path)}
	orm.migrate()
	return orm
}

func initializeDatabase(path string) *storm.DB {
	db, err := storm.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (self *ORM) Where(field string, value interface{}, instance interface{}) error {
	err := self.Find(field, value, instance)
	if err == storm.ErrNotFound {
		emptySlice(instance)
		return nil
	}
	return err
}

func (self *ORM) InitBucket(model interface{}) error {
	return self.Init(model)
}

func (self *ORM) JobsWithCron() ([]Job, error) {
	initrs := []Initiator{}
	self.Where("Type", "cron", &initrs)
	jobIDs := []string{}
	for _, initr := range initrs {
		jobIDs = append(jobIDs, initr.JobID)
	}
	jobs := []Job{}
	err := self.Select(q.In("ID", jobIDs)).Find(&jobs)
	if err == storm.ErrNotFound {
		return jobs, nil
	}

	return jobs, err
}

func (self *ORM) JobRunsFor(job Job) ([]JobRun, error) {
	var runs []JobRun
	err := self.Where("JobID", job.ID, &runs)
	return runs, err
}

func emptySlice(to interface{}) {
	ref := reflect.ValueOf(to)
	results := reflect.MakeSlice(reflect.Indirect(ref).Type(), 0, 0)
	reflect.Indirect(ref).Set(results)
}

func (self *ORM) SaveJob(job Job) error {
	tx, err := self.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := tx.Save(&job); err != nil {
		return err
	}
	for _, initr := range job.Initiators {
		initr.JobID = job.ID
		if err := tx.Save(&initr); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (self *ORM) PendingJobRuns() ([]JobRun, error) {
	var runs []JobRun
	err := self.Where("Status", "pending", &runs)
	return runs, err
}

func (self *ORM) CreateTx(
	from common.Address,
	nonce uint64,
	to common.Address,
	data []byte,
	value *big.Int,
	gasLimit *big.Int,
) (*Tx, error) {
	tx := Tx{
		From:     from,
		To:       to,
		Nonce:    nonce,
		Data:     data,
		Value:    value,
		GasLimit: gasLimit,
	}
	return &tx, self.Save(&tx)
}

func (self *ORM) ConfirmTx(tx *Tx, txat *TxAttempt) error {
	dbtx, err := self.Begin(true)
	if err != nil {
		return err
	}
	defer dbtx.Rollback()

	txat.Confirmed = true
	tx.TxAttempt = *txat
	if err := dbtx.Save(tx); err != nil {
		return err
	}
	if err := dbtx.Save(txat); err != nil {
		return err
	}
	return dbtx.Commit()
}

func (self *ORM) AttemptsFor(id uint64) ([]*TxAttempt, error) {
	attempts := []*TxAttempt{}
	if err := self.Where("TxID", id, &attempts); err != nil {
		return attempts, err
	}
	return attempts, nil
}

func (self *ORM) AddAttempt(
	tx *Tx,
	etx *types.Transaction,
	blkNum uint64,
) (*TxAttempt, error) {
	hex, err := utils.EncodeTxToHex(etx)
	if err != nil {
		return nil, err
	}
	attempt := &TxAttempt{
		Hash:     etx.Hash(),
		GasPrice: etx.GasPrice(),
		Hex:      hex,
		TxID:     tx.ID,
		SentAt:   blkNum,
	}
	if !tx.Confirmed {
		tx.TxAttempt = *attempt
	}
	dbtx, err := self.Begin(true)
	if err != nil {
		return nil, err
	}
	defer dbtx.Rollback()
	if err = dbtx.Save(tx); err != nil {
		return nil, err
	}
	if err = dbtx.Save(attempt); err != nil {
		return nil, err
	}

	return attempt, dbtx.Commit()
}
