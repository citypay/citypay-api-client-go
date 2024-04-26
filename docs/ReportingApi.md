# \ReportingApi

All URIs are relative to *https://api.citypay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchedTransactionReportRequest**](ReportingApi.md#BatchedTransactionReportRequest) | **Post** /v6/merchant-batch/{merchantid}/{batch_no}/transactions | Batch Transaction Report Request
[**MerchantBatchReportRequest**](ReportingApi.md#MerchantBatchReportRequest) | **Post** /v6/merchant-batch/report | Merchant Batch Report Request
[**MerchantBatchRequest**](ReportingApi.md#MerchantBatchRequest) | **Get** /v6/merchant-batch/{merchantid}/{batch_no} | Merchant Batch Request
[**RemittanceRangeReport**](ReportingApi.md#RemittanceRangeReport) | **Post** /v6/remittance/report/{clientid} | Remittance Report Request
[**RemittanceReportRequest**](ReportingApi.md#RemittanceReportRequest) | **Get** /v6/remittance/report/{clientid}/{date} | Remittance Date Report Request



## BatchedTransactionReportRequest

> BatchTransactionReportResponse BatchedTransactionReportRequest(ctx, merchantid, batchNo).BatchTransactionReportRequest(batchTransactionReportRequest).Execute()

Batch Transaction Report Request



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
	merchantid := int32(56) // int32 | A merchant ID (MID) for which data is requested. This field allows for filtering of the request by a specific merchant account.
	batchNo := "batchNo_example" // string | The batch number that is being requested.
	batchTransactionReportRequest := *openapiclient.NewBatchTransactionReportRequest() // BatchTransactionReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingApi.BatchedTransactionReportRequest(context.Background(), merchantid, batchNo).BatchTransactionReportRequest(batchTransactionReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.BatchedTransactionReportRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchedTransactionReportRequest`: BatchTransactionReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportingApi.BatchedTransactionReportRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantid** | **int32** | A merchant ID (MID) for which data is requested. This field allows for filtering of the request by a specific merchant account. | 
**batchNo** | **string** | The batch number that is being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchedTransactionReportRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchTransactionReportRequest** | [**BatchTransactionReportRequest**](BatchTransactionReportRequest.md) |  | 

### Return type

[**BatchTransactionReportResponse**](BatchTransactionReportResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantBatchReportRequest

> MerchantBatchReportResponse MerchantBatchReportRequest(ctx).MerchantBatchReportRequest(merchantBatchReportRequest).Execute()

Merchant Batch Report Request



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
	merchantBatchReportRequest := *openapiclient.NewMerchantBatchReportRequest() // MerchantBatchReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingApi.MerchantBatchReportRequest(context.Background()).MerchantBatchReportRequest(merchantBatchReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.MerchantBatchReportRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantBatchReportRequest`: MerchantBatchReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportingApi.MerchantBatchReportRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantBatchReportRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantBatchReportRequest** | [**MerchantBatchReportRequest**](MerchantBatchReportRequest.md) |  | 

### Return type

[**MerchantBatchReportResponse**](MerchantBatchReportResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantBatchRequest

> MerchantBatchResponse MerchantBatchRequest(ctx, merchantid, batchNo).Execute()

Merchant Batch Request



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
	merchantid := int32(56) // int32 | A merchant ID (MID) for which data is requested. This field allows for filtering of the request by a specific merchant account.
	batchNo := "batchNo_example" // string | The batch number that is being requested.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingApi.MerchantBatchRequest(context.Background(), merchantid, batchNo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.MerchantBatchRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantBatchRequest`: MerchantBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportingApi.MerchantBatchRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantid** | **int32** | A merchant ID (MID) for which data is requested. This field allows for filtering of the request by a specific merchant account. | 
**batchNo** | **string** | The batch number that is being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMerchantBatchRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MerchantBatchResponse**](MerchantBatchResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemittanceRangeReport

> RemittanceReportResponse RemittanceRangeReport(ctx, clientid).RemittanceReportRequest(remittanceReportRequest).Execute()

Remittance Report Request



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
	clientid := "clientid_example" // string | A client Id for which data is requested.
	remittanceReportRequest := *openapiclient.NewRemittanceReportRequest() // RemittanceReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingApi.RemittanceRangeReport(context.Background(), clientid).RemittanceReportRequest(remittanceReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.RemittanceRangeReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemittanceRangeReport`: RemittanceReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportingApi.RemittanceRangeReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientid** | **string** | A client Id for which data is requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemittanceRangeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remittanceReportRequest** | [**RemittanceReportRequest**](RemittanceReportRequest.md) |  | 

### Return type

[**RemittanceReportResponse**](RemittanceReportResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemittanceReportRequest

> RemittedClientData RemittanceReportRequest(ctx, clientid, date).Execute()

Remittance Date Report Request



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
	clientid := "clientid_example" // string | A client Id for which data is requested.
	date := "date_example" // string | Date (YYYY-MM-DD) to filter the request for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingApi.RemittanceReportRequest(context.Background(), clientid, date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingApi.RemittanceReportRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemittanceReportRequest`: RemittedClientData
	fmt.Fprintf(os.Stdout, "Response from `ReportingApi.RemittanceReportRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientid** | **string** | A client Id for which data is requested. | 
**date** | **string** | Date (YYYY-MM-DD) to filter the request for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemittanceReportRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemittedClientData**](RemittedClientData.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

