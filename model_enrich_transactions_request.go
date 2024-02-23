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

// checks if the EnrichTransactionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrichTransactionsRequest{}

// EnrichTransactionsRequest struct for EnrichTransactionsRequest
type EnrichTransactionsRequest struct {
	Transactions []TransactionToEnrich `json:"transactions"`
}

type _EnrichTransactionsRequest EnrichTransactionsRequest

// NewEnrichTransactionsRequest instantiates a new EnrichTransactionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrichTransactionsRequest(transactions []TransactionToEnrich) *EnrichTransactionsRequest {
	this := EnrichTransactionsRequest{}
	this.Transactions = transactions
	return &this
}

// NewEnrichTransactionsRequestWithDefaults instantiates a new EnrichTransactionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrichTransactionsRequestWithDefaults() *EnrichTransactionsRequest {
	this := EnrichTransactionsRequest{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *EnrichTransactionsRequest) GetTransactions() []TransactionToEnrich {
	if o == nil {
		var ret []TransactionToEnrich
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *EnrichTransactionsRequest) GetTransactionsOk() ([]TransactionToEnrich, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *EnrichTransactionsRequest) SetTransactions(v []TransactionToEnrich) {
	o.Transactions = v
}

func (o EnrichTransactionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrichTransactionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

func (o *EnrichTransactionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transactions",
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

	varEnrichTransactionsRequest := _EnrichTransactionsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnrichTransactionsRequest)

	if err != nil {
		return err
	}

	*o = EnrichTransactionsRequest(varEnrichTransactionsRequest)

	return err
}

type NullableEnrichTransactionsRequest struct {
	value *EnrichTransactionsRequest
	isSet bool
}

func (v NullableEnrichTransactionsRequest) Get() *EnrichTransactionsRequest {
	return v.value
}

func (v *NullableEnrichTransactionsRequest) Set(val *EnrichTransactionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrichTransactionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrichTransactionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrichTransactionsRequest(val *EnrichTransactionsRequest) *NullableEnrichTransactionsRequest {
	return &NullableEnrichTransactionsRequest{value: val, isSet: true}
}

func (v NullableEnrichTransactionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrichTransactionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


