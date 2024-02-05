# PoolDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 
**ActiveSubscription** | Pointer to **bool** |  | [optional] 
**SourceEntitlement** | Pointer to [**NestedEntitlementDTO**](NestedEntitlementDTO.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Attributes** | Pointer to [**[]AttributeDTO**](AttributeDTO.md) |  | [optional] 
**RestrictedToUsername** | Pointer to **string** |  | [optional] 
**ContractNumber** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**OrderNumber** | Pointer to **string** |  | [optional] 
**Consumed** | Pointer to **int64** |  | [optional] 
**Exported** | Pointer to **int64** |  | [optional] 
**Branding** | Pointer to [**[]BrandingDTO**](BrandingDTO.md) |  | [optional] 
**CalculatedAttributes** | Pointer to **map[string]string** |  | [optional] 
**UpstreamPoolId** | Pointer to **string** |  | [optional] 
**UpstreamEntitlementId** | Pointer to **string** |  | [optional] 
**UpstreamConsumerId** | Pointer to **string** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**ProductAttributes** | Pointer to [**[]AttributeDTO**](AttributeDTO.md) |  | [optional] 
**StackId** | Pointer to **string** |  | [optional] 
**Stacked** | Pointer to **bool** |  | [optional] 
**SourceStackId** | Pointer to **string** |  | [optional] 
**DevelopmentPool** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**DerivedProductAttributes** | Pointer to [**[]AttributeDTO**](AttributeDTO.md) |  | [optional] 
**DerivedProductId** | Pointer to **string** |  | [optional] 
**DerivedProductName** | Pointer to **string** |  | [optional] 
**ProvidedProducts** | Pointer to [**[]ProvidedProductDTO**](ProvidedProductDTO.md) |  | [optional] 
**DerivedProvidedProducts** | Pointer to [**[]ProvidedProductDTO**](ProvidedProductDTO.md) |  | [optional] 
**SubscriptionSubKey** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 

## Methods

### NewPoolDTO

`func NewPoolDTO() *PoolDTO`

NewPoolDTO instantiates a new PoolDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDTOWithDefaults

`func NewPoolDTOWithDefaults() *PoolDTO`

NewPoolDTOWithDefaults instantiates a new PoolDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PoolDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PoolDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PoolDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PoolDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *PoolDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PoolDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PoolDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PoolDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *PoolDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PoolDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PoolDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoolDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoolDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PoolDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *PoolDTO) GetOwner() NestedOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PoolDTO) GetOwnerOk() (*NestedOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PoolDTO) SetOwner(v NestedOwnerDTO)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PoolDTO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetActiveSubscription

`func (o *PoolDTO) GetActiveSubscription() bool`

GetActiveSubscription returns the ActiveSubscription field if non-nil, zero value otherwise.

### GetActiveSubscriptionOk

`func (o *PoolDTO) GetActiveSubscriptionOk() (*bool, bool)`

GetActiveSubscriptionOk returns a tuple with the ActiveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSubscription

`func (o *PoolDTO) SetActiveSubscription(v bool)`

SetActiveSubscription sets ActiveSubscription field to given value.

### HasActiveSubscription

`func (o *PoolDTO) HasActiveSubscription() bool`

HasActiveSubscription returns a boolean if a field has been set.

### GetSourceEntitlement

`func (o *PoolDTO) GetSourceEntitlement() NestedEntitlementDTO`

GetSourceEntitlement returns the SourceEntitlement field if non-nil, zero value otherwise.

### GetSourceEntitlementOk

`func (o *PoolDTO) GetSourceEntitlementOk() (*NestedEntitlementDTO, bool)`

GetSourceEntitlementOk returns a tuple with the SourceEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntitlement

`func (o *PoolDTO) SetSourceEntitlement(v NestedEntitlementDTO)`

SetSourceEntitlement sets SourceEntitlement field to given value.

### HasSourceEntitlement

`func (o *PoolDTO) HasSourceEntitlement() bool`

HasSourceEntitlement returns a boolean if a field has been set.

### GetQuantity

`func (o *PoolDTO) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PoolDTO) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PoolDTO) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PoolDTO) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartDate

`func (o *PoolDTO) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PoolDTO) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PoolDTO) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PoolDTO) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PoolDTO) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PoolDTO) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PoolDTO) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PoolDTO) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAttributes

`func (o *PoolDTO) GetAttributes() []AttributeDTO`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PoolDTO) GetAttributesOk() (*[]AttributeDTO, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PoolDTO) SetAttributes(v []AttributeDTO)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PoolDTO) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRestrictedToUsername

