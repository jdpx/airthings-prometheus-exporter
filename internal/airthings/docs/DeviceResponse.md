# DeviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **string** |  | [optional] 
**Home** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Sensors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeviceResponse

`func NewDeviceResponse() *DeviceResponse`

NewDeviceResponse instantiates a new DeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceResponseWithDefaults

`func NewDeviceResponseWithDefaults() *DeviceResponse`

NewDeviceResponseWithDefaults instantiates a new DeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *DeviceResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DeviceResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DeviceResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DeviceResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetHome

`func (o *DeviceResponse) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *DeviceResponse) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *DeviceResponse) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *DeviceResponse) HasHome() bool`

HasHome returns a boolean if a field has been set.

### SetHomeNil

`func (o *DeviceResponse) SetHomeNil(b bool)`

 SetHomeNil sets the value for Home to be an explicit nil

### UnsetHome
`func (o *DeviceResponse) UnsetHome()`

UnsetHome ensures that no value is present for Home, not even an explicit nil
### GetName

`func (o *DeviceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DeviceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSensors

`func (o *DeviceResponse) GetSensors() []string`

GetSensors returns the Sensors field if non-nil, zero value otherwise.

### GetSensorsOk

`func (o *DeviceResponse) GetSensorsOk() (*[]string, bool)`

GetSensorsOk returns a tuple with the Sensors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensors

`func (o *DeviceResponse) SetSensors(v []string)`

SetSensors sets Sensors field to given value.

### HasSensors

`func (o *DeviceResponse) HasSensors() bool`

HasSensors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


