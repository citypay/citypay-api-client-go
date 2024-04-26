# \OperationalFunctionsApi

All URIs are relative to *https://api.citypay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AclCheckRequest**](OperationalFunctionsApi.md#AclCheckRequest) | **Post** /v6/acl/check | ACL Check Request
[**DomainKeyCheckRequest**](OperationalFunctionsApi.md#DomainKeyCheckRequest) | **Post** /dk/check | Domain Key Check Request
[**DomainKeyGenRequest**](OperationalFunctionsApi.md#DomainKeyGenRequest) | **Post** /dk/gen | Domain Key Generation Request
[**ListMerchantsRequest**](OperationalFunctionsApi.md#ListMerchantsRequest) | **Get** /v6/merchants/{clientid} | List Merchants Request
[**PingRequest**](OperationalFunctionsApi.md#PingRequest) | **Post** /v6/ping | Ping Request



## AclCheckRequest

> AclCheckResponseModel AclCheckRequest(ctx).AclCheckRequest(aclCheckRequest).Execute()

ACL Check Request



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
	aclCheckRequest := *openapiclient.NewAclCheckRequest("8.8.8.8") // AclCheckRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationalFunctionsApi.AclCheckRequest(context.Background()).AclCheckRequest(aclCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationalFunctionsApi.AclCheckRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AclCheckRequest`: AclCheckResponseModel
	fmt.Fprintf(os.Stdout, "Response from `OperationalFunctionsApi.AclCheckRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAclCheckRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aclCheckRequest** | [**AclCheckRequest**](AclCheckRequest.md) |  | 

### Return type

[**AclCheckResponseModel**](AclCheckResponseModel.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainKeyCheckRequest

> DomainKeyResponse DomainKeyCheckRequest(ctx).DomainKeyCheckRequest(domainKeyCheckRequest).Execute()

Domain Key Check Request



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
	domainKeyCheckRequest := *openapiclient.NewDomainKeyCheckRequest("3MEcU8cEf...QMeebACxcQVejmT1Wi") // DomainKeyCheckRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationalFunctionsApi.DomainKeyCheckRequest(context.Background()).DomainKeyCheckRequest(domainKeyCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationalFunctionsApi.DomainKeyCheckRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainKeyCheckRequest`: DomainKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationalFunctionsApi.DomainKeyCheckRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainKeyCheckRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainKeyCheckRequest** | [**DomainKeyCheckRequest**](DomainKeyCheckRequest.md) |  | 

### Return type

[**DomainKeyResponse**](DomainKeyResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainKeyGenRequest

> DomainKeyResponse DomainKeyGenRequest(ctx).DomainKeyRequest(domainKeyRequest).Execute()

Domain Key Generation Request



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
	domainKeyRequest := *openapiclient.NewDomainKeyRequest([]string{"Domain_example"}, int32(11223344)) // DomainKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationalFunctionsApi.DomainKeyGenRequest(context.Background()).DomainKeyRequest(domainKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationalFunctionsApi.DomainKeyGenRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainKeyGenRequest`: DomainKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationalFunctionsApi.DomainKeyGenRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainKeyGenRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainKeyRequest** | [**DomainKeyRequest**](DomainKeyRequest.md) |  | 

### Return type

[**DomainKeyResponse**](DomainKeyResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMerchantsRequest

> ListMerchantsResponse ListMerchantsRequest(ctx, clientid).Execute()

List Merchants Request



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
	clientid := "clientid_example" // string | The client id to return merchants for, specifying \"default\" will use the value in your api key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationalFunctionsApi.ListMerchantsRequest(context.Background(), clientid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationalFunctionsApi.ListMerchantsRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMerchantsRequest`: ListMerchantsResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationalFunctionsApi.ListMerchantsRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientid** | **string** | The client id to return merchants for, specifying \&quot;default\&quot; will use the value in your api key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMerchantsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMerchantsResponse**](ListMerchantsResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingRequest

> Acknowledgement PingRequest(ctx).Ping(ping).Execute()

Ping Request



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
	ping := *openapiclient.NewPing() // Ping | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationalFunctionsApi.PingRequest(context.Background()).Ping(ping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationalFunctionsApi.PingRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PingRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `OperationalFunctionsApi.PingRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPingRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ping** | [**Ping**](Ping.md) |  | 

### Return type

[**Acknowledgement**](Acknowledgement.md)

### Authorization

[cp-domain-key](../README.md#cp-domain-key), [cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, text/xml
- **Accept**: application/x-www-form-urlencoded, application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

