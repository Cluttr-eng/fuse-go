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

// checks if the FinancialConnectionsInvestmentSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialConnectionsInvestmentSecurity{}

// FinancialConnectionsInvestmentSecurity struct for FinancialConnectionsInvestmentSecurity
type FinancialConnectionsInvestmentSecurity struct {
	// Remote Id of the security, ie Plaid or Snaptrade security id
	RemoteId string `json:"remote_id"`
	// The trading symbol for publicly traded securities, or a short identifier if available.
	Symbol string `json:"symbol"`
	// The International Securities Identification Number (ISIN) uniquely identifies the security.
	Isin *string `json:"isin,omitempty"`
	// The Stock Exchange Daily Official List (SEDOL) code uniquely identifies the security, primarily used in the United Kingdom and Ireland.
	Sedol *string `json:"sedol,omitempty"`
	// The Committee on Uniform Securities Identification Procedures (CUSIP) number uniquely identifies the security, primarily used in the United States and Canada.
	Cusip *string `json:"cusip,omitempty"`
	// The closing price of the security, in cents, at the end of the most recent trading day. The format of this value is a double.
	ClosePrice *float32 `json:"close_price,omitempty"`
	Currency Currency `json:"currency"`
	// A descriptive name for the security, suitable for display.
	Name *string `json:"name,omitempty"`
	Type *FinancialConnectionsInvestmentSecurityType `json:"type,omitempty"`
	Exchange *FinancialConnectionsInvestmentSecurityExchange `json:"exchange,omitempty"`
}

// NewFinancialConnectionsInvestmentSecurity instantiates a new FinancialConnectionsInvestmentSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialConnectionsInvestmentSecurity(remoteId string, symbol string, currency Currency) *FinancialConnectionsInvestmentSecurity {
	this := FinancialConnectionsInvestmentSecurity{}
	this.RemoteId = remoteId
	this.Symbol = symbol
	this.Currency = currency
	return &this
}

// NewFinancialConnectionsInvestmentSecurityWithDefaults instantiates a new FinancialConnectionsInvestmentSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialConnectionsInvestmentSecurityWithDefaults() *FinancialConnectionsInvestmentSecurity {
	this := FinancialConnectionsInvestmentSecurity{}
	return &this
}

// GetRemoteId returns the RemoteId field value
func (o *FinancialConnectionsInvestmentSecurity) GetRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurity) GetRemoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteId, true
}

// SetRemoteId sets field value
func (o *FinancialConnectionsInvestmentSecurity) SetRemoteId(v string) {
	o.RemoteId = v
}

// GetSymbol returns the Symbol field value
func (o *FinancialConnectionsInvestmentSecurity) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurity) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *FinancialConnectionsInvestmentSecurity) SetSymbol(v string) {
	o.Symbol = v
}

// GetIsin returns the Isin field value if set, zero value otherwise.
func (o *FinancialConnectionsInvestmentSecurity) GetIsin() string {
	if o == nil || IsNil(o.Isin) {
		var ret string
		return ret
	}
	return *o.Isin
}

// GetIsinOk returns a tuple with the Isin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurity) GetIsinOk() (*string, bool) {
	if o == nil || IsNil(o.Isin) {
		return nil, false
	}
	return o.Isin, true
}

// HasIsin returns a boolean if a field has been set.
func (o *FinancialConnectionsInvestmentSecurity) HasIsin() bool {
	if o != nil && !IsNil(o.Isin) {
		return true
	}

	return false
}

// SetIsin gets a reference to the given string and assigns it to the Isin field.
func (o *FinancialConnectionsInvestmentSecurity) SetIsin(v string) {
	o.Isin = &v
}

// GetSedol returns the Sedol field value if set, zero value otherwise.
func (o *FinancialConnectionsInvestmentSecurity) GetSedol() string {
	if o == nil || IsNil(o.Sedol) {
		var ret string
		return ret
	}
	return *o.Sedol
}

// GetSedolOk returns a tuple with the Sedol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurity) GetSedolOk() (*string, bool) {
	if o == nil || IsNil(o.Sedol) {
		return nil, false
	}
	return o.Sedol, true
}

// HasSedol returns a boolean if a field has been set.
func (o *FinancialConnectionsInvestmentSecurity) HasSedol() bool {
	if o != nil && !IsNil(o.Sedol) {
		return true
	}

	return false
}

