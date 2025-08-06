# \AuthorisationAndPaymentApi

All URIs are relative to *https://api.citypay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorisationRequest**](AuthorisationAndPaymentApi.md#AuthorisationRequest) | **Post** /v6/authorise | Authorisation
[**BinRangeLookupRequest**](AuthorisationAndPaymentApi.md#BinRangeLookupRequest) | **Post** /v6/bin | Bin Lookup
[**CResRequest**](AuthorisationAndPaymentApi.md#CResRequest) | **Post** /v6/cres | CRes
[**CaptureRequest**](AuthorisationAndPaymentApi.md#CaptureRequest) | **Post** /v6/capture | Capture
[**CardTokenisationRequest**](AuthorisationAndPaymentApi.md#CardTokenisationRequest) | **Post** /v6/tokenise | Card Tokenisation Request
[**RefundRequest**](AuthorisationAndPaymentApi.md#RefundRequest) | **Post** /v6/refund | Refund
[**RetrievalRequest**](AuthorisationAndPaymentApi.md#RetrievalRequest) | **Post** /v6/retrieve | Transaction Retrieval
[**VerificationRequest**](AuthorisationAndPaymentApi.md#VerificationRequest) | **Post** /v6/verify | Verification
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
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
)

func main() {
	authRequest := *openapiclient.NewAuthRequest(int32(19995), "95b857a1-5955-4b86-963c-5a6dbfc4fb95", int32(11223344)) // AuthRequest | 

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
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
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
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
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
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
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


## CardTokenisationRequest

> CardTokenisationResponse CardTokenisationRequest(ctx).CardTokenisationRequest(cardTokenisationRequest).Execute()

Card Tokenisation Request



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
	cardTokenisationRequest := *openapiclient.NewCardTokenisationRequest() // CardTokenisationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.CardTokenisationRequest(context.Background()).CardTokenisationRequest(cardTokenisationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.CardTokenisationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CardTokenisationRequest`: CardTokenisationResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.CardTokenisationRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCardTokenisationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardTokenisationRequest** | [**CardTokenisationRequest**](CardTokenisationRequest.md) |  | 

### Return type

[**CardTokenisationResponse**](CardTokenisationResponse.md)

### Authorization

[cp-domain-key](../README.md#cp-domain-key), [cp-api-key](../README.md#cp-api-key)

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
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
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

Transaction Retrieval



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


## VerificationRequest

> Decision VerificationRequest(ctx).VerificationRequest(verificationRequest).Execute()

Verification



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
	verificationRequest := *openapiclient.NewVerificationRequest(int32(19995), "95b857a1-5955-4b86-963c-5a6dbfc4fb95", int32(11223344)) // VerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorisationAndPaymentApi.VerificationRequest(context.Background()).VerificationRequest(verificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorisationAndPaymentApi.VerificationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationRequest`: Decision
	fmt.Fprintf(os.Stdout, "Response from `AuthorisationAndPaymentApi.VerificationRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerificationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verificationRequest** | [**VerificationRequest**](VerificationRequest.md) |  | 

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
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
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

