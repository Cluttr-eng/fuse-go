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

// checks if the AssetReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetReportResponse{}

// AssetReportResponse struct for AssetReportResponse
type AssetReportResponse struct {
	Report *AssetReport `json:"report,omitempty"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId *string `json:"request_id,omitempty"`
}

// NewAssetReportResponse instantiates a new AssetReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportResponse() *AssetReportResponse {
	this := AssetReportResponse{}
	return &this
}

// NewAssetReportResponseWithDefaults instantiates a new AssetReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportResponseWithDefaults() *AssetReportResponse {
	this := AssetReportResponse{}
	return &this
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *AssetReportResponse) GetReport() AssetReport {
	if o == nil || IsNil(o.Report) {
		var ret AssetReport
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportResponse) GetReportOk() (*AssetReport, bool) {
	if o == nil || IsNil(o.Report) {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *AssetReportResponse) HasReport() bool {
	if o != nil && !IsNil(o.Report) {
		return true
	}

	return false
}

// SetReport gets a reference to the given AssetReport and assigns it to the Report field.
func (o *AssetReportResponse) SetReport(v AssetReport) {
	o.Report = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *AssetReportResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *AssetReportResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *AssetReportResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o AssetReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Report) {
		toSerialize["report"] = o.Report
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableAssetReportResponse struct {
	value *AssetReportResponse
	isSet bool
}

func (v NullableAssetReportResponse) Get() *AssetReportResponse {
	return v.value
}

func (v *NullableAssetReportResponse) Set(val *AssetReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportResponse(val *AssetReportResponse) *NullableAssetReportResponse {
	return &NullableAssetReportResponse{value: val, isSet: true}
}

func (v NullableAssetReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


