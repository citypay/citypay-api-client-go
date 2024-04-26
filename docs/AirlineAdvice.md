# AirlineAdvice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierName** | **string** | The name of the airline carrier that generated the tickets for airline travel. | 
**ConjunctionTicketIndicator** | Pointer to **bool** | true if a conjunction ticket (with additional coupons) was issued for an itinerary with more than four segments. Defaults to false.  | [optional] 
**EticketIndicator** | Pointer to **bool** | The Electronic Ticket Indicator, a code that indicates if an electronic ticket was issued.  Defaults to true. | [optional] 
**NoAirSegments** | Pointer to **int32** | A value that indicates the number of air travel segments included on this ticket.  Valid entries include the numerals “0” through “4”. Required only if the transaction type is TKT or EXC.  | [optional] 
**NumberInParty** | Pointer to **int32** | The number of people in the party. | [optional] 
**OriginalTicketNo** | Pointer to **string** | Required if transaction type is EXC. | [optional] 
**PassengerName** | Pointer to **string** | The name of the passenger when the traveller is not the card member that purchased the ticket. Required only if the transaction type is TKT or EXC. | [optional] 
**Segment1** | [**AirlineSegment**](AirlineSegment.md) |  | 
**Segment2** | Pointer to [**AirlineSegment**](AirlineSegment.md) |  | [optional] 
**Segment3** | Pointer to [**AirlineSegment**](AirlineSegment.md) |  | [optional] 
**Segment4** | Pointer to [**AirlineSegment**](AirlineSegment.md) |  | [optional] 
**TicketIssueCity** | **string** | The name of the city town or village where the transaction took place. | 
**TicketIssueDate** | **string** | The date the ticket was issued in ISO Date format (yyyy-MM-dd). | 
**TicketIssueName** | **string** | The name of the agency generating the ticket. | 
**TicketNo** | **string** | This must be a valid ticket number, i.e. numeric (the first 3 digits must represent the valid IATA plate carrier code). The final check digit should be validated prior to submission. On credit charges, this field should contain the number of the original ticket, and not of a replacement.  | 
**TransactionType** | **string** | This field contains the Transaction Type code assigned to this transaction. Valid codes include:   - &#x60;TKT&#x60; &#x3D; Ticket Purchase   - &#x60;REF&#x60; &#x3D; Refund   - &#x60;EXC&#x60; &#x3D; Exchange Ticket   - &#x60;MSC&#x60; &#x3D; Miscellaneous (non-Ticket Purchase- and non-Exchange Ticket-related transactions only).  | 

## Methods

### NewAirlineAdvice

`func NewAirlineAdvice(carrierName string, segment1 AirlineSegment, ticketIssueCity string, ticketIssueDate string, ticketIssueName string, ticketNo string, transactionType string, ) *AirlineAdvice`

NewAirlineAdvice instantiates a new AirlineAdvice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirlineAdviceWithDefaults

`func NewAirlineAdviceWithDefaults() *AirlineAdvice`

NewAirlineAdviceWithDefaults instantiates a new AirlineAdvice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierName

