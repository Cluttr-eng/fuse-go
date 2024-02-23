# FinancialConnectionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The fuse financial connection id. | 
**ConnectionStatus** | **string** | Connection status of the current financial connection | 
**ConnectionStatusUpdatedAt** | **string** | Last time the connection status was updated in ISO-8601 format. | 
**IsOauth** | **bool** | Whether this is an oauth connection | 
**Aggregator** | [**Aggregator**](Aggregator.md) |  | 
**Plaid** | Pointer to [**FinancialConnectionDetailsPlaid**](FinancialConnectionDetailsPlaid.md) |  | [optional] 
**Teller** | Pointer to [**FinancialConnectionDetailsTeller**](FinancialConnectionDetailsTeller.md) |  | [optional] 
**Mx** | Pointer to [**FinancialConnectionDetailsMx**](FinancialConnectionDetailsMx.md) |  | [optional] 
**Snaptrade** | Pointer to [**FinancialConnectionDetailsSnaptrade**](FinancialConnectionDetailsSnaptrade.md) |  | [optional] 
**Flinks** | Pointer to [**FinancialConnectionDetailsFlinks**](FinancialConnectionDetailsFlinks.md) |  | [optional] 
**Mono** | Pointer to [**FinancialConnectionDetailsMono**](FinancialConnectionDetailsMono.md) |  | [optional] 
**Truelayer** | Pointer to [**FinancialConnectionDetailsTruelayer**](FinancialConnectionDetailsTruelayer.md) |  | [optional] 
**Finverse** | Pointer to [**FinancialConnectionDetailsFinverse**](FinancialConnectionDetailsFinverse.md) |  | [optional] 
**Basiq** | Pointer to [**FinancialConnectionDetailsBasiq**](FinancialConnectionDetailsBasiq.md) |  | [optional] 
**Belvo** | Pointer to [**FinancialConnectionDetailsBelvo**](FinancialConnectionDetailsBelvo.md) |  | [optional] 
**Finicity** | Pointer to [**FinancialConnectionDetailsFinicity**](FinancialConnectionDetailsFinicity.md) |  | [optional] 
**Akoya** | Pointer to [**FinancialConnectionDetailsAkoya**](FinancialConnectionDetailsAkoya.md) |  | [optional] 
**Saltedge** | Pointer to [**FinancialConnectionDetailsSaltedge**](FinancialConnectionDetailsSaltedge.md) |  | [optional] 
**Sophtron** | Pointer to [**FinancialConnectionDetailsSophtron**](FinancialConnectionDetailsSophtron.md) |  | [optional] 

## Methods

### NewFinancialConnectionDetails

`func NewFinancialConnectionDetails(id string, connectionStatus string, connectionStatusUpdatedAt string, isOauth bool, aggregator Aggregator, ) *FinancialConnectionDetails`

NewFinancialConnectionDetails instantiates a new FinancialConnectionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialConnectionDetailsWithDefaults

`func NewFinancialConnectionDetailsWithDefaults() *FinancialConnectionDetails`

NewFinancialConnectionDetailsWithDefaults instantiates a new FinancialConnectionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FinancialConnectionDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialConnectionDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialConnectionDetails) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionStatus

