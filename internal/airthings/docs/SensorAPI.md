# \SensorAPI

All URIs are relative to *https://consumer-api.airthings.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMultipleSensors**](SensorAPI.md#GetMultipleSensors) | **Get** /v1/accounts/{accountId}/sensors | Get sensors for a set of devices



## GetMultipleSensors

> GetMultipleSensors200Response GetMultipleSensors(ctx, accountId).Sn(sn).PageNumber(pageNumber).Unit(unit).Execute()

Get sensors for a set of devices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountId := "accountId_example" // string | The account ID associated with the user
	sn := []string{"Inner_example"} // []string | The serial numbers of the devices (optional)
	pageNumber := int32(56) // int32 | The number of a page (of 50 records) to fetch (optional) (default to 1)
	unit := "unit_example" // string | The units type sensors values will be returned in (optional) (default to "metric")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SensorAPI.GetMultipleSensors(context.Background(), accountId).Sn(sn).PageNumber(pageNumber).Unit(unit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetMultipleSensors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipleSensors`: GetMultipleSensors200Response
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetMultipleSensors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The account ID associated with the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleSensorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sn** | **[]string** | The serial numbers of the devices | 
 **pageNumber** | **int32** | The number of a page (of 50 records) to fetch | [default to 1]
 **unit** | **string** | The units type sensors values will be returned in | [default to &quot;metric&quot;]

### Return type

[**GetMultipleSensors200Response**](GetMultipleSensors200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

