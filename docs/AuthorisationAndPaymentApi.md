# \AuthorisationAndPaymentApi

All URIs are relative to *https://api.citypay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorisationRequest**](AuthorisationAndPaymentApi.md#AuthorisationRequest) | **Post** /v6/authorise | Authorisation
[**BinRangeLookupRequest**](AuthorisationAndPaymentApi.md#BinRangeLookupRequest) | **Post** /v6/bin | Bin Lookup
[**CResRequest**](AuthorisationAndPaymentApi.md#CResRequest) | **Post** /v6/cres | CRes
[**CaptureRequest**](AuthorisationAndPaymentApi.md#CaptureRequest) | **Post** /v6/capture | Capture
[**CreatePaymentIntent**](AuthorisationAndPaymentApi.md#CreatePaymentIntent) | **Post** /v6/intent/create | Create a Payment Intent
[**PaResRequest**](AuthorisationAndPaymentApi.md#PaResRequest) | **Post** /v6/pares | PaRes
[**RefundRequest**](AuthorisationAndPaymentApi.md#RefundRequest) | **Post** /v6/refund | Refund
[**RetrievalRequest**](AuthorisationAndPaymentApi.md#RetrievalRequest) | **Post** /v6/retrieve | Retrieval
[**VoidRequest**](AuthorisationAndPaymentApi.md#VoidRequest) | **Post** /v6/void | Void



## AuthorisationRequest

> Decision AuthorisationRequest(ctx).AuthRequest(authRequest).Execute()

Authorisation



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
	authRequest := *openapiclient.NewAuthRequest(int32(19995), "4000 0000 0000 0002", int32(9), int32(2027), "95b857a1-5955-4b86-963c-5a6dbfc4fb95", int32(11223344)) // AuthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.AuthorisationRequest(context.Background()).AuthRequest(authRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.AuthorisationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorisationRequest`: Decision
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.AuthorisationRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorisationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authRequest** | [**AuthRequest**](AuthRequest.md) |  | 

### Return type

[**Decision**](Decision.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BinRangeLookupRequest

> Bin BinRangeLookupRequest(ctx).BinLookup(binLookup).Execute()

Bin Lookup



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
	binLookup := *openapiclient.NewBinLookup(int32(543712)) // BinLookup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.BinRangeLookupRequest(context.Background()).BinLookup(binLookup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.BinRangeLookupRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BinRangeLookupRequest`: Bin
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.BinRangeLookupRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBinRangeLookupRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **binLookup** | [**BinLookup**](BinLookup.md) |  | 

### Return type

[**Bin**](Bin.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CResRequest

> AuthResponse CResRequest(ctx).CResAuthRequest(cResAuthRequest).Execute()

CRes



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
	cResAuthRequest := *openapiclient.NewCResAuthRequest() // CResAuthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.CResRequest(context.Background()).CResAuthRequest(cResAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.CResRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CResRequest`: AuthResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.CResRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCResRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cResAuthRequest** | [**CResAuthRequest**](CResAuthRequest.md) |  | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CaptureRequest

> Acknowledgement CaptureRequest(ctx).CaptureRequest(captureRequest).Execute()

Capture



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
	captureRequest := *openapiclient.NewCaptureRequest(int32(11223344)) // CaptureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.CaptureRequest(context.Background()).CaptureRequest(captureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.CaptureRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CaptureRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.CaptureRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCaptureRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captureRequest** | [**CaptureRequest**](CaptureRequest.md) |  | 

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


## CreatePaymentIntent

> PaymentIntentReference CreatePaymentIntent(ctx).PaymentIntent(paymentIntent).Execute()

Create a Payment Intent



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
	paymentIntent := *openapiclient.NewPaymentIntent(int32(19995), "95b857a1-5955-4b86-963c-5a6dbfc4fb95") // PaymentIntent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.CreatePaymentIntent(context.Background()).PaymentIntent(paymentIntent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.CreatePaymentIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentIntent`: PaymentIntentReference
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.CreatePaymentIntent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentIntent** | [**PaymentIntent**](PaymentIntent.md) |  | 

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


## PaResRequest

> AuthResponse PaResRequest(ctx).PaResAuthRequest(paResAuthRequest).Execute()

PaRes



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
	paResAuthRequest := *openapiclient.NewPaResAuthRequest("Md_example", "v66ycfSp8jNlvy7PkHbx44NEt3vox90+vZ/7Ll05Vid/jPfQn8adw+4D/vRDUGT19kndW97Hfirb...") // PaResAuthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.PaResRequest(context.Background()).PaResAuthRequest(paResAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.PaResRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaResRequest`: AuthResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.PaResRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaResRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paResAuthRequest** | [**PaResAuthRequest**](PaResAuthRequest.md) |  | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundRequest

> AuthResponse RefundRequest(ctx).RefundRequest(refundRequest).Execute()

Refund



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
	refundRequest := *openapiclient.NewRefundRequest(int32(19995), "95b857a1-5955-4b86-963c-5a6dbfc4fb95", int32(11223344), int32(8322)) // RefundRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.RefundRequest(context.Background()).RefundRequest(refundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.RefundRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundRequest`: AuthResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.RefundRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundRequest** | [**RefundRequest**](RefundRequest.md) |  | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievalRequest

> AuthReferences RetrievalRequest(ctx).RetrieveRequest(retrieveRequest).Execute()

Retrieval



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
	retrieveRequest := *openapiclient.NewRetrieveRequest(int32(11223344)) // RetrieveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.RetrievalRequest(context.Background()).RetrieveRequest(retrieveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.RetrievalRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrievalRequest`: AuthReferences
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.RetrievalRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retrieveRequest** | [**RetrieveRequest**](RetrieveRequest.md) |  | 

### Return type

[**AuthReferences**](AuthReferences.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidRequest

> Acknowledgement VoidRequest(ctx).VoidRequest(voidRequest).Execute()

Void



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
	voidRequest := *openapiclient.NewVoidRequest(int32(11223344)) // VoidRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.VoidRequest(context.Background()).VoidRequest(voidRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.VoidRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.VoidRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVoidRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voidRequest** | [**VoidRequest**](VoidRequest.md) |  | 

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

