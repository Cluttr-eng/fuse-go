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

// checks if the FinancialConnectionDetailsBelvo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionDetailsBelvo{}

// FinancialConnectionDetailsBelvo Data needed to query data from Belvo
type FinancialConnectionDetailsBelvo struct {
	// The identifier of the link for Belvo.
	LinkId string `json:"link_id"`
	// The identifier of the institution for Belvo.
	Institution *string `json:"institution,omitempty"`
}

type _FinancialConnectionDetailsBelvo FinancialConnectionDetailsBelvo

// NewFinancialConnectionDetailsBelvo instantiates a new FinancialConnectionDetailsBelvo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionDetailsBelvo(linkId string) *FinancialConnectionDetailsBelvo {
	this := FinancialConnectionDetailsBelvo{}
	this.LinkId = linkId
	return &this
}

// NewFinancialConnectionDetailsBelvoWithDefaults instantiates a new FinancialConnectionDetailsBelvo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionDetailsBelvoWithDefaults() *FinancialConnectionDetailsBelvo {
	this := FinancialConnectionDetailsBelvo{}
	return &this
}

// GetLinkId returns the LinkId field value
func (o *FinancialConnectionDetailsBelvo) GetLinkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsBelvo) GetLinkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkId, true
}

// SetLinkId sets field value
func (o *FinancialConnectionDetailsBelvo) SetLinkId(v string) {
	o.LinkId = v
}

// GetInstitution returns the Institution field value if set, zero value otherwise.
func (o *FinancialConnectionDetailsBelvo) GetInstitution() string {
	if o == nil || IsNil(o.Institution) {
		var ret string
		return ret
	}
	return *o.Institution
}

// GetInstitutionOk returns a tuple with the Institution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionDetailsBelvo) GetInstitutionOk() (*string, bool) {
	if o == nil || IsNil(o.Institution) {
		return nil, false
	}
	return o.Institution, true
}

// HasInstitution returns a boolean if a field has been set.
func (o *FinancialConnectionDetailsBelvo) HasInstitution() bool {
	if o != nil && !IsNil(o.Institution) {
		return true
	}

	return false
}

// SetInstitution gets a reference to the given string and assigns it to the Institution field.
func (o *FinancialConnectionDetailsBelvo) SetInstitution(v string) {
	o.Institution = &v
}

func (o FinancialConnectionDetailsBelvo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionDetailsBelvo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["link_id"] = o.LinkId
	if !IsNil(o.Institution) {
		toSerialize["institution"] = o.Institution
	}
	return toSerialize, nil
}

func (o *FinancialConnectionDetailsBelvo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"link_id",
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

	varFinancialConnectionDetailsBelvo := _FinancialConnectionDetailsBelvo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinancialConnectionDetailsBelvo)

	if err != nil {
		return err
	}

	*o = FinancialConnectionDetailsBelvo(varFinancialConnectionDetailsBelvo)

	return err
}

type NullableFinancialConnectionDetailsBelvo struct {
	value *FinancialConnectionDetailsBelvo
	isSet bool
}

func (v NullableFinancialConnectionDetailsBelvo) Get() *FinancialConnectionDetailsBelvo {
	return v.value
}

func (v *NullableFinancialConnectionDetailsBelvo) Set(val *FinancialConnectionDetailsBelvo) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionDetailsBelvo) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionDetailsBelvo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionDetailsBelvo(val *FinancialConnectionDetailsBelvo) *NullableFinancialConnectionDetailsBelvo {
	return &NullableFinancialConnectionDetailsBelvo{value: val, isSet: true}
}

func (v NullableFinancialConnectionDetailsBelvo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionDetailsBelvo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


