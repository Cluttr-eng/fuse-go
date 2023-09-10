# UpdateEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email of the entity | [optional] 
**Aggregators** | Pointer to [**[]Aggregator**](Aggregator.md) | These will force the user to connect through all of these aggregators | [optional] 
**InstitutionIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateEntityRequest

`func NewUpdateEntityRequest() *UpdateEntityRequest`

NewUpdateEntityRequest instantiates a new UpdateEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEntityRequestWithDefaults

`func NewUpdateEntityRequestWithDefaults() *UpdateEntityRequest`

NewUpdateEntityRequestWithDefaults instantiates a new UpdateEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateEntityRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateEntityRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateEntityRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateEntityRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAggregators

`func (o *UpdateEntityRequest) GetAggregators() []Aggregator`

GetAggregators returns the Aggregators field if non-nil, zero value otherwise.

### GetAggregatorsOk

`func (o *UpdateEntityRequest) GetAggregatorsOk() (*[]Aggregator, bool)`

GetAggregatorsOk returns a tuple with the Aggregators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregators

`func (o *UpdateEntityRequest) SetAggregators(v []Aggregator)`

SetAggregators sets Aggregators field to given value.

### HasAggregators

`func (o *UpdateEntityRequest) HasAggregators() bool`

HasAggregators returns a boolean if a field has been set.

### GetInstitutionIds

`func (o *UpdateEntityRequest) GetInstitutionIds() []string`

GetInstitutionIds returns the InstitutionIds field if non-nil, zero value otherwise.

### GetInstitutionIdsOk

`func (o *UpdateEntityRequest) GetInstitutionIdsOk() (*[]string, bool)`

GetInstitutionIdsOk returns a tuple with the InstitutionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionIds

`func (o *UpdateEntityRequest) SetInstitutionIds(v []string)`

SetInstitutionIds sets InstitutionIds field to given value.

### HasInstitutionIds

`func (o *UpdateEntityRequest) HasInstitutionIds() bool`

HasInstitutionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


