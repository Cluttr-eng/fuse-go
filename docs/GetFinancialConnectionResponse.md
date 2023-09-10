# GetFinancialConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinancialConnection** | [**FinancialConnectionDetails**](FinancialConnectionDetails.md) |  | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetFinancialConnectionResponse

`func NewGetFinancialConnectionResponse(financialConnection FinancialConnectionDetails, requestId string, ) *GetFinancialConnectionResponse`

NewGetFinancialConnectionResponse instantiates a new GetFinancialConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinancialConnectionResponseWithDefaults

`func NewGetFinancialConnectionResponseWithDefaults() *GetFinancialConnectionResponse`

NewGetFinancialConnectionResponseWithDefaults instantiates a new GetFinancialConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinancialConnection

`func (o *GetFinancialConnectionResponse) GetFinancialConnection() FinancialConnectionDetails`

GetFinancialConnection returns the FinancialConnection field if non-nil, zero value otherwise.

### GetFinancialConnectionOk

`func (o *GetFinancialConnectionResponse) GetFinancialConnectionOk() (*FinancialConnectionDetails, bool)`

GetFinancialConnectionOk returns a tuple with the FinancialConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialConnection

`func (o *GetFinancialConnectionResponse) SetFinancialConnection(v FinancialConnectionDetails)`

SetFinancialConnection sets FinancialConnection field to given value.


### GetRequestId

`func (o *GetFinancialConnectionResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetFinancialConnectionResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetFinancialConnectionResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


