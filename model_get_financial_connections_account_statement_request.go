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

// checks if the GetFinancialConnectionsAccountStatementRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinancialConnectionsAccountStatementRequest{}

// GetFinancialConnectionsAccountStatementRequest struct for GetFinancialConnectionsAccountStatementRequest
type GetFinancialConnectionsAccountStatementRequest struct {
	// Access token for authentication
	AccessToken string `json:"access_token"`
	// The remote account id to retrieve the statement for.
	RemoteAccountId string `json:"remote_account_id"`
	// The year and month for the account statement to be retrieved in YYYY-MM.
	Date *string `json:"date,omitempty"`
}

// NewGetFinancialConnectionsAccountStatementRequest instantiates a new GetFinancialConnectionsAccountStatementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinancialConnectionsAccountStatementRequest(accessToken string, remoteAccountId string) *GetFinancialConnectionsAccountStatementRequest {
	this := GetFinancialConnectionsAccountStatementRequest{}
	this.AccessToken = accessToken
	this.RemoteAccountId = remoteAccountId
	return &this
}

// NewGetFinancialConnectionsAccountStatementRequestWithDefaults instantiates a new GetFinancialConnectionsAccountStatementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinancialConnectionsAccountStatementRequestWithDefaults() *GetFinancialConnectionsAccountStatementRequest {
	this := GetFinancialConnectionsAccountStatementRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *GetFinancialConnectionsAccountStatementRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *GetFinancialConnectionsAccountStatementRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *GetFinancialConnectionsAccountStatementRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetRemoteAccountId returns the RemoteAccountId field value
func (o *GetFinancialConnectionsAccountStatementRequest) GetRemoteAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteAccountId
}

// GetRemoteAccountIdOk returns a tuple with the RemoteAccountId field value
// and a boolean to check if the value has been set.
func (o *GetFinancialConnectionsAccountStatementRequest) GetRemoteAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteAccountId, true
}

// SetRemoteAccountId sets field value
func (o *GetFinancialConnectionsAccountStatementRequest) SetRemoteAccountId(v string) {
	o.RemoteAccountId = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *GetFinancialConnectionsAccountStatementRequest) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinancialConnectionsAccountStatementRequest) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *GetFinancialConnectionsAccountStatementRequest) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *GetFinancialConnectionsAccountStatementRequest) SetDate(v string) {
	o.Date = &v
}

func (o GetFinancialConnectionsAccountStatementRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinancialConnectionsAccountStatementRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["remote_account_id"] = o.RemoteAccountId
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return toSerialize, nil
}

type NullableGetFinancialConnectionsAccountStatementRequest struct {
	value *GetFinancialConnectionsAccountStatementRequest
	isSet bool
}

func (v NullableGetFinancialConnectionsAccountStatementRequest) Get() *GetFinancialConnectionsAccountStatementRequest {
	return v.value
}

func (v *NullableGetFinancialConnectionsAccountStatementRequest) Set(val *GetFinancialConnectionsAccountStatementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinancialConnectionsAccountStatementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinancialConnectionsAccountStatementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinancialConnectionsAccountStatementRequest(val *GetFinancialConnectionsAccountStatementRequest) *NullableGetFinancialConnectionsAccountStatementRequest {
	return &NullableGetFinancialConnectionsAccountStatementRequest{value: val, isSet: true}
}

func (v NullableGetFinancialConnectionsAccountStatementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinancialConnectionsAccountStatementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