`func (o *PoolDTO) GetRestrictedToUsername() string`

GetRestrictedToUsername returns the RestrictedToUsername field if non-nil, zero value otherwise.

### GetRestrictedToUsernameOk

`func (o *PoolDTO) GetRestrictedToUsernameOk() (*string, bool)`

GetRestrictedToUsernameOk returns a tuple with the RestrictedToUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToUsername

`func (o *PoolDTO) SetRestrictedToUsername(v string)`

SetRestrictedToUsername sets RestrictedToUsername field to given value.

### HasRestrictedToUsername

`func (o *PoolDTO) HasRestrictedToUsername() bool`

HasRestrictedToUsername returns a boolean if a field has been set.

### GetContractNumber

`func (o *PoolDTO) GetContractNumber() string`

GetContractNumber returns the ContractNumber field if non-nil, zero value otherwise.

### GetContractNumberOk

`func (o *PoolDTO) GetContractNumberOk() (*string, bool)`

GetContractNumberOk returns a tuple with the ContractNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractNumber

`func (o *PoolDTO) SetContractNumber(v string)`

SetContractNumber sets ContractNumber field to given value.

### HasContractNumber

`func (o *PoolDTO) HasContractNumber() bool`

HasContractNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *PoolDTO) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PoolDTO) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PoolDTO) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PoolDTO) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetOrderNumber

`func (o *PoolDTO) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *PoolDTO) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *PoolDTO) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *PoolDTO) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetConsumed

`func (o *PoolDTO) GetConsumed() int64`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *PoolDTO) GetConsumedOk() (*int64, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *PoolDTO) SetConsumed(v int64)`

SetConsumed sets Consumed field to given value.

### HasConsumed

`func (o *PoolDTO) HasConsumed() bool`

HasConsumed returns a boolean if a field has been set.

### GetExported

`func (o *PoolDTO) GetExported() int64`

GetExported returns the Exported field if non-nil, zero value otherwise.

### GetExportedOk

`func (o *PoolDTO) GetExportedOk() (*int64, bool)`

GetExportedOk returns a tuple with the Exported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExported

`func (o *PoolDTO) SetExported(v int64)`

SetExported sets Exported field to given value.

### HasExported

`func (o *PoolDTO) HasExported() bool`

HasExported returns a boolean if a field has been set.

### GetBranding

`func (o *PoolDTO) GetBranding() []BrandingDTO`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *PoolDTO) GetBrandingOk() (*[]BrandingDTO, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *PoolDTO) SetBranding(v []BrandingDTO)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *PoolDTO) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetCalculatedAttributes

`func (o *PoolDTO) GetCalculatedAttributes() map[string]string`

GetCalculatedAttributes returns the CalculatedAttributes field if non-nil, zero value otherwise.

### GetCalculatedAttributesOk

`func (o *PoolDTO) GetCalculatedAttributesOk() (*map[string]string, bool)`

GetCalculatedAttributesOk returns a tuple with the CalculatedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedAttributes

`func (o *PoolDTO) SetCalculatedAttributes(v map[string]string)`

SetCalculatedAttributes sets CalculatedAttributes field to given value.

### HasCalculatedAttributes

`func (o *PoolDTO) HasCalculatedAttributes() bool`

HasCalculatedAttributes returns a boolean if a field has been set.

### GetUpstreamPoolId

`func (o *PoolDTO) GetUpstreamPoolId() string`

GetUpstreamPoolId returns the UpstreamPoolId field if non-nil, zero value otherwise.

### GetUpstreamPoolIdOk

`func (o *PoolDTO) GetUpstreamPoolIdOk() (*string, bool)`

GetUpstreamPoolIdOk returns a tuple with the UpstreamPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamPoolId

`func (o *PoolDTO) SetUpstreamPoolId(v string)`

SetUpstreamPoolId sets UpstreamPoolId field to given value.

### HasUpstreamPoolId

`func (o *PoolDTO) HasUpstreamPoolId() bool`

HasUpstreamPoolId returns a boolean if a field has been set.

### GetUpstreamEntitlementId

`func (o *PoolDTO) GetUpstreamEntitlementId() string`

GetUpstreamEntitlementId returns the UpstreamEntitlementId field if non-nil, zero value otherwise.

### GetUpstreamEntitlementIdOk

