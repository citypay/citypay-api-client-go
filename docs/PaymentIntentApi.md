# \PaymentIntentApi

All URIs are relative to *https://api.citypay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentIntent**](PaymentIntentApi.md#CreatePaymentIntent) | **Post** /v6/intent/create | Create a Payment Intent
[**GetPaymentIntent**](PaymentIntentApi.md#GetPaymentIntent) | **Post** /v6/intent/retrieve | Retrieves a Payment Intent



## CreatePaymentIntent

> PaymentIntentReference CreatePaymentIntent(ctx).PaymentIntentRequestModel(paymentIntentRequestModel).Execute()

Create a Payment Intent



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
	paymentIntentRequestModel := *openapiclient.NewPaymentIntentRequestModel(int32(19995), "95b857a1-5955-4b86-963c-5a6dbfc4fb95") // PaymentIntentRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentIntentApi.CreatePaymentIntent(context.Background()).PaymentIntentRequestModel(paymentIntentRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentApi.CreatePaymentIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentIntent`: PaymentIntentReference
	fmt.Fprintf(os.Stdout, "Response from `PaymentIntentApi.CreatePaymentIntent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentIntentRequestModel** | [**PaymentIntentRequestModel**](PaymentIntentRequestModel.md) |  | 

### Return type

[**PaymentIntentReference**](PaymentIntentReference.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentIntent

> PaymentIntentResponseModel GetPaymentIntent(ctx).FindPaymentIntentRequest(findPaymentIntentRequest).Execute()

Retrieves a Payment Intent



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
	findPaymentIntentRequest := *openapiclient.NewFindPaymentIntentRequest() // FindPaymentIntentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentIntentApi.GetPaymentIntent(context.Background()).FindPaymentIntentRequest(findPaymentIntentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentApi.GetPaymentIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentIntent`: PaymentIntentResponseModel
	fmt.Fprintf(os.Stdout, "Response from `PaymentIntentApi.GetPaymentIntent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **findPaymentIntentRequest** | [**FindPaymentIntentRequest**](FindPaymentIntentRequest.md) |  | 

### Return type

[**PaymentIntentResponseModel**](PaymentIntentResponseModel.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

