package models

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/utils"
)

// BridgeTypeRequest is the incoming record used to create a BridgeType
type BridgeTypeRequest struct {
	Name                   TaskType     `json:"name"`
	URL                    WebURL       `json:"url"`
	Confirmations          uint32       `json:"confirmations"`
	MinimumContractPayment *assets.Link `json:"minimumContractPayment"`
}

// GetID returns the ID of this structure for jsonapi serialization.
func (bt BridgeTypeRequest) GetID() string {
	return bt.Name.String()
}

// GetName returns the pluralized "type" of this structure for jsonapi serialization.
func (bt BridgeTypeRequest) GetName() string {
	return "bridges"
}

// SetID is used to set the ID of this structure when deserializing from jsonapi documents.
func (bt *BridgeTypeRequest) SetID(value string) error {
	name, err := NewTaskType(value)
	bt.Name = name
	return err
}

// BridgeTypeAuthentication is the record returned in response to a request to create a BridgeType
type BridgeTypeAuthentication struct {
	Name                   TaskType     `json:"name"`
	URL                    WebURL       `json:"url"`
	Confirmations          uint32       `json:"confirmations"`
	IncomingToken          string       `json:"incomingToken"`
	OutgoingToken          string       `json:"outgoingToken"`
	MinimumContractPayment *assets.Link `json:"minimumContractPayment"`
}

// GetID returns the ID of this structure for jsonapi serialization.
func (bt BridgeTypeAuthentication) GetID() string {
	return bt.Name.String()
}

// GetName returns the pluralized "type" of this structure for jsonapi serialization.
func (bt BridgeTypeAuthentication) GetName() string {
	return "bridges"
}

// SetID is used to set the ID of this structure when deserializing from jsonapi documents.
func (bt *BridgeTypeAuthentication) SetID(value string) error {
	name, err := NewTaskType(value)
	bt.Name = name
	return err
}

// BridgeType is used for external adapters and has fields for
// the name of the adapter and its URL.
type BridgeType struct {
	Name                   TaskType     `json:"name" gorm:"primary_key"`
	URL                    WebURL       `json:"url"`
	Confirmations          uint32       `json:"confirmations"`
	IncomingTokenHash      string       `json:"-"`
	Salt                   string       `json:"-"`
	OutgoingToken          string       `json:"outgoingToken"`
	MinimumContractPayment *assets.Link `json:"minimumContractPayment" gorm:"type:varchar(255)"`
	CreatedAt              time.Time    `json:"-"`
	UpdatedAt              time.Time    `json:"-"`
}

// GetID returns the ID of this structure for jsonapi serialization.
func (bt BridgeType) GetID() string {
	return bt.Name.String()
}

// GetName returns the pluralized "type" of this structure for jsonapi serialization.
func (bt BridgeType) GetName() string {
	return "bridges"
}

// SetID is used to set the ID of this structure when deserializing from jsonapi documents.
func (bt *BridgeType) SetID(value string) error {
	name, err := NewTaskType(value)
	bt.Name = name
	return err
}

// NewBridgeType returns a bridge bridge type authentication (with plaintext
// password) and a bridge type (with hashed password, for persisting)
func NewBridgeType(btr *BridgeTypeRequest) (*BridgeTypeAuthentication,
	*BridgeType, error) {
	incomingToken := utils.NewSecret(24)
	outgoingToken := utils.NewSecret(24)
	salt := utils.NewSecret(24)

	hash, err := incomingTokenHash(incomingToken, salt)
	if err != nil {
		return nil, nil, err
	}

	return &BridgeTypeAuthentication{
			Name:                   btr.Name,
			URL:                    btr.URL,
			Confirmations:          btr.Confirmations,
			IncomingToken:          incomingToken,
			OutgoingToken:          outgoingToken,
			MinimumContractPayment: btr.MinimumContractPayment,
		}, &BridgeType{
			Name:                   btr.Name,
			URL:                    btr.URL,
			Confirmations:          btr.Confirmations,
			IncomingTokenHash:      hash,
			Salt:                   salt,
			OutgoingToken:          outgoingToken,
			MinimumContractPayment: btr.MinimumContractPayment,
		}, nil
}

// AuthenticateBridgeType returns true if the passed token matches its
// IncomingToken, or returns false with an error.
func AuthenticateBridgeType(bt *BridgeType, token string) (bool, error) {
	hash, err := incomingTokenHash(token, bt.Salt)
	if err != nil {
		return false, err
	}
	return subtle.ConstantTimeCompare([]byte(hash), []byte(bt.IncomingTokenHash)) == 1, nil
}

func incomingTokenHash(token, salt string) (string, error) {
	input := fmt.Sprintf("%s-%s", token, salt)
	hash, err := utils.Sha256(input)
	if err != nil {
		return "", err
	}
	return hash, nil
}

// NOTE: latestAnswer and updatedAt is the only metadata used.
// Currently market closer adapter and outlier detection depend latestAnswer.
// https://github.com/smartcontractkit/external-adapters-js/tree/f474bd2e2de13ebe5c9dc3df36ebb7018817005e/composite/market-closure
// https://github.com/smartcontractkit/external-adapters-js/tree/5abb8e5ec2024f724fd39122897baa63c3cd0167/composite/outlier-detection
func BridgeMetaData(latestAnswer *big.Int, updatedAt *big.Int) (map[string]interface{}, error) {
	type m struct {
		LatestAnswer *big.Int `json:"latestAnswer"`
		UpdatedAt    *big.Int `json:"updatedAt"` // A unix timestamp
	}
	b, err := json.Marshal(&m{LatestAnswer: latestAnswer, UpdatedAt: updatedAt})
	if err != nil {
		return nil, err
	}
	var mp map[string]interface{}
	err = json.Unmarshal(b, &mp)
	if err != nil {
		return nil, err
	}
	return mp, nil
}
