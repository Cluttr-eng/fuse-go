# CreateEntityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the entity | [optional] 
**Email** | Pointer to **string** | Email of the entity | [optional] 
**Aggregators** | Pointer to [**[]Aggregator**](Aggregator.md) | These will force the user to connect through all of these aggregators | [optional] 
**InstitutionIds** | Pointer to **[]string** |  | [optional] 
**RequestId** | Pointer to **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | [optional] 

## Methods

### NewCreateEntityResponse

`func NewCreateEntityResponse() *CreateEntityResponse`

NewCreateEntityResponse instantiates a new CreateEntityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntityResponseWithDefaults

`func NewCreateEntityResponseWithDefaults() *CreateEntityResponse`

NewCreateEntityResponseWithDefaults instantiates a new CreateEntityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateEntityResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateEntityResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateEntityResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateEntityResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *CreateEntityResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateEntityResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateEntityResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateEntityResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAggregators

`func (o *CreateEntityResponse) GetAggregators() []Aggregator`

GetAggregators returns the Aggregators field if non-nil, zero value otherwise.

### GetAggregatorsOk

`func (o *CreateEntityResponse) GetAggregatorsOk() (*[]Aggregator, bool)`

GetAggregatorsOk returns a tuple with the Aggregators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregators

`func (o *CreateEntityResponse) SetAggregators(v []Aggregator)`

SetAggregators sets Aggregators field to given value.

### HasAggregators

`func (o *CreateEntityResponse) HasAggregators() bool`

HasAggregators returns a boolean if a field has been set.

### GetInstitutionIds

`func (o *CreateEntityResponse) GetInstitutionIds() []string`

GetInstitutionIds returns the InstitutionIds field if non-nil, zero value otherwise.

### GetInstitutionIdsOk

`func (o *CreateEntityResponse) GetInstitutionIdsOk() (*[]string, bool)`

GetInstitutionIdsOk returns a tuple with the InstitutionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionIds

`func (o *CreateEntityResponse) SetInstitutionIds(v []string)`

SetInstitutionIds sets InstitutionIds field to given value.

### HasInstitutionIds

`func (o *CreateEntityResponse) HasInstitutionIds() bool`

HasInstitutionIds returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateEntityResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateEntityResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateEntityResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateEntityResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


