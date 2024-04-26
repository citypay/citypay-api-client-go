# \CardHolderAccountApi

All URIs are relative to *https://api.citypay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountCardDeleteRequest**](CardHolderAccountApi.md#AccountCardDeleteRequest) | **Delete** /v6/account/{accountid}/card/{cardId} | Card Deletion
[**AccountCardRegisterRequest**](CardHolderAccountApi.md#AccountCardRegisterRequest) | **Post** /v6/account/{accountid}/register | Card Registration
[**AccountCardStatusRequest**](CardHolderAccountApi.md#AccountCardStatusRequest) | **Post** /v6/account/{accountid}/card/{cardId}/status | Card Status
[**AccountChangeContactRequest**](CardHolderAccountApi.md#AccountChangeContactRequest) | **Post** /v6/account/{accountid}/contact | Contact Details Update
[**AccountCreate**](CardHolderAccountApi.md#AccountCreate) | **Post** /v6/account/create | Account Create
[**AccountDeleteRequest**](CardHolderAccountApi.md#AccountDeleteRequest) | **Delete** /v6/account/{accountid} | Account Deletion
[**AccountExistsRequest**](CardHolderAccountApi.md#AccountExistsRequest) | **Get** /v6/account-exists/{accountid} | Account Exists
[**AccountRetrieveRequest**](CardHolderAccountApi.md#AccountRetrieveRequest) | **Get** /v6/account/{accountid} | Account Retrieval
[**AccountStatusRequest**](CardHolderAccountApi.md#AccountStatusRequest) | **Post** /v6/account/{accountid}/status | Account Status
[**ChargeRequest**](CardHolderAccountApi.md#ChargeRequest) | **Post** /v6/charge | Charge



## AccountCardDeleteRequest

> Acknowledgement AccountCardDeleteRequest(ctx, accountid, cardId).Force(force).Execute()

Card Deletion



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
	accountid := "accountid_example" // string | The account id that refers to the customer's account no. This value will have been provided when setting up the card holder account.
	cardId := "cardId_example" // string | The id of the card that is presented by a call to retrieve a card holder account.
	force := true // bool | Requests that the item is forced immediately. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardHolderAccountApi.AccountCardDeleteRequest(context.Background(), accountid, cardId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardHolderAccountApi.AccountCardDeleteRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountCardDeleteRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `CardHolderAccountApi.AccountCardDeleteRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | The account id that refers to the customer&#39;s account no. This value will have been provided when setting up the card holder account. | 
**cardId** | **string** | The id of the card that is presented by a call to retrieve a card holder account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountCardDeleteRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | Requests that the item is forced immediately. | 

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


## AccountCardRegisterRequest

> CardHolderAccount AccountCardRegisterRequest(ctx, accountid).RegisterCard(registerCard).Execute()

Card Registration



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
	accountid := "accountid_example" // string | The account id that refers to the customer's account no. This value will have been provided when setting up the card holder account.
	registerCard := *openapiclient.NewRegisterCard("4000 0000 0000 0002", int32(9), int32(2027)) // RegisterCard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardHolderAccountApi.AccountCardRegisterRequest(context.Background(), accountid).RegisterCard(registerCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardHolderAccountApi.AccountCardRegisterRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountCardRegisterRequest`: CardHolderAccount
	fmt.Fprintf(os.Stdout, "Response from `CardHolderAccountApi.AccountCardRegisterRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | The account id that refers to the customer&#39;s account no. This value will have been provided when setting up the card holder account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountCardRegisterRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerCard** | [**RegisterCard**](RegisterCard.md) |  | 

### Return type

[**CardHolderAccount**](CardHolderAccount.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountCardStatusRequest

> Acknowledgement AccountCardStatusRequest(ctx, accountid, cardId).CardStatus(cardStatus).Execute()

Card Status



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
	accountid := "accountid_example" // string | The account id that refers to the customer's account no. This value will have been provided when setting up the card holder account.
	cardId := "cardId_example" // string | The id of the card that is presented by a call to retrieve a card holder account.
	cardStatus := *openapiclient.NewCardStatus() // CardStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardHolderAccountApi.AccountCardStatusRequest(context.Background(), accountid, cardId).CardStatus(cardStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardHolderAccountApi.AccountCardStatusRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountCardStatusRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `CardHolderAccountApi.AccountCardStatusRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | The account id that refers to the customer&#39;s account no. This value will have been provided when setting up the card holder account. | 
**cardId** | **string** | The id of the card that is presented by a call to retrieve a card holder account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountCardStatusRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cardStatus** | [**CardStatus**](CardStatus.md) |  | 

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


## AccountChangeContactRequest

> CardHolderAccount AccountChangeContactRequest(ctx, accountid).ContactDetails(contactDetails).Execute()

Contact Details Update



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
	accountid := "accountid_example" // string | The account id that refers to the customer's account no. This value will have been provided when setting up the card holder account.
	contactDetails := *openapiclient.NewContactDetails() // ContactDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardHolderAccountApi.AccountChangeContactRequest(context.Background(), accountid).ContactDetails(contactDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardHolderAccountApi.AccountChangeContactRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountChangeContactRequest`: CardHolderAccount
	fmt.Fprintf(os.Stdout, "Response from `CardHolderAccountApi.AccountChangeContactRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | The account id that refers to the customer&#39;s account no. This value will have been provided when setting up the card holder account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountChangeContactRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactDetails** | [**ContactDetails**](ContactDetails.md) |  | 

### Return type

[**CardHolderAccount**](CardHolderAccount.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountCreate

> CardHolderAccount AccountCreate(ctx).AccountCreate(accountCreate).Execute()

Account Create



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
	accountCreate := *openapiclient.NewAccountCreate("aaabbb-cccddd-eee") // AccountCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardHolderAccountApi.AccountCreate(context.Background()).AccountCreate(accountCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardHolderAccountApi.AccountCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountCreate`: CardHolderAccount
	fmt.Fprintf(os.Stdout, "Response from `CardHolderAccountApi.AccountCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountCreate** | [**AccountCreate**](AccountCreate.md) |  | 

### Return type

[**CardHolderAccount**](CardHolderAccount.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: application/json, text/xml
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountDeleteRequest

> Acknowledgement AccountDeleteRequest(ctx, accountid).Execute()

Account Deletion



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
	accountid := "accountid_example" // string | The account id that refers to the customer's account no. This value will have been provided when setting up the card holder account.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardHolderAccountApi.AccountDeleteRequest(context.Background(), accountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardHolderAccountApi.AccountDeleteRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountDeleteRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `CardHolderAccountApi.AccountDeleteRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | The account id that refers to the customer&#39;s account no. This value will have been provided when setting up the card holder account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountDeleteRequestRequest struct via the builder pattern


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


## AccountExistsRequest

> Exists AccountExistsRequest(ctx, accountid).Execute()

Account Exists



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
	accountid := "accountid_example" // string | The account id that refers to the customer's account no. This value will have been provided when setting up the card holder account.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardHolderAccountApi.AccountExistsRequest(context.Background(), accountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardHolderAccountApi.AccountExistsRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountExistsRequest`: Exists
	fmt.Fprintf(os.Stdout, "Response from `CardHolderAccountApi.AccountExistsRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | The account id that refers to the customer&#39;s account no. This value will have been provided when setting up the card holder account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountExistsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Exists**](Exists.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountRetrieveRequest

> CardHolderAccount AccountRetrieveRequest(ctx, accountid).Execute()

Account Retrieval



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
	accountid := "accountid_example" // string | The account id that refers to the customer's account no. This value will have been provided when setting up the card holder account.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardHolderAccountApi.AccountRetrieveRequest(context.Background(), accountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardHolderAccountApi.AccountRetrieveRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountRetrieveRequest`: CardHolderAccount
	fmt.Fprintf(os.Stdout, "Response from `CardHolderAccountApi.AccountRetrieveRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | The account id that refers to the customer&#39;s account no. This value will have been provided when setting up the card holder account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountRetrieveRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CardHolderAccount**](CardHolderAccount.md)

### Authorization

[cp-api-key](../README.md#cp-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountStatusRequest

> Acknowledgement AccountStatusRequest(ctx, accountid).AccountStatus(accountStatus).Execute()

Account Status



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
	accountid := "accountid_example" // string | The account id that refers to the customer's account no. This value will have been provided when setting up the card holder account.
	accountStatus := *openapiclient.NewAccountStatus() // AccountStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardHolderAccountApi.AccountStatusRequest(context.Background(), accountid).AccountStatus(accountStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardHolderAccountApi.AccountStatusRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountStatusRequest`: Acknowledgement
	fmt.Fprintf(os.Stdout, "Response from `CardHolderAccountApi.AccountStatusRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | The account id that refers to the customer&#39;s account no. This value will have been provided when setting up the card holder account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountStatusRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountStatus** | [**AccountStatus**](AccountStatus.md) |  | 

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


## ChargeRequest

> Decision ChargeRequest(ctx).ChargeRequest(chargeRequest).Execute()

Charge



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
	chargeRequest := *openapiclient.NewChargeRequest(int32(19995), "95b857a1-5955-4b86-963c-5a6dbfc4fb95", int32(11223344), "ctPCAPyNyCkx3Ry8wGyv8khC3ch2hUSB3Db..Qzr") // ChargeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardHolderAccountApi.ChargeRequest(context.Background()).ChargeRequest(chargeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardHolderAccountApi.ChargeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChargeRequest`: Decision
	fmt.Fprintf(os.Stdout, "Response from `CardHolderAccountApi.ChargeRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChargeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeRequest** | [**ChargeRequest**](ChargeRequest.md) |  | 

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

