# Go API client for citypay


Welcome to the CityPay API, a robust HTTP API payment solution designed for seamless server-to-server 
transactional processing. Our API facilitates a wide array of payment operations, catering to diverse business needs. 
Whether you're integrating Internet payments, handling Mail Order/Telephone Order (MOTO) transactions, managing 
Subscriptions with Recurring and Continuous Authority payments, or navigating the complexities of 3-D Secure 
authentication, our API is equipped to support your requirements. Additionally, we offer functionalities for 
Authorisation, Refunding, Pre-Authorisation, Cancellation/Voids, and Completion processing, alongside the capability 
for tokenised payments.

## Compliance and Security Overview
<aside class=\"notice\">
  Ensuring the security of payment transactions and compliance with industry standards is paramount. Our API is 
  designed with stringent security measures and compliance protocols to safeguard sensitive information and meet 
  the rigorous requirements of Visa, MasterCard, and the PCI Security Standards Council.
</aside>

### Key Compliance and Security Measures

* **TLS Encryption**: All data transmissions must utilise TLS version 1.2 or higher, employing [strong cryptography](#enabled-tls-ciphers). Our infrastructure strictly enforces this requirement to maintain the integrity and confidentiality of data in transit. We conduct regular scans and assessments of our TLS endpoints to identify and mitigate vulnerabilities.
* **Data Storage Prohibitions**: Storing sensitive cardholder data (CHD), such as the card security code (CSC) or primary account number (PAN), is strictly prohibited. Our API is designed to minimize your exposure to sensitive data, thereby reducing your compliance burden.
* **Data Masking**: For consumer protection and compliance, full card numbers must not be displayed on receipts or any customer-facing materials. Our API automatically masks PANs, displaying only the last four digits to facilitate safe receipt generation.
* **Network Scans**: If your application is web-based, regular scans of your hosting environment are mandatory to identify and rectify potential vulnerabilities. This proactive measure is crucial for maintaining a secure and compliant online presence.
* **PCI Compliance**: Adherence to PCI DSS standards is not optional; it's a requirement for operating securely and legally in the payments ecosystem. For detailed information on compliance requirements and resources, please visit the PCI Security Standards Council website [https://www.pcisecuritystandards.org/](https://www.pcisecuritystandards.org/).
* **Request Validation**: Our API includes mechanisms to verify the legitimacy of each request, ensuring it pertains to a valid account and originates from a trusted source. We leverage remote IP address verification alongside sophisticated application firewall technologies to thwart a wide array of common security threats.

## Getting Started
Before integrating with the CityPay API, ensure your application and development practices align with the outlined compliance and security measures. This preparatory step is crucial for a smooth integration process and the long-term success of your payment processing operations.

For further details on API endpoints, request/response formats, and code examples, proceed to the subsequent sections of our documentation. Our aim is to provide you with all the necessary tools and information to integrate our payment processing capabilities seamlessly into your application.

Thank you for choosing CityPay API. We look forward to supporting your payment processing needs with our secure, compliant, and versatile API solution.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 6.6.40
- Package version: 1.1.0
- Build date: 2024-04-26T15:24:37.497360709Z[Etc/UTC]
- Generator version: 7.5.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.citypay.com/contacts/](https://www.citypay.com/contacts/)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import citypay "github.com/citypay/citypay-api-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `citypay.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), citypay.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `citypay.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), citypay.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `citypay.ContextOperationServerIndices` and `citypay.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), citypay.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), citypay.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.citypay.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthorisationAndPaymentApi* | [**AuthorisationRequest**](docs/AuthorisationAndPaymentApi.md#authorisationrequest) | **Post** /v6/authorise | Authorisation
*AuthorisationAndPaymentApi* | [**BinRangeLookupRequest**](docs/AuthorisationAndPaymentApi.md#binrangelookuprequest) | **Post** /v6/bin | Bin Lookup
*AuthorisationAndPaymentApi* | [**CResRequest**](docs/AuthorisationAndPaymentApi.md#cresrequest) | **Post** /v6/cres | CRes
*AuthorisationAndPaymentApi* | [**CaptureRequest**](docs/AuthorisationAndPaymentApi.md#capturerequest) | **Post** /v6/capture | Capture
*AuthorisationAndPaymentApi* | [**CreatePaymentIntent**](docs/AuthorisationAndPaymentApi.md#createpaymentintent) | **Post** /v6/intent/create | Create a Payment Intent
*AuthorisationAndPaymentApi* | [**PaResRequest**](docs/AuthorisationAndPaymentApi.md#paresrequest) | **Post** /v6/pares | PaRes
*AuthorisationAndPaymentApi* | [**RefundRequest**](docs/AuthorisationAndPaymentApi.md#refundrequest) | **Post** /v6/refund | Refund
*AuthorisationAndPaymentApi* | [**RetrievalRequest**](docs/AuthorisationAndPaymentApi.md#retrievalrequest) | **Post** /v6/retrieve | Retrieval
*AuthorisationAndPaymentApi* | [**VoidRequest**](docs/AuthorisationAndPaymentApi.md#voidrequest) | **Post** /v6/void | Void
*BatchProcessingApi* | [**BatchProcessRequest**](docs/BatchProcessingApi.md#batchprocessrequest) | **Post** /v6/batch/process | Batch Process Request
*BatchProcessingApi* | [**BatchRetrieveRequest**](docs/BatchProcessingApi.md#batchretrieverequest) | **Post** /v6/batch/retrieve | Batch Retrieve Request
*BatchProcessingApi* | [**CheckBatchStatusRequest**](docs/BatchProcessingApi.md#checkbatchstatusrequest) | **Post** /v6/batch/status | Check Batch Status
*CardHolderAccountApi* | [**AccountCardDeleteRequest**](docs/CardHolderAccountApi.md#accountcarddeleterequest) | **Delete** /v6/account/{accountid}/card/{cardId} | Card Deletion
*CardHolderAccountApi* | [**AccountCardRegisterRequest**](docs/CardHolderAccountApi.md#accountcardregisterrequest) | **Post** /v6/account/{accountid}/register | Card Registration
*CardHolderAccountApi* | [**AccountCardStatusRequest**](docs/CardHolderAccountApi.md#accountcardstatusrequest) | **Post** /v6/account/{accountid}/card/{cardId}/status | Card Status
*CardHolderAccountApi* | [**AccountChangeContactRequest**](docs/CardHolderAccountApi.md#accountchangecontactrequest) | **Post** /v6/account/{accountid}/contact | Contact Details Update
*CardHolderAccountApi* | [**AccountCreate**](docs/CardHolderAccountApi.md#accountcreate) | **Post** /v6/account/create | Account Create
*CardHolderAccountApi* | [**AccountDeleteRequest**](docs/CardHolderAccountApi.md#accountdeleterequest) | **Delete** /v6/account/{accountid} | Account Deletion
*CardHolderAccountApi* | [**AccountExistsRequest**](docs/CardHolderAccountApi.md#accountexistsrequest) | **Get** /v6/account-exists/{accountid} | Account Exists
*CardHolderAccountApi* | [**AccountRetrieveRequest**](docs/CardHolderAccountApi.md#accountretrieverequest) | **Get** /v6/account/{accountid} | Account Retrieval
*CardHolderAccountApi* | [**AccountStatusRequest**](docs/CardHolderAccountApi.md#accountstatusrequest) | **Post** /v6/account/{accountid}/status | Account Status
*CardHolderAccountApi* | [**ChargeRequest**](docs/CardHolderAccountApi.md#chargerequest) | **Post** /v6/charge | Charge
*DirectPostApi* | [**DirectCResAuthRequest**](docs/DirectPostApi.md#directcresauthrequest) | **Post** /direct/cres/auth/{uuid} | Handles a CRes response from ACS, returning back the result of authorisation
*DirectPostApi* | [**DirectCResTokeniseRequest**](docs/DirectPostApi.md#directcrestokeniserequest) | **Post** /direct/cres/tokenise/{uuid} | Handles a CRes response from ACS, returning back a token for future authorisation
*DirectPostApi* | [**DirectPostAuthRequest**](docs/DirectPostApi.md#directpostauthrequest) | **Post** /direct/auth | Direct Post Auth Request
*DirectPostApi* | [**DirectPostTokeniseRequest**](docs/DirectPostApi.md#directposttokeniserequest) | **Post** /direct/tokenise | Direct Post Tokenise Request
*DirectPostApi* | [**TokenRequest**](docs/DirectPostApi.md#tokenrequest) | **Post** /direct/token | Direct Post Token Request
*OperationalFunctionsApi* | [**AclCheckRequest**](docs/OperationalFunctionsApi.md#aclcheckrequest) | **Post** /v6/acl/check | ACL Check Request
*OperationalFunctionsApi* | [**DomainKeyCheckRequest**](docs/OperationalFunctionsApi.md#domainkeycheckrequest) | **Post** /dk/check | Domain Key Check Request
*OperationalFunctionsApi* | [**DomainKeyGenRequest**](docs/OperationalFunctionsApi.md#domainkeygenrequest) | **Post** /dk/gen | Domain Key Generation Request
*OperationalFunctionsApi* | [**ListMerchantsRequest**](docs/OperationalFunctionsApi.md#listmerchantsrequest) | **Get** /v6/merchants/{clientid} | List Merchants Request
*OperationalFunctionsApi* | [**PingRequest**](docs/OperationalFunctionsApi.md#pingrequest) | **Post** /v6/ping | Ping Request
*PaylinkApi* | [**TokenAdjustmentRequest**](docs/PaylinkApi.md#tokenadjustmentrequest) | **Post** /paylink/{token}/adjustment | Paylink Token Adjustment
*PaylinkApi* | [**TokenCancelRequest**](docs/PaylinkApi.md#tokencancelrequest) | **Put** /paylink/{token}/cancel | Cancel a Paylink Token
*PaylinkApi* | [**TokenChangesRequest**](docs/PaylinkApi.md#tokenchangesrequest) | **Post** /paylink/token/changes | Paylink Token Audit
*PaylinkApi* | [**TokenCloseRequest**](docs/PaylinkApi.md#tokencloserequest) | **Put** /paylink/{token}/close | Close Paylink Token
*PaylinkApi* | [**TokenCreateBillPaymentRequest**](docs/PaylinkApi.md#tokencreatebillpaymentrequest) | **Post** /paylink/bill-payment | Create Bill Payment Paylink Token
*PaylinkApi* | [**TokenCreateRequest**](docs/PaylinkApi.md#tokencreaterequest) | **Post** /paylink/create | Create Paylink Token
*PaylinkApi* | [**TokenPurgeAttachmentsRequest**](docs/PaylinkApi.md#tokenpurgeattachmentsrequest) | **Put** /paylink/{token}/purge-attachments | Purges any attachments for a Paylink Token
*PaylinkApi* | [**TokenReconciledRequest**](docs/PaylinkApi.md#tokenreconciledrequest) | **Put** /paylink/{token}/reconciled | Reconcile Paylink Token
*PaylinkApi* | [**TokenReopenRequest**](docs/PaylinkApi.md#tokenreopenrequest) | **Put** /paylink/{token}/reopen | Reopen Paylink Token
*PaylinkApi* | [**TokenResendNotificationRequest**](docs/PaylinkApi.md#tokenresendnotificationrequest) | **Post** /paylink/{token}/resend-notification | Resend a notification for Paylink Token
*PaylinkApi* | [**TokenStatusRequest**](docs/PaylinkApi.md#tokenstatusrequest) | **Get** /paylink/{token}/status | Paylink Token Status
*ReportingApi* | [**BatchedTransactionReportRequest**](docs/ReportingApi.md#batchedtransactionreportrequest) | **Post** /v6/merchant-batch/{merchantid}/{batch_no}/transactions | Batch Transaction Report Request
*ReportingApi* | [**MerchantBatchReportRequest**](docs/ReportingApi.md#merchantbatchreportrequest) | **Post** /v6/merchant-batch/report | Merchant Batch Report Request
*ReportingApi* | [**MerchantBatchRequest**](docs/ReportingApi.md#merchantbatchrequest) | **Get** /v6/merchant-batch/{merchantid}/{batch_no} | Merchant Batch Request
*ReportingApi* | [**RemittanceRangeReport**](docs/ReportingApi.md#remittancerangereport) | **Post** /v6/remittance/report/{clientid} | Remittance Report Request
*ReportingApi* | [**RemittanceReportRequest**](docs/ReportingApi.md#remittancereportrequest) | **Get** /v6/remittance/report/{clientid}/{date} | Remittance Date Report Request


## Documentation For Models

 - [AccountCreate](docs/AccountCreate.md)
 - [AccountStatus](docs/AccountStatus.md)
 - [Acknowledgement](docs/Acknowledgement.md)
 - [AclCheckRequest](docs/AclCheckRequest.md)
 - [AclCheckResponseModel](docs/AclCheckResponseModel.md)
 - [AirlineAdvice](docs/AirlineAdvice.md)
 - [AirlineSegment](docs/AirlineSegment.md)
 - [AuthReference](docs/AuthReference.md)
 - [AuthReferences](docs/AuthReferences.md)
 - [AuthRequest](docs/AuthRequest.md)
 - [AuthResponse](docs/AuthResponse.md)
 - [Batch](docs/Batch.md)
 - [BatchReportRequest](docs/BatchReportRequest.md)
 - [BatchReportResponseModel](docs/BatchReportResponseModel.md)
 - [BatchTransaction](docs/BatchTransaction.md)
 - [BatchTransactionReportRequest](docs/BatchTransactionReportRequest.md)
 - [BatchTransactionReportResponse](docs/BatchTransactionReportResponse.md)
 - [BatchTransactionResultModel](docs/BatchTransactionResultModel.md)
 - [Bin](docs/Bin.md)
 - [BinLookup](docs/BinLookup.md)
 - [CResAuthRequest](docs/CResAuthRequest.md)
 - [CaptureRequest](docs/CaptureRequest.md)
 - [Card](docs/Card.md)
 - [CardHolderAccount](docs/CardHolderAccount.md)
 - [CardStatus](docs/CardStatus.md)
 - [ChargeRequest](docs/ChargeRequest.md)
 - [CheckBatchStatus](docs/CheckBatchStatus.md)
 - [CheckBatchStatusResponse](docs/CheckBatchStatusResponse.md)
 - [ContactDetails](docs/ContactDetails.md)
 - [Decision](docs/Decision.md)
 - [DirectPostRequest](docs/DirectPostRequest.md)
 - [DirectTokenAuthRequest](docs/DirectTokenAuthRequest.md)
 - [DomainKeyCheckRequest](docs/DomainKeyCheckRequest.md)
 - [DomainKeyRequest](docs/DomainKeyRequest.md)
 - [DomainKeyResponse](docs/DomainKeyResponse.md)
 - [Error](docs/Error.md)
 - [EventDataModel](docs/EventDataModel.md)
 - [Exists](docs/Exists.md)
 - [ExternalMPI](docs/ExternalMPI.md)
 - [ListMerchantsResponse](docs/ListMerchantsResponse.md)
 - [MCC6012](docs/MCC6012.md)
 - [Merchant](docs/Merchant.md)
 - [MerchantBatchReportRequest](docs/MerchantBatchReportRequest.md)
 - [MerchantBatchReportResponse](docs/MerchantBatchReportResponse.md)
 - [MerchantBatchResponse](docs/MerchantBatchResponse.md)
 - [NetSummaryResponse](docs/NetSummaryResponse.md)
 - [PaResAuthRequest](docs/PaResAuthRequest.md)
 - [PaylinkAddress](docs/PaylinkAddress.md)
 - [PaylinkAdjustmentRequest](docs/PaylinkAdjustmentRequest.md)
 - [PaylinkAttachmentRequest](docs/PaylinkAttachmentRequest.md)
 - [PaylinkAttachmentResult](docs/PaylinkAttachmentResult.md)
 - [PaylinkBillPaymentTokenRequest](docs/PaylinkBillPaymentTokenRequest.md)
 - [PaylinkCardHolder](docs/PaylinkCardHolder.md)
 - [PaylinkCart](docs/PaylinkCart.md)
 - [PaylinkCartItemModel](docs/PaylinkCartItemModel.md)
 - [PaylinkConfig](docs/PaylinkConfig.md)
 - [PaylinkCustomParam](docs/PaylinkCustomParam.md)
 - [PaylinkEmailNotificationPath](docs/PaylinkEmailNotificationPath.md)
 - [PaylinkErrorCode](docs/PaylinkErrorCode.md)
 - [PaylinkFieldGuardModel](docs/PaylinkFieldGuardModel.md)
 - [PaylinkPartPayments](docs/PaylinkPartPayments.md)
 - [PaylinkResendNotificationRequest](docs/PaylinkResendNotificationRequest.md)
 - [PaylinkSMSNotificationPath](docs/PaylinkSMSNotificationPath.md)
 - [PaylinkStateEvent](docs/PaylinkStateEvent.md)
 - [PaylinkTokenCreated](docs/PaylinkTokenCreated.md)
 - [PaylinkTokenRequestModel](docs/PaylinkTokenRequestModel.md)
 - [PaylinkTokenStatus](docs/PaylinkTokenStatus.md)
 - [PaylinkTokenStatusChangeRequest](docs/PaylinkTokenStatusChangeRequest.md)
 - [PaylinkTokenStatusChangeResponse](docs/PaylinkTokenStatusChangeResponse.md)
 - [PaylinkUI](docs/PaylinkUI.md)
 - [PaymentIntent](docs/PaymentIntent.md)
 - [PaymentIntentReference](docs/PaymentIntentReference.md)
 - [Ping](docs/Ping.md)
 - [ProcessBatchRequest](docs/ProcessBatchRequest.md)
 - [ProcessBatchResponse](docs/ProcessBatchResponse.md)
 - [RefundRequest](docs/RefundRequest.md)
 - [RegisterCard](docs/RegisterCard.md)
 - [RemittanceData](docs/RemittanceData.md)
 - [RemittanceReportRequest](docs/RemittanceReportRequest.md)
 - [RemittanceReportResponse](docs/RemittanceReportResponse.md)
 - [RemittedClientData](docs/RemittedClientData.md)
 - [RequestChallenged](docs/RequestChallenged.md)
 - [RetrieveRequest](docs/RetrieveRequest.md)
 - [ThreeDSecure](docs/ThreeDSecure.md)
 - [TokenisationResponseModel](docs/TokenisationResponseModel.md)
 - [VoidRequest](docs/VoidRequest.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### cp-api-key

- **Type**: API key
- **API key parameter name**: cp-api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: cp-api-key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		citypay.ContextAPIKeys,
		map[string]citypay.APIKey{
			"cp-api-key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### cp-domain-key

- **Type**: API key
- **API key parameter name**: cp-domain-key
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: cp-domain-key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		citypay.ContextAPIKeys,
		map[string]citypay.APIKey{
			"cp-domain-key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@citypay.com
