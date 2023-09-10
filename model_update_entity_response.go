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

// checks if the UpdateEntityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntityResponse{}

// UpdateEntityResponse struct for UpdateEntityResponse
type UpdateEntityResponse struct {
	// Id of the entity
	Id *string `json:"id,omitempty"`
	// Email of the entity
	Email *string `json:"email,omitempty"`
	// These will force the user to connect through all of these aggregators
	Aggregators []Aggregator `json:"aggregators,omitempty"`
	InstitutionIds []string `json:"institution_ids,omitempty"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId *string `json:"request_id,omitempty"`
}

// NewUpdateEntityResponse instantiates a new UpdateEntityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntityResponse() *UpdateEntityResponse {
	this := UpdateEntityResponse{}
	return &this
}

// NewUpdateEntityResponseWithDefaults instantiates a new UpdateEntityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntityResponseWithDefaults() *UpdateEntityResponse {
	this := UpdateEntityResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateEntityResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateEntityResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateEntityResponse) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateEntityResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateEntityResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateEntityResponse) SetEmail(v string) {
	o.Email = &v
}

// GetAggregators returns the Aggregators field value if set, zero value otherwise.
func (o *UpdateEntityResponse) GetAggregators() []Aggregator {
	if o == nil || IsNil(o.Aggregators) {
		var ret []Aggregator
		return ret
	}
	return o.Aggregators
}

// GetAggregatorsOk returns a tuple with the Aggregators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityResponse) GetAggregatorsOk() ([]Aggregator, bool) {
	if o == nil || IsNil(o.Aggregators) {
		return nil, false
	}
	return o.Aggregators, true
}

// HasAggregators returns a boolean if a field has been set.
func (o *UpdateEntityResponse) HasAggregators() bool {
	if o != nil && !IsNil(o.Aggregators) {
		return true
	}

	return false
}

// SetAggregators gets a reference to the given []Aggregator and assigns it to the Aggregators field.
func (o *UpdateEntityResponse) SetAggregators(v []Aggregator) {
	o.Aggregators = v
}

// GetInstitutionIds returns the InstitutionIds field value if set, zero value otherwise.
func (o *UpdateEntityResponse) GetInstitutionIds() []string {
	if o == nil || IsNil(o.InstitutionIds) {
		var ret []string
		return ret
	}
	return o.InstitutionIds
}

// GetInstitutionIdsOk returns a tuple with the InstitutionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityResponse) GetInstitutionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.InstitutionIds) {
		return nil, false
	}
	return o.InstitutionIds, true
}

// HasInstitutionIds returns a boolean if a field has been set.
func (o *UpdateEntityResponse) HasInstitutionIds() bool {
	if o != nil && !IsNil(o.InstitutionIds) {
		return true
	}

	return false
}

// SetInstitutionIds gets a reference to the given []string and assigns it to the InstitutionIds field.
func (o *UpdateEntityResponse) SetInstitutionIds(v []string) {
	o.InstitutionIds = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *UpdateEntityResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *UpdateEntityResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *UpdateEntityResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o UpdateEntityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Aggregators) {
		toSerialize["aggregators"] = o.Aggregators
	}
	if !IsNil(o.InstitutionIds) {
		toSerialize["institution_ids"] = o.InstitutionIds
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableUpdateEntityResponse struct {
	value *UpdateEntityResponse
	isSet bool
}

func (v NullableUpdateEntityResponse) Get() *UpdateEntityResponse {
	return v.value
}

func (v *NullableUpdateEntityResponse) Set(val *UpdateEntityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntityResponse(val *UpdateEntityResponse) *NullableUpdateEntityResponse {
	return &NullableUpdateEntityResponse{value: val, isSet: true}
}

func (v NullableUpdateEntityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