`func (o *PoolDTO) GetUpstreamEntitlementIdOk() (*string, bool)`

GetUpstreamEntitlementIdOk returns a tuple with the UpstreamEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamEntitlementId

`func (o *PoolDTO) SetUpstreamEntitlementId(v string)`

SetUpstreamEntitlementId sets UpstreamEntitlementId field to given value.

### HasUpstreamEntitlementId

`func (o *PoolDTO) HasUpstreamEntitlementId() bool`

HasUpstreamEntitlementId returns a boolean if a field has been set.

### GetUpstreamConsumerId

`func (o *PoolDTO) GetUpstreamConsumerId() string`

GetUpstreamConsumerId returns the UpstreamConsumerId field if non-nil, zero value otherwise.

### GetUpstreamConsumerIdOk

`func (o *PoolDTO) GetUpstreamConsumerIdOk() (*string, bool)`

GetUpstreamConsumerIdOk returns a tuple with the UpstreamConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamConsumerId

`func (o *PoolDTO) SetUpstreamConsumerId(v string)`

SetUpstreamConsumerId sets UpstreamConsumerId field to given value.

### HasUpstreamConsumerId

`func (o *PoolDTO) HasUpstreamConsumerId() bool`

HasUpstreamConsumerId returns a boolean if a field has been set.

### GetProductName

`func (o *PoolDTO) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *PoolDTO) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *PoolDTO) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *PoolDTO) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductId

`func (o *PoolDTO) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PoolDTO) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PoolDTO) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PoolDTO) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductAttributes

`func (o *PoolDTO) GetProductAttributes() []AttributeDTO`

GetProductAttributes returns the ProductAttributes field if non-nil, zero value otherwise.

### GetProductAttributesOk

`func (o *PoolDTO) GetProductAttributesOk() (*[]AttributeDTO, bool)`

GetProductAttributesOk returns a tuple with the ProductAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAttributes

`func (o *PoolDTO) SetProductAttributes(v []AttributeDTO)`

SetProductAttributes sets ProductAttributes field to given value.

### HasProductAttributes

`func (o *PoolDTO) HasProductAttributes() bool`

HasProductAttributes returns a boolean if a field has been set.

### GetStackId

`func (o *PoolDTO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *PoolDTO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *PoolDTO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *PoolDTO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStacked

`func (o *PoolDTO) GetStacked() bool`

GetStacked returns the Stacked field if non-nil, zero value otherwise.

### GetStackedOk

`func (o *PoolDTO) GetStackedOk() (*bool, bool)`

GetStackedOk returns a tuple with the Stacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacked

`func (o *PoolDTO) SetStacked(v bool)`

SetStacked sets Stacked field to given value.

### HasStacked

`func (o *PoolDTO) HasStacked() bool`

HasStacked returns a boolean if a field has been set.

### GetSourceStackId

`func (o *PoolDTO) GetSourceStackId() string`

GetSourceStackId returns the SourceStackId field if non-nil, zero value otherwise.

### GetSourceStackIdOk

`func (o *PoolDTO) GetSourceStackIdOk() (*string, bool)`

GetSourceStackIdOk returns a tuple with the SourceStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStackId

`func (o *PoolDTO) SetSourceStackId(v string)`

SetSourceStackId sets SourceStackId field to given value.

### HasSourceStackId

`func (o *PoolDTO) HasSourceStackId() bool`

HasSourceStackId returns a boolean if a field has been set.

### GetDevelopmentPool

`func (o *PoolDTO) GetDevelopmentPool() bool`

GetDevelopmentPool returns the DevelopmentPool field if non-nil, zero value otherwise.

### GetDevelopmentPoolOk

`func (o *PoolDTO) GetDevelopmentPoolOk() (*bool, bool)`

GetDevelopmentPoolOk returns a tuple with the DevelopmentPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopmentPool

`func (o *PoolDTO) SetDevelopmentPool(v bool)`

SetDevelopmentPool sets DevelopmentPool field to given value.

### HasDevelopmentPool

`func (o *PoolDTO) HasDevelopmentPool() bool`

HasDevelopmentPool returns a boolean if a field has been set.

### GetHref

`func (o *PoolDTO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PoolDTO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PoolDTO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PoolDTO) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetDerivedProductAttributes

`func (o *PoolDTO) GetDerivedProductAttributes() []AttributeDTO`

GetDerivedProductAttributes returns the DerivedProductAttributes field if non-nil, zero value otherwise.

### GetDerivedProductAttributesOk

