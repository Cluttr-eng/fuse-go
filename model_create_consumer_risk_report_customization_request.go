/*
Fuse

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fuse

import (
	"encoding/json"
)

// checks if the CreateConsumerRiskReportCustomizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConsumerRiskReportCustomizationRequest{}

// CreateConsumerRiskReportCustomizationRequest struct for CreateConsumerRiskReportCustomizationRequest
type CreateConsumerRiskReportCustomizationRequest struct {
	Timeframe ConsumerRiskReportTimeFrame `json:"timeframe"`
	// The minimum allowed limit, in cents.
	MinLimit float32 `json:"min_limit"`
	// The maximum allowed limit, in cents.
	MaxLimit float32 `json:"max_limit"`
	// This parameter indicates the risk tolerance associated with spend limits. A high risk tolerance allow for higher limits, increasing both potential gains and losses. A Lower risk tolerance enforces strict limits, reducing the potential for loss but also limiting transaction volume for reliable users.
	RiskTolerance float32 `json:"risk_tolerance"`
}

// NewCreateConsumerRiskReportCustomizationRequest instantiates a new CreateConsumerRiskReportCustomizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConsumerRiskReportCustomizationRequest(timeframe ConsumerRiskReportTimeFrame, minLimit float32, maxLimit float32, riskTolerance float32) *CreateConsumerRiskReportCustomizationRequest {
	this := CreateConsumerRiskReportCustomizationRequest{}
	this.Timeframe = timeframe
	this.MinLimit = minLimit
	this.MaxLimit = maxLimit
	this.RiskTolerance = riskTolerance
	return &this
}

// NewCreateConsumerRiskReportCustomizationRequestWithDefaults instantiates a new CreateConsumerRiskReportCustomizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConsumerRiskReportCustomizationRequestWithDefaults() *CreateConsumerRiskReportCustomizationRequest {
	this := CreateConsumerRiskReportCustomizationRequest{}
	return &this
}

// GetTimeframe returns the Timeframe field value
func (o *CreateConsumerRiskReportCustomizationRequest) GetTimeframe() ConsumerRiskReportTimeFrame {
	if o == nil {
		var ret ConsumerRiskReportTimeFrame
		return ret
	}

	return o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
func (o *CreateConsumerRiskReportCustomizationRequest) GetTimeframeOk() (*ConsumerRiskReportTimeFrame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframe, true
}

// SetTimeframe sets field value
func (o *CreateConsumerRiskReportCustomizationRequest) SetTimeframe(v ConsumerRiskReportTimeFrame) {
	o.Timeframe = v
}

// GetMinLimit returns the MinLimit field value
func (o *CreateConsumerRiskReportCustomizationRequest) GetMinLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinLimit
}

// GetMinLimitOk returns a tuple with the MinLimit field value
// and a boolean to check if the value has been set.
func (o *CreateConsumerRiskReportCustomizationRequest) GetMinLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinLimit, true
}

// SetMinLimit sets field value
func (o *CreateConsumerRiskReportCustomizationRequest) SetMinLimit(v float32) {
	o.MinLimit = v
}

// GetMaxLimit returns the MaxLimit field value
func (o *CreateConsumerRiskReportCustomizationRequest) GetMaxLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value
// and a boolean to check if the value has been set.
func (o *CreateConsumerRiskReportCustomizationRequest) GetMaxLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLimit, true
}

// SetMaxLimit sets field value
func (o *CreateConsumerRiskReportCustomizationRequest) SetMaxLimit(v float32) {
	o.MaxLimit = v
}

// GetRiskTolerance returns the RiskTolerance field value
func (o *CreateConsumerRiskReportCustomizationRequest) GetRiskTolerance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RiskTolerance
}

// GetRiskToleranceOk returns a tuple with the RiskTolerance field value
// and a boolean to check if the value has been set.
func (o *CreateConsumerRiskReportCustomizationRequest) GetRiskToleranceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskTolerance, true
}

// SetRiskTolerance sets field value
func (o *CreateConsumerRiskReportCustomizationRequest) SetRiskTolerance(v float32) {
	o.RiskTolerance = v
}

func (o CreateConsumerRiskReportCustomizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateConsumerRiskReportCustomizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeframe"] = o.Timeframe
	toSerialize["min_limit"] = o.MinLimit
	toSerialize["max_limit"] = o.MaxLimit
	toSerialize["risk_tolerance"] = o.RiskTolerance
	return toSerialize, nil
}

type NullableCreateConsumerRiskReportCustomizationRequest struct {
	value *CreateConsumerRiskReportCustomizationRequest
	isSet bool
}

func (v NullableCreateConsumerRiskReportCustomizationRequest) Get() *CreateConsumerRiskReportCustomizationRequest {
	return v.value
}

func (v *NullableCreateConsumerRiskReportCustomizationRequest) Set(val *CreateConsumerRiskReportCustomizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConsumerRiskReportCustomizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConsumerRiskReportCustomizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConsumerRiskReportCustomizationRequest(val *CreateConsumerRiskReportCustomizationRequest) *NullableCreateConsumerRiskReportCustomizationRequest {
	return &NullableCreateConsumerRiskReportCustomizationRequest{value: val, isSet: true}
}

func (v NullableCreateConsumerRiskReportCustomizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConsumerRiskReportCustomizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


