# \WebHooks

All URIs are relative to *https://api.citypay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebHookChannelCreateRequest**](WebHooks.md#WebHookChannelCreateRequest) | **Post** /hooks/channel/create | Web Hook Channel Create Request
[**WebHookChannelDeleteRequest**](WebHooks.md#WebHookChannelDeleteRequest) | **Post** /hooks/channel/delete | Web Hook Channel Delete Request
[**WebHookSubscriptionRequest**](WebHooks.md#WebHookSubscriptionRequest) | **Post** /hooks/subscribe | Web Hook Subscription Request
[**WebHookUnsubscribeRequest**](WebHooks.md#WebHookUnsubscribeRequest) | **Post** /hooks/unsubscribe | Web Hook Unsubscribe Request



## WebHookChannelCreateRequest

> WebHookChannelCreateResponse WebHookChannelCreateRequest(ctx).WebHookChannelCreateRequest(webHookChannelCreateRequest).Execute()

Web Hook Channel Create Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
)

func main() {
	webHookChannelCreateRequest := *openapiclient.NewWebHookChannelCreateRequest("ChannelName_example", "PC12345", openapiclient.WebHookChannelCreateRequest_config{HttpConfig: openapiclient.NewHttpConfig("https://yoursite.com/path")}, "EndpointId_example") // WebHookChannelCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebHooks.WebHookChannelCreateRequest(context.Background()).WebHookChannelCreateRequest(webHookChannelCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebHooks.WebHookChannelCreateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebHookChannelCreateRequest`: WebHookChannelCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `WebHooks.WebHookChannelCreateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebHookChannelCreateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webHookChannelCreateRequest** | [**WebHookChannelCreateRequest**](WebHookChannelCreateRequest.md) |  | 

### Return type

[**WebHookChannelCreateResponse**](WebHookChannelCreateResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebHookChannelDeleteRequest

> Acknowledgement WebHookChannelDeleteRequest(ctx).WebHookChannelDeleteRequest(webHookChannelDeleteRequest).Execute()

Web Hook Channel Delete Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
)

func main() {
	webHookChannelDeleteRequest := *openapiclient.NewWebHookChannelDeleteRequest("wc_4ERJ5kV7nX2hT9m") // WebHookChannelDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebHooks.WebHookChannelDeleteRequest(context.Background()).WebHookChannelDeleteRequest(webHookChannelDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebHooks.WebHookChannelDeleteRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebHookChannelDeleteRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `WebHooks.WebHookChannelDeleteRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebHookChannelDeleteRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webHookChannelDeleteRequest** | [**WebHookChannelDeleteRequest**](WebHookChannelDeleteRequest.md) |  | 

### Return type

[**Acknowledgement**](Acknowledgement.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebHookSubscriptionRequest

> WebHookSubscriptionResponse WebHookSubscriptionRequest(ctx).WebHookSubscriptionRequest(webHookSubscriptionRequest).Execute()

Web Hook Subscription Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
)

func main() {
	webHookSubscriptionRequest := *openapiclient.NewWebHookSubscriptionRequest("PC12345") // WebHookSubscriptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebHooks.WebHookSubscriptionRequest(context.Background()).WebHookSubscriptionRequest(webHookSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebHooks.WebHookSubscriptionRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebHookSubscriptionRequest`: WebHookSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `WebHooks.WebHookSubscriptionRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebHookSubscriptionRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webHookSubscriptionRequest** | [**WebHookSubscriptionRequest**](WebHookSubscriptionRequest.md) |  | 

### Return type

[**WebHookSubscriptionResponse**](WebHookSubscriptionResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebHookUnsubscribeRequest

> Acknowledgement WebHookUnsubscribeRequest(ctx).WebHookUnsubscribeRequest(webHookUnsubscribeRequest).Execute()

Web Hook Unsubscribe Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
)

func main() {
	webHookUnsubscribeRequest := *openapiclient.NewWebHookUnsubscribeRequest("PC12345", "hk_Zf3UbQp8cY6LsRwY1") // WebHookUnsubscribeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebHooks.WebHookUnsubscribeRequest(context.Background()).WebHookUnsubscribeRequest(webHookUnsubscribeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebHooks.WebHookUnsubscribeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebHookUnsubscribeRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `WebHooks.WebHookUnsubscribeRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebHookUnsubscribeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webHookUnsubscribeRequest** | [**WebHookUnsubscribeRequest**](WebHookUnsubscribeRequest.md) |  | 

### Return type

[**Acknowledgement**](Acknowledgement.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

