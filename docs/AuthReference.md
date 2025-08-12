# AuthReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address of the card holder. | [optional] 
**Amount** | Pointer to **string** | The amount of the transaction in decimal currency format. | [optional] 
**AmountValue** | Pointer to **int32** | The amount of the transaction in integer/request format. | [optional] 
**Atrn** | Pointer to **string** | A reference number provided by the acquiring services. | [optional] 
**Authcode** | Pointer to **string** | The authorisation code of the transaction returned by the acquirer or card issuer. | [optional] 
**AuthenResult** | Pointer to **string** | The authentication result if an ecommerce transaction. &#39;Y&#39;. Authentication Successful, &#39;N&#39;. Authentication Failed, &#39;A&#39;. Attempts Processing Performed, &#39;U&#39;. Authentication Could Not Be Performed, &#39;C&#39;. Challenge Required. | [optional] 
**Batchno** | Pointer to **string** | A batch number which the transaction has been end of day batched towards. | [optional] 
**BinCommercial** | Pointer to **bool** | Whether the card is a commercial card. | [optional] 
**BinConsumer** | Pointer to **bool** | Whether the card is a consumer card. | [optional] 
**BinCorporate** | Pointer to **bool** | Whether the card is a corporate card. | [optional] 
**BinCredit** | Pointer to **bool** | Whether the card is a credit card. | [optional] 
**BinDebit** | Pointer to **bool** | Whether the card is a debit card. | [optional] 
**CardholderAgreement** | Pointer to **string** | Merchant-initiated transactions (MITs) are payments you trigger, where the cardholder has previously consented to you carrying out such payments. These may be scheduled (such as recurring payments and installments) or unscheduled (like account top-ups triggered by balance thresholds and no-show charges).  Scheduled These are regular payments using stored card details, like installments or a monthly subscription fee.  - &#x60;I&#x60; Instalment - A single purchase of goods or services billed to a cardholder in multiple transactions, over a period of time agreed by the cardholder and you.  - &#x60;R&#x60; Recurring - Transactions processed at fixed, regular intervals not to exceed one year between transactions, representing an agreement between a cardholder and you to purchase goods or services provided over a period of time.  Unscheduled These are payments using stored card details that do not occur on a regular schedule, like top-ups for a digital wallet triggered by the balance falling below a certain threshold.  - &#x60;A&#x60; Reauthorisation - a purchase made after the original purchase. A common scenario is delayed/split shipments.  - &#x60;C&#x60; Unscheduled Payment - A transaction using a stored credential for a fixed or variable amount that does not occur on a scheduled or regularly occurring transaction date. This includes account top-ups triggered by balance thresholds.  - &#x60;D&#x60; Delayed Charge - A delayed charge is typically used in hotel, cruise lines and vehicle rental environments to perform a supplemental account charge after original services are rendered.  - &#x60;L&#x60; Incremental - An incremental authorisation is typically found in hotel and car rental environments, where the cardholder has agreed to pay for any service incurred during the duration of the contract. An incremental authorisation is where you need to seek authorisation of further funds in addition to what you have originally requested. A common scenario is additional services charged to the contract, such as extending a stay in a hotel.  - &#x60;S&#x60; Resubmission - When the original purchase occurred, but you were not able to get authorisation at the time the goods or services were provided. It should be only used where the goods or services have already been provided, but the authorisation request is declined for insufficient funds.  - &#x60;X&#x60; No-show - A no-show is a transaction where you are enabled to charge for services which the cardholder entered into an agreement to purchase, but the cardholder did not meet the terms of the agreement.  - &#x60;N&#x60; Not Applicable - For all other transactions the value will be not applicable.  | [optional] 
**Currency** | Pointer to **string** | The currency of the transaction in ISO 4217 code format. | [optional] 
**Datetime** | Pointer to **time.Time** | The date and time of the transaction. | [optional] 
**Eci** | Pointer to **string** | The ECI if an ecommerce transaction. | [optional] 
**Email** | Pointer to **string** | The email address of the card holder. | [optional] 
**Env** | Pointer to **string** | The environment that the transaction is process within based on the transaction type. | [optional] 
**Identifier** | Pointer to **string** | The identifier of the transaction used to process the transaction. | [optional] 
**Initiation** | Pointer to **string** | The initiation of the payment. The value will be C for Card holder initiated and M for a merchant initiated transaction. | [optional] 
**Instrument** | Pointer to **string** | The payment instrument used such as Card, Cash, Bank, Crypto, ApplePay, GooglePay, Click2Pay, PayPal, OpenBankingPayment. | [optional] 
**Maskedpan** | Pointer to **string** | A masking of the card number which was used to process the tranasction. | [optional] 
**Merchantid** | Pointer to **int32** | The merchant id of the transaction result. | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**NameOnCard** | Pointer to **string** | The name of the card holder. | [optional] 
**Postcode** | Pointer to **string** | The postcode of the card holder. | [optional] 
**Result** | Pointer to **string** | The result of the transaction. | [optional] 
**ResultId** | Pointer to **string** | The id of the result of the transaction. | [optional] 
**Scheme** | Pointer to **string** | The card scheme of any card used. | [optional] 
**SchemeLogo** | Pointer to **string** | The card scheme logo of any card used. | [optional] 
**TransStatus** | Pointer to **string** | The current status of the transaction through it&#39;s lifecycle. | [optional] 
**TransType** | Pointer to **string** | The type code of transaction that was processed. | [optional] 
**Transno** | Pointer to **int32** | The transaction number of the transaction. | [optional] 
**Type** | Pointer to **string** | Defines whether the transaction is a sale, refund or verification. | [optional] 
**Utc** | Pointer to **int64** | The date and time of the transaction in UTC milli seconds since the epoc. | [optional] 