`func (o *PoolDTO) GetDerivedProductAttributesOk() (*[]AttributeDTO, bool)`

GetDerivedProductAttributesOk returns a tuple with the DerivedProductAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedProductAttributes

`func (o *PoolDTO) SetDerivedProductAttributes(v []AttributeDTO)`

SetDerivedProductAttributes sets DerivedProductAttributes field to given value.

### HasDerivedProductAttributes

`func (o *PoolDTO) HasDerivedProductAttributes() bool`

HasDerivedProductAttributes returns a boolean if a field has been set.

### GetDerivedProductId

`func (o *PoolDTO) GetDerivedProductId() string`

GetDerivedProductId returns the DerivedProductId field if non-nil, zero value otherwise.

### GetDerivedProductIdOk

`func (o *PoolDTO) GetDerivedProductIdOk() (*string, bool)`

GetDerivedProductIdOk returns a tuple with the DerivedProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedProductId

`func (o *PoolDTO) SetDerivedProductId(v string)`

SetDerivedProductId sets DerivedProductId field to given value.

### HasDerivedProductId

`func (o *PoolDTO) HasDerivedProductId() bool`

HasDerivedProductId returns a boolean if a field has been set.

### GetDerivedProductName

`func (o *PoolDTO) GetDerivedProductName() string`

GetDerivedProductName returns the DerivedProductName field if non-nil, zero value otherwise.

### GetDerivedProductNameOk

`func (o *PoolDTO) GetDerivedProductNameOk() (*string, bool)`

GetDerivedProductNameOk returns a tuple with the DerivedProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedProductName

`func (o *PoolDTO) SetDerivedProductName(v string)`

SetDerivedProductName sets DerivedProductName field to given value.

### HasDerivedProductName

`func (o *PoolDTO) HasDerivedProductName() bool`

HasDerivedProductName returns a boolean if a field has been set.

### GetProvidedProducts

`func (o *PoolDTO) GetProvidedProducts() []ProvidedProductDTO`

GetProvidedProducts returns the ProvidedProducts field if non-nil, zero value otherwise.

### GetProvidedProductsOk

`func (o *PoolDTO) GetProvidedProductsOk() (*[]ProvidedProductDTO, bool)`

GetProvidedProductsOk returns a tuple with the ProvidedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidedProducts

`func (o *PoolDTO) SetProvidedProducts(v []ProvidedProductDTO)`

SetProvidedProducts sets ProvidedProducts field to given value.

### HasProvidedProducts

`func (o *PoolDTO) HasProvidedProducts() bool`

HasProvidedProducts returns a boolean if a field has been set.

### GetDerivedProvidedProducts

`func (o *PoolDTO) GetDerivedProvidedProducts() []ProvidedProductDTO`

GetDerivedProvidedProducts returns the DerivedProvidedProducts field if non-nil, zero value otherwise.

### GetDerivedProvidedProductsOk

`func (o *PoolDTO) GetDerivedProvidedProductsOk() (*[]ProvidedProductDTO, bool)`

GetDerivedProvidedProductsOk returns a tuple with the DerivedProvidedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedProvidedProducts

`func (o *PoolDTO) SetDerivedProvidedProducts(v []ProvidedProductDTO)`

SetDerivedProvidedProducts sets DerivedProvidedProducts field to given value.

### HasDerivedProvidedProducts

`func (o *PoolDTO) HasDerivedProvidedProducts() bool`

HasDerivedProvidedProducts returns a boolean if a field has been set.

### GetSubscriptionSubKey

`func (o *PoolDTO) GetSubscriptionSubKey() string`

GetSubscriptionSubKey returns the SubscriptionSubKey field if non-nil, zero value otherwise.

### GetSubscriptionSubKeyOk

`func (o *PoolDTO) GetSubscriptionSubKeyOk() (*string, bool)`

GetSubscriptionSubKeyOk returns a tuple with the SubscriptionSubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionSubKey

`func (o *PoolDTO) SetSubscriptionSubKey(v string)`

SetSubscriptionSubKey sets SubscriptionSubKey field to given value.

### HasSubscriptionSubKey

`func (o *PoolDTO) HasSubscriptionSubKey() bool`

HasSubscriptionSubKey returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *PoolDTO) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PoolDTO) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PoolDTO) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *PoolDTO) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetManaged

`func (o *PoolDTO) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PoolDTO) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PoolDTO) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PoolDTO) HasManaged() bool`

HasManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


