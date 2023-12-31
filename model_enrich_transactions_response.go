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

// checks if the EnrichTransactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrichTransactionsResponse{}

// EnrichTransactionsResponse struct for EnrichTransactionsResponse
type EnrichTransactionsResponse struct {
	// The enriched transactions.
	EnrichedTransactions []EnrichedTransaction `json:"enriched_transactions,omitempty"`
}

// NewEnrichTransactionsResponse instantiates a new EnrichTransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrichTransactionsResponse() *EnrichTransactionsResponse {
	this := EnrichTransactionsResponse{}
	return &this
}

// NewEnrichTransactionsResponseWithDefaults instantiates a new EnrichTransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrichTransactionsResponseWithDefaults() *EnrichTransactionsResponse {
	this := EnrichTransactionsResponse{}
	return &this
}

// GetEnrichedTransactions returns the EnrichedTransactions field value if set, zero value otherwise.
func (o *EnrichTransactionsResponse) GetEnrichedTransactions() []EnrichedTransaction {
	if o == nil || IsNil(o.EnrichedTransactions) {
		var ret []EnrichedTransaction
		return ret
	}
	return o.EnrichedTransactions
}

// GetEnrichedTransactionsOk returns a tuple with the EnrichedTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrichTransactionsResponse) GetEnrichedTransactionsOk() ([]EnrichedTransaction, bool) {
	if o == nil || IsNil(o.EnrichedTransactions) {
		return nil, false
	}
	return o.EnrichedTransactions, true
}

// HasEnrichedTransactions returns a boolean if a field has been set.
func (o *EnrichTransactionsResponse) HasEnrichedTransactions() bool {
	if o != nil && !IsNil(o.EnrichedTransactions) {
		return true
	}

	return false
}

// SetEnrichedTransactions gets a reference to the given []EnrichedTransaction and assigns it to the EnrichedTransactions field.
func (o *EnrichTransactionsResponse) SetEnrichedTransactions(v []EnrichedTransaction) {
	o.EnrichedTransactions = v
}

func (o EnrichTransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrichTransactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnrichedTransactions) {
		toSerialize["enriched_transactions"] = o.EnrichedTransactions
	}
	return toSerialize, nil
}

type NullableEnrichTransactionsResponse struct {
	value *EnrichTransactionsResponse
	isSet bool
}

func (v NullableEnrichTransactionsResponse) Get() *EnrichTransactionsResponse {
	return v.value
}

func (v *NullableEnrichTransactionsResponse) Set(val *EnrichTransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrichTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrichTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrichTransactionsResponse(val *EnrichTransactionsResponse) *NullableEnrichTransactionsResponse {
	return &NullableEnrichTransactionsResponse{value: val, isSet: true}
}

func (v NullableEnrichTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrichTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


