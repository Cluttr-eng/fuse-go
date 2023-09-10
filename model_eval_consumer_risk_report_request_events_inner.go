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

// checks if the EvalConsumerRiskReportRequestEventsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvalConsumerRiskReportRequestEventsInner{}

// EvalConsumerRiskReportRequestEventsInner struct for EvalConsumerRiskReportRequestEventsInner
type EvalConsumerRiskReportRequestEventsInner struct {
	// The id of the account that event belongs to
	AccountId string `json:"account_id"`
	Event AddAccountEventsRequestEventsInner `json:"event"`
}

// NewEvalConsumerRiskReportRequestEventsInner instantiates a new EvalConsumerRiskReportRequestEventsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvalConsumerRiskReportRequestEventsInner(accountId string, event AddAccountEventsRequestEventsInner) *EvalConsumerRiskReportRequestEventsInner {
	this := EvalConsumerRiskReportRequestEventsInner{}
	this.AccountId = accountId
	this.Event = event
	return &this
}

// NewEvalConsumerRiskReportRequestEventsInnerWithDefaults instantiates a new EvalConsumerRiskReportRequestEventsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvalConsumerRiskReportRequestEventsInnerWithDefaults() *EvalConsumerRiskReportRequestEventsInner {
	this := EvalConsumerRiskReportRequestEventsInner{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *EvalConsumerRiskReportRequestEventsInner) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *EvalConsumerRiskReportRequestEventsInner) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *EvalConsumerRiskReportRequestEventsInner) SetAccountId(v string) {
	o.AccountId = v
}

// GetEvent returns the Event field value
func (o *EvalConsumerRiskReportRequestEventsInner) GetEvent() AddAccountEventsRequestEventsInner {
	if o == nil {
		var ret AddAccountEventsRequestEventsInner
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EvalConsumerRiskReportRequestEventsInner) GetEventOk() (*AddAccountEventsRequestEventsInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *EvalConsumerRiskReportRequestEventsInner) SetEvent(v AddAccountEventsRequestEventsInner) {
	o.Event = v
}

func (o EvalConsumerRiskReportRequestEventsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvalConsumerRiskReportRequestEventsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["event"] = o.Event
	return toSerialize, nil
}

type NullableEvalConsumerRiskReportRequestEventsInner struct {
	value *EvalConsumerRiskReportRequestEventsInner
	isSet bool
}

func (v NullableEvalConsumerRiskReportRequestEventsInner) Get() *EvalConsumerRiskReportRequestEventsInner {
	return v.value
}

func (v *NullableEvalConsumerRiskReportRequestEventsInner) Set(val *EvalConsumerRiskReportRequestEventsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEvalConsumerRiskReportRequestEventsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEvalConsumerRiskReportRequestEventsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvalConsumerRiskReportRequestEventsInner(val *EvalConsumerRiskReportRequestEventsInner) *NullableEvalConsumerRiskReportRequestEventsInner {
	return &NullableEvalConsumerRiskReportRequestEventsInner{value: val, isSet: true}
}

func (v NullableEvalConsumerRiskReportRequestEventsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvalConsumerRiskReportRequestEventsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

