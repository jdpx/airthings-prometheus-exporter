# SensorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **string** |  | [optional] 
**Sensors** | Pointer to [**[]SensorResponse**](SensorResponse.md) |  | [optional] 
**Recorded** | Pointer to **NullableString** |  | [optional] 
**BatteryPercentage** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewSensorsResponse

`func NewSensorsResponse() *SensorsResponse`

NewSensorsResponse instantiates a new SensorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorsResponseWithDefaults

`func NewSensorsResponseWithDefaults() *SensorsResponse`

NewSensorsResponseWithDefaults instantiates a new SensorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *SensorsResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SensorsResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SensorsResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SensorsResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSensors

`func (o *SensorsResponse) GetSensors() []SensorResponse`

GetSensors returns the Sensors field if non-nil, zero value otherwise.

### GetSensorsOk

`func (o *SensorsResponse) GetSensorsOk() (*[]SensorResponse, bool)`

GetSensorsOk returns a tuple with the Sensors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensors

`func (o *SensorsResponse) SetSensors(v []SensorResponse)`

SetSensors sets Sensors field to given value.

### HasSensors

`func (o *SensorsResponse) HasSensors() bool`

HasSensors returns a boolean if a field has been set.

### GetRecorded

`func (o *SensorsResponse) GetRecorded() string`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *SensorsResponse) GetRecordedOk() (*string, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *SensorsResponse) SetRecorded(v string)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *SensorsResponse) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### SetRecordedNil

`func (o *SensorsResponse) SetRecordedNil(b bool)`

 SetRecordedNil sets the value for Recorded to be an explicit nil

### UnsetRecorded
`func (o *SensorsResponse) UnsetRecorded()`

UnsetRecorded ensures that no value is present for Recorded, not even an explicit nil
### GetBatteryPercentage

`func (o *SensorsResponse) GetBatteryPercentage() int32`

GetBatteryPercentage returns the BatteryPercentage field if non-nil, zero value otherwise.

### GetBatteryPercentageOk

`func (o *SensorsResponse) GetBatteryPercentageOk() (*int32, bool)`

GetBatteryPercentageOk returns a tuple with the BatteryPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryPercentage

`func (o *SensorsResponse) SetBatteryPercentage(v int32)`

SetBatteryPercentage sets BatteryPercentage field to given value.

### HasBatteryPercentage

`func (o *SensorsResponse) HasBatteryPercentage() bool`

HasBatteryPercentage returns a boolean if a field has been set.

### SetBatteryPercentageNil

`func (o *SensorsResponse) SetBatteryPercentageNil(b bool)`

 SetBatteryPercentageNil sets the value for BatteryPercentage to be an explicit nil

### UnsetBatteryPercentage
`func (o *SensorsResponse) UnsetBatteryPercentage()`

UnsetBatteryPercentage ensures that no value is present for BatteryPercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


