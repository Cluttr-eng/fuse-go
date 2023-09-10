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

// checks if the MigrateFinancialConnectionsAggregatorConnectionDataPlaid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrateFinancialConnectionsAggregatorConnectionDataPlaid{}

// MigrateFinancialConnectionsAggregatorConnectionDataPlaid Details of the Plaid connection to migrate into the unified Fuse API.
type MigrateFinancialConnectionsAggregatorConnectionDataPlaid struct {
	// The Plaid access token associated with the user's financial accounts.
	AccessToken string `json:"access_token"`
	// If true, any webhooks received for this new financial connection will be sent to the webhook url used when the item was created. If false, the webhook url set at the organization sandbox/production environment level will be used instead.
	UseItemWebhook *bool `json:"use_item_webhook,omitempty"`
}

// NewMigrateFinancialConnectionsAggregatorConnectionDataPlaid instantiates a new MigrateFinancialConnectionsAggregatorConnectionDataPlaid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrateFinancialConnectionsAggregatorConnectionDataPlaid(accessToken string) *MigrateFinancialConnectionsAggregatorConnectionDataPlaid {
	this := MigrateFinancialConnectionsAggregatorConnectionDataPlaid{}
	this.AccessToken = accessToken
	return &this
}

// NewMigrateFinancialConnectionsAggregatorConnectionDataPlaidWithDefaults instantiates a new MigrateFinancialConnectionsAggregatorConnectionDataPlaid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrateFinancialConnectionsAggregatorConnectionDataPlaidWithDefaults() *MigrateFinancialConnectionsAggregatorConnectionDataPlaid {
	this := MigrateFinancialConnectionsAggregatorConnectionDataPlaid{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetUseItemWebhook returns the UseItemWebhook field value if set, zero value otherwise.
func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) GetUseItemWebhook() bool {
	if o == nil || IsNil(o.UseItemWebhook) {
		var ret bool
		return ret
	}
	return *o.UseItemWebhook
}

// GetUseItemWebhookOk returns a tuple with the UseItemWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) GetUseItemWebhookOk() (*bool, bool) {
	if o == nil || IsNil(o.UseItemWebhook) {
		return nil, false
	}
	return o.UseItemWebhook, true
}

// HasUseItemWebhook returns a boolean if a field has been set.
func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) HasUseItemWebhook() bool {
	if o != nil && !IsNil(o.UseItemWebhook) {
		return true
	}

	return false
}

// SetUseItemWebhook gets a reference to the given bool and assigns it to the UseItemWebhook field.
func (o *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) SetUseItemWebhook(v bool) {
	o.UseItemWebhook = &v
}

func (o MigrateFinancialConnectionsAggregatorConnectionDataPlaid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrateFinancialConnectionsAggregatorConnectionDataPlaid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	if !IsNil(o.UseItemWebhook) {
		toSerialize["use_item_webhook"] = o.UseItemWebhook
	}
	return toSerialize, nil
}

type NullableMigrateFinancialConnectionsAggregatorConnectionDataPlaid struct {
	value *MigrateFinancialConnectionsAggregatorConnectionDataPlaid
	isSet bool
}

func (v NullableMigrateFinancialConnectionsAggregatorConnectionDataPlaid) Get() *MigrateFinancialConnectionsAggregatorConnectionDataPlaid {
	return v.value
}

func (v *NullableMigrateFinancialConnectionsAggregatorConnectionDataPlaid) Set(val *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrateFinancialConnectionsAggregatorConnectionDataPlaid) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrateFinancialConnectionsAggregatorConnectionDataPlaid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrateFinancialConnectionsAggregatorConnectionDataPlaid(val *MigrateFinancialConnectionsAggregatorConnectionDataPlaid) *NullableMigrateFinancialConnectionsAggregatorConnectionDataPlaid {
	return &NullableMigrateFinancialConnectionsAggregatorConnectionDataPlaid{value: val, isSet: true}
}

func (v NullableMigrateFinancialConnectionsAggregatorConnectionDataPlaid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrateFinancialConnectionsAggregatorConnectionDataPlaid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