## Methods

### NewAuthReference

`func NewAuthReference() *AuthReference`

NewAuthReference instantiates a new AuthReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthReferenceWithDefaults

`func NewAuthReferenceWithDefaults() *AuthReference`

NewAuthReferenceWithDefaults instantiates a new AuthReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AuthReference) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AuthReference) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AuthReference) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AuthReference) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAmount

`func (o *AuthReference) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AuthReference) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AuthReference) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AuthReference) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountValue

`func (o *AuthReference) GetAmountValue() int32`

GetAmountValue returns the AmountValue field if non-nil, zero value otherwise.

### GetAmountValueOk

`func (o *AuthReference) GetAmountValueOk() (*int32, bool)`

GetAmountValueOk returns a tuple with the AmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountValue

`func (o *AuthReference) SetAmountValue(v int32)`

SetAmountValue sets AmountValue field to given value.

### HasAmountValue

`func (o *AuthReference) HasAmountValue() bool`

HasAmountValue returns a boolean if a field has been set.

### GetAtrn

`func (o *AuthReference) GetAtrn() string`

GetAtrn returns the Atrn field if non-nil, zero value otherwise.

### GetAtrnOk

`func (o *AuthReference) GetAtrnOk() (*string, bool)`

GetAtrnOk returns a tuple with the Atrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtrn

`func (o *AuthReference) SetAtrn(v string)`

SetAtrn sets Atrn field to given value.

### HasAtrn

`func (o *AuthReference) HasAtrn() bool`

HasAtrn returns a boolean if a field has been set.

### GetAuthcode

`func (o *AuthReference) GetAuthcode() string`

GetAuthcode returns the Authcode field if non-nil, zero value otherwise.

### GetAuthcodeOk

`func (o *AuthReference) GetAuthcodeOk() (*string, bool)`

GetAuthcodeOk returns a tuple with the Authcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthcode

`func (o *AuthReference) SetAuthcode(v string)`

SetAuthcode sets Authcode field to given value.

### HasAuthcode

`func (o *AuthReference) HasAuthcode() bool`

HasAuthcode returns a boolean if a field has been set.

### GetAuthenResult

`func (o *AuthReference) GetAuthenResult() string`

GetAuthenResult returns the AuthenResult field if non-nil, zero value otherwise.

### GetAuthenResultOk

`func (o *AuthReference) GetAuthenResultOk() (*string, bool)`

GetAuthenResultOk returns a tuple with the AuthenResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenResult

`func (o *AuthReference) SetAuthenResult(v string)`

SetAuthenResult sets AuthenResult field to given value.

### HasAuthenResult

`func (o *AuthReference) HasAuthenResult() bool`

HasAuthenResult returns a boolean if a field has been set.

### GetBatchno

`func (o *AuthReference) GetBatchno() string`

GetBatchno returns the Batchno field if non-nil, zero value otherwise.

### GetBatchnoOk

`func (o *AuthReference) GetBatchnoOk() (*string, bool)`

GetBatchnoOk returns a tuple with the Batchno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchno

`func (o *AuthReference) SetBatchno(v string)`

SetBatchno sets Batchno field to given value.

### HasBatchno

`func (o *AuthReference) HasBatchno() bool`

HasBatchno returns a boolean if a field has been set.

### GetBinCommercial

`func (o *AuthReference) GetBinCommercial() bool`

GetBinCommercial returns the BinCommercial field if non-nil, zero value otherwise.

### GetBinCommercialOk

`func (o *AuthReference) GetBinCommercialOk() (*bool, bool)`

GetBinCommercialOk returns a tuple with the BinCommercial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCommercial

`func (o *AuthReference) SetBinCommercial(v bool)`

SetBinCommercial sets BinCommercial field to given value.

### HasBinCommercial

`func (o *AuthReference) HasBinCommercial() bool`

HasBinCommercial returns a boolean if a field has been set.

### GetBinConsumer

`func (o *AuthReference) GetBinConsumer() bool`

GetBinConsumer returns the BinConsumer field if non-nil, zero value otherwise.

### GetBinConsumerOk

`func (o *AuthReference) GetBinConsumerOk() (*bool, bool)`

GetBinConsumerOk returns a tuple with the BinConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinConsumer

`func (o *AuthReference) SetBinConsumer(v bool)`

SetBinConsumer sets BinConsumer field to given value.

### HasBinConsumer

`func (o *AuthReference) HasBinConsumer() bool`

HasBinConsumer returns a boolean if a field has been set.

### GetBinCorporate

`func (o *AuthReference) GetBinCorporate() bool`

GetBinCorporate returns the BinCorporate field if non-nil, zero value otherwise.

### GetBinCorporateOk

`func (o *AuthReference) GetBinCorporateOk() (*bool, bool)`

GetBinCorporateOk returns a tuple with the BinCorporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCorporate

`func (o *AuthReference) SetBinCorporate(v bool)`

SetBinCorporate sets BinCorporate field to given value.

### HasBinCorporate

`func (o *AuthReference) HasBinCorporate() bool`

HasBinCorporate returns a boolean if a field has been set.

### GetBinCredit

`func (o *AuthReference) GetBinCredit() bool`

GetBinCredit returns the BinCredit field if non-nil, zero value otherwise.

### GetBinCreditOk

`func (o *AuthReference) GetBinCreditOk() (*bool, bool)`

GetBinCreditOk returns a tuple with the BinCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCredit

`func (o *AuthReference) SetBinCredit(v bool)`

SetBinCredit sets BinCredit field to given value.

### HasBinCredit

`func (o *AuthReference) HasBinCredit() bool`

HasBinCredit returns a boolean if a field has been set.

### GetBinDebit

`func (o *AuthReference) GetBinDebit() bool`

GetBinDebit returns the BinDebit field if non-nil, zero value otherwise.

### GetBinDebitOk

`func (o *AuthReference) GetBinDebitOk() (*bool, bool)`

GetBinDebitOk returns a tuple with the BinDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinDebit

`func (o *AuthReference) SetBinDebit(v bool)`

SetBinDebit sets BinDebit field to given value.

### HasBinDebit

`func (o *AuthReference) HasBinDebit() bool`

HasBinDebit returns a boolean if a field has been set.

### GetCardholderAgreement

`func (o *AuthReference) GetCardholderAgreement() string`

GetCardholderAgreement returns the CardholderAgreement field if non-nil, zero value otherwise.

### GetCardholderAgreementOk

`func (o *AuthReference) GetCardholderAgreementOk() (*string, bool)`

GetCardholderAgreementOk returns a tuple with the CardholderAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderAgreement

`func (o *AuthReference) SetCardholderAgreement(v string)`

SetCardholderAgreement sets CardholderAgreement field to given value.

### HasCardholderAgreement

`func (o *AuthReference) HasCardholderAgreement() bool`

HasCardholderAgreement returns a boolean if a field has been set.

### GetCurrency

`func (o *AuthReference) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AuthReference) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AuthReference) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AuthReference) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDatetime

