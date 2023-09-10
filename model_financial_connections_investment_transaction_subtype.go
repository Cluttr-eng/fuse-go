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

// FinancialConnectionsInvestmentTransactionSubtype the model 'FinancialConnectionsInvestmentTransactionSubtype'
type FinancialConnectionsInvestmentTransactionSubtype string

// List of FinancialConnectionsInvestmentTransactionSubtype
const (
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_ACCOUNT_FEE FinancialConnectionsInvestmentTransactionSubtype = "account_fee"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_ADJUSTMENT FinancialConnectionsInvestmentTransactionSubtype = "adjustment"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_ASSIGNMENT FinancialConnectionsInvestmentTransactionSubtype = "assignment"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_BUY FinancialConnectionsInvestmentTransactionSubtype = "buy"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_BUY_TO_COVER FinancialConnectionsInvestmentTransactionSubtype = "buy_to_cover"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_CONTRIBUTION FinancialConnectionsInvestmentTransactionSubtype = "contribution"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_DEPOSIT FinancialConnectionsInvestmentTransactionSubtype = "deposit"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_DISTRIBUTION FinancialConnectionsInvestmentTransactionSubtype = "distribution"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_DIVIDEND FinancialConnectionsInvestmentTransactionSubtype = "dividend"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_DIVIDEND_REINVESTMENT FinancialConnectionsInvestmentTransactionSubtype = "dividend_reinvestment"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_EXERCISE FinancialConnectionsInvestmentTransactionSubtype = "exercise"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_EXPIRE FinancialConnectionsInvestmentTransactionSubtype = "expire"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_FUND_FEE FinancialConnectionsInvestmentTransactionSubtype = "fund_fee"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_INTEREST FinancialConnectionsInvestmentTransactionSubtype = "interest"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_INTEREST_RECEIVABLE FinancialConnectionsInvestmentTransactionSubtype = "interest_receivable"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_INTEREST_REINVESTMENT FinancialConnectionsInvestmentTransactionSubtype = "interest_reinvestment"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_LEGAL_FEE FinancialConnectionsInvestmentTransactionSubtype = "legal_fee"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_LOAN_PAYMENT FinancialConnectionsInvestmentTransactionSubtype = "loan_payment"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_LONG_TERM_CAPITAL_GAIN FinancialConnectionsInvestmentTransactionSubtype = "long_term_capital_gain"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_LONG_TERM_CAPITAL_GAIN_REINVESTMENT FinancialConnectionsInvestmentTransactionSubtype = "long_term_capital_gain_reinvestment"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_MANAGEMENT_FEE FinancialConnectionsInvestmentTransactionSubtype = "management_fee"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_MARGIN_EXPENSE FinancialConnectionsInvestmentTransactionSubtype = "margin_expense"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_MERGER FinancialConnectionsInvestmentTransactionSubtype = "merger"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_MISCELLANEOUS_FEE FinancialConnectionsInvestmentTransactionSubtype = "miscellaneous_fee"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_NON_QUALIFIED_DIVIDEND FinancialConnectionsInvestmentTransactionSubtype = "non_qualified_dividend"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_NON_RESIDENT_TAX FinancialConnectionsInvestmentTransactionSubtype = "non_resident_tax"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_PENDING_CREDIT FinancialConnectionsInvestmentTransactionSubtype = "pending_credit"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_PENDING_DEBIT FinancialConnectionsInvestmentTransactionSubtype = "pending_debit"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_QUALIFIED_DIVIDEND FinancialConnectionsInvestmentTransactionSubtype = "qualified_dividend"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_REBALANCE FinancialConnectionsInvestmentTransactionSubtype = "rebalance"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_RETURN_OF_PRINCIPAL FinancialConnectionsInvestmentTransactionSubtype = "return_of_principal"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_REQUEST FinancialConnectionsInvestmentTransactionSubtype = "request"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_SELL FinancialConnectionsInvestmentTransactionSubtype = "sell"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_SELL_SHORT FinancialConnectionsInvestmentTransactionSubtype = "sell_short"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_SEND FinancialConnectionsInvestmentTransactionSubtype = "send"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_SHORT_TERM_CAPITAL_GAIN FinancialConnectionsInvestmentTransactionSubtype = "short_term_capital_gain"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_SHORT_TERM_CAPITAL_GAIN_REINVESTMENT FinancialConnectionsInvestmentTransactionSubtype = "short_term_capital_gain_reinvestment"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_SPIN_OFF FinancialConnectionsInvestmentTransactionSubtype = "spin_off"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_SPLIT FinancialConnectionsInvestmentTransactionSubtype = "split"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_STOCK_DISTRIBUTION FinancialConnectionsInvestmentTransactionSubtype = "stock_distribution"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_TAX FinancialConnectionsInvestmentTransactionSubtype = "tax"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_TAX_WITHHELD FinancialConnectionsInvestmentTransactionSubtype = "tax_withheld"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_TRADE FinancialConnectionsInvestmentTransactionSubtype = "trade"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_TRANSFER FinancialConnectionsInvestmentTransactionSubtype = "transfer"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_TRANSFER_FEE FinancialConnectionsInvestmentTransactionSubtype = "transfer_fee"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_TRUST_FEE FinancialConnectionsInvestmentTransactionSubtype = "trust_fee"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_UNQUALIFIED_GAIN FinancialConnectionsInvestmentTransactionSubtype = "unqualified_gain"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_WITHDRAWAL FinancialConnectionsInvestmentTransactionSubtype = "withdrawal"
	FINANCIALCONNECTIONSINVESTMENTTRANSACTIONSUBTYPE_MINUS FinancialConnectionsInvestmentTransactionSubtype = "-"
)

