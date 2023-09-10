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

// checks if the CreateAssetReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssetReportResponse{}

// CreateAssetReportResponse struct for CreateAssetReportResponse
type CreateAssetReportResponse struct {
	// A token that can be provided to endpoints such as /asset_report to fetch an asset report.
	AssetReportToken *string `json:"asset_report_token,omitempty"`
	// A unique ID identifying an Asset Report.
	AssetReportId *string `json:"asset_report_id,omitempty"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId *string `json:"request_id,omitempty"`
}

// NewCreateAssetReportResponse instantiates a new CreateAssetReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssetReportResponse() *CreateAssetReportResponse {
	this := CreateAssetReportResponse{}
	return &this
}

// NewCreateAssetReportResponseWithDefaults instantiates a new CreateAssetReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssetReportResponseWithDefaults() *CreateAssetReportResponse {
	this := CreateAssetReportResponse{}
	return &this
}

// GetAssetReportToken returns the AssetReportToken field value if set, zero value otherwise.
func (o *CreateAssetReportResponse) GetAssetReportToken() string {
	if o == nil || IsNil(o.AssetReportToken) {
		var ret string
		return ret
	}
	return *o.AssetReportToken
}

// GetAssetReportTokenOk returns a tuple with the AssetReportToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetReportResponse) GetAssetReportTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AssetReportToken) {
		return nil, false
	}
	return o.AssetReportToken, true
}

// HasAssetReportToken returns a boolean if a field has been set.
func (o *CreateAssetReportResponse) HasAssetReportToken() bool {
	if o != nil && !IsNil(o.AssetReportToken) {
		return true
	}

	return false
}

// SetAssetReportToken gets a reference to the given string and assigns it to the AssetReportToken field.
func (o *CreateAssetReportResponse) SetAssetReportToken(v string) {
	o.AssetReportToken = &v
}

// GetAssetReportId returns the AssetReportId field value if set, zero value otherwise.
func (o *CreateAssetReportResponse) GetAssetReportId() string {
	if o == nil || IsNil(o.AssetReportId) {
		var ret string
		return ret
	}
	return *o.AssetReportId
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetReportResponse) GetAssetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetReportId) {
		return nil, false
	}
	return o.AssetReportId, true
}

// HasAssetReportId returns a boolean if a field has been set.
func (o *CreateAssetReportResponse) HasAssetReportId() bool {
	if o != nil && !IsNil(o.AssetReportId) {
		return true
	}

	return false
}

// SetAssetReportId gets a reference to the given string and assigns it to the AssetReportId field.
func (o *CreateAssetReportResponse) SetAssetReportId(v string) {
	o.AssetReportId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateAssetReportResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetReportResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateAssetReportResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateAssetReportResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o CreateAssetReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssetReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetReportToken) {
		toSerialize["asset_report_token"] = o.AssetReportToken
	}
	if !IsNil(o.AssetReportId) {
		toSerialize["asset_report_id"] = o.AssetReportId
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableCreateAssetReportResponse struct {
	value *CreateAssetReportResponse
	isSet bool
}

func (v NullableCreateAssetReportResponse) Get() *CreateAssetReportResponse {
	return v.value
}

func (v *NullableCreateAssetReportResponse) Set(val *CreateAssetReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssetReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssetReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssetReportResponse(val *CreateAssetReportResponse) *NullableCreateAssetReportResponse {
	return &NullableCreateAssetReportResponse{value: val, isSet: true}
}

func (v NullableCreateAssetReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssetReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

