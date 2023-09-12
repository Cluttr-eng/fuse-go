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

// CountryCode the model 'CountryCode'
type CountryCode string

// List of CountryCode
const (
	COUNTRYCODE_AE CountryCode = "AE"
	COUNTRYCODE_AM CountryCode = "AM"
	COUNTRYCODE_AR CountryCode = "AR"
	COUNTRYCODE_AT CountryCode = "AT"
	COUNTRYCODE_AU CountryCode = "AU"
	COUNTRYCODE_BD CountryCode = "BD"
	COUNTRYCODE_BE CountryCode = "BE"
	COUNTRYCODE_BG CountryCode = "BG"
	COUNTRYCODE_BH CountryCode = "BH"
	COUNTRYCODE_BM CountryCode = "BM"
	COUNTRYCODE_BN CountryCode = "BN"
	COUNTRYCODE_BR CountryCode = "BR"
	COUNTRYCODE_BW CountryCode = "BW"
	COUNTRYCODE_BY CountryCode = "BY"
	COUNTRYCODE_CA CountryCode = "CA"
	COUNTRYCODE_CH CountryCode = "CH"
	COUNTRYCODE_CI CountryCode = "CI"
	COUNTRYCODE_CL CountryCode = "CL"
	COUNTRYCODE_CM CountryCode = "CM"
	COUNTRYCODE_CN CountryCode = "CN"
	COUNTRYCODE_CO CountryCode = "CO"
	COUNTRYCODE_CZ CountryCode = "CZ"
	COUNTRYCODE_DE CountryCode = "DE"
	COUNTRYCODE_DO CountryCode = "DO"
	COUNTRYCODE_DZ CountryCode = "DZ"
	COUNTRYCODE_EC CountryCode = "EC"
	COUNTRYCODE_EG CountryCode = "EG"
	COUNTRYCODE_ES CountryCode = "ES"
	COUNTRYCODE_FI CountryCode = "FI"
	COUNTRYCODE_FK CountryCode = "FK"
	COUNTRYCODE_FR CountryCode = "FR"
	COUNTRYCODE_GB CountryCode = "GB"
	COUNTRYCODE_GG CountryCode = "GG"
	COUNTRYCODE_GH CountryCode = "GH"
	COUNTRYCODE_GM CountryCode = "GM"
	COUNTRYCODE_GR CountryCode = "GR"
	COUNTRYCODE_HK CountryCode = "HK"
	COUNTRYCODE_HU CountryCode = "HU"
	COUNTRYCODE_ID CountryCode = "ID"
	COUNTRYCODE_IE CountryCode = "IE"
	COUNTRYCODE_IL CountryCode = "IL"
	COUNTRYCODE_IM CountryCode = "IM"
	COUNTRYCODE_IN CountryCode = "IN"
	COUNTRYCODE_IT CountryCode = "IT"
	COUNTRYCODE_JE CountryCode = "JE"
	COUNTRYCODE_JO CountryCode = "JO"
	COUNTRYCODE_JP CountryCode = "JP"
	COUNTRYCODE_KE CountryCode = "KE"
	COUNTRYCODE_KH CountryCode = "KH"
	COUNTRYCODE_KR CountryCode = "KR"
	COUNTRYCODE_KW CountryCode = "KW"
	COUNTRYCODE_LA CountryCode = "LA"
	COUNTRYCODE_LB CountryCode = "LB"
	COUNTRYCODE_LK CountryCode = "LK"
	COUNTRYCODE_LT CountryCode = "LT"
	COUNTRYCODE_LU CountryCode = "LU"
	COUNTRYCODE_MC CountryCode = "MC"
	COUNTRYCODE_MD CountryCode = "MD"
	COUNTRYCODE_MK CountryCode = "MK"
	COUNTRYCODE_MO CountryCode = "MO"
	COUNTRYCODE_MT CountryCode = "MT"
	COUNTRYCODE_MU CountryCode = "MU"
	COUNTRYCODE_MV CountryCode = "MV"
	COUNTRYCODE_MX CountryCode = "MX"
	COUNTRYCODE_MY CountryCode = "MY"
	COUNTRYCODE_NG CountryCode = "NG"
	COUNTRYCODE_NL CountryCode = "NL"
	COUNTRYCODE_NP CountryCode = "NP"
	COUNTRYCODE_NZ CountryCode = "NZ"
	COUNTRYCODE_OM CountryCode = "OM"
	COUNTRYCODE_PE CountryCode = "PE"
	COUNTRYCODE_PH CountryCode = "PH"
	COUNTRYCODE_PK CountryCode = "PK"
	COUNTRYCODE_PL CountryCode = "PL"
	COUNTRYCODE_PT CountryCode = "PT"
	COUNTRYCODE_QA CountryCode = "QA"
	COUNTRYCODE_RO CountryCode = "RO"
	COUNTRYCODE_RU CountryCode = "RU"
	COUNTRYCODE_SA CountryCode = "SA"
	COUNTRYCODE_SE CountryCode = "SE"
	COUNTRYCODE_SG CountryCode = "SG"
	COUNTRYCODE_SK CountryCode = "SK"
	COUNTRYCODE_SL CountryCode = "SL"
	COUNTRYCODE_TH CountryCode = "TH"
	COUNTRYCODE_TR CountryCode = "TR"
	COUNTRYCODE_TW CountryCode = "TW"
	COUNTRYCODE_TZ CountryCode = "TZ"
	COUNTRYCODE_UA CountryCode = "UA"
	COUNTRYCODE_UG CountryCode = "UG"
	COUNTRYCODE_US CountryCode = "US"
	COUNTRYCODE_UY CountryCode = "UY"
	COUNTRYCODE_VN CountryCode = "VN"
	COUNTRYCODE_ZA CountryCode = "ZA"
	COUNTRYCODE_ZM CountryCode = "ZM"
	COUNTRYCODE_ZW CountryCode = "ZW"
)