// SetSedol gets a reference to the given string and assigns it to the Sedol field.
func (o *FinancialConnectionsInvestmentSecurity) SetSedol(v string) {
	o.Sedol = &v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *FinancialConnectionsInvestmentSecurity) GetCusip() string {
	if o == nil || IsNil(o.Cusip) {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurity) GetCusipOk() (*string, bool) {
	if o == nil || IsNil(o.Cusip) {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *FinancialConnectionsInvestmentSecurity) HasCusip() bool {
	if o != nil && !IsNil(o.Cusip) {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *FinancialConnectionsInvestmentSecurity) SetCusip(v string) {
	o.Cusip = &v
}

// GetClosePrice returns the ClosePrice field value if set, zero value otherwise.
func (o *FinancialConnectionsInvestmentSecurity) GetClosePrice() float32 {
	if o == nil || IsNil(o.ClosePrice) {
		var ret float32
		return ret
	}
	return *o.ClosePrice
}

// GetClosePriceOk returns a tuple with the ClosePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurity) GetClosePriceOk() (*float32, bool) {
	if o == nil || IsNil(o.ClosePrice) {
		return nil, false
	}
	return o.ClosePrice, true
}

// HasClosePrice returns a boolean if a field has been set.
func (o *FinancialConnectionsInvestmentSecurity) HasClosePrice() bool {
	if o != nil && !IsNil(o.ClosePrice) {
		return true
	}

	return false
}

// SetClosePrice gets a reference to the given float32 and assigns it to the ClosePrice field.
func (o *FinancialConnectionsInvestmentSecurity) SetClosePrice(v float32) {
	o.ClosePrice = &v
}

// GetCurrency returns the Currency field value
func (o *FinancialConnectionsInvestmentSecurity) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurity) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *FinancialConnectionsInvestmentSecurity) SetCurrency(v Currency) {
	o.Currency = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FinancialConnectionsInvestmentSecurity) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurity) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FinancialConnectionsInvestmentSecurity) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FinancialConnectionsInvestmentSecurity) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FinancialConnectionsInvestmentSecurity) GetType() FinancialConnectionsInvestmentSecurityType {
	if o == nil || IsNil(o.Type) {
		var ret FinancialConnectionsInvestmentSecurityType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurity) GetTypeOk() (*FinancialConnectionsInvestmentSecurityType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FinancialConnectionsInvestmentSecurity) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FinancialConnectionsInvestmentSecurityType and assigns it to the Type field.
func (o *FinancialConnectionsInvestmentSecurity) SetType(v FinancialConnectionsInvestmentSecurityType) {
	o.Type = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *FinancialConnectionsInvestmentSecurity) GetExchange() FinancialConnectionsInvestmentSecurityExchange {
	if o == nil || IsNil(o.Exchange) {
		var ret FinancialConnectionsInvestmentSecurityExchange
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialConnectionsInvestmentSecurity) GetExchangeOk() (*FinancialConnectionsInvestmentSecurityExchange, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *FinancialConnectionsInvestmentSecurity) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given FinancialConnectionsInvestmentSecurityExchange and assigns it to the Exchange field.
func (o *FinancialConnectionsInvestmentSecurity) SetExchange(v FinancialConnectionsInvestmentSecurityExchange) {
	o.Exchange = &v
}

func (o FinancialConnectionsInvestmentSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialConnectionsInvestmentSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["remote_id"] = o.RemoteId
	toSerialize["symbol"] = o.Symbol
	if !IsNil(o.Isin) {
		toSerialize["isin"] = o.Isin
	}
	if !IsNil(o.Sedol) {
		toSerialize["sedol"] = o.Sedol
	}
	if !IsNil(o.Cusip) {
		toSerialize["cusip"] = o.Cusip
	}
	if !IsNil(o.ClosePrice) {
		toSerialize["close_price"] = o.ClosePrice
	}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	return toSerialize, nil
}

type NullableFinancialConnectionsInvestmentSecurity struct {
	value *FinancialConnectionsInvestmentSecurity
	isSet bool
}

func (v NullableFinancialConnectionsInvestmentSecurity) Get() *FinancialConnectionsInvestmentSecurity {
	return v.value
}

func (v *NullableFinancialConnectionsInvestmentSecurity) Set(val *FinancialConnectionsInvestmentSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionsInvestmentSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionsInvestmentSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionsInvestmentSecurity(val *FinancialConnectionsInvestmentSecurity) *NullableFinancialConnectionsInvestmentSecurity {
	return &NullableFinancialConnectionsInvestmentSecurity{value: val, isSet: true}
}

func (v NullableFinancialConnectionsInvestmentSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionsInvestmentSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

