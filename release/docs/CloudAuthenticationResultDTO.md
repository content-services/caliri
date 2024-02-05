# CloudAuthenticationResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerKey** | Pointer to **string** |  | [optional] 
**AnonymousConsumerUuid** | Pointer to **string** |  | [optional] 
**Token** | **string** |  | 
**TokenType** | **string** |  | 

## Methods

### NewCloudAuthenticationResultDTO

`func NewCloudAuthenticationResultDTO(token string, tokenType string, ) *CloudAuthenticationResultDTO`

NewCloudAuthenticationResultDTO instantiates a new CloudAuthenticationResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAuthenticationResultDTOWithDefaults

`func NewCloudAuthenticationResultDTOWithDefaults() *CloudAuthenticationResultDTO`

NewCloudAuthenticationResultDTOWithDefaults instantiates a new CloudAuthenticationResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerKey

`func (o *CloudAuthenticationResultDTO) GetOwnerKey() string`

GetOwnerKey returns the OwnerKey field if non-nil, zero value otherwise.

### GetOwnerKeyOk

`func (o *CloudAuthenticationResultDTO) GetOwnerKeyOk() (*string, bool)`

GetOwnerKeyOk returns a tuple with the OwnerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerKey

`func (o *CloudAuthenticationResultDTO) SetOwnerKey(v string)`

SetOwnerKey sets OwnerKey field to given value.

### HasOwnerKey

`func (o *CloudAuthenticationResultDTO) HasOwnerKey() bool`

HasOwnerKey returns a boolean if a field has been set.

### GetAnonymousConsumerUuid

`func (o *CloudAuthenticationResultDTO) GetAnonymousConsumerUuid() string`

GetAnonymousConsumerUuid returns the AnonymousConsumerUuid field if non-nil, zero value otherwise.

### GetAnonymousConsumerUuidOk

`func (o *CloudAuthenticationResultDTO) GetAnonymousConsumerUuidOk() (*string, bool)`

GetAnonymousConsumerUuidOk returns a tuple with the AnonymousConsumerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousConsumerUuid

`func (o *CloudAuthenticationResultDTO) SetAnonymousConsumerUuid(v string)`

SetAnonymousConsumerUuid sets AnonymousConsumerUuid field to given value.

### HasAnonymousConsumerUuid

`func (o *CloudAuthenticationResultDTO) HasAnonymousConsumerUuid() bool`

HasAnonymousConsumerUuid returns a boolean if a field has been set.

### GetToken

`func (o *CloudAuthenticationResultDTO) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CloudAuthenticationResultDTO) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CloudAuthenticationResultDTO) SetToken(v string)`

SetToken sets Token field to given value.


### GetTokenType

`func (o *CloudAuthenticationResultDTO) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *CloudAuthenticationResultDTO) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *CloudAuthenticationResultDTO) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