// All allowed values of CountryCode enum
var AllowedCountryCodeEnumValues = []CountryCode{
	"AE",
	"AM",
	"AR",
	"AT",
	"AU",
	"BD",
	"BE",
	"BG",
	"BH",
	"BM",
	"BN",
	"BR",
	"BW",
	"BY",
	"CA",
	"CH",
	"CI",
	"CL",
	"CM",
	"CN",
	"CO",
	"CZ",
	"DE",
	"DO",
	"DZ",
	"EC",
	"EG",
	"ES",
	"FI",
	"FK",
	"FR",
	"GB",
	"GG",
	"GH",
	"GM",
	"GR",
	"HK",
	"HU",
	"ID",
	"IE",
	"IL",
	"IM",
	"IN",
	"IT",
	"JE",
	"JO",
	"JP",
	"KE",
	"KH",
	"KR",
	"KW",
	"LA",
	"LB",
	"LK",
	"LT",
	"LU",
	"MC",
	"MD",
	"MK",
	"MO",
	"MT",
	"MU",
	"MV",
	"MX",
	"MY",
	"NG",
	"NL",
	"NP",
	"NZ",
	"OM",
	"PE",
	"PH",
	"PK",
	"PL",
	"PT",
	"QA",
	"RO",
	"RU",
	"SA",
	"SE",
	"SG",
	"SK",
	"SL",
	"TH",
	"TR",
	"TW",
	"TZ",
	"UA",
	"UG",
	"US",
	"UY",
	"VN",
	"ZA",
	"ZM",
	"ZW",
}

func (v *CountryCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CountryCode(value)
	for _, existing := range AllowedCountryCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CountryCode", value)
}

// NewCountryCodeFromValue returns a pointer to a valid CountryCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCountryCodeFromValue(v string) (*CountryCode, error) {
	ev := CountryCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CountryCode: valid values are %v", v, AllowedCountryCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CountryCode) IsValid() bool {
	for _, existing := range AllowedCountryCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CountryCode value
func (v CountryCode) Ptr() *CountryCode {
	return &v
}

type NullableCountryCode struct {
	value *CountryCode
	isSet bool
}

func (v NullableCountryCode) Get() *CountryCode {
	return v.value
}

func (v *NullableCountryCode) Set(val *CountryCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryCode(val *CountryCode) *NullableCountryCode {
	return &NullableCountryCode{value: val, isSet: true}
}

func (v NullableCountryCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

