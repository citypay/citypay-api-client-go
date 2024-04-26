# \BatchProcessingApi

All URIs are relative to *https://api.citypay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchProcessRequest**](BatchProcessingApi.md#BatchProcessRequest) | **Post** /v6/batch/process | Batch Process Request
[**BatchRetrieveRequest**](BatchProcessingApi.md#BatchRetrieveRequest) | **Post** /v6/batch/retrieve | Batch Retrieve Request
[**CheckBatchStatusRequest**](BatchProcessingApi.md#CheckBatchStatusRequest) | **Post** /v6/batch/status | Check Batch Status



## BatchProcessRequest

> ProcessBatchResponse BatchProcessRequest(ctx).ProcessBatchRequest(processBatchRequest).Execute()

Batch Process Request



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
	processBatchRequest := *openapiclient.NewProcessBatchRequest(time.Now(), int32(35), []openapiclient.BatchTransaction{*openapiclient.NewBatchTransaction("aaabbb-cccddd-eee", int32(19995))}) // ProcessBatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchProcessingApi.BatchProcessRequest(context.Background()).ProcessBatchRequest(processBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchProcessingApi.BatchProcessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchProcessRequest`: ProcessBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchProcessingApi.BatchProcessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchProcessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processBatchRequest** | [**ProcessBatchRequest**](ProcessBatchRequest.md) |  | 

### Return type

[**ProcessBatchResponse**](ProcessBatchResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRetrieveRequest

> BatchReportResponseModel BatchRetrieveRequest(ctx).BatchReportRequest(batchReportRequest).Execute()

Batch Retrieve Request



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
	batchReportRequest := *openapiclient.NewBatchReportRequest(int32(35)) // BatchReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchProcessingApi.BatchRetrieveRequest(context.Background()).BatchReportRequest(batchReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchProcessingApi.BatchRetrieveRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchRetrieveRequest`: BatchReportResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BatchProcessingApi.BatchRetrieveRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchRetrieveRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchReportRequest** | [**BatchReportRequest**](BatchReportRequest.md) |  | 

### Return type

[**BatchReportResponseModel**](BatchReportResponseModel.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckBatchStatusRequest

> CheckBatchStatusResponse CheckBatchStatusRequest(ctx).CheckBatchStatus(checkBatchStatus).Execute()

Check Batch Status



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
	checkBatchStatus := *openapiclient.NewCheckBatchStatus([]int32{int32(123)}) // CheckBatchStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchProcessingApi.CheckBatchStatusRequest(context.Background()).CheckBatchStatus(checkBatchStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchProcessingApi.CheckBatchStatusRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckBatchStatusRequest`: CheckBatchStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchProcessingApi.CheckBatchStatusRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckBatchStatusRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkBatchStatus** | [**CheckBatchStatus**](CheckBatchStatus.md) |  | 

### Return type

[**CheckBatchStatusResponse**](CheckBatchStatusResponse.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

