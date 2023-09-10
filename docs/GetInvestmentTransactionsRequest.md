# GetInvestmentTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access token for authentication | 
**StartDate** | **string** | The earliest date for which data should be returned. Dates should be formatted as YYYY-MM-DD. | 
**EndDate** | **string** | The latest date for which data should be returned. Dates should be formatted as YYYY-MM-DD. | 
**Page** | **int32** | Specify current page. | 
**RecordsPerPage** | **int32** | Number of items per page. | [default to 25]
**Options** | Pointer to [**GetInvestmentTransactionsRequestOptions**](GetInvestmentTransactionsRequestOptions.md) |  | [optional] 

## Methods

### NewGetInvestmentTransactionsRequest

`func NewGetInvestmentTransactionsRequest(accessToken string, startDate string, endDate string, page int32, recordsPerPage int32, ) *GetInvestmentTransactionsRequest`

NewGetInvestmentTransactionsRequest instantiates a new GetInvestmentTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvestmentTransactionsRequestWithDefaults

`func NewGetInvestmentTransactionsRequestWithDefaults() *GetInvestmentTransactionsRequest`

NewGetInvestmentTransactionsRequestWithDefaults instantiates a new GetInvestmentTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GetInvestmentTransactionsRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GetInvestmentTransactionsRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GetInvestmentTransactionsRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetStartDate

`func (o *GetInvestmentTransactionsRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetInvestmentTransactionsRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetInvestmentTransactionsRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *GetInvestmentTransactionsRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetInvestmentTransactionsRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetInvestmentTransactionsRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetPage

`func (o *GetInvestmentTransactionsRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetInvestmentTransactionsRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetInvestmentTransactionsRequest) SetPage(v int32)`

SetPage sets Page field to given value.


### GetRecordsPerPage

`func (o *GetInvestmentTransactionsRequest) GetRecordsPerPage() int32`

GetRecordsPerPage returns the RecordsPerPage field if non-nil, zero value otherwise.

### GetRecordsPerPageOk

`func (o *GetInvestmentTransactionsRequest) GetRecordsPerPageOk() (*int32, bool)`

GetRecordsPerPageOk returns a tuple with the RecordsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsPerPage

`func (o *GetInvestmentTransactionsRequest) SetRecordsPerPage(v int32)`

SetRecordsPerPage sets RecordsPerPage field to given value.


### GetOptions

`func (o *GetInvestmentTransactionsRequest) GetOptions() GetInvestmentTransactionsRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetInvestmentTransactionsRequest) GetOptionsOk() (*GetInvestmentTransactionsRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetInvestmentTransactionsRequest) SetOptions(v GetInvestmentTransactionsRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetInvestmentTransactionsRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


