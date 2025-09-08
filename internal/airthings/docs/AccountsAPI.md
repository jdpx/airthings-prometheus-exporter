# \AccountsAPI

All URIs are relative to *https://consumer-api.airthings.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountsIds**](AccountsAPI.md#GetAccountsIds) | **Get** /v1/accounts | List all accounts the current user is member of



## GetAccountsIds

> AccountsResponse GetAccountsIds(ctx).Execute()

List all accounts the current user is member of



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountsIds(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountsIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountsIds`: AccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountsIds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsIdsRequest struct via the builder pattern


### Return type

[**AccountsResponse**](AccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