// All allowed values of FinancialConnectionsInvestmentTransactionSubtype enum
var AllowedFinancialConnectionsInvestmentTransactionSubtypeEnumValues = []FinancialConnectionsInvestmentTransactionSubtype{
	"account_fee",
	"adjustment",
	"assignment",
	"buy",
	"buy_to_cover",
	"contribution",
	"deposit",
	"distribution",
	"dividend",
	"dividend_reinvestment",
	"exercise",
	"expire",
	"fund_fee",
	"interest",
	"interest_receivable",
	"interest_reinvestment",
	"legal_fee",
	"loan_payment",
	"long_term_capital_gain",
	"long_term_capital_gain_reinvestment",
	"management_fee",
	"margin_expense",
	"merger",
	"miscellaneous_fee",
	"non_qualified_dividend",
	"non_resident_tax",
	"pending_credit",
	"pending_debit",
	"qualified_dividend",
	"rebalance",
	"return_of_principal",
	"request",
	"sell",
	"sell_short",
	"send",
	"short_term_capital_gain",
	"short_term_capital_gain_reinvestment",
	"spin_off",
	"split",
	"stock_distribution",
	"tax",
	"tax_withheld",
	"trade",
	"transfer",
	"transfer_fee",
	"trust_fee",
	"unqualified_gain",
	"withdrawal",
	"-",
}

func (v *FinancialConnectionsInvestmentTransactionSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FinancialConnectionsInvestmentTransactionSubtype(value)
	for _, existing := range AllowedFinancialConnectionsInvestmentTransactionSubtypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FinancialConnectionsInvestmentTransactionSubtype", value)
}

// NewFinancialConnectionsInvestmentTransactionSubtypeFromValue returns a pointer to a valid FinancialConnectionsInvestmentTransactionSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFinancialConnectionsInvestmentTransactionSubtypeFromValue(v string) (*FinancialConnectionsInvestmentTransactionSubtype, error) {
	ev := FinancialConnectionsInvestmentTransactionSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FinancialConnectionsInvestmentTransactionSubtype: valid values are %v", v, AllowedFinancialConnectionsInvestmentTransactionSubtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FinancialConnectionsInvestmentTransactionSubtype) IsValid() bool {
	for _, existing := range AllowedFinancialConnectionsInvestmentTransactionSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FinancialConnectionsInvestmentTransactionSubtype value
func (v FinancialConnectionsInvestmentTransactionSubtype) Ptr() *FinancialConnectionsInvestmentTransactionSubtype {
	return &v
}

type NullableFinancialConnectionsInvestmentTransactionSubtype struct {
	value *FinancialConnectionsInvestmentTransactionSubtype
	isSet bool
}

func (v NullableFinancialConnectionsInvestmentTransactionSubtype) Get() *FinancialConnectionsInvestmentTransactionSubtype {
	return v.value
}

func (v *NullableFinancialConnectionsInvestmentTransactionSubtype) Set(val *FinancialConnectionsInvestmentTransactionSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialConnectionsInvestmentTransactionSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialConnectionsInvestmentTransactionSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialConnectionsInvestmentTransactionSubtype(val *FinancialConnectionsInvestmentTransactionSubtype) *NullableFinancialConnectionsInvestmentTransactionSubtype {
	return &NullableFinancialConnectionsInvestmentTransactionSubtype{value: val, isSet: true}
}

func (v NullableFinancialConnectionsInvestmentTransactionSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialConnectionsInvestmentTransactionSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

