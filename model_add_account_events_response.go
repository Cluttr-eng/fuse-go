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

// checks if the AddAccountEventsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAccountEventsResponse{}

// AddAccountEventsResponse struct for AddAccountEventsResponse
type AddAccountEventsResponse struct {
	// Response message
	Message string `json:"message"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId string `json:"request_id"`
}

// NewAddAccountEventsResponse instantiates a new AddAccountEventsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAccountEventsResponse(message string, requestId string) *AddAccountEventsResponse {
	this := AddAccountEventsResponse{}
	this.Message = message
	this.RequestId = requestId
	return &this
}

// NewAddAccountEventsResponseWithDefaults instantiates a new AddAccountEventsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAccountEventsResponseWithDefaults() *AddAccountEventsResponse {
	this := AddAccountEventsResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *AddAccountEventsResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AddAccountEventsResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AddAccountEventsResponse) SetMessage(v string) {
	o.Message = v
}

// GetRequestId returns the RequestId field value
func (o *AddAccountEventsResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AddAccountEventsResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AddAccountEventsResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AddAccountEventsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddAccountEventsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["request_id"] = o.RequestId
	return toSerialize, nil
}

type NullableAddAccountEventsResponse struct {
	value *AddAccountEventsResponse
	isSet bool
}

func (v NullableAddAccountEventsResponse) Get() *AddAccountEventsResponse {
	return v.value
}

func (v *NullableAddAccountEventsResponse) Set(val *AddAccountEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAccountEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAccountEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAccountEventsResponse(val *AddAccountEventsResponse) *NullableAddAccountEventsResponse {
	return &NullableAddAccountEventsResponse{value: val, isSet: true}
}

func (v NullableAddAccountEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAccountEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

