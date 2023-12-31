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

// checks if the UpdateConsumerRiskReportCustomizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateConsumerRiskReportCustomizationRequest{}

// UpdateConsumerRiskReportCustomizationRequest struct for UpdateConsumerRiskReportCustomizationRequest
type UpdateConsumerRiskReportCustomizationRequest struct {
	Timeframe *ConsumerRiskReportTimeFrame `json:"timeframe,omitempty"`
	// The minimum allowed limit, in cents.
	MinLimit *float32 `json:"min_limit,omitempty"`
	// The maximum allowed limit, in cents.
	MaxLimit *float32 `json:"max_limit,omitempty"`
	// This parameter indicates the risk tolerance associated with spend limits. A high risk tolerance allow for higher limits, increasing both potential gains and losses. A Lower risk tolerance enforces strict limits, reducing the potential for loss but also limiting transaction volume for reliable users.
	RiskTolerance *float32 `json:"risk_tolerance,omitempty"`
}

// NewUpdateConsumerRiskReportCustomizationRequest instantiates a new UpdateConsumerRiskReportCustomizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConsumerRiskReportCustomizationRequest() *UpdateConsumerRiskReportCustomizationRequest {
	this := UpdateConsumerRiskReportCustomizationRequest{}
	return &this
}

// NewUpdateConsumerRiskReportCustomizationRequestWithDefaults instantiates a new UpdateConsumerRiskReportCustomizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConsumerRiskReportCustomizationRequestWithDefaults() *UpdateConsumerRiskReportCustomizationRequest {
	this := UpdateConsumerRiskReportCustomizationRequest{}
	return &this
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *UpdateConsumerRiskReportCustomizationRequest) GetTimeframe() ConsumerRiskReportTimeFrame {
	if o == nil || IsNil(o.Timeframe) {
		var ret ConsumerRiskReportTimeFrame
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConsumerRiskReportCustomizationRequest) GetTimeframeOk() (*ConsumerRiskReportTimeFrame, bool) {
	if o == nil || IsNil(o.Timeframe) {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *UpdateConsumerRiskReportCustomizationRequest) HasTimeframe() bool {
	if o != nil && !IsNil(o.Timeframe) {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given ConsumerRiskReportTimeFrame and assigns it to the Timeframe field.
func (o *UpdateConsumerRiskReportCustomizationRequest) SetTimeframe(v ConsumerRiskReportTimeFrame) {
	o.Timeframe = &v
}

// GetMinLimit returns the MinLimit field value if set, zero value otherwise.
func (o *UpdateConsumerRiskReportCustomizationRequest) GetMinLimit() float32 {
	if o == nil || IsNil(o.MinLimit) {
		var ret float32
		return ret
	}
	return *o.MinLimit
}

// GetMinLimitOk returns a tuple with the MinLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConsumerRiskReportCustomizationRequest) GetMinLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.MinLimit) {
		return nil, false
	}
	return o.MinLimit, true
}

// HasMinLimit returns a boolean if a field has been set.
func (o *UpdateConsumerRiskReportCustomizationRequest) HasMinLimit() bool {
	if o != nil && !IsNil(o.MinLimit) {
		return true
	}

	return false
}

// SetMinLimit gets a reference to the given float32 and assigns it to the MinLimit field.
func (o *UpdateConsumerRiskReportCustomizationRequest) SetMinLimit(v float32) {
	o.MinLimit = &v
}

// GetMaxLimit returns the MaxLimit field value if set, zero value otherwise.
func (o *UpdateConsumerRiskReportCustomizationRequest) GetMaxLimit() float32 {
	if o == nil || IsNil(o.MaxLimit) {
		var ret float32
		return ret
	}
	return *o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConsumerRiskReportCustomizationRequest) GetMaxLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxLimit) {
		return nil, false
	}
	return o.MaxLimit, true
}

// HasMaxLimit returns a boolean if a field has been set.
func (o *UpdateConsumerRiskReportCustomizationRequest) HasMaxLimit() bool {
	if o != nil && !IsNil(o.MaxLimit) {
		return true
	}

	return false
}

// SetMaxLimit gets a reference to the given float32 and assigns it to the MaxLimit field.
func (o *UpdateConsumerRiskReportCustomizationRequest) SetMaxLimit(v float32) {
	o.MaxLimit = &v
}

// GetRiskTolerance returns the RiskTolerance field value if set, zero value otherwise.
func (o *UpdateConsumerRiskReportCustomizationRequest) GetRiskTolerance() float32 {
	if o == nil || IsNil(o.RiskTolerance) {
		var ret float32
		return ret
	}
	return *o.RiskTolerance
}

// GetRiskToleranceOk returns a tuple with the RiskTolerance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConsumerRiskReportCustomizationRequest) GetRiskToleranceOk() (*float32, bool) {
	if o == nil || IsNil(o.RiskTolerance) {
		return nil, false
	}
	return o.RiskTolerance, true
}

// HasRiskTolerance returns a boolean if a field has been set.
func (o *UpdateConsumerRiskReportCustomizationRequest) HasRiskTolerance() bool {
	if o != nil && !IsNil(o.RiskTolerance) {
		return true
	}

	return false
}

// SetRiskTolerance gets a reference to the given float32 and assigns it to the RiskTolerance field.
func (o *UpdateConsumerRiskReportCustomizationRequest) SetRiskTolerance(v float32) {
	o.RiskTolerance = &v
}

func (o UpdateConsumerRiskReportCustomizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateConsumerRiskReportCustomizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timeframe) {
		toSerialize["timeframe"] = o.Timeframe
	}
	if !IsNil(o.MinLimit) {
		toSerialize["min_limit"] = o.MinLimit
	}
	if !IsNil(o.MaxLimit) {
		toSerialize["max_limit"] = o.MaxLimit
	}
	if !IsNil(o.RiskTolerance) {
		toSerialize["risk_tolerance"] = o.RiskTolerance
	}
	return toSerialize, nil
}

type NullableUpdateConsumerRiskReportCustomizationRequest struct {
	value *UpdateConsumerRiskReportCustomizationRequest
	isSet bool
}

func (v NullableUpdateConsumerRiskReportCustomizationRequest) Get() *UpdateConsumerRiskReportCustomizationRequest {
	return v.value
}

func (v *NullableUpdateConsumerRiskReportCustomizationRequest) Set(val *UpdateConsumerRiskReportCustomizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConsumerRiskReportCustomizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConsumerRiskReportCustomizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConsumerRiskReportCustomizationRequest(val *UpdateConsumerRiskReportCustomizationRequest) *NullableUpdateConsumerRiskReportCustomizationRequest {
	return &NullableUpdateConsumerRiskReportCustomizationRequest{value: val, isSet: true}
}

func (v NullableUpdateConsumerRiskReportCustomizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConsumerRiskReportCustomizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


