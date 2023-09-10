# UpdateEntityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the entity | [optional] 
**Email** | Pointer to **string** | Email of the entity | [optional] 
**Aggregators** | Pointer to [**[]Aggregator**](Aggregator.md) | These will force the user to connect through all of these aggregators | [optional] 
**InstitutionIds** | Pointer to **[]string** |  | [optional] 
**RequestId** | Pointer to **string** | An identifier that is exclusive to the request and can serve as a means for investigating and resolving issues. | [optional] 

## Methods

### NewUpdateEntityResponse

`func NewUpdateEntityResponse() *UpdateEntityResponse`

NewUpdateEntityResponse instantiates a new UpdateEntityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEntityResponseWithDefaults

`func NewUpdateEntityResponseWithDefaults() *UpdateEntityResponse`

NewUpdateEntityResponseWithDefaults instantiates a new UpdateEntityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateEntityResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateEntityResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateEntityResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateEntityResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateEntityResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateEntityResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateEntityResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateEntityResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAggregators

`func (o *UpdateEntityResponse) GetAggregators() []Aggregator`

GetAggregators returns the Aggregators field if non-nil, zero value otherwise.

### GetAggregatorsOk

`func (o *UpdateEntityResponse) GetAggregatorsOk() (*[]Aggregator, bool)`

GetAggregatorsOk returns a tuple with the Aggregators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregators

`func (o *UpdateEntityResponse) SetAggregators(v []Aggregator)`

SetAggregators sets Aggregators field to given value.

### HasAggregators

`func (o *UpdateEntityResponse) HasAggregators() bool`

HasAggregators returns a boolean if a field has been set.

### GetInstitutionIds

`func (o *UpdateEntityResponse) GetInstitutionIds() []string`

GetInstitutionIds returns the InstitutionIds field if non-nil, zero value otherwise.

### GetInstitutionIdsOk

`func (o *UpdateEntityResponse) GetInstitutionIdsOk() (*[]string, bool)`

GetInstitutionIdsOk returns a tuple with the InstitutionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionIds

`func (o *UpdateEntityResponse) SetInstitutionIds(v []string)`

SetInstitutionIds sets InstitutionIds field to given value.

### HasInstitutionIds

`func (o *UpdateEntityResponse) HasInstitutionIds() bool`

HasInstitutionIds returns a boolean if a field has been set.

### GetRequestId

`func (o *UpdateEntityResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UpdateEntityResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UpdateEntityResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *UpdateEntityResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


