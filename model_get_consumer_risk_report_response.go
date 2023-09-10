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

// checks if the GetConsumerRiskReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConsumerRiskReportResponse{}

// GetConsumerRiskReportResponse struct for GetConsumerRiskReportResponse
type GetConsumerRiskReportResponse struct {
	ConsumerRiskReport ConsumerRiskReport `json:"consumer_risk_report"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId string `json:"request_id"`
}

// NewGetConsumerRiskReportResponse instantiates a new GetConsumerRiskReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConsumerRiskReportResponse(consumerRiskReport ConsumerRiskReport, requestId string) *GetConsumerRiskReportResponse {
	this := GetConsumerRiskReportResponse{}
	this.ConsumerRiskReport = consumerRiskReport
	this.RequestId = requestId
	return &this
}

// NewGetConsumerRiskReportResponseWithDefaults instantiates a new GetConsumerRiskReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConsumerRiskReportResponseWithDefaults() *GetConsumerRiskReportResponse {
	this := GetConsumerRiskReportResponse{}
	return &this
}

// GetConsumerRiskReport returns the ConsumerRiskReport field value
func (o *GetConsumerRiskReportResponse) GetConsumerRiskReport() ConsumerRiskReport {
	if o == nil {
		var ret ConsumerRiskReport
		return ret
	}

	return o.ConsumerRiskReport
}

// GetConsumerRiskReportOk returns a tuple with the ConsumerRiskReport field value
// and a boolean to check if the value has been set.
func (o *GetConsumerRiskReportResponse) GetConsumerRiskReportOk() (*ConsumerRiskReport, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerRiskReport, true
}

// SetConsumerRiskReport sets field value
func (o *GetConsumerRiskReportResponse) SetConsumerRiskReport(v ConsumerRiskReport) {
	o.ConsumerRiskReport = v
}

// GetRequestId returns the RequestId field value
func (o *GetConsumerRiskReportResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *GetConsumerRiskReportResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *GetConsumerRiskReportResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o GetConsumerRiskReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConsumerRiskReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["consumer_risk_report"] = o.ConsumerRiskReport
	toSerialize["request_id"] = o.RequestId
	return toSerialize, nil
}

type NullableGetConsumerRiskReportResponse struct {
	value *GetConsumerRiskReportResponse
	isSet bool
}

func (v NullableGetConsumerRiskReportResponse) Get() *GetConsumerRiskReportResponse {
	return v.value
}

func (v *NullableGetConsumerRiskReportResponse) Set(val *GetConsumerRiskReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConsumerRiskReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConsumerRiskReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConsumerRiskReportResponse(val *GetConsumerRiskReportResponse) *NullableGetConsumerRiskReportResponse {
	return &NullableGetConsumerRiskReportResponse{value: val, isSet: true}
}

func (v NullableGetConsumerRiskReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConsumerRiskReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


