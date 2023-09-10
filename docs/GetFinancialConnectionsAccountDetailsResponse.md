# GetFinancialConnectionsAccountDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountDetails** | [**[]FinancialConnectionsAccountDetails**](FinancialConnectionsAccountDetails.md) |  | 
**FinancialConnection** | [**FinancialConnectionData**](FinancialConnectionData.md) |  | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetFinancialConnectionsAccountDetailsResponse

`func NewGetFinancialConnectionsAccountDetailsResponse(accountDetails []FinancialConnectionsAccountDetails, financialConnection FinancialConnectionData, requestId string, ) *GetFinancialConnectionsAccountDetailsResponse`

NewGetFinancialConnectionsAccountDetailsResponse instantiates a new GetFinancialConnectionsAccountDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialConnectionsAccountDetailsResponseWithDefaults

`func NewGetFinancialConnectionsAccountDetailsResponseWithDefaults() *GetFinancialConnectionsAccountDetailsResponse`

NewGetFinancialConnectionsAccountDetailsResponseWithDefaults instantiates a new GetFinancialConnectionsAccountDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountDetails

`func (o *GetFinancialConnectionsAccountDetailsResponse) GetAccountDetails() []FinancialConnectionsAccountDetails`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *GetFinancialConnectionsAccountDetailsResponse) GetAccountDetailsOk() (*[]FinancialConnectionsAccountDetails, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *GetFinancialConnectionsAccountDetailsResponse) SetAccountDetails(v []FinancialConnectionsAccountDetails)`

SetAccountDetails sets AccountDetails field to given value.


### GetFinancialConnection

`func (o *GetFinancialConnectionsAccountDetailsResponse) GetFinancialConnection() FinancialConnectionData`

GetFinancialConnection returns the FinancialConnection field if non-nil, zero value otherwise.

### GetFinancialConnectionOk

`func (o *GetFinancialConnectionsAccountDetailsResponse) GetFinancialConnectionOk() (*FinancialConnectionData, bool)`

GetFinancialConnectionOk returns a tuple with the FinancialConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialConnection

`func (o *GetFinancialConnectionsAccountDetailsResponse) SetFinancialConnection(v FinancialConnectionData)`

SetFinancialConnection sets FinancialConnection field to given value.


### GetRequestId

`func (o *GetFinancialConnectionsAccountDetailsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetFinancialConnectionsAccountDetailsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetFinancialConnectionsAccountDetailsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


