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

// checks if the GetEntityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntityResponse{}

// GetEntityResponse struct for GetEntityResponse
type GetEntityResponse struct {
	// Id of the entity
	Id string `json:"id"`
	// Email of the entity
	Email *string `json:"email,omitempty"`
	// These will force the user to connect through all of these aggregators
	Aggregators []Aggregator `json:"aggregators,omitempty"`
	InstitutionIds []string `json:"institution_ids,omitempty"`
	// Data needed to query data from the various aggregators
	FinancialConnections []FinancialConnectionDetails `json:"financial_connections"`
	// An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues.
	RequestId string `json:"request_id"`
}

// NewGetEntityResponse instantiates a new GetEntityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntityResponse(id string, financialConnections []FinancialConnectionDetails, requestId string) *GetEntityResponse {
	this := GetEntityResponse{}
	this.Id = id
	this.FinancialConnections = financialConnections
	this.RequestId = requestId
	return &this
}

// NewGetEntityResponseWithDefaults instantiates a new GetEntityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntityResponseWithDefaults() *GetEntityResponse {
	this := GetEntityResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GetEntityResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetEntityResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetEntityResponse) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetEntityResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntityResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetEntityResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetEntityResponse) SetEmail(v string) {
	o.Email = &v
}

// GetAggregators returns the Aggregators field value if set, zero value otherwise.
func (o *GetEntityResponse) GetAggregators() []Aggregator {
	if o == nil || IsNil(o.Aggregators) {
		var ret []Aggregator
		return ret
	}
	return o.Aggregators
}

// GetAggregatorsOk returns a tuple with the Aggregators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntityResponse) GetAggregatorsOk() ([]Aggregator, bool) {
	if o == nil || IsNil(o.Aggregators) {
		return nil, false
	}
	return o.Aggregators, true
}

// HasAggregators returns a boolean if a field has been set.
func (o *GetEntityResponse) HasAggregators() bool {
	if o != nil && !IsNil(o.Aggregators) {
		return true
	}

	return false
}

// SetAggregators gets a reference to the given []Aggregator and assigns it to the Aggregators field.
func (o *GetEntityResponse) SetAggregators(v []Aggregator) {
	o.Aggregators = v
}

// GetInstitutionIds returns the InstitutionIds field value if set, zero value otherwise.
func (o *GetEntityResponse) GetInstitutionIds() []string {
	if o == nil || IsNil(o.InstitutionIds) {
		var ret []string
		return ret
	}
	return o.InstitutionIds
}

// GetInstitutionIdsOk returns a tuple with the InstitutionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEntityResponse) GetInstitutionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.InstitutionIds) {
		return nil, false
	}
	return o.InstitutionIds, true
}

// HasInstitutionIds returns a boolean if a field has been set.
func (o *GetEntityResponse) HasInstitutionIds() bool {
	if o != nil && !IsNil(o.InstitutionIds) {
		return true
	}

	return false
}

// SetInstitutionIds gets a reference to the given []string and assigns it to the InstitutionIds field.
func (o *GetEntityResponse) SetInstitutionIds(v []string) {
	o.InstitutionIds = v
}

// GetFinancialConnections returns the FinancialConnections field value
func (o *GetEntityResponse) GetFinancialConnections() []FinancialConnectionDetails {
	if o == nil {
		var ret []FinancialConnectionDetails
		return ret
	}

	return o.FinancialConnections
}

// GetFinancialConnectionsOk returns a tuple with the FinancialConnections field value
// and a boolean to check if the value has been set.
func (o *GetEntityResponse) GetFinancialConnectionsOk() ([]FinancialConnectionDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinancialConnections, true
}

// SetFinancialConnections sets field value
func (o *GetEntityResponse) SetFinancialConnections(v []FinancialConnectionDetails) {
	o.FinancialConnections = v
}

// GetRequestId returns the RequestId field value
func (o *GetEntityResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *GetEntityResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *GetEntityResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o GetEntityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Aggregators) {
		toSerialize["aggregators"] = o.Aggregators
	}
	if !IsNil(o.InstitutionIds) {
		toSerialize["institution_ids"] = o.InstitutionIds
	}
	toSerialize["financial_connections"] = o.FinancialConnections
	toSerialize["request_id"] = o.RequestId
	return toSerialize, nil
}

type NullableGetEntityResponse struct {
	value *GetEntityResponse
	isSet bool
}

func (v NullableGetEntityResponse) Get() *GetEntityResponse {
	return v.value
}

func (v *NullableGetEntityResponse) Set(val *GetEntityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntityResponse(val *GetEntityResponse) *NullableGetEntityResponse {
	return &NullableGetEntityResponse{value: val, isSet: true}
}

func (v NullableGetEntityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


