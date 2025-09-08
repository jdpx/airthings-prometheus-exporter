# SensorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SensorType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewSensorResponse

`func NewSensorResponse() *SensorResponse`

NewSensorResponse instantiates a new SensorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorResponseWithDefaults

`func NewSensorResponseWithDefaults() *SensorResponse`

NewSensorResponseWithDefaults instantiates a new SensorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSensorType

`func (o *SensorResponse) GetSensorType() string`

GetSensorType returns the SensorType field if non-nil, zero value otherwise.

### GetSensorTypeOk

`func (o *SensorResponse) GetSensorTypeOk() (*string, bool)`

GetSensorTypeOk returns a tuple with the SensorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorType

`func (o *SensorResponse) SetSensorType(v string)`

SetSensorType sets SensorType field to given value.

### HasSensorType

`func (o *SensorResponse) HasSensorType() bool`

HasSensorType returns a boolean if a field has been set.

### GetValue

`func (o *SensorResponse) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SensorResponse) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SensorResponse) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *SensorResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnit

`func (o *SensorResponse) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SensorResponse) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SensorResponse) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SensorResponse) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


