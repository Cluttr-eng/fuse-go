/*
Fuse

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fuse

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetFinancialConnectionsAccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinancialConnectionsAccountsResponse{}

// GetFinancialConnectionsAccountsResponse struct for GetFinancialConnectionsAccountsResponse
type GetFinancialConnectionsAccountsResponse struct {
	Accounts []FinancialConnectionsAccount `json:"accounts"`
	FinancialConnection FinancialConnectionData `json:"financial_connection"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId string `json:"request_id"`
}

type _GetFinancialConnectionsAccountsResponse GetFinancialConnectionsAccountsResponse

// NewGetFinancialConnectionsAccountsResponse instantiates a new GetFinancialConnectionsAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinancialConnectionsAccountsResponse(accounts []FinancialConnectionsAccount, financialConnection FinancialConnectionData, requestId string) *GetFinancialConnectionsAccountsResponse {
	this := GetFinancialConnectionsAccountsResponse{}
	this.Accounts = accounts
	this.FinancialConnection = financialConnection
	this.RequestId = requestId
	return &this
}

// NewGetFinancialConnectionsAccountsResponseWithDefaults instantiates a new GetFinancialConnectionsAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinancialConnectionsAccountsResponseWithDefaults() *GetFinancialConnectionsAccountsResponse {
	this := GetFinancialConnectionsAccountsResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *GetFinancialConnectionsAccountsResponse) GetAccounts() []FinancialConnectionsAccount {
	if o == nil {
		var ret []FinancialConnectionsAccount
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *GetFinancialConnectionsAccountsResponse) GetAccountsOk() ([]FinancialConnectionsAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *GetFinancialConnectionsAccountsResponse) SetAccounts(v []FinancialConnectionsAccount) {
	o.Accounts = v
}

// GetFinancialConnection returns the FinancialConnection field value
func (o *GetFinancialConnectionsAccountsResponse) GetFinancialConnection() FinancialConnectionData {
	if o == nil {
		var ret FinancialConnectionData
		return ret
	}

	return o.FinancialConnection
}

// GetFinancialConnectionOk returns a tuple with the FinancialConnection field value
// and a boolean to check if the value has been set.
func (o *GetFinancialConnectionsAccountsResponse) GetFinancialConnectionOk() (*FinancialConnectionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinancialConnection, true
}

// SetFinancialConnection sets field value
func (o *GetFinancialConnectionsAccountsResponse) SetFinancialConnection(v FinancialConnectionData) {
	o.FinancialConnection = v
}

// GetRequestId returns the RequestId field value
func (o *GetFinancialConnectionsAccountsResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *GetFinancialConnectionsAccountsResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *GetFinancialConnectionsAccountsResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o GetFinancialConnectionsAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinancialConnectionsAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
	toSerialize["financial_connection"] = o.FinancialConnection
	toSerialize["request_id"] = o.RequestId
	return toSerialize, nil
}

func (o *GetFinancialConnectionsAccountsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accounts",
		"financial_connection",
		"request_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetFinancialConnectionsAccountsResponse := _GetFinancialConnectionsAccountsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFinancialConnectionsAccountsResponse)

	if err != nil {
		return err
	}

	*o = GetFinancialConnectionsAccountsResponse(varGetFinancialConnectionsAccountsResponse)

	return err
}

type NullableGetFinancialConnectionsAccountsResponse struct {
	value *GetFinancialConnectionsAccountsResponse
	isSet bool
}

func (v NullableGetFinancialConnectionsAccountsResponse) Get() *GetFinancialConnectionsAccountsResponse {
	return v.value
}

func (v *NullableGetFinancialConnectionsAccountsResponse) Set(val *GetFinancialConnectionsAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinancialConnectionsAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinancialConnectionsAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinancialConnectionsAccountsResponse(val *GetFinancialConnectionsAccountsResponse) *NullableGetFinancialConnectionsAccountsResponse {
	return &NullableGetFinancialConnectionsAccountsResponse{value: val, isSet: true}
}

func (v NullableGetFinancialConnectionsAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinancialConnectionsAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


