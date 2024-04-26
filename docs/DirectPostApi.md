# \DirectPostApi

All URIs are relative to *https://api.citypay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectCResAuthRequest**](DirectPostApi.md#DirectCResAuthRequest) | **Post** /direct/cres/auth/{uuid} | Handles a CRes response from ACS, returning back the result of authorisation
[**DirectCResTokeniseRequest**](DirectPostApi.md#DirectCResTokeniseRequest) | **Post** /direct/cres/tokenise/{uuid} | Handles a CRes response from ACS, returning back a token for future authorisation
[**DirectPostAuthRequest**](DirectPostApi.md#DirectPostAuthRequest) | **Post** /direct/auth | Direct Post Auth Request
[**DirectPostTokeniseRequest**](DirectPostApi.md#DirectPostTokeniseRequest) | **Post** /direct/tokenise | Direct Post Tokenise Request
[**TokenRequest**](DirectPostApi.md#TokenRequest) | **Post** /direct/token | Direct Post Token Request



## DirectCResAuthRequest

> AuthResponse DirectCResAuthRequest(ctx, uuid).Cres(cres).ThreeDSSessionData(threeDSSessionData).Execute()

Handles a CRes response from ACS, returning back the result of authorisation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citypay/citypay-api-client-go"
)

func main() {
	uuid := "uuid_example" // string | An identifier used to track the CReq/CRes cycle.
	cres := "cres_example" // string | The CRES from the ACS. (optional)
	threeDSSessionData := "threeDSSessionData_example" // string | The session data from the ACS. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectPostApi.DirectCResAuthRequest(context.Background(), uuid).Cres(cres).ThreeDSSessionData(threeDSSessionData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectPostApi.DirectCResAuthRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DirectCResAuthRequest`: AuthResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectPostApi.DirectCResAuthRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | An identifier used to track the CReq/CRes cycle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectCResAuthRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cres** | **string** | The CRES from the ACS. | 
 **threeDSSessionData** | **string** | The session data from the ACS. | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json, application/xml, application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectCResTokeniseRequest

> TokenisationResponseModel DirectCResTokeniseRequest(ctx, uuid).Cres(cres).ThreeDSSessionData(threeDSSessionData).Execute()

Handles a CRes response from ACS, returning back a token for future authorisation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citypay/citypay-api-client-go"
)

func main() {
	uuid := "uuid_example" // string | An identifier used to track the CReq/CRes cycle.
	cres := "cres_example" // string | The CRES from the ACS. (optional)
	threeDSSessionData := "threeDSSessionData_example" // string | The session data from the ACS. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectPostApi.DirectCResTokeniseRequest(context.Background(), uuid).Cres(cres).ThreeDSSessionData(threeDSSessionData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectPostApi.DirectCResTokeniseRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DirectCResTokeniseRequest`: TokenisationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `DirectPostApi.DirectCResTokeniseRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | An identifier used to track the CReq/CRes cycle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectCResTokeniseRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cres** | **string** | The CRES from the ACS. | 
 **threeDSSessionData** | **string** | The session data from the ACS. | 

### Return type

[**TokenisationResponseModel**](TokenisationResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json, application/xml, application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectPostAuthRequest

> AuthResponse DirectPostAuthRequest(ctx).DirectPostRequest(directPostRequest).Execute()

Direct Post Auth Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citypay/citypay-api-client-go"
)

func main() {
	directPostRequest := *openapiclient.NewDirectPostRequest(int32(19995), "4000 0000 0000 0002", int32(9), int32(2027), "95b857a1-5955-4b86-963c-5a6dbfc4fb95", "3896FBC43674AF59478DAF7F546FA4D4CB89981A936E6AAE997E43B55DF6C39D") // DirectPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectPostApi.DirectPostAuthRequest(context.Background()).DirectPostRequest(directPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectPostApi.DirectPostAuthRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DirectPostAuthRequest`: AuthResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectPostApi.DirectPostAuthRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectPostAuthRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directPostRequest** | [**DirectPostRequest**](DirectPostRequest.md) |  | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[cp-domain-key](../README.md#cp-domain-key), [cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, text/xml
- **Accept**: application/json, application/xml, application/x-www-form-urlencoded, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectPostTokeniseRequest

> AuthResponse DirectPostTokeniseRequest(ctx).DirectPostRequest(directPostRequest).Execute()

Direct Post Tokenise Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citypay/citypay-api-client-go"
)

func main() {
	directPostRequest := *openapiclient.NewDirectPostRequest(int32(19995), "4000 0000 0000 0002", int32(9), int32(2027), "95b857a1-5955-4b86-963c-5a6dbfc4fb95", "3896FBC43674AF59478DAF7F546FA4D4CB89981A936E6AAE997E43B55DF6C39D") // DirectPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectPostApi.DirectPostTokeniseRequest(context.Background()).DirectPostRequest(directPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectPostApi.DirectPostTokeniseRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DirectPostTokeniseRequest`: AuthResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectPostApi.DirectPostTokeniseRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectPostTokeniseRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directPostRequest** | [**DirectPostRequest**](DirectPostRequest.md) |  | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[cp-domain-key](../README.md#cp-domain-key), [cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, text/xml
- **Accept**: application/json, application/xml, application/x-www-form-urlencoded, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenRequest

> AuthResponse TokenRequest(ctx).DirectTokenAuthRequest(directTokenAuthRequest).Execute()

Direct Post Token Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citypay/citypay-api-client-go"
)

func main() {
	directTokenAuthRequest := *openapiclient.NewDirectTokenAuthRequest() // DirectTokenAuthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectPostApi.TokenRequest(context.Background()).DirectTokenAuthRequest(directTokenAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectPostApi.TokenRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenRequest`: AuthResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectPostApi.TokenRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directTokenAuthRequest** | [**DirectTokenAuthRequest**](DirectTokenAuthRequest.md) |  | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[cp-domain-key](../README.md#cp-domain-key), [cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, text/xml
- **Accept**: application/json, application/xml, application/x-www-form-urlencoded, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

