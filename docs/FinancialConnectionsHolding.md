# FinancialConnectionsHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteAccountId** | **string** | The remote account ID associated with this holding. | 
**CostBasis** | **float32** | The original total value of the holding, in cents, when it was purchased. The format of this value is a double. | 
**Value** | **float32** | The current market value of the holding, in cents. The format of this value is a double. | 
**Quantity** | **float32** | The number of units of the security held in this holding. | 
**InstitutionPrice** | **float32** | The price of the security, in cents, as provided by the financial institution managing the holding. The format of this value is a double. | 
**Security** | [**FinancialConnectionsInvestmentSecurity**](FinancialConnectionsInvestmentSecurity.md) |  | 

## Methods

### NewFinancialConnectionsHolding

`func NewFinancialConnectionsHolding(remoteAccountId string, costBasis float32, value float32, quantity float32, institutionPrice float32, security FinancialConnectionsInvestmentSecurity, ) *FinancialConnectionsHolding`

NewFinancialConnectionsHolding instantiates a new FinancialConnectionsHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialConnectionsHoldingWithDefaults

`func NewFinancialConnectionsHoldingWithDefaults() *FinancialConnectionsHolding`

NewFinancialConnectionsHoldingWithDefaults instantiates a new FinancialConnectionsHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteAccountId

`func (o *FinancialConnectionsHolding) GetRemoteAccountId() string`

GetRemoteAccountId returns the RemoteAccountId field if non-nil, zero value otherwise.

### GetRemoteAccountIdOk

`func (o *FinancialConnectionsHolding) GetRemoteAccountIdOk() (*string, bool)`

GetRemoteAccountIdOk returns a tuple with the RemoteAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccountId

`func (o *FinancialConnectionsHolding) SetRemoteAccountId(v string)`

SetRemoteAccountId sets RemoteAccountId field to given value.


### GetCostBasis

`func (o *FinancialConnectionsHolding) GetCostBasis() float32`

GetCostBasis returns the CostBasis field if non-nil, zero value otherwise.

### GetCostBasisOk

`func (o *FinancialConnectionsHolding) GetCostBasisOk() (*float32, bool)`

GetCostBasisOk returns a tuple with the CostBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBasis

`func (o *FinancialConnectionsHolding) SetCostBasis(v float32)`

SetCostBasis sets CostBasis field to given value.


### GetValue

`func (o *FinancialConnectionsHolding) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FinancialConnectionsHolding) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FinancialConnectionsHolding) SetValue(v float32)`

SetValue sets Value field to given value.


### GetQuantity

`func (o *FinancialConnectionsHolding) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *FinancialConnectionsHolding) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *FinancialConnectionsHolding) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetInstitutionPrice

`func (o *FinancialConnectionsHolding) GetInstitutionPrice() float32`

GetInstitutionPrice returns the InstitutionPrice field if non-nil, zero value otherwise.

### GetInstitutionPriceOk

`func (o *FinancialConnectionsHolding) GetInstitutionPriceOk() (*float32, bool)`

GetInstitutionPriceOk returns a tuple with the InstitutionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionPrice

`func (o *FinancialConnectionsHolding) SetInstitutionPrice(v float32)`

SetInstitutionPrice sets InstitutionPrice field to given value.


### GetSecurity

`func (o *FinancialConnectionsHolding) GetSecurity() FinancialConnectionsInvestmentSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *FinancialConnectionsHolding) GetSecurityOk() (*FinancialConnectionsInvestmentSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *FinancialConnectionsHolding) SetSecurity(v FinancialConnectionsInvestmentSecurity)`

SetSecurity sets Security field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


