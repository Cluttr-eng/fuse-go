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

// checks if the GetConsumerRiskReportCustomizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConsumerRiskReportCustomizationResponse{}

// GetConsumerRiskReportCustomizationResponse struct for GetConsumerRiskReportCustomizationResponse
type GetConsumerRiskReportCustomizationResponse struct {
	ConsumerRiskReportCustomization ConsumerRiskReportCustomization `json:"consumer_risk_report_customization"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId string `json:"request_id"`
}

// NewGetConsumerRiskReportCustomizationResponse instantiates a new GetConsumerRiskReportCustomizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConsumerRiskReportCustomizationResponse(consumerRiskReportCustomization ConsumerRiskReportCustomization, requestId string) *GetConsumerRiskReportCustomizationResponse {
	this := GetConsumerRiskReportCustomizationResponse{}
	this.ConsumerRiskReportCustomization = consumerRiskReportCustomization
	this.RequestId = requestId
	return &this
}

// NewGetConsumerRiskReportCustomizationResponseWithDefaults instantiates a new GetConsumerRiskReportCustomizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConsumerRiskReportCustomizationResponseWithDefaults() *GetConsumerRiskReportCustomizationResponse {
	this := GetConsumerRiskReportCustomizationResponse{}
	return &this
}

// GetConsumerRiskReportCustomization returns the ConsumerRiskReportCustomization field value
func (o *GetConsumerRiskReportCustomizationResponse) GetConsumerRiskReportCustomization() ConsumerRiskReportCustomization {
	if o == nil {
		var ret ConsumerRiskReportCustomization
		return ret
	}

	return o.ConsumerRiskReportCustomization
}

// GetConsumerRiskReportCustomizationOk returns a tuple with the ConsumerRiskReportCustomization field value
// and a boolean to check if the value has been set.
func (o *GetConsumerRiskReportCustomizationResponse) GetConsumerRiskReportCustomizationOk() (*ConsumerRiskReportCustomization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerRiskReportCustomization, true
}

// SetConsumerRiskReportCustomization sets field value
func (o *GetConsumerRiskReportCustomizationResponse) SetConsumerRiskReportCustomization(v ConsumerRiskReportCustomization) {
	o.ConsumerRiskReportCustomization = v
}

// GetRequestId returns the RequestId field value
func (o *GetConsumerRiskReportCustomizationResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *GetConsumerRiskReportCustomizationResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *GetConsumerRiskReportCustomizationResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o GetConsumerRiskReportCustomizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConsumerRiskReportCustomizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["consumer_risk_report_customization"] = o.ConsumerRiskReportCustomization
	toSerialize["request_id"] = o.RequestId
	return toSerialize, nil
}

type NullableGetConsumerRiskReportCustomizationResponse struct {
	value *GetConsumerRiskReportCustomizationResponse
	isSet bool
}

func (v NullableGetConsumerRiskReportCustomizationResponse) Get() *GetConsumerRiskReportCustomizationResponse {
	return v.value
}

func (v *NullableGetConsumerRiskReportCustomizationResponse) Set(val *GetConsumerRiskReportCustomizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConsumerRiskReportCustomizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConsumerRiskReportCustomizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConsumerRiskReportCustomizationResponse(val *GetConsumerRiskReportCustomizationResponse) *NullableGetConsumerRiskReportCustomizationResponse {
	return &NullableGetConsumerRiskReportCustomizationResponse{value: val, isSet: true}
}

func (v NullableGetConsumerRiskReportCustomizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConsumerRiskReportCustomizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


