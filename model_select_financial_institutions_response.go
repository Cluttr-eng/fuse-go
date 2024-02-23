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

// checks if the SelectFinancialInstitutionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectFinancialInstitutionsResponse{}

// SelectFinancialInstitutionsResponse struct for SelectFinancialInstitutionsResponse
type SelectFinancialInstitutionsResponse struct {
	FinancialInstitution FinancialInstitution `json:"financial_institution"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId *string `json:"request_id,omitempty"`
}

type _SelectFinancialInstitutionsResponse SelectFinancialInstitutionsResponse

// NewSelectFinancialInstitutionsResponse instantiates a new SelectFinancialInstitutionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectFinancialInstitutionsResponse(financialInstitution FinancialInstitution) *SelectFinancialInstitutionsResponse {
	this := SelectFinancialInstitutionsResponse{}
	this.FinancialInstitution = financialInstitution
	return &this
}

// NewSelectFinancialInstitutionsResponseWithDefaults instantiates a new SelectFinancialInstitutionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectFinancialInstitutionsResponseWithDefaults() *SelectFinancialInstitutionsResponse {
	this := SelectFinancialInstitutionsResponse{}
	return &this
}

// GetFinancialInstitution returns the FinancialInstitution field value
func (o *SelectFinancialInstitutionsResponse) GetFinancialInstitution() FinancialInstitution {
	if o == nil {
		var ret FinancialInstitution
		return ret
	}

	return o.FinancialInstitution
}

// GetFinancialInstitutionOk returns a tuple with the FinancialInstitution field value
// and a boolean to check if the value has been set.
func (o *SelectFinancialInstitutionsResponse) GetFinancialInstitutionOk() (*FinancialInstitution, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinancialInstitution, true
}

// SetFinancialInstitution sets field value
func (o *SelectFinancialInstitutionsResponse) SetFinancialInstitution(v FinancialInstitution) {
	o.FinancialInstitution = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *SelectFinancialInstitutionsResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectFinancialInstitutionsResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *SelectFinancialInstitutionsResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *SelectFinancialInstitutionsResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o SelectFinancialInstitutionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectFinancialInstitutionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["financial_institution"] = o.FinancialInstitution
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	return toSerialize, nil
}

func (o *SelectFinancialInstitutionsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"financial_institution",
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

	varSelectFinancialInstitutionsResponse := _SelectFinancialInstitutionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSelectFinancialInstitutionsResponse)

	if err != nil {
		return err
	}

	*o = SelectFinancialInstitutionsResponse(varSelectFinancialInstitutionsResponse)

	return err
}

type NullableSelectFinancialInstitutionsResponse struct {
	value *SelectFinancialInstitutionsResponse
	isSet bool
}

func (v NullableSelectFinancialInstitutionsResponse) Get() *SelectFinancialInstitutionsResponse {
	return v.value
}

func (v *NullableSelectFinancialInstitutionsResponse) Set(val *SelectFinancialInstitutionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectFinancialInstitutionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectFinancialInstitutionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectFinancialInstitutionsResponse(val *SelectFinancialInstitutionsResponse) *NullableSelectFinancialInstitutionsResponse {
	return &NullableSelectFinancialInstitutionsResponse{value: val, isSet: true}
}

func (v NullableSelectFinancialInstitutionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectFinancialInstitutionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

