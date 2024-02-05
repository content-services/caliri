# AsyncJobStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**Executor** | Pointer to **string** |  | [optional] 
**Principal** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**PreviousState** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Attempts** | Pointer to **int32** |  | [optional] 
**MaxAttempts** | Pointer to **int32** |  | [optional] 
**StatusPath** | Pointer to **string** |  | [optional] 
**ResultData** | Pointer to **map[string]interface{}** | May contain a single message or job data as JSON | [optional] 

## Methods

### NewAsyncJobStatusDTO

`func NewAsyncJobStatusDTO() *AsyncJobStatusDTO`

NewAsyncJobStatusDTO instantiates a new AsyncJobStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncJobStatusDTOWithDefaults

`func NewAsyncJobStatusDTOWithDefaults() *AsyncJobStatusDTO`

NewAsyncJobStatusDTOWithDefaults instantiates a new AsyncJobStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AsyncJobStatusDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AsyncJobStatusDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AsyncJobStatusDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AsyncJobStatusDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *AsyncJobStatusDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AsyncJobStatusDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AsyncJobStatusDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AsyncJobStatusDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *AsyncJobStatusDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AsyncJobStatusDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AsyncJobStatusDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AsyncJobStatusDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *AsyncJobStatusDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AsyncJobStatusDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AsyncJobStatusDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AsyncJobStatusDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *AsyncJobStatusDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AsyncJobStatusDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AsyncJobStatusDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AsyncJobStatusDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroup

`func (o *AsyncJobStatusDTO) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AsyncJobStatusDTO) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AsyncJobStatusDTO) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AsyncJobStatusDTO) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetOrigin

`func (o *AsyncJobStatusDTO) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *AsyncJobStatusDTO) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *AsyncJobStatusDTO) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *AsyncJobStatusDTO) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetExecutor

`func (o *AsyncJobStatusDTO) GetExecutor() string`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *AsyncJobStatusDTO) GetExecutorOk() (*string, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *AsyncJobStatusDTO) SetExecutor(v string)`

SetExecutor sets Executor field to given value.

### HasExecutor

`func (o *AsyncJobStatusDTO) HasExecutor() bool`

HasExecutor returns a boolean if a field has been set.

### GetPrincipal

`func (o *AsyncJobStatusDTO) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *AsyncJobStatusDTO) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *AsyncJobStatusDTO) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *AsyncJobStatusDTO) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetState

`func (o *AsyncJobStatusDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AsyncJobStatusDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AsyncJobStatusDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AsyncJobStatusDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPreviousState

`func (o *AsyncJobStatusDTO) GetPreviousState() string`

GetPreviousState returns the PreviousState field if non-nil, zero value otherwise.

### GetPreviousStateOk

`func (o *AsyncJobStatusDTO) GetPreviousStateOk() (*string, bool)`

GetPreviousStateOk returns a tuple with the PreviousState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousState

`func (o *AsyncJobStatusDTO) SetPreviousState(v string)`

SetPreviousState sets PreviousState field to given value.

### HasPreviousState

`func (o *AsyncJobStatusDTO) HasPreviousState() bool`

HasPreviousState returns a boolean if a field has been set.

### GetStartTime

`func (o *AsyncJobStatusDTO) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AsyncJobStatusDTO) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AsyncJobStatusDTO) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AsyncJobStatusDTO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *AsyncJobStatusDTO) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AsyncJobStatusDTO) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AsyncJobStatusDTO) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AsyncJobStatusDTO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetAttempts

`func (o *AsyncJobStatusDTO) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *AsyncJobStatusDTO) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *AsyncJobStatusDTO) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *AsyncJobStatusDTO) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *AsyncJobStatusDTO) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *AsyncJobStatusDTO) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *AsyncJobStatusDTO) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *AsyncJobStatusDTO) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetStatusPath

`func (o *AsyncJobStatusDTO) GetStatusPath() string`

GetStatusPath returns the StatusPath field if non-nil, zero value otherwise.

### GetStatusPathOk

`func (o *AsyncJobStatusDTO) GetStatusPathOk() (*string, bool)`

GetStatusPathOk returns a tuple with the StatusPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPath

`func (o *AsyncJobStatusDTO) SetStatusPath(v string)`

SetStatusPath sets StatusPath field to given value.

### HasStatusPath

`func (o *AsyncJobStatusDTO) HasStatusPath() bool`

HasStatusPath returns a boolean if a field has been set.

### GetResultData

`func (o *AsyncJobStatusDTO) GetResultData() map[string]interface{}`

GetResultData returns the ResultData field if non-nil, zero value otherwise.

### GetResultDataOk

`func (o *AsyncJobStatusDTO) GetResultDataOk() (*map[string]interface{}, bool)`

GetResultDataOk returns a tuple with the ResultData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultData

`func (o *AsyncJobStatusDTO) SetResultData(v map[string]interface{})`

SetResultData sets ResultData field to given value.

### HasResultData

`func (o *AsyncJobStatusDTO) HasResultData() bool`

HasResultData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


