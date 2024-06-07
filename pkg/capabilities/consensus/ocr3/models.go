package ocr3

import (
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

type config struct {
	AggregationMethod string      `mapstructure:"aggregation_method" json:"aggregation_method" jsonschema:"enum=data_feeds"`
	AggregationConfig *values.Map `mapstructure:"aggregation_config" json:"aggregation_config"`
	Encoder           string      `mapstructure:"encoder" json:"encoder"`
	EncoderConfig     *values.Map `mapstructure:"encoder_config" json:"encoder_config"`
	ReportID          string      `mapstructure:"report_id" json:"report_id" jsonschema:"required,pattern=^[a-f0-9]{4}$"`
}

type inputs struct {
	Observations *values.List `json:"observations"`
}

type outputs struct {
	WorkflowExecutionID string
	capabilities.CapabilityResponse
}

type request struct {
	Observations *values.List `mapstructure:"-"`
	ExpiresAt    time.Time

	// CallbackCh is a channel to send a response back to the requester
	// after the request has been processed or timed out.
	CallbackCh chan capabilities.CapabilityResponse
	StopCh     services.StopChan

	WorkflowExecutionID string
	WorkflowID          string
	WorkflowOwner       string
	WorkflowName        string
	WorkflowDonID       string
	ReportID            string
}