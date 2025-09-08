# GetMultipleSensors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SensorsResponse**](SensorsResponse.md) |  | [optional] 
**HasNext** | Pointer to **bool** | True if next pages can be fetched, false otherwise. | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetMultipleSensors200Response

`func NewGetMultipleSensors200Response() *GetMultipleSensors200Response`

NewGetMultipleSensors200Response instantiates a new GetMultipleSensors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMultipleSensors200ResponseWithDefaults

`func NewGetMultipleSensors200ResponseWithDefaults() *GetMultipleSensors200Response`

NewGetMultipleSensors200ResponseWithDefaults instantiates a new GetMultipleSensors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetMultipleSensors200Response) GetResults() []SensorsResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetMultipleSensors200Response) GetResultsOk() (*[]SensorsResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetMultipleSensors200Response) SetResults(v []SensorsResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetMultipleSensors200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetHasNext

`func (o *GetMultipleSensors200Response) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *GetMultipleSensors200Response) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *GetMultipleSensors200Response) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *GetMultipleSensors200Response) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetTotalPages

`func (o *GetMultipleSensors200Response) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetMultipleSensors200Response) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetMultipleSensors200Response) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *GetMultipleSensors200Response) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


