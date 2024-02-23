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

// checks if the GetFinancialInstitutionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinancialInstitutionResponse{}

// GetFinancialInstitutionResponse struct for GetFinancialInstitutionResponse
type GetFinancialInstitutionResponse struct {
	FinancialInstitution FinancialInstitution `json:"financial_institution"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId string `json:"request_id"`
}

type _GetFinancialInstitutionResponse GetFinancialInstitutionResponse

// NewGetFinancialInstitutionResponse instantiates a new GetFinancialInstitutionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinancialInstitutionResponse(financialInstitution FinancialInstitution, requestId string) *GetFinancialInstitutionResponse {
	this := GetFinancialInstitutionResponse{}
	this.FinancialInstitution = financialInstitution
	this.RequestId = requestId
	return &this
}

// NewGetFinancialInstitutionResponseWithDefaults instantiates a new GetFinancialInstitutionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinancialInstitutionResponseWithDefaults() *GetFinancialInstitutionResponse {
	this := GetFinancialInstitutionResponse{}
	return &this
}

// GetFinancialInstitution returns the FinancialInstitution field value
func (o *GetFinancialInstitutionResponse) GetFinancialInstitution() FinancialInstitution {
	if o == nil {
		var ret FinancialInstitution
		return ret
	}

	return o.FinancialInstitution
}

// GetFinancialInstitutionOk returns a tuple with the FinancialInstitution field value
// and a boolean to check if the value has been set.
func (o *GetFinancialInstitutionResponse) GetFinancialInstitutionOk() (*FinancialInstitution, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinancialInstitution, true
}

// SetFinancialInstitution sets field value
func (o *GetFinancialInstitutionResponse) SetFinancialInstitution(v FinancialInstitution) {
	o.FinancialInstitution = v
}

// GetRequestId returns the RequestId field value
func (o *GetFinancialInstitutionResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *GetFinancialInstitutionResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *GetFinancialInstitutionResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o GetFinancialInstitutionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinancialInstitutionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["financial_institution"] = o.FinancialInstitution
	toSerialize["request_id"] = o.RequestId
	return toSerialize, nil
}

func (o *GetFinancialInstitutionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"financial_institution",
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

	varGetFinancialInstitutionResponse := _GetFinancialInstitutionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFinancialInstitutionResponse)

	if err != nil {
		return err
	}

	*o = GetFinancialInstitutionResponse(varGetFinancialInstitutionResponse)

	return err
}

type NullableGetFinancialInstitutionResponse struct {
	value *GetFinancialInstitutionResponse
	isSet bool
}

func (v NullableGetFinancialInstitutionResponse) Get() *GetFinancialInstitutionResponse {
	return v.value
}

func (v *NullableGetFinancialInstitutionResponse) Set(val *GetFinancialInstitutionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinancialInstitutionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinancialInstitutionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinancialInstitutionResponse(val *GetFinancialInstitutionResponse) *NullableGetFinancialInstitutionResponse {
	return &NullableGetFinancialInstitutionResponse{value: val, isSet: true}
}

func (v NullableGetFinancialInstitutionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinancialInstitutionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


