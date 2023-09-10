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

// checks if the MigrateFinancialConnectionsTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrateFinancialConnectionsTokenResponse{}

// MigrateFinancialConnectionsTokenResponse struct for MigrateFinancialConnectionsTokenResponse
type MigrateFinancialConnectionsTokenResponse struct {
	ConnectionData MigrateFinancialConnectionsAggregatorConnectionData `json:"connection_data"`
	// Fuse access token for the fuse connection
	FuseAccessToken string `json:"fuse_access_token"`
	// Financial connection id for the fuse connection
	FuseFinancialConnectionId string `json:"fuse_financial_connection_id"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId *string `json:"request_id,omitempty"`
}

// NewMigrateFinancialConnectionsTokenResponse instantiates a new MigrateFinancialConnectionsTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrateFinancialConnectionsTokenResponse(connectionData MigrateFinancialConnectionsAggregatorConnectionData, fuseAccessToken string, fuseFinancialConnectionId string) *MigrateFinancialConnectionsTokenResponse {
	this := MigrateFinancialConnectionsTokenResponse{}
	this.ConnectionData = connectionData
	this.FuseAccessToken = fuseAccessToken
	this.FuseFinancialConnectionId = fuseFinancialConnectionId
	return &this
}

// NewMigrateFinancialConnectionsTokenResponseWithDefaults instantiates a new MigrateFinancialConnectionsTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrateFinancialConnectionsTokenResponseWithDefaults() *MigrateFinancialConnectionsTokenResponse {
	this := MigrateFinancialConnectionsTokenResponse{}
	return &this
}

// GetConnectionData returns the ConnectionData field value
func (o *MigrateFinancialConnectionsTokenResponse) GetConnectionData() MigrateFinancialConnectionsAggregatorConnectionData {
	if o == nil {
		var ret MigrateFinancialConnectionsAggregatorConnectionData
		return ret
	}

	return o.ConnectionData
}

// GetConnectionDataOk returns a tuple with the ConnectionData field value
// and a boolean to check if the value has been set.
func (o *MigrateFinancialConnectionsTokenResponse) GetConnectionDataOk() (*MigrateFinancialConnectionsAggregatorConnectionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionData, true
}

// SetConnectionData sets field value
func (o *MigrateFinancialConnectionsTokenResponse) SetConnectionData(v MigrateFinancialConnectionsAggregatorConnectionData) {
	o.ConnectionData = v
}

// GetFuseAccessToken returns the FuseAccessToken field value
func (o *MigrateFinancialConnectionsTokenResponse) GetFuseAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FuseAccessToken
}

// GetFuseAccessTokenOk returns a tuple with the FuseAccessToken field value
// and a boolean to check if the value has been set.
func (o *MigrateFinancialConnectionsTokenResponse) GetFuseAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FuseAccessToken, true
}

// SetFuseAccessToken sets field value
func (o *MigrateFinancialConnectionsTokenResponse) SetFuseAccessToken(v string) {
	o.FuseAccessToken = v
}

// GetFuseFinancialConnectionId returns the FuseFinancialConnectionId field value
func (o *MigrateFinancialConnectionsTokenResponse) GetFuseFinancialConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FuseFinancialConnectionId
}

// GetFuseFinancialConnectionIdOk returns a tuple with the FuseFinancialConnectionId field value
// and a boolean to check if the value has been set.
func (o *MigrateFinancialConnectionsTokenResponse) GetFuseFinancialConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FuseFinancialConnectionId, true
}

// SetFuseFinancialConnectionId sets field value
func (o *MigrateFinancialConnectionsTokenResponse) SetFuseFinancialConnectionId(v string) {
	o.FuseFinancialConnectionId = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *MigrateFinancialConnectionsTokenResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateFinancialConnectionsTokenResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *MigrateFinancialConnectionsTokenResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *MigrateFinancialConnectionsTokenResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o MigrateFinancialConnectionsTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrateFinancialConnectionsTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connection_data"] = o.ConnectionData
	toSerialize["fuse_access_token"] = o.FuseAccessToken
	toSerialize["fuse_financial_connection_id"] = o.FuseFinancialConnectionId
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableMigrateFinancialConnectionsTokenResponse struct {
	value *MigrateFinancialConnectionsTokenResponse
	isSet bool
}

func (v NullableMigrateFinancialConnectionsTokenResponse) Get() *MigrateFinancialConnectionsTokenResponse {
	return v.value
}

func (v *NullableMigrateFinancialConnectionsTokenResponse) Set(val *MigrateFinancialConnectionsTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrateFinancialConnectionsTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrateFinancialConnectionsTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrateFinancialConnectionsTokenResponse(val *MigrateFinancialConnectionsTokenResponse) *NullableMigrateFinancialConnectionsTokenResponse {
	return &NullableMigrateFinancialConnectionsTokenResponse{value: val, isSet: true}
}

func (v NullableMigrateFinancialConnectionsTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrateFinancialConnectionsTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


