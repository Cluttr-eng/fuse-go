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

// checks if the FinancialConnectionsHolding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionsHolding{}

// FinancialConnectionsHolding struct for FinancialConnectionsHolding
type FinancialConnectionsHolding struct {
	// The remote account ID associated with this holding.
	RemoteAccountId string `json:"remote_account_id"`
	// The original total value of the holding, in cents, when it was purchased. The format of this value is a double.
	CostBasis float32 `json:"cost_basis"`
	// The current market value of the holding, in cents. The format of this value is a double.
	Value float32 `json:"value"`
	// The number of units of the security held in this holding.
	Quantity float32 `json:"quantity"`
	// The price of the security, in cents, as provided by the financial institution managing the holding. The format of this value is a double.
	InstitutionPrice float32 `json:"institution_price"`
	Security FinancialConnectionsInvestmentSecurity `json:"security"`
}

// NewFinancialConnectionsHolding instantiates a new FinancialConnectionsHolding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionsHolding(remoteAccountId string, costBasis float32, value float32, quantity float32, institutionPrice float32, security FinancialConnectionsInvestmentSecurity) *FinancialConnectionsHolding {
	this := FinancialConnectionsHolding{}
	this.RemoteAccountId = remoteAccountId
	this.CostBasis = costBasis
	this.Value = value
	this.Quantity = quantity
	this.InstitutionPrice = institutionPrice
	this.Security = security
	return &this
}

// NewFinancialConnectionsHoldingWithDefaults instantiates a new FinancialConnectionsHolding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionsHoldingWithDefaults() *FinancialConnectionsHolding {
	this := FinancialConnectionsHolding{}
	return &this
}

// GetRemoteAccountId returns the RemoteAccountId field value
func (o *FinancialConnectionsHolding) GetRemoteAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteAccountId
}

// GetRemoteAccountIdOk returns a tuple with the RemoteAccountId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsHolding) GetRemoteAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteAccountId, true
}

// SetRemoteAccountId sets field value
func (o *FinancialConnectionsHolding) SetRemoteAccountId(v string) {
	o.RemoteAccountId = v
}

// GetCostBasis returns the CostBasis field value
func (o *FinancialConnectionsHolding) GetCostBasis() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CostBasis
}

// GetCostBasisOk returns a tuple with the CostBasis field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsHolding) GetCostBasisOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostBasis, true
}

// SetCostBasis sets field value
func (o *FinancialConnectionsHolding) SetCostBasis(v float32) {
	o.CostBasis = v
}

// GetValue returns the Value field value
func (o *FinancialConnectionsHolding) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsHolding) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FinancialConnectionsHolding) SetValue(v float32) {
	o.Value = v
}

// GetQuantity returns the Quantity field value
func (o *FinancialConnectionsHolding) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsHolding) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *FinancialConnectionsHolding) SetQuantity(v float32) {
	o.Quantity = v
}

// GetInstitutionPrice returns the InstitutionPrice field value
func (o *FinancialConnectionsHolding) GetInstitutionPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InstitutionPrice
}

// GetInstitutionPriceOk returns a tuple with the InstitutionPrice field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsHolding) GetInstitutionPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstitutionPrice, true
}

// SetInstitutionPrice sets field value
func (o *FinancialConnectionsHolding) SetInstitutionPrice(v float32) {
	o.InstitutionPrice = v
}

// GetSecurity returns the Security field value
func (o *FinancialConnectionsHolding) GetSecurity() FinancialConnectionsInvestmentSecurity {
	if o == nil {
		var ret FinancialConnectionsInvestmentSecurity
		return ret
	}

	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsHolding) GetSecurityOk() (*FinancialConnectionsInvestmentSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Security, true
}

// SetSecurity sets field value
func (o *FinancialConnectionsHolding) SetSecurity(v FinancialConnectionsInvestmentSecurity) {
	o.Security = v
}

func (o FinancialConnectionsHolding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionsHolding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["remote_account_id"] = o.RemoteAccountId
	toSerialize["cost_basis"] = o.CostBasis
	toSerialize["value"] = o.Value
	toSerialize["quantity"] = o.Quantity
	toSerialize["institution_price"] = o.InstitutionPrice
	toSerialize["security"] = o.Security
	return toSerialize, nil
}

type NullableFinancialConnectionsHolding struct {
	value *FinancialConnectionsHolding
	isSet bool
}

func (v NullableFinancialConnectionsHolding) Get() *FinancialConnectionsHolding {
	return v.value
}

func (v *NullableFinancialConnectionsHolding) Set(val *FinancialConnectionsHolding) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionsHolding) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionsHolding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionsHolding(val *FinancialConnectionsHolding) *NullableFinancialConnectionsHolding {
	return &NullableFinancialConnectionsHolding{value: val, isSet: true}
}

func (v NullableFinancialConnectionsHolding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionsHolding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


