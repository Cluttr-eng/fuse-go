/*
Fuse

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fuse

import (
	"encoding/json"
	"fmt"
)

// WebhookType the model 'WebhookType'
type WebhookType string

// List of WebhookType
const (
	WEBHOOKTYPE_FINANCIAL_CONNECTION_SYNC_DATA WebhookType = "financial_connection.sync_data"
	WEBHOOKTYPE_FINANCIAL_CONNECTION_DISCONNECTED WebhookType = "financial_connection.disconnected"
	WEBHOOKTYPE_FINANCIAL_CONNECTION_FINISHED WebhookType = "financial_connection.finished"
	WEBHOOKTYPE_TRANSACTIONS_UPDATES WebhookType = "transactions.updates"
	WEBHOOKTYPE_ASSETS_REPORT_READY WebhookType = "assets.report_ready"
)

// All allowed values of WebhookType enum
var AllowedWebhookTypeEnumValues = []WebhookType{
	"financial_connection.sync_data",
	"financial_connection.disconnected",
	"financial_connection.finished",
	"transactions.updates",
	"assets.report_ready",
}

func (v *WebhookType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookType(value)
	for _, existing := range AllowedWebhookTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookType", value)
}

// NewWebhookTypeFromValue returns a pointer to a valid WebhookType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookTypeFromValue(v string) (*WebhookType, error) {
	ev := WebhookType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookType: valid values are %v", v, AllowedWebhookTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookType) IsValid() bool {
	for _, existing := range AllowedWebhookTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookType value
func (v WebhookType) Ptr() *WebhookType {
	return &v
}

type NullableWebhookType struct {
	value *WebhookType
	isSet bool
}

func (v NullableWebhookType) Get() *WebhookType {
	return v.value
}

func (v *NullableWebhookType) Set(val *WebhookType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookType(val *WebhookType) *NullableWebhookType {
	return &NullableWebhookType{value: val, isSet: true}
}

func (v NullableWebhookType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
