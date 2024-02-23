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

// checks if the SearchFinancialInstitutionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchFinancialInstitutionsRequest{}

// SearchFinancialInstitutionsRequest struct for SearchFinancialInstitutionsRequest
type SearchFinancialInstitutionsRequest struct {
	// The session client secret created.
	SessionClientSecret string `json:"session_client_secret"`
	// The search term, ie wells fargo.
	SearchTerm string `json:"search_term"`
}

type _SearchFinancialInstitutionsRequest SearchFinancialInstitutionsRequest

// NewSearchFinancialInstitutionsRequest instantiates a new SearchFinancialInstitutionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchFinancialInstitutionsRequest(sessionClientSecret string, searchTerm string) *SearchFinancialInstitutionsRequest {
	this := SearchFinancialInstitutionsRequest{}
	this.SessionClientSecret = sessionClientSecret
	this.SearchTerm = searchTerm
	return &this
}

// NewSearchFinancialInstitutionsRequestWithDefaults instantiates a new SearchFinancialInstitutionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchFinancialInstitutionsRequestWithDefaults() *SearchFinancialInstitutionsRequest {
	this := SearchFinancialInstitutionsRequest{}
	return &this
}

// GetSessionClientSecret returns the SessionClientSecret field value
func (o *SearchFinancialInstitutionsRequest) GetSessionClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionClientSecret
}

// GetSessionClientSecretOk returns a tuple with the SessionClientSecret field value
// and a boolean to check if the value has been set.
func (o *SearchFinancialInstitutionsRequest) GetSessionClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionClientSecret, true
}

// SetSessionClientSecret sets field value
func (o *SearchFinancialInstitutionsRequest) SetSessionClientSecret(v string) {
	o.SessionClientSecret = v
}

// GetSearchTerm returns the SearchTerm field value
func (o *SearchFinancialInstitutionsRequest) GetSearchTerm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchTerm
}

// GetSearchTermOk returns a tuple with the SearchTerm field value
// and a boolean to check if the value has been set.
func (o *SearchFinancialInstitutionsRequest) GetSearchTermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchTerm, true
}

// SetSearchTerm sets field value
func (o *SearchFinancialInstitutionsRequest) SetSearchTerm(v string) {
	o.SearchTerm = v
}

func (o SearchFinancialInstitutionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchFinancialInstitutionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session_client_secret"] = o.SessionClientSecret
	toSerialize["search_term"] = o.SearchTerm
	return toSerialize, nil
}

func (o *SearchFinancialInstitutionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"session_client_secret",
		"search_term",
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

	varSearchFinancialInstitutionsRequest := _SearchFinancialInstitutionsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchFinancialInstitutionsRequest)

	if err != nil {
		return err
	}

	*o = SearchFinancialInstitutionsRequest(varSearchFinancialInstitutionsRequest)

	return err
}

type NullableSearchFinancialInstitutionsRequest struct {
	value *SearchFinancialInstitutionsRequest
	isSet bool
}

func (v NullableSearchFinancialInstitutionsRequest) Get() *SearchFinancialInstitutionsRequest {
	return v.value
}

func (v *NullableSearchFinancialInstitutionsRequest) Set(val *SearchFinancialInstitutionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchFinancialInstitutionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchFinancialInstitutionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchFinancialInstitutionsRequest(val *SearchFinancialInstitutionsRequest) *NullableSearchFinancialInstitutionsRequest {
	return &NullableSearchFinancialInstitutionsRequest{value: val, isSet: true}
}

func (v NullableSearchFinancialInstitutionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchFinancialInstitutionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