`func (o *AuthReference) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AuthReference) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *AuthReference) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *AuthReference) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetEci

`func (o *AuthReference) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *AuthReference) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *AuthReference) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *AuthReference) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetEmail

`func (o *AuthReference) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthReference) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthReference) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthReference) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnv

`func (o *AuthReference) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *AuthReference) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *AuthReference) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *AuthReference) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetIdentifier

`func (o *AuthReference) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AuthReference) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AuthReference) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AuthReference) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetInitiation

`func (o *AuthReference) GetInitiation() string`

GetInitiation returns the Initiation field if non-nil, zero value otherwise.

### GetInitiationOk

`func (o *AuthReference) GetInitiationOk() (*string, bool)`

GetInitiationOk returns a tuple with the Initiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiation

`func (o *AuthReference) SetInitiation(v string)`

SetInitiation sets Initiation field to given value.

### HasInitiation

`func (o *AuthReference) HasInitiation() bool`

HasInitiation returns a boolean if a field has been set.

### GetInstrument

`func (o *AuthReference) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *AuthReference) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *AuthReference) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *AuthReference) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetMaskedpan

`func (o *AuthReference) GetMaskedpan() string`

GetMaskedpan returns the Maskedpan field if non-nil, zero value otherwise.

### GetMaskedpanOk