`func (o *AirlineAdvice) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *AirlineAdvice) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *AirlineAdvice) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.


### GetConjunctionTicketIndicator

`func (o *AirlineAdvice) GetConjunctionTicketIndicator() bool`

GetConjunctionTicketIndicator returns the ConjunctionTicketIndicator field if non-nil, zero value otherwise.

### GetConjunctionTicketIndicatorOk

`func (o *AirlineAdvice) GetConjunctionTicketIndicatorOk() (*bool, bool)`

GetConjunctionTicketIndicatorOk returns a tuple with the ConjunctionTicketIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConjunctionTicketIndicator

`func (o *AirlineAdvice) SetConjunctionTicketIndicator(v bool)`

SetConjunctionTicketIndicator sets ConjunctionTicketIndicator field to given value.

### HasConjunctionTicketIndicator

`func (o *AirlineAdvice) HasConjunctionTicketIndicator() bool`

HasConjunctionTicketIndicator returns a boolean if a field has been set.

### GetEticketIndicator

`func (o *AirlineAdvice) GetEticketIndicator() bool`

GetEticketIndicator returns the EticketIndicator field if non-nil, zero value otherwise.

### GetEticketIndicatorOk

`func (o *AirlineAdvice) GetEticketIndicatorOk() (*bool, bool)`

GetEticketIndicatorOk returns a tuple with the EticketIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEticketIndicator

`func (o *AirlineAdvice) SetEticketIndicator(v bool)`

SetEticketIndicator sets EticketIndicator field to given value.

### HasEticketIndicator

`func (o *AirlineAdvice) HasEticketIndicator() bool`

HasEticketIndicator returns a boolean if a field has been set.

### GetNoAirSegments

`func (o *AirlineAdvice) GetNoAirSegments() int32`

GetNoAirSegments returns the NoAirSegments field if non-nil, zero value otherwise.

### GetNoAirSegmentsOk

`func (o *AirlineAdvice) GetNoAirSegmentsOk() (*int32, bool)`

GetNoAirSegmentsOk returns a tuple with the NoAirSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAirSegments

`func (o *AirlineAdvice) SetNoAirSegments(v int32)`

SetNoAirSegments sets NoAirSegments field to given value.

### HasNoAirSegments

`func (o *AirlineAdvice) HasNoAirSegments() bool`

HasNoAirSegments returns a boolean if a field has been set.

### GetNumberInParty

`func (o *AirlineAdvice) GetNumberInParty() int32`

GetNumberInParty returns the NumberInParty field if non-nil, zero value otherwise.

### GetNumberInPartyOk

`func (o *AirlineAdvice) GetNumberInPartyOk() (*int32, bool)`

GetNumberInPartyOk returns a tuple with the NumberInParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberInParty

`func (o *AirlineAdvice) SetNumberInParty(v int32)`

SetNumberInParty sets NumberInParty field to given value.

### HasNumberInParty

`func (o *AirlineAdvice) HasNumberInParty() bool`

HasNumberInParty returns a boolean if a field has been set.

### GetOriginalTicketNo

`func (o *AirlineAdvice) GetOriginalTicketNo() string`

GetOriginalTicketNo returns the OriginalTicketNo field if non-nil, zero value otherwise.

### GetOriginalTicketNoOk

`func (o *AirlineAdvice) GetOriginalTicketNoOk() (*string, bool)`

GetOriginalTicketNoOk returns a tuple with the OriginalTicketNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTicketNo

`func (o *AirlineAdvice) SetOriginalTicketNo(v string)`

SetOriginalTicketNo sets OriginalTicketNo field to given value.

### HasOriginalTicketNo

`func (o *AirlineAdvice) HasOriginalTicketNo() bool`

HasOriginalTicketNo returns a boolean if a field has been set.

### GetPassengerName

`func (o *AirlineAdvice) GetPassengerName() string`

GetPassengerName returns the PassengerName field if non-nil, zero value otherwise.

### GetPassengerNameOk

`func (o *AirlineAdvice) GetPassengerNameOk() (*string, bool)`

GetPassengerNameOk returns a tuple with the PassengerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengerName

`func (o *AirlineAdvice) SetPassengerName(v string)`

SetPassengerName sets PassengerName field to given value.

### HasPassengerName

`func (o *AirlineAdvice) HasPassengerName() bool`

HasPassengerName returns a boolean if a field has been set.

### GetSegment1

`func (o *AirlineAdvice) GetSegment1() AirlineSegment`

GetSegment1 returns the Segment1 field if non-nil, zero value otherwise.

### GetSegment1Ok

`func (o *AirlineAdvice) GetSegment1Ok() (*AirlineSegment, bool)`

GetSegment1Ok returns a tuple with the Segment1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1

`func (o *AirlineAdvice) SetSegment1(v AirlineSegment)`

SetSegment1 sets Segment1 field to given value.


### GetSegment2

`func (o *AirlineAdvice) GetSegment2() AirlineSegment`

GetSegment2 returns the Segment2 field if non-nil, zero value otherwise.

### GetSegment2Ok

`func (o *AirlineAdvice) GetSegment2Ok() (*AirlineSegment, bool)`

GetSegment2Ok returns a tuple with the Segment2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2

`func (o *AirlineAdvice) SetSegment2(v AirlineSegment)`

SetSegment2 sets Segment2 field to given value.

### HasSegment2

`func (o *AirlineAdvice) HasSegment2() bool`

HasSegment2 returns a boolean if a field has been set.

### GetSegment3

`func (o *AirlineAdvice) GetSegment3() AirlineSegment`

GetSegment3 returns the Segment3 field if non-nil, zero value otherwise.

### GetSegment3Ok

`func (o *AirlineAdvice) GetSegment3Ok() (*AirlineSegment, bool)`

GetSegment3Ok returns a tuple with the Segment3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3

`func (o *AirlineAdvice) SetSegment3(v AirlineSegment)`

SetSegment3 sets Segment3 field to given value.

### HasSegment3

`func (o *AirlineAdvice) HasSegment3() bool`

HasSegment3 returns a boolean if a field has been set.

### GetSegment4

`func (o *AirlineAdvice) GetSegment4() AirlineSegment`

GetSegment4 returns the Segment4 field if non-nil, zero value otherwise.

### GetSegment4Ok

`func (o *AirlineAdvice) GetSegment4Ok() (*AirlineSegment, bool)`

GetSegment4Ok returns a tuple with the Segment4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment4

`func (o *AirlineAdvice) SetSegment4(v AirlineSegment)`

SetSegment4 sets Segment4 field to given value.

### HasSegment4

`func (o *AirlineAdvice) HasSegment4() bool`

HasSegment4 returns a boolean if a field has been set.

### GetTicketIssueCity

`func (o *AirlineAdvice) GetTicketIssueCity() string`

GetTicketIssueCity returns the TicketIssueCity field if non-nil, zero value otherwise.

### GetTicketIssueCityOk

`func (o *AirlineAdvice) GetTicketIssueCityOk() (*string, bool)`

GetTicketIssueCityOk returns a tuple with the TicketIssueCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketIssueCity

`func (o *AirlineAdvice) SetTicketIssueCity(v string)`

SetTicketIssueCity sets TicketIssueCity field to given value.


### GetTicketIssueDate

`func (o *AirlineAdvice) GetTicketIssueDate() string`

GetTicketIssueDate returns the TicketIssueDate field if non-nil, zero value otherwise.

### GetTicketIssueDateOk

`func (o *AirlineAdvice) GetTicketIssueDateOk() (*string, bool)`

GetTicketIssueDateOk returns a tuple with the TicketIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketIssueDate

`func (o *AirlineAdvice) SetTicketIssueDate(v string)`

SetTicketIssueDate sets TicketIssueDate field to given value.


### GetTicketIssueName

`func (o *AirlineAdvice) GetTicketIssueName() string`

GetTicketIssueName returns the TicketIssueName field if non-nil, zero value otherwise.

### GetTicketIssueNameOk

`func (o *AirlineAdvice) GetTicketIssueNameOk() (*string, bool)`

GetTicketIssueNameOk returns a tuple with the TicketIssueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketIssueName

`func (o *AirlineAdvice) SetTicketIssueName(v string)`

SetTicketIssueName sets TicketIssueName field to given value.


### GetTicketNo

`func (o *AirlineAdvice) GetTicketNo() string`

GetTicketNo returns the TicketNo field if non-nil, zero value otherwise.

### GetTicketNoOk

`func (o *AirlineAdvice) GetTicketNoOk() (*string, bool)`

GetTicketNoOk returns a tuple with the TicketNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketNo

`func (o *AirlineAdvice) SetTicketNo(v string)`

SetTicketNo sets TicketNo field to given value.


### GetTransactionType

`func (o *AirlineAdvice) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *AirlineAdvice) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *AirlineAdvice) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


