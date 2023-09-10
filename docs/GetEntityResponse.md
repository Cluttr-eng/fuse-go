# GetEntityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the entity | 
**Email** | Pointer to **string** | Email of the entity | [optional] 
**Aggregators** | Pointer to [**[]Aggregator**](Aggregator.md) | These will force the user to connect through all of these aggregators | [optional] 
**InstitutionIds** | Pointer to **[]string** |  | [optional] 
**FinancialConnections** | [**[]FinancialConnectionDetails**](FinancialConnectionDetails.md) | Data needed to query data from the various aggregators | 
**RequestId** | **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | 

## Methods

### NewGetEntityResponse

`func NewGetEntityResponse(id string, financialConnections []FinancialConnectionDetails, requestId string, ) *GetEntityResponse`

NewGetEntityResponse instantiates a new GetEntityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntityResponseWithDefaults

`func NewGetEntityResponseWithDefaults() *GetEntityResponse`

NewGetEntityResponseWithDefaults instantiates a new GetEntityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEntityResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEntityResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEntityResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *GetEntityResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetEntityResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetEntityResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetEntityResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAggregators

`func (o *GetEntityResponse) GetAggregators() []Aggregator`

GetAggregators returns the Aggregators field if non-nil, zero value otherwise.

### GetAggregatorsOk

`func (o *GetEntityResponse) GetAggregatorsOk() (*[]Aggregator, bool)`

GetAggregatorsOk returns a tuple with the Aggregators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregators

`func (o *GetEntityResponse) SetAggregators(v []Aggregator)`

SetAggregators sets Aggregators field to given value.

### HasAggregators

`func (o *GetEntityResponse) HasAggregators() bool`

HasAggregators returns a boolean if a field has been set.

### GetInstitutionIds

`func (o *GetEntityResponse) GetInstitutionIds() []string`

GetInstitutionIds returns the InstitutionIds field if non-nil, zero value otherwise.

### GetInstitutionIdsOk

`func (o *GetEntityResponse) GetInstitutionIdsOk() (*[]string, bool)`

GetInstitutionIdsOk returns a tuple with the InstitutionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionIds

`func (o *GetEntityResponse) SetInstitutionIds(v []string)`

SetInstitutionIds sets InstitutionIds field to given value.

### HasInstitutionIds

`func (o *GetEntityResponse) HasInstitutionIds() bool`

HasInstitutionIds returns a boolean if a field has been set.

### GetFinancialConnections

`func (o *GetEntityResponse) GetFinancialConnections() []FinancialConnectionDetails`

GetFinancialConnections returns the FinancialConnections field if non-nil, zero value otherwise.

### GetFinancialConnectionsOk

`func (o *GetEntityResponse) GetFinancialConnectionsOk() (*[]FinancialConnectionDetails, bool)`

GetFinancialConnectionsOk returns a tuple with the FinancialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialConnections

`func (o *GetEntityResponse) SetFinancialConnections(v []FinancialConnectionDetails)`

SetFinancialConnections sets FinancialConnections field to given value.


### GetRequestId

`func (o *GetEntityResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetEntityResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetEntityResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