`func (o *AuthReference) GetMaskedpanOk() (*string, bool)`

GetMaskedpanOk returns a tuple with the Maskedpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedpan

`func (o *AuthReference) SetMaskedpan(v string)`

SetMaskedpan sets Maskedpan field to given value.

### HasMaskedpan

`func (o *AuthReference) HasMaskedpan() bool`

HasMaskedpan returns a boolean if a field has been set.

### GetMerchantid

`func (o *AuthReference) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *AuthReference) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *AuthReference) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.

### HasMerchantid

`func (o *AuthReference) HasMerchantid() bool`

HasMerchantid returns a boolean if a field has been set.

### GetMeta

`func (o *AuthReference) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuthReference) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuthReference) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuthReference) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNameOnCard

`func (o *AuthReference) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *AuthReference) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *AuthReference) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *AuthReference) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.

### GetPostcode

`func (o *AuthReference) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *AuthReference) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *AuthReference) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *AuthReference) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetResult

`func (o *AuthReference) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AuthReference) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AuthReference) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AuthReference) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetResultId

`func (o *AuthReference) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *AuthReference) GetResultIdOk() (*string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *AuthReference) SetResultId(v string)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *AuthReference) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### GetScheme

`func (o *AuthReference) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *AuthReference) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *AuthReference) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *AuthReference) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetSchemeLogo

`func (o *AuthReference) GetSchemeLogo() string`

GetSchemeLogo returns the SchemeLogo field if non-nil, zero value otherwise.

### GetSchemeLogoOk

`func (o *AuthReference) GetSchemeLogoOk() (*string, bool)`

GetSchemeLogoOk returns a tuple with the SchemeLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeLogo

`func (o *AuthReference) SetSchemeLogo(v string)`

SetSchemeLogo sets SchemeLogo field to given value.

### HasSchemeLogo

`func (o *AuthReference) HasSchemeLogo() bool`

HasSchemeLogo returns a boolean if a field has been set.

### GetTransStatus

`func (o *AuthReference) GetTransStatus() string`

GetTransStatus returns the TransStatus field if non-nil, zero value otherwise.

### GetTransStatusOk

`func (o *AuthReference) GetTransStatusOk() (*string, bool)`

GetTransStatusOk returns a tuple with the TransStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransStatus

`func (o *AuthReference) SetTransStatus(v string)`

SetTransStatus sets TransStatus field to given value.

### HasTransStatus

`func (o *AuthReference) HasTransStatus() bool`

HasTransStatus returns a boolean if a field has been set.

### GetTransType

`func (o *AuthReference) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *AuthReference) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *AuthReference) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *AuthReference) HasTransType() bool`

HasTransType returns a boolean if a field has been set.

### GetTransno

`func (o *AuthReference) GetTransno() int32`

GetTransno returns the Transno field if non-nil, zero value otherwise.

### GetTransnoOk

`func (o *AuthReference) GetTransnoOk() (*int32, bool)`

GetTransnoOk returns a tuple with the Transno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransno

`func (o *AuthReference) SetTransno(v int32)`

SetTransno sets Transno field to given value.

### HasTransno

`func (o *AuthReference) HasTransno() bool`

HasTransno returns a boolean if a field has been set.

### GetType

`func (o *AuthReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthReference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthReference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUtc

`func (o *AuthReference) GetUtc() int64`

GetUtc returns the Utc field if non-nil, zero value otherwise.

### GetUtcOk

`func (o *AuthReference) GetUtcOk() (*int64, bool)`

GetUtcOk returns a tuple with the Utc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtc

`func (o *AuthReference) SetUtc(v int64)`

SetUtc sets Utc field to given value.

### HasUtc

`func (o *AuthReference) HasUtc() bool`

HasUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


