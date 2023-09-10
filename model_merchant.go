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

// checks if the Merchant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Merchant{}

// Merchant struct for Merchant
type Merchant struct {
	// A Fuse defined, unique ID for the merchant associated with this transaction.
	Id string `json:"id"`
	// The name of the merchant.
	Name string `json:"name"`
	Logo *MerchantLogo `json:"logo,omitempty"`
}

// NewMerchant instantiates a new Merchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchant(id string, name string) *Merchant {
	this := Merchant{}
	this.Id = id
	this.Name = name
	return &this
}

// NewMerchantWithDefaults instantiates a new Merchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantWithDefaults() *Merchant {
	this := Merchant{}
	return &this
}

// GetId returns the Id field value
func (o *Merchant) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Merchant) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Merchant) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Merchant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Merchant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Merchant) SetName(v string) {
	o.Name = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *Merchant) GetLogo() MerchantLogo {
	if o == nil || IsNil(o.Logo) {
		var ret MerchantLogo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetLogoOk() (*MerchantLogo, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *Merchant) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given MerchantLogo and assigns it to the Logo field.
func (o *Merchant) SetLogo(v MerchantLogo) {
	o.Logo = &v
}

func (o Merchant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Merchant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	return toSerialize, nil
}

type NullableMerchant struct {
	value *Merchant
	isSet bool
}

func (v NullableMerchant) Get() *Merchant {
	return v.value
}

func (v *NullableMerchant) Set(val *Merchant) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchant(val *Merchant) *NullableMerchant {
	return &NullableMerchant{value: val, isSet: true}
}

func (v NullableMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

