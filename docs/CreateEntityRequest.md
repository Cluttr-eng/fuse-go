# CreateEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the entity | [optional] 
**Email** | Pointer to **string** | Email of the entity | [optional] 
**Aggregators** | Pointer to [**[]Aggregator**](Aggregator.md) | These will force the user to connect through all of these aggregators | [optional] 
**InstitutionIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateEntityRequest

`func NewCreateEntityRequest() *CreateEntityRequest`

NewCreateEntityRequest instantiates a new CreateEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntityRequestWithDefaults

`func NewCreateEntityRequestWithDefaults() *CreateEntityRequest`

NewCreateEntityRequestWithDefaults instantiates a new CreateEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateEntityRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateEntityRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateEntityRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateEntityRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *CreateEntityRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateEntityRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateEntityRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateEntityRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAggregators

`func (o *CreateEntityRequest) GetAggregators() []Aggregator`

GetAggregators returns the Aggregators field if non-nil, zero value otherwise.

### GetAggregatorsOk

`func (o *CreateEntityRequest) GetAggregatorsOk() (*[]Aggregator, bool)`

GetAggregatorsOk returns a tuple with the Aggregators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregators

`func (o *CreateEntityRequest) SetAggregators(v []Aggregator)`

SetAggregators sets Aggregators field to given value.

### HasAggregators

`func (o *CreateEntityRequest) HasAggregators() bool`

HasAggregators returns a boolean if a field has been set.

### GetInstitutionIds

`func (o *CreateEntityRequest) GetInstitutionIds() []string`

GetInstitutionIds returns the InstitutionIds field if non-nil, zero value otherwise.

### GetInstitutionIdsOk

`func (o *CreateEntityRequest) GetInstitutionIdsOk() (*[]string, bool)`

GetInstitutionIdsOk returns a tuple with the InstitutionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionIds

`func (o *CreateEntityRequest) SetInstitutionIds(v []string)`

SetInstitutionIds sets InstitutionIds field to given value.

### HasInstitutionIds

`func (o *CreateEntityRequest) HasInstitutionIds() bool`

HasInstitutionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


