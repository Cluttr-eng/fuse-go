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

// checks if the MigrateFinancialConnectionsTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrateFinancialConnectionsTokenRequest{}

// MigrateFinancialConnectionsTokenRequest struct for MigrateFinancialConnectionsTokenRequest
type MigrateFinancialConnectionsTokenRequest struct {
	ConnectionData MigrateFinancialConnectionsAggregatorConnectionData `json:"connection_data"`
	// The aggregator being migrated (either 'plaid', 'mx' or 'teller).
	Aggregator string `json:"aggregator"`
	Entity MigrateFinancialConnectionsTokenRequestEntity `json:"entity"`
	// A list of Fuse products that the migrated connection will have access to.
	FuseProducts []Product `json:"fuse_products"`
}

// NewMigrateFinancialConnectionsTokenRequest instantiates a new MigrateFinancialConnectionsTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrateFinancialConnectionsTokenRequest(connectionData MigrateFinancialConnectionsAggregatorConnectionData, aggregator string, entity MigrateFinancialConnectionsTokenRequestEntity, fuseProducts []Product) *MigrateFinancialConnectionsTokenRequest {
	this := MigrateFinancialConnectionsTokenRequest{}
	this.ConnectionData = connectionData
	this.Aggregator = aggregator
	this.Entity = entity
	this.FuseProducts = fuseProducts
	return &this
}

// NewMigrateFinancialConnectionsTokenRequestWithDefaults instantiates a new MigrateFinancialConnectionsTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrateFinancialConnectionsTokenRequestWithDefaults() *MigrateFinancialConnectionsTokenRequest {
	this := MigrateFinancialConnectionsTokenRequest{}
	return &this
}

// GetConnectionData returns the ConnectionData field value
func (o *MigrateFinancialConnectionsTokenRequest) GetConnectionData() MigrateFinancialConnectionsAggregatorConnectionData {
	if o == nil {
		var ret MigrateFinancialConnectionsAggregatorConnectionData
		return ret
	}

	return o.ConnectionData
}

// GetConnectionDataOk returns a tuple with the ConnectionData field value
// and a boolean to check if the value has been set.
func (o *MigrateFinancialConnectionsTokenRequest) GetConnectionDataOk() (*MigrateFinancialConnectionsAggregatorConnectionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionData, true
}

// SetConnectionData sets field value
func (o *MigrateFinancialConnectionsTokenRequest) SetConnectionData(v MigrateFinancialConnectionsAggregatorConnectionData) {
	o.ConnectionData = v
}

// GetAggregator returns the Aggregator field value
func (o *MigrateFinancialConnectionsTokenRequest) GetAggregator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value
// and a boolean to check if the value has been set.
func (o *MigrateFinancialConnectionsTokenRequest) GetAggregatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregator, true
}

// SetAggregator sets field value
func (o *MigrateFinancialConnectionsTokenRequest) SetAggregator(v string) {
	o.Aggregator = v
}

// GetEntity returns the Entity field value
func (o *MigrateFinancialConnectionsTokenRequest) GetEntity() MigrateFinancialConnectionsTokenRequestEntity {
	if o == nil {
		var ret MigrateFinancialConnectionsTokenRequestEntity
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *MigrateFinancialConnectionsTokenRequest) GetEntityOk() (*MigrateFinancialConnectionsTokenRequestEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *MigrateFinancialConnectionsTokenRequest) SetEntity(v MigrateFinancialConnectionsTokenRequestEntity) {
	o.Entity = v
}

// GetFuseProducts returns the FuseProducts field value
func (o *MigrateFinancialConnectionsTokenRequest) GetFuseProducts() []Product {
	if o == nil {
		var ret []Product
		return ret
	}

	return o.FuseProducts
}

// GetFuseProductsOk returns a tuple with the FuseProducts field value
// and a boolean to check if the value has been set.
func (o *MigrateFinancialConnectionsTokenRequest) GetFuseProductsOk() ([]Product, bool) {
	if o == nil {
		return nil, false
	}
	return o.FuseProducts, true
}

// SetFuseProducts sets field value
func (o *MigrateFinancialConnectionsTokenRequest) SetFuseProducts(v []Product) {
	o.FuseProducts = v
}

func (o MigrateFinancialConnectionsTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrateFinancialConnectionsTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connection_data"] = o.ConnectionData
	toSerialize["aggregator"] = o.Aggregator
	toSerialize["entity"] = o.Entity
	toSerialize["fuse_products"] = o.FuseProducts
	return toSerialize, nil
}

type NullableMigrateFinancialConnectionsTokenRequest struct {
	value *MigrateFinancialConnectionsTokenRequest
	isSet bool
}

func (v NullableMigrateFinancialConnectionsTokenRequest) Get() *MigrateFinancialConnectionsTokenRequest {
	return v.value
}

func (v *NullableMigrateFinancialConnectionsTokenRequest) Set(val *MigrateFinancialConnectionsTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrateFinancialConnectionsTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrateFinancialConnectionsTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrateFinancialConnectionsTokenRequest(val *MigrateFinancialConnectionsTokenRequest) *NullableMigrateFinancialConnectionsTokenRequest {
	return &NullableMigrateFinancialConnectionsTokenRequest{value: val, isSet: true}
}

func (v NullableMigrateFinancialConnectionsTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrateFinancialConnectionsTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


