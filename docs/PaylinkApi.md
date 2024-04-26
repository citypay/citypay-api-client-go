# \PaylinkApi

All URIs are relative to *https://api.citypay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokenAdjustmentRequest**](PaylinkApi.md#TokenAdjustmentRequest) | **Post** /paylink/{token}/adjustment | Paylink Token Adjustment
[**TokenCancelRequest**](PaylinkApi.md#TokenCancelRequest) | **Put** /paylink/{token}/cancel | Cancel a Paylink Token
[**TokenChangesRequest**](PaylinkApi.md#TokenChangesRequest) | **Post** /paylink/token/changes | Paylink Token Audit
[**TokenCloseRequest**](PaylinkApi.md#TokenCloseRequest) | **Put** /paylink/{token}/close | Close Paylink Token
[**TokenCreateBillPaymentRequest**](PaylinkApi.md#TokenCreateBillPaymentRequest) | **Post** /paylink/bill-payment | Create Bill Payment Paylink Token
[**TokenCreateRequest**](PaylinkApi.md#TokenCreateRequest) | **Post** /paylink/create | Create Paylink Token
[**TokenPurgeAttachmentsRequest**](PaylinkApi.md#TokenPurgeAttachmentsRequest) | **Put** /paylink/{token}/purge-attachments | Purges any attachments for a Paylink Token
[**TokenReconciledRequest**](PaylinkApi.md#TokenReconciledRequest) | **Put** /paylink/{token}/reconciled | Reconcile Paylink Token
[**TokenReopenRequest**](PaylinkApi.md#TokenReopenRequest) | **Put** /paylink/{token}/reopen | Reopen Paylink Token
[**TokenResendNotificationRequest**](PaylinkApi.md#TokenResendNotificationRequest) | **Post** /paylink/{token}/resend-notification | Resend a notification for Paylink Token
[**TokenStatusRequest**](PaylinkApi.md#TokenStatusRequest) | **Get** /paylink/{token}/status | Paylink Token Status



## TokenAdjustmentRequest

> Acknowledgement TokenAdjustmentRequest(ctx, token).PaylinkAdjustmentRequest(paylinkAdjustmentRequest).Execute()

Paylink Token Adjustment



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
	token := "token_example" // string | The token returned by the create token process.
	paylinkAdjustmentRequest := *openapiclient.NewPaylinkAdjustmentRequest() // PaylinkAdjustmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenAdjustmentRequest(context.Background(), token).PaylinkAdjustmentRequest(paylinkAdjustmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenAdjustmentRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenAdjustmentRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenAdjustmentRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token returned by the create token process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenAdjustmentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paylinkAdjustmentRequest** | [**PaylinkAdjustmentRequest**](PaylinkAdjustmentRequest.md) |  | 

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


## TokenCancelRequest

> Acknowledgement TokenCancelRequest(ctx, token).Execute()

Cancel a Paylink Token



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
	token := "token_example" // string | The token returned by the create token process.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenCancelRequest(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenCancelRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenCancelRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenCancelRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token returned by the create token process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenCancelRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Acknowledgement**](Acknowledgement.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenChangesRequest

> PaylinkTokenStatusChangeResponse TokenChangesRequest(ctx).PaylinkTokenStatusChangeRequest(paylinkTokenStatusChangeRequest).Execute()

Paylink Token Audit



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/citypay/citypay-api-client-go"
)

func main() {
	paylinkTokenStatusChangeRequest := *openapiclient.NewPaylinkTokenStatusChangeRequest(time.Now(), int32(11223344)) // PaylinkTokenStatusChangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenChangesRequest(context.Background()).PaylinkTokenStatusChangeRequest(paylinkTokenStatusChangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenChangesRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenChangesRequest`: PaylinkTokenStatusChangeResponse
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenChangesRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenChangesRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paylinkTokenStatusChangeRequest** | [**PaylinkTokenStatusChangeRequest**](PaylinkTokenStatusChangeRequest.md) |  | 

### Return type

[**PaylinkTokenStatusChangeResponse**](PaylinkTokenStatusChangeResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenCloseRequest

> Acknowledgement TokenCloseRequest(ctx, token).Execute()

Close Paylink Token



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
	token := "token_example" // string | The token returned by the create token process.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenCloseRequest(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenCloseRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenCloseRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenCloseRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token returned by the create token process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenCloseRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Acknowledgement**](Acknowledgement.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenCreateBillPaymentRequest

> PaylinkTokenCreated TokenCreateBillPaymentRequest(ctx).PaylinkBillPaymentTokenRequest(paylinkBillPaymentTokenRequest).Execute()

Create Bill Payment Paylink Token



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
	paylinkBillPaymentTokenRequest := *openapiclient.NewPaylinkBillPaymentTokenRequest(*openapiclient.NewPaylinkTokenRequestModel(int32(123), "95b857a1-5955-4b86-963c-5a6dbfc4fb95", int32(11223344))) // PaylinkBillPaymentTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenCreateBillPaymentRequest(context.Background()).PaylinkBillPaymentTokenRequest(paylinkBillPaymentTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenCreateBillPaymentRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenCreateBillPaymentRequest`: PaylinkTokenCreated
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenCreateBillPaymentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenCreateBillPaymentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paylinkBillPaymentTokenRequest** | [**PaylinkBillPaymentTokenRequest**](PaylinkBillPaymentTokenRequest.md) |  | 

### Return type

[**PaylinkTokenCreated**](PaylinkTokenCreated.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenCreateRequest

> PaylinkTokenCreated TokenCreateRequest(ctx).PaylinkTokenRequestModel(paylinkTokenRequestModel).Execute()

Create Paylink Token



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
	paylinkTokenRequestModel := *openapiclient.NewPaylinkTokenRequestModel(int32(123), "95b857a1-5955-4b86-963c-5a6dbfc4fb95", int32(11223344)) // PaylinkTokenRequestModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenCreateRequest(context.Background()).PaylinkTokenRequestModel(paylinkTokenRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenCreateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenCreateRequest`: PaylinkTokenCreated
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenCreateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenCreateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paylinkTokenRequestModel** | [**PaylinkTokenRequestModel**](PaylinkTokenRequestModel.md) |  | 

### Return type

[**PaylinkTokenCreated**](PaylinkTokenCreated.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenPurgeAttachmentsRequest

> Acknowledgement TokenPurgeAttachmentsRequest(ctx, token).Execute()

Purges any attachments for a Paylink Token



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
	token := "token_example" // string | The token returned by the create token process.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenPurgeAttachmentsRequest(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenPurgeAttachmentsRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenPurgeAttachmentsRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenPurgeAttachmentsRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token returned by the create token process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenPurgeAttachmentsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Acknowledgement**](Acknowledgement.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenReconciledRequest

> Acknowledgement TokenReconciledRequest(ctx, token).Execute()

Reconcile Paylink Token



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
	token := "token_example" // string | The token returned by the create token process.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenReconciledRequest(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenReconciledRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenReconciledRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenReconciledRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token returned by the create token process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenReconciledRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Acknowledgement**](Acknowledgement.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenReopenRequest

> Acknowledgement TokenReopenRequest(ctx, token).Execute()

Reopen Paylink Token



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
	token := "token_example" // string | The token returned by the create token process.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenReopenRequest(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenReopenRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenReopenRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenReopenRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token returned by the create token process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenReopenRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Acknowledgement**](Acknowledgement.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenResendNotificationRequest

> Acknowledgement TokenResendNotificationRequest(ctx, token).PaylinkResendNotificationRequest(paylinkResendNotificationRequest).Execute()

Resend a notification for Paylink Token



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
	token := "token_example" // string | The token returned by the create token process.
	paylinkResendNotificationRequest := *openapiclient.NewPaylinkResendNotificationRequest() // PaylinkResendNotificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenResendNotificationRequest(context.Background(), token).PaylinkResendNotificationRequest(paylinkResendNotificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenResendNotificationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenResendNotificationRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenResendNotificationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token returned by the create token process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenResendNotificationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paylinkResendNotificationRequest** | [**PaylinkResendNotificationRequest**](PaylinkResendNotificationRequest.md) |  | 

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


## TokenStatusRequest

> PaylinkTokenStatus TokenStatusRequest(ctx, token).Execute()

Paylink Token Status



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
	token := "token_example" // string | The token returned by the create token process.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaylinkApi.TokenStatusRequest(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaylinkApi.TokenStatusRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenStatusRequest`: PaylinkTokenStatus
	fmt.Fprintf(os.Stdout, "Response from `PaylinkApi.TokenStatusRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token returned by the create token process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenStatusRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaylinkTokenStatus**](PaylinkTokenStatus.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

