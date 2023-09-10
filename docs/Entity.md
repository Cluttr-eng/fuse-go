# Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the user or business account that is connecting to an institution. Use this id when calling the GET /entities/${entity_id} endpoint. | 
**Name** | Pointer to **string** | Name for the user or business account. Required for EU connections. | [optional] 
**Email** | Pointer to **string** | Email address associated with the user or business account. One of email/phone is required for EU connections. | [optional] 
**Phone** | Pointer to **string** | Phone number associated with the user or business account. One of email/phone is required for EU connections. | [optional] 

## Methods

### NewEntity

`func NewEntity(id string, ) *Entity`

NewEntity instantiates a new Entity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithDefaults

`func NewEntityWithDefaults() *Entity`

NewEntityWithDefaults instantiates a new Entity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Entity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entity) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Entity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Entity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *Entity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Entity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Entity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Entity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *Entity) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Entity) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Entity) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Entity) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


