# QueueStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueName** | Pointer to **string** |  | [optional] 
**PendingMessageCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueueStatus

`func NewQueueStatus() *QueueStatus`

NewQueueStatus instantiates a new QueueStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueStatusWithDefaults

`func NewQueueStatusWithDefaults() *QueueStatus`

NewQueueStatusWithDefaults instantiates a new QueueStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueName

`func (o *QueueStatus) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *QueueStatus) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *QueueStatus) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *QueueStatus) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetPendingMessageCount

`func (o *QueueStatus) GetPendingMessageCount() int64`

GetPendingMessageCount returns the PendingMessageCount field if non-nil, zero value otherwise.

### GetPendingMessageCountOk

`func (o *QueueStatus) GetPendingMessageCountOk() (*int64, bool)`

GetPendingMessageCountOk returns a tuple with the PendingMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingMessageCount

`func (o *QueueStatus) SetPendingMessageCount(v int64)`

SetPendingMessageCount sets PendingMessageCount field to given value.

### HasPendingMessageCount

`func (o *QueueStatus) HasPendingMessageCount() bool`

HasPendingMessageCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


