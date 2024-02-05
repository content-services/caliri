# ExceptionMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayMessage** | Pointer to **string** |  | [optional] 
**RequestUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewExceptionMessage

`func NewExceptionMessage() *ExceptionMessage`

NewExceptionMessage instantiates a new ExceptionMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionMessageWithDefaults

`func NewExceptionMessageWithDefaults() *ExceptionMessage`

NewExceptionMessageWithDefaults instantiates a new ExceptionMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayMessage

`func (o *ExceptionMessage) GetDisplayMessage() string`

GetDisplayMessage returns the DisplayMessage field if non-nil, zero value otherwise.

### GetDisplayMessageOk

`func (o *ExceptionMessage) GetDisplayMessageOk() (*string, bool)`

GetDisplayMessageOk returns a tuple with the DisplayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMessage

`func (o *ExceptionMessage) SetDisplayMessage(v string)`

SetDisplayMessage sets DisplayMessage field to given value.

### HasDisplayMessage

`func (o *ExceptionMessage) HasDisplayMessage() bool`

HasDisplayMessage returns a boolean if a field has been set.

### GetRequestUuid

`func (o *ExceptionMessage) GetRequestUuid() string`

GetRequestUuid returns the RequestUuid field if non-nil, zero value otherwise.

### GetRequestUuidOk

`func (o *ExceptionMessage) GetRequestUuidOk() (*string, bool)`

GetRequestUuidOk returns a tuple with the RequestUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUuid

`func (o *ExceptionMessage) SetRequestUuid(v string)`

SetRequestUuid sets RequestUuid field to given value.

### HasRequestUuid

`func (o *ExceptionMessage) HasRequestUuid() bool`

HasRequestUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


