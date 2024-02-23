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

// checks if the FinancialInstitution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialInstitution{}

// FinancialInstitution struct for FinancialInstitution
type FinancialInstitution struct {
	// Unique identifier for the financial institution id.
	Id string `json:"id"`
	// Name for the financial institution.
	Name string `json:"name"`
	Logo *FinancialInstitutionLogo `json:"logo,omitempty"`
	// Website of the financial institution.
	Website *string `json:"website,omitempty"`
	// List of country codes supported by this institution
	CountryCodes []CountryCode `json:"country_codes"`
}

type _FinancialInstitution FinancialInstitution

// NewFinancialInstitution instantiates a new FinancialInstitution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialInstitution(id string, name string, countryCodes []CountryCode) *FinancialInstitution {
	this := FinancialInstitution{}
	this.Id = id
	this.Name = name
	this.CountryCodes = countryCodes
	return &this
}

// NewFinancialInstitutionWithDefaults instantiates a new FinancialInstitution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialInstitutionWithDefaults() *FinancialInstitution {
	this := FinancialInstitution{}
	return &this
}

// GetId returns the Id field value
func (o *FinancialInstitution) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FinancialInstitution) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FinancialInstitution) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FinancialInstitution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FinancialInstitution) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FinancialInstitution) SetName(v string) {
	o.Name = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *FinancialInstitution) GetLogo() FinancialInstitutionLogo {
	if o == nil || IsNil(o.Logo) {
		var ret FinancialInstitutionLogo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialInstitution) GetLogoOk() (*FinancialInstitutionLogo, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *FinancialInstitution) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given FinancialInstitutionLogo and assigns it to the Logo field.
func (o *FinancialInstitution) SetLogo(v FinancialInstitutionLogo) {
	o.Logo = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *FinancialInstitution) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialInstitution) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *FinancialInstitution) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *FinancialInstitution) SetWebsite(v string) {
	o.Website = &v
}

// GetCountryCodes returns the CountryCodes field value
func (o *FinancialInstitution) GetCountryCodes() []CountryCode {
	if o == nil {
		var ret []CountryCode
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *FinancialInstitution) GetCountryCodesOk() ([]CountryCode, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *FinancialInstitution) SetCountryCodes(v []CountryCode) {
	o.CountryCodes = v
}

func (o FinancialInstitution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialInstitution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	toSerialize["country_codes"] = o.CountryCodes
	return toSerialize, nil
}

func (o *FinancialInstitution) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"country_codes",
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

	varFinancialInstitution := _FinancialInstitution{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinancialInstitution)

	if err != nil {
		return err
	}

	*o = FinancialInstitution(varFinancialInstitution)

	return err
}

type NullableFinancialInstitution struct {
	value *FinancialInstitution
	isSet bool
}

func (v NullableFinancialInstitution) Get() *FinancialInstitution {
	return v.value
}

func (v *NullableFinancialInstitution) Set(val *FinancialInstitution) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialInstitution) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialInstitution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialInstitution(val *FinancialInstitution) *NullableFinancialInstitution {
	return &NullableFinancialInstitution{value: val, isSet: true}
}

func (v NullableFinancialInstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialInstitution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