`func (o *FinancialConnectionDetails) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *FinancialConnectionDetails) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *FinancialConnectionDetails) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetConnectionStatusUpdatedAt

`func (o *FinancialConnectionDetails) GetConnectionStatusUpdatedAt() string`

GetConnectionStatusUpdatedAt returns the ConnectionStatusUpdatedAt field if non-nil, zero value otherwise.

### GetConnectionStatusUpdatedAtOk

`func (o *FinancialConnectionDetails) GetConnectionStatusUpdatedAtOk() (*string, bool)`

GetConnectionStatusUpdatedAtOk returns a tuple with the ConnectionStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatusUpdatedAt

`func (o *FinancialConnectionDetails) SetConnectionStatusUpdatedAt(v string)`

SetConnectionStatusUpdatedAt sets ConnectionStatusUpdatedAt field to given value.


### GetIsOauth

`func (o *FinancialConnectionDetails) GetIsOauth() bool`

GetIsOauth returns the IsOauth field if non-nil, zero value otherwise.

### GetIsOauthOk

`func (o *FinancialConnectionDetails) GetIsOauthOk() (*bool, bool)`

GetIsOauthOk returns a tuple with the IsOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOauth

`func (o *FinancialConnectionDetails) SetIsOauth(v bool)`

SetIsOauth sets IsOauth field to given value.


### GetAggregator

`func (o *FinancialConnectionDetails) GetAggregator() Aggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *FinancialConnectionDetails) GetAggregatorOk() (*Aggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *FinancialConnectionDetails) SetAggregator(v Aggregator)`

SetAggregator sets Aggregator field to given value.


### GetPlaid

`func (o *FinancialConnectionDetails) GetPlaid() FinancialConnectionDetailsPlaid`

GetPlaid returns the Plaid field if non-nil, zero value otherwise.

### GetPlaidOk

`func (o *FinancialConnectionDetails) GetPlaidOk() (*FinancialConnectionDetailsPlaid, bool)`

GetPlaidOk returns a tuple with the Plaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaid

`func (o *FinancialConnectionDetails) SetPlaid(v FinancialConnectionDetailsPlaid)`

SetPlaid sets Plaid field to given value.

### HasPlaid

`func (o *FinancialConnectionDetails) HasPlaid() bool`

HasPlaid returns a boolean if a field has been set.

### GetTeller

`func (o *FinancialConnectionDetails) GetTeller() FinancialConnectionDetailsTeller`

GetTeller returns the Teller field if non-nil, zero value otherwise.

### GetTellerOk

`func (o *FinancialConnectionDetails) GetTellerOk() (*FinancialConnectionDetailsTeller, bool)`

GetTellerOk returns a tuple with the Teller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeller

`func (o *FinancialConnectionDetails) SetTeller(v FinancialConnectionDetailsTeller)`

SetTeller sets Teller field to given value.

### HasTeller

`func (o *FinancialConnectionDetails) HasTeller() bool`

HasTeller returns a boolean if a field has been set.

### GetMx

`func (o *FinancialConnectionDetails) GetMx() FinancialConnectionDetailsMx`

GetMx returns the Mx field if non-nil, zero value otherwise.

### GetMxOk

`func (o *FinancialConnectionDetails) GetMxOk() (*FinancialConnectionDetailsMx, bool)`

GetMxOk returns a tuple with the Mx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMx

`func (o *FinancialConnectionDetails) SetMx(v FinancialConnectionDetailsMx)`

SetMx sets Mx field to given value.

### HasMx

`func (o *FinancialConnectionDetails) HasMx() bool`

HasMx returns a boolean if a field has been set.

### GetSnaptrade

`func (o *FinancialConnectionDetails) GetSnaptrade() FinancialConnectionDetailsSnaptrade`

GetSnaptrade returns the Snaptrade field if non-nil, zero value otherwise.

### GetSnaptradeOk

`func (o *FinancialConnectionDetails) GetSnaptradeOk() (*FinancialConnectionDetailsSnaptrade, bool)`

GetSnaptradeOk returns a tuple with the Snaptrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnaptrade

`func (o *FinancialConnectionDetails) SetSnaptrade(v FinancialConnectionDetailsSnaptrade)`

SetSnaptrade sets Snaptrade field to given value.

### HasSnaptrade

`func (o *FinancialConnectionDetails) HasSnaptrade() bool`

HasSnaptrade returns a boolean if a field has been set.

### GetFlinks

`func (o *FinancialConnectionDetails) GetFlinks() FinancialConnectionDetailsFlinks`

GetFlinks returns the Flinks field if non-nil, zero value otherwise.

### GetFlinksOk

`func (o *FinancialConnectionDetails) GetFlinksOk() (*FinancialConnectionDetailsFlinks, bool)`

GetFlinksOk returns a tuple with the Flinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinks

`func (o *FinancialConnectionDetails) SetFlinks(v FinancialConnectionDetailsFlinks)`

SetFlinks sets Flinks field to given value.

### HasFlinks

`func (o *FinancialConnectionDetails) HasFlinks() bool`

HasFlinks returns a boolean if a field has been set.

### GetMono

`func (o *FinancialConnectionDetails) GetMono() FinancialConnectionDetailsMono`

GetMono returns the Mono field if non-nil, zero value otherwise.

### GetMonoOk

`func (o *FinancialConnectionDetails) GetMonoOk() (*FinancialConnectionDetailsMono, bool)`

GetMonoOk returns a tuple with the Mono field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMono

`func (o *FinancialConnectionDetails) SetMono(v FinancialConnectionDetailsMono)`

SetMono sets Mono field to given value.

### HasMono

`func (o *FinancialConnectionDetails) HasMono() bool`

HasMono returns a boolean if a field has been set.

### GetTruelayer

`func (o *FinancialConnectionDetails) GetTruelayer() FinancialConnectionDetailsTruelayer`

GetTruelayer returns the Truelayer field if non-nil, zero value otherwise.

### GetTruelayerOk

`func (o *FinancialConnectionDetails) GetTruelayerOk() (*FinancialConnectionDetailsTruelayer, bool)`

GetTruelayerOk returns a tuple with the Truelayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruelayer

`func (o *FinancialConnectionDetails) SetTruelayer(v FinancialConnectionDetailsTruelayer)`

SetTruelayer sets Truelayer field to given value.

### HasTruelayer

`func (o *FinancialConnectionDetails) HasTruelayer() bool`

HasTruelayer returns a boolean if a field has been set.

### GetFinverse

`func (o *FinancialConnectionDetails) GetFinverse() FinancialConnectionDetailsFinverse`

GetFinverse returns the Finverse field if non-nil, zero value otherwise.

### GetFinverseOk

`func (o *FinancialConnectionDetails) GetFinverseOk() (*FinancialConnectionDetailsFinverse, bool)`

GetFinverseOk returns a tuple with the Finverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinverse

`func (o *FinancialConnectionDetails) SetFinverse(v FinancialConnectionDetailsFinverse)`

SetFinverse sets Finverse field to given value.

### HasFinverse

`func (o *FinancialConnectionDetails) HasFinverse() bool`

HasFinverse returns a boolean if a field has been set.

### GetBasiq

`func (o *FinancialConnectionDetails) GetBasiq() FinancialConnectionDetailsBasiq`

GetBasiq returns the Basiq field if non-nil, zero value otherwise.

### GetBasiqOk

`func (o *FinancialConnectionDetails) GetBasiqOk() (*FinancialConnectionDetailsBasiq, bool)`

GetBasiqOk returns a tuple with the Basiq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasiq

`func (o *FinancialConnectionDetails) SetBasiq(v FinancialConnectionDetailsBasiq)`

SetBasiq sets Basiq field to given value.

### HasBasiq

`func (o *FinancialConnectionDetails) HasBasiq() bool`

HasBasiq returns a boolean if a field has been set.

### GetBelvo

`func (o *FinancialConnectionDetails) GetBelvo() FinancialConnectionDetailsBelvo`

GetBelvo returns the Belvo field if non-nil, zero value otherwise.

### GetBelvoOk

`func (o *FinancialConnectionDetails) GetBelvoOk() (*FinancialConnectionDetailsBelvo, bool)`

GetBelvoOk returns a tuple with the Belvo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelvo

`func (o *FinancialConnectionDetails) SetBelvo(v FinancialConnectionDetailsBelvo)`

SetBelvo sets Belvo field to given value.

### HasBelvo

`func (o *FinancialConnectionDetails) HasBelvo() bool`

HasBelvo returns a boolean if a field has been set.

### GetFinicity

`func (o *FinancialConnectionDetails) GetFinicity() FinancialConnectionDetailsFinicity`

GetFinicity returns the Finicity field if non-nil, zero value otherwise.

### GetFinicityOk

`func (o *FinancialConnectionDetails) GetFinicityOk() (*FinancialConnectionDetailsFinicity, bool)`

GetFinicityOk returns a tuple with the Finicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinicity

`func (o *FinancialConnectionDetails) SetFinicity(v FinancialConnectionDetailsFinicity)`

SetFinicity sets Finicity field to given value.

### HasFinicity

`func (o *FinancialConnectionDetails) HasFinicity() bool`

HasFinicity returns a boolean if a field has been set.

### GetAkoya

`func (o *FinancialConnectionDetails) GetAkoya() FinancialConnectionDetailsAkoya`

GetAkoya returns the Akoya field if non-nil, zero value otherwise.

### GetAkoyaOk

`func (o *FinancialConnectionDetails) GetAkoyaOk() (*FinancialConnectionDetailsAkoya, bool)`

GetAkoyaOk returns a tuple with the Akoya field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkoya

`func (o *FinancialConnectionDetails) SetAkoya(v FinancialConnectionDetailsAkoya)`

SetAkoya sets Akoya field to given value.

### HasAkoya

`func (o *FinancialConnectionDetails) HasAkoya() bool`

HasAkoya returns a boolean if a field has been set.

### GetSaltedge

`func (o *FinancialConnectionDetails) GetSaltedge() FinancialConnectionDetailsSaltedge`

GetSaltedge returns the Saltedge field if non-nil, zero value otherwise.

### GetSaltedgeOk

`func (o *FinancialConnectionDetails) GetSaltedgeOk() (*FinancialConnectionDetailsSaltedge, bool)`

GetSaltedgeOk returns a tuple with the Saltedge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltedge

`func (o *FinancialConnectionDetails) SetSaltedge(v FinancialConnectionDetailsSaltedge)`

SetSaltedge sets Saltedge field to given value.

### HasSaltedge

`func (o *FinancialConnectionDetails) HasSaltedge() bool`

HasSaltedge returns a boolean if a field has been set.

### GetSophtron

`func (o *FinancialConnectionDetails) GetSophtron() FinancialConnectionDetailsSophtron`

GetSophtron returns the Sophtron field if non-nil, zero value otherwise.

### GetSophtronOk

`func (o *FinancialConnectionDetails) GetSophtronOk() (*FinancialConnectionDetailsSophtron, bool)`

GetSophtronOk returns a tuple with the Sophtron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSophtron

`func (o *FinancialConnectionDetails) SetSophtron(v FinancialConnectionDetailsSophtron)`

SetSophtron sets Sophtron field to given value.

### HasSophtron

`func (o *FinancialConnectionDetails) HasSophtron() bool`

HasSophtron returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


