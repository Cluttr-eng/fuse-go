# \FuseAPI

All URIs are relative to *https://sandbox-api.letsfuse.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccountEvents**](FuseAPI.md#AddAccountEvents) | **Post** /v1/accounts/{account_id}/events | 
[**CreateAssetReport**](FuseAPI.md#CreateAssetReport) | **Post** /v1/financial_connections/asset_report/create | 
[**CreateConsumerRiskReport**](FuseAPI.md#CreateConsumerRiskReport) | **Post** /v1/risk_report/consumer | 
[**CreateConsumerRiskReportCustomization**](FuseAPI.md#CreateConsumerRiskReportCustomization) | **Post** /v1/risk_report/consumer/customization | 
[**CreateLinkToken**](FuseAPI.md#CreateLinkToken) | **Post** /v1/link/token | 
[**CreateSession**](FuseAPI.md#CreateSession) | **Post** /v1/session | 
[**DeleteFinancialConnection**](FuseAPI.md#DeleteFinancialConnection) | **Delete** /v1/financial_connections/{financial_connection_id_to_delete} | Delete a financial connection
[**EnrichTransactions**](FuseAPI.md#EnrichTransactions) | **Post** /v1/transactions/enrich | 
[**ExchangeFinancialConnectionsPublicToken**](FuseAPI.md#ExchangeFinancialConnectionsPublicToken) | **Post** /v1/financial_connections/public_token/exchange | 
[**GetAssetReport**](FuseAPI.md#GetAssetReport) | **Post** /v1/financial_connections/asset_report | 
[**GetConsumerRiskReport**](FuseAPI.md#GetConsumerRiskReport) | **Get** /v1/risk_report/consumer/{consumer_risk_report_id} | Get consumer risk report
[**GetEntity**](FuseAPI.md#GetEntity) | **Get** /v1/entities/{entity_id} | Get entity
[**GetFinanceScore**](FuseAPI.md#GetFinanceScore) | **Get** /v1/accounts/{account_id}/finance_score | Get finance score
[**GetFinancialConnection**](FuseAPI.md#GetFinancialConnection) | **Get** /v1/financial_connections/{financial_connection_id} | Get financial connection details
[**GetFinancialConnectionsAccountDetails**](FuseAPI.md#GetFinancialConnectionsAccountDetails) | **Post** /v1/financial_connections/accounts/details | Get account details
[**GetFinancialConnectionsAccountStatement**](FuseAPI.md#GetFinancialConnectionsAccountStatement) | **Post** /v1/financial_connections/accounts/statement | 
[**GetFinancialConnectionsAccounts**](FuseAPI.md#GetFinancialConnectionsAccounts) | **Post** /v1/financial_connections/accounts | Get accounts
[**GetFinancialConnectionsBalances**](FuseAPI.md#GetFinancialConnectionsBalances) | **Post** /v1/financial_connections/balances | Get balances
[**GetFinancialConnectionsOwners**](FuseAPI.md#GetFinancialConnectionsOwners) | **Post** /v1/financial_connections/owners | Get account owners
[**GetFinancialConnectionsTransactions**](FuseAPI.md#GetFinancialConnectionsTransactions) | **Post** /v1/financial_connections/transactions | Get transactions
[**GetFinancialInstitution**](FuseAPI.md#GetFinancialInstitution) | **Get** /v1/financial_connections/institutions/{institution_id} | Get a financial institution
[**GetInvestmentHoldings**](FuseAPI.md#GetInvestmentHoldings) | **Post** /v1/financial_connections/investments/holdings | Get investment holdings
[**GetInvestmentTransactions**](FuseAPI.md#GetInvestmentTransactions) | **Post** /v1/financial_connections/investments/transactions | Get investment transactions
[**MigrateFinancialConnection**](FuseAPI.md#MigrateFinancialConnection) | **Post** /v1/financial_connections/migrate | Migrate financial connection
[**RefreshAssetReport**](FuseAPI.md#RefreshAssetReport) | **Post** /v1/financial_connections/asset_report/refresh | 
[**SyncFinancialConnectionsData**](FuseAPI.md#SyncFinancialConnectionsData) | **Post** /v1/financial_connections/sync | Sync financial connections data
[**UpdateConsumerRiskReportCustomization**](FuseAPI.md#UpdateConsumerRiskReportCustomization) | **Post** /v1/risk_report/consumer/customization/{consumer_risk_report_customization_id} | Update consumer risk report customization
[**V1FinancialConnectionsLiabilitiesPost**](FuseAPI.md#V1FinancialConnectionsLiabilitiesPost) | **Post** /v1/financial_connections/liabilities | Get liabilities



## AddAccountEvents

> AddAccountEventsResponse AddAccountEvents(ctx, accountId).AddAccountEventsRequest(addAccountEventsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountId := "accountId_example" // string | 
    addAccountEventsRequest := *openapiclient.NewAddAccountEventsRequest([]openapiclient.AddAccountEventsRequestEventsInner{openapiclient.AddAccountEventsRequest_events_inner{ExternalTransactionEvent: openapiclient.NewExternalTransactionEvent("Id_example", "EventType_example", openapiclient.ExternalTransactionEventStatus("pending"), float32(123), "IsoCurrencyCode_example", "MerchantName_example", "Timestamp_example")}}) // AddAccountEventsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.AddAccountEvents(context.Background(), accountId).AddAccountEventsRequest(addAccountEventsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.AddAccountEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccountEvents`: AddAccountEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.AddAccountEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAccountEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addAccountEventsRequest** | [**AddAccountEventsRequest**](AddAccountEventsRequest.md) |  | 

### Return type

[**AddAccountEventsResponse**](AddAccountEventsResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetReport

> CreateAssetReportResponse CreateAssetReport(ctx).CreateAssetReportRequest(createAssetReportRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createAssetReportRequest := *openapiclient.NewCreateAssetReportRequest("AccessToken_example", float32(123)) // CreateAssetReportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.CreateAssetReport(context.Background()).CreateAssetReportRequest(createAssetReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.CreateAssetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssetReport`: CreateAssetReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.CreateAssetReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAssetReportRequest** | [**CreateAssetReportRequest**](CreateAssetReportRequest.md) |  | 

### Return type

[**CreateAssetReportResponse**](CreateAssetReportResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConsumerRiskReport

> CreateConsumerRiskReportResponse CreateConsumerRiskReport(ctx).CreateConsumerRiskReportRequest(createConsumerRiskReportRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createConsumerRiskReportRequest := *openapiclient.NewCreateConsumerRiskReportRequest("AccountId_example", "IsoCurrencyCode_example", "CustomizationId_example") // CreateConsumerRiskReportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.CreateConsumerRiskReport(context.Background()).CreateConsumerRiskReportRequest(createConsumerRiskReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.CreateConsumerRiskReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConsumerRiskReport`: CreateConsumerRiskReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.CreateConsumerRiskReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConsumerRiskReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConsumerRiskReportRequest** | [**CreateConsumerRiskReportRequest**](CreateConsumerRiskReportRequest.md) |  | 

### Return type

[**CreateConsumerRiskReportResponse**](CreateConsumerRiskReportResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConsumerRiskReportCustomization

> CreateConsumerRiskReportCustomizationResponse CreateConsumerRiskReportCustomization(ctx).CreateConsumerRiskReportCustomizationRequest(createConsumerRiskReportCustomizationRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createConsumerRiskReportCustomizationRequest := *openapiclient.NewCreateConsumerRiskReportCustomizationRequest(openapiclient.ConsumerRiskReportTimeFrame("daily"), float32(123), float32(123), float32(123)) // CreateConsumerRiskReportCustomizationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.CreateConsumerRiskReportCustomization(context.Background()).CreateConsumerRiskReportCustomizationRequest(createConsumerRiskReportCustomizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.CreateConsumerRiskReportCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConsumerRiskReportCustomization`: CreateConsumerRiskReportCustomizationResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.CreateConsumerRiskReportCustomization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConsumerRiskReportCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConsumerRiskReportCustomizationRequest** | [**CreateConsumerRiskReportCustomizationRequest**](CreateConsumerRiskReportCustomizationRequest.md) |  | 

### Return type

[**CreateConsumerRiskReportCustomizationResponse**](CreateConsumerRiskReportCustomizationResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLinkToken

> CreateLinkTokenResponse CreateLinkToken(ctx).CreateLinkTokenRequest(createLinkTokenRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createLinkTokenRequest := *openapiclient.NewCreateLinkTokenRequest("InstitutionId_example", *openapiclient.NewEntity("Id_example"), "ClientName_example", "SessionClientSecret_example") // CreateLinkTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.CreateLinkToken(context.Background()).CreateLinkTokenRequest(createLinkTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.CreateLinkToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLinkToken`: CreateLinkTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.CreateLinkToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLinkTokenRequest** | [**CreateLinkTokenRequest**](CreateLinkTokenRequest.md) |  | 

### Return type

[**CreateLinkTokenResponse**](CreateLinkTokenResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> CreateSessionResponse CreateSession(ctx).CreateSessionRequest(createSessionRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createSessionRequest := *openapiclient.NewCreateSessionRequest([]openapiclient.Aggregator{openapiclient.Aggregator("akoya")}, []openapiclient.Product{openapiclient.Product("account_details")}, *openapiclient.NewEntity("Id_example")) // CreateSessionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.CreateSession(context.Background()).CreateSessionRequest(createSessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.CreateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSession`: CreateSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.CreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSessionRequest** | [**CreateSessionRequest**](CreateSessionRequest.md) |  | 

### Return type

[**CreateSessionResponse**](CreateSessionResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFinancialConnection

> DeleteFinancialConnectionResponse DeleteFinancialConnection(ctx, financialConnectionIdToDelete).Execute()

Delete a financial connection

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    financialConnectionIdToDelete := "financialConnectionIdToDelete_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.DeleteFinancialConnection(context.Background(), financialConnectionIdToDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.DeleteFinancialConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFinancialConnection`: DeleteFinancialConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.DeleteFinancialConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**financialConnectionIdToDelete** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinancialConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteFinancialConnectionResponse**](DeleteFinancialConnectionResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrichTransactions

> EnrichTransactionsResponse EnrichTransactions(ctx).FuseClientId(fuseClientId).FuseApiKey(fuseApiKey).EnrichTransactionsRequest(enrichTransactionsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    fuseClientId := "fuseClientId_example" // string | 
    fuseApiKey := "fuseApiKey_example" // string | 
    enrichTransactionsRequest := *openapiclient.NewEnrichTransactionsRequest([]openapiclient.TransactionToEnrich{*openapiclient.NewTransactionToEnrich("Id_example", "Description_example", float32(123), "Direction_example")}) // EnrichTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.EnrichTransactions(context.Background()).FuseClientId(fuseClientId).FuseApiKey(fuseApiKey).EnrichTransactionsRequest(enrichTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.EnrichTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrichTransactions`: EnrichTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.EnrichTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrichTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fuseClientId** | **string** |  | 
 **fuseApiKey** | **string** |  | 
 **enrichTransactionsRequest** | [**EnrichTransactionsRequest**](EnrichTransactionsRequest.md) |  | 

### Return type

[**EnrichTransactionsResponse**](EnrichTransactionsResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExchangeFinancialConnectionsPublicToken

> ExchangeFinancialConnectionsPublicTokenResponse ExchangeFinancialConnectionsPublicToken(ctx).ExchangeFinancialConnectionsPublicTokenRequest(exchangeFinancialConnectionsPublicTokenRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    exchangeFinancialConnectionsPublicTokenRequest := *openapiclient.NewExchangeFinancialConnectionsPublicTokenRequest("PublicToken_example") // ExchangeFinancialConnectionsPublicTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.ExchangeFinancialConnectionsPublicToken(context.Background()).ExchangeFinancialConnectionsPublicTokenRequest(exchangeFinancialConnectionsPublicTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.ExchangeFinancialConnectionsPublicToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExchangeFinancialConnectionsPublicToken`: ExchangeFinancialConnectionsPublicTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.ExchangeFinancialConnectionsPublicToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangeFinancialConnectionsPublicTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchangeFinancialConnectionsPublicTokenRequest** | [**ExchangeFinancialConnectionsPublicTokenRequest**](ExchangeFinancialConnectionsPublicTokenRequest.md) |  | 

### Return type

[**ExchangeFinancialConnectionsPublicTokenResponse**](ExchangeFinancialConnectionsPublicTokenResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetReport

> RefreshAssetReportResponse GetAssetReport(ctx).GetAssetReportRequest(getAssetReportRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getAssetReportRequest := *openapiclient.NewGetAssetReportRequest("AccessToken_example", "AssetReportToken_example") // GetAssetReportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetAssetReport(context.Background()).GetAssetReportRequest(getAssetReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetAssetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetReport`: RefreshAssetReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetAssetReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getAssetReportRequest** | [**GetAssetReportRequest**](GetAssetReportRequest.md) |  | 

### Return type

[**RefreshAssetReportResponse**](RefreshAssetReportResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsumerRiskReport

> GetConsumerRiskReportResponse GetConsumerRiskReport(ctx, consumerRiskReportId).Recalculate(recalculate).Execute()

Get consumer risk report

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    consumerRiskReportId := "consumerRiskReportId_example" // string | 
    recalculate := true // bool | An optional boolean parameter. If set to true, the system will recalculate before returning the risk report. If omitted or set to false, the current risk report will be returned without recalculation. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetConsumerRiskReport(context.Background(), consumerRiskReportId).Recalculate(recalculate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetConsumerRiskReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerRiskReport`: GetConsumerRiskReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetConsumerRiskReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerRiskReportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerRiskReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recalculate** | **bool** | An optional boolean parameter. If set to true, the system will recalculate before returning the risk report. If omitted or set to false, the current risk report will be returned without recalculation. | 

### Return type

[**GetConsumerRiskReportResponse**](GetConsumerRiskReportResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntity

> GetEntityResponse GetEntity(ctx, entityId).Execute()

Get entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityId := "entityId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetEntity(context.Background(), entityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntity`: GetEntityResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEntityResponse**](GetEntityResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceScore

> GetFinanceScoreResponse GetFinanceScore(ctx, accountId).Execute()

Get finance score

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetFinanceScore(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetFinanceScore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinanceScore`: GetFinanceScoreResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetFinanceScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFinanceScoreResponse**](GetFinanceScoreResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialConnection

> GetFinancialConnectionResponse GetFinancialConnection(ctx, financialConnectionId).Execute()

Get financial connection details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    financialConnectionId := "financialConnectionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetFinancialConnection(context.Background(), financialConnectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetFinancialConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinancialConnection`: GetFinancialConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetFinancialConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**financialConnectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFinancialConnectionResponse**](GetFinancialConnectionResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialConnectionsAccountDetails

> GetFinancialConnectionsAccountDetailsResponse GetFinancialConnectionsAccountDetails(ctx).GetFinancialConnectionsAccountDetailsRequest(getFinancialConnectionsAccountDetailsRequest).Execute()

Get account details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getFinancialConnectionsAccountDetailsRequest := *openapiclient.NewGetFinancialConnectionsAccountDetailsRequest("AccessToken_example") // GetFinancialConnectionsAccountDetailsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetFinancialConnectionsAccountDetails(context.Background()).GetFinancialConnectionsAccountDetailsRequest(getFinancialConnectionsAccountDetailsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetFinancialConnectionsAccountDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinancialConnectionsAccountDetails`: GetFinancialConnectionsAccountDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetFinancialConnectionsAccountDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialConnectionsAccountDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getFinancialConnectionsAccountDetailsRequest** | [**GetFinancialConnectionsAccountDetailsRequest**](GetFinancialConnectionsAccountDetailsRequest.md) |  | 

### Return type

[**GetFinancialConnectionsAccountDetailsResponse**](GetFinancialConnectionsAccountDetailsResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialConnectionsAccountStatement

> GetFinancialConnectionsAccountStatementResponse GetFinancialConnectionsAccountStatement(ctx).GetFinancialConnectionsAccountStatementRequest(getFinancialConnectionsAccountStatementRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getFinancialConnectionsAccountStatementRequest := *openapiclient.NewGetFinancialConnectionsAccountStatementRequest("AccessToken_example", "RemoteAccountId_example") // GetFinancialConnectionsAccountStatementRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetFinancialConnectionsAccountStatement(context.Background()).GetFinancialConnectionsAccountStatementRequest(getFinancialConnectionsAccountStatementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetFinancialConnectionsAccountStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinancialConnectionsAccountStatement`: GetFinancialConnectionsAccountStatementResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetFinancialConnectionsAccountStatement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialConnectionsAccountStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getFinancialConnectionsAccountStatementRequest** | [**GetFinancialConnectionsAccountStatementRequest**](GetFinancialConnectionsAccountStatementRequest.md) |  | 

### Return type

[**GetFinancialConnectionsAccountStatementResponse**](GetFinancialConnectionsAccountStatementResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialConnectionsAccounts

> GetFinancialConnectionsAccountsResponse GetFinancialConnectionsAccounts(ctx).GetFinancialConnectionsAccountsRequest(getFinancialConnectionsAccountsRequest).Execute()

Get accounts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getFinancialConnectionsAccountsRequest := *openapiclient.NewGetFinancialConnectionsAccountsRequest("AccessToken_example") // GetFinancialConnectionsAccountsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetFinancialConnectionsAccounts(context.Background()).GetFinancialConnectionsAccountsRequest(getFinancialConnectionsAccountsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetFinancialConnectionsAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinancialConnectionsAccounts`: GetFinancialConnectionsAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetFinancialConnectionsAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialConnectionsAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getFinancialConnectionsAccountsRequest** | [**GetFinancialConnectionsAccountsRequest**](GetFinancialConnectionsAccountsRequest.md) |  | 

### Return type

[**GetFinancialConnectionsAccountsResponse**](GetFinancialConnectionsAccountsResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialConnectionsBalances

> GetFinancialConnectionsBalanceResponse GetFinancialConnectionsBalances(ctx).GetFinancialConnectionsBalanceRequest(getFinancialConnectionsBalanceRequest).Execute()

Get balances

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getFinancialConnectionsBalanceRequest := *openapiclient.NewGetFinancialConnectionsBalanceRequest("AccessToken_example") // GetFinancialConnectionsBalanceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetFinancialConnectionsBalances(context.Background()).GetFinancialConnectionsBalanceRequest(getFinancialConnectionsBalanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetFinancialConnectionsBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinancialConnectionsBalances`: GetFinancialConnectionsBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetFinancialConnectionsBalances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialConnectionsBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getFinancialConnectionsBalanceRequest** | [**GetFinancialConnectionsBalanceRequest**](GetFinancialConnectionsBalanceRequest.md) |  | 

### Return type

[**GetFinancialConnectionsBalanceResponse**](GetFinancialConnectionsBalanceResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialConnectionsOwners

> GetFinancialConnectionsOwnersResponse GetFinancialConnectionsOwners(ctx).GetFinancialConnectionsOwnersRequest(getFinancialConnectionsOwnersRequest).Execute()

Get account owners

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getFinancialConnectionsOwnersRequest := *openapiclient.NewGetFinancialConnectionsOwnersRequest("AccessToken_example") // GetFinancialConnectionsOwnersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetFinancialConnectionsOwners(context.Background()).GetFinancialConnectionsOwnersRequest(getFinancialConnectionsOwnersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetFinancialConnectionsOwners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinancialConnectionsOwners`: GetFinancialConnectionsOwnersResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetFinancialConnectionsOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialConnectionsOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getFinancialConnectionsOwnersRequest** | [**GetFinancialConnectionsOwnersRequest**](GetFinancialConnectionsOwnersRequest.md) |  | 

### Return type

[**GetFinancialConnectionsOwnersResponse**](GetFinancialConnectionsOwnersResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialConnectionsTransactions

> GetFinancialConnectionsTransactionsResponse GetFinancialConnectionsTransactions(ctx).GetFinancialConnectionsTransactionsRequest(getFinancialConnectionsTransactionsRequest).Execute()

Get transactions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getFinancialConnectionsTransactionsRequest := *openapiclient.NewGetFinancialConnectionsTransactionsRequest("AccessToken_example", "StartDate_example", "EndDate_example", int32(123), int32(123)) // GetFinancialConnectionsTransactionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetFinancialConnectionsTransactions(context.Background()).GetFinancialConnectionsTransactionsRequest(getFinancialConnectionsTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetFinancialConnectionsTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinancialConnectionsTransactions`: GetFinancialConnectionsTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetFinancialConnectionsTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialConnectionsTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getFinancialConnectionsTransactionsRequest** | [**GetFinancialConnectionsTransactionsRequest**](GetFinancialConnectionsTransactionsRequest.md) |  | 

### Return type

[**GetFinancialConnectionsTransactionsResponse**](GetFinancialConnectionsTransactionsResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialInstitution

> GetFinancialInstitutionResponse GetFinancialInstitution(ctx, institutionId).Execute()

Get a financial institution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    institutionId := "institutionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetFinancialInstitution(context.Background(), institutionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetFinancialInstitution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinancialInstitution`: GetFinancialInstitutionResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetFinancialInstitution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**institutionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialInstitutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFinancialInstitutionResponse**](GetFinancialInstitutionResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvestmentHoldings

> GetInvestmentHoldingsResponse GetInvestmentHoldings(ctx).GetInvestmentHoldingsRequest(getInvestmentHoldingsRequest).Execute()

Get investment holdings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getInvestmentHoldingsRequest := *openapiclient.NewGetInvestmentHoldingsRequest("AccessToken_example") // GetInvestmentHoldingsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetInvestmentHoldings(context.Background()).GetInvestmentHoldingsRequest(getInvestmentHoldingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetInvestmentHoldings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvestmentHoldings`: GetInvestmentHoldingsResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetInvestmentHoldings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvestmentHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getInvestmentHoldingsRequest** | [**GetInvestmentHoldingsRequest**](GetInvestmentHoldingsRequest.md) |  | 

### Return type

[**GetInvestmentHoldingsResponse**](GetInvestmentHoldingsResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvestmentTransactions

> GetInvestmentTransactionsResponse GetInvestmentTransactions(ctx).GetInvestmentTransactionsRequest(getInvestmentTransactionsRequest).Execute()

Get investment transactions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getInvestmentTransactionsRequest := *openapiclient.NewGetInvestmentTransactionsRequest("AccessToken_example", "StartDate_example", "EndDate_example", int32(123), int32(123)) // GetInvestmentTransactionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.GetInvestmentTransactions(context.Background()).GetInvestmentTransactionsRequest(getInvestmentTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.GetInvestmentTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvestmentTransactions`: GetInvestmentTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.GetInvestmentTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvestmentTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getInvestmentTransactionsRequest** | [**GetInvestmentTransactionsRequest**](GetInvestmentTransactionsRequest.md) |  | 

### Return type

[**GetInvestmentTransactionsResponse**](GetInvestmentTransactionsResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateFinancialConnection

> MigrateFinancialConnectionsTokenResponse MigrateFinancialConnection(ctx).MigrateFinancialConnectionsTokenRequest(migrateFinancialConnectionsTokenRequest).Execute()

Migrate financial connection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    migrateFinancialConnectionsTokenRequest := *openapiclient.NewMigrateFinancialConnectionsTokenRequest(*openapiclient.NewMigrateFinancialConnectionsAggregatorConnectionData(), "Aggregator_example", *openapiclient.NewMigrateFinancialConnectionsTokenRequestEntity(), []openapiclient.Product{openapiclient.Product("account_details")}) // MigrateFinancialConnectionsTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.MigrateFinancialConnection(context.Background()).MigrateFinancialConnectionsTokenRequest(migrateFinancialConnectionsTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.MigrateFinancialConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrateFinancialConnection`: MigrateFinancialConnectionsTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.MigrateFinancialConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrateFinancialConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrateFinancialConnectionsTokenRequest** | [**MigrateFinancialConnectionsTokenRequest**](MigrateFinancialConnectionsTokenRequest.md) |  | 

### Return type

[**MigrateFinancialConnectionsTokenResponse**](MigrateFinancialConnectionsTokenResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshAssetReport

> AssetReportResponse RefreshAssetReport(ctx).RefreshAssetReportRequest(refreshAssetReportRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    refreshAssetReportRequest := *openapiclient.NewRefreshAssetReportRequest("AccessToken_example", "AssetReportToken_example") // RefreshAssetReportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.RefreshAssetReport(context.Background()).RefreshAssetReportRequest(refreshAssetReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.RefreshAssetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshAssetReport`: AssetReportResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.RefreshAssetReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshAssetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshAssetReportRequest** | [**RefreshAssetReportRequest**](RefreshAssetReportRequest.md) |  | 

### Return type

[**AssetReportResponse**](AssetReportResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncFinancialConnectionsData

> SyncFinancialConnectionsDataResponse SyncFinancialConnectionsData(ctx).FuseVerification(fuseVerification).Body(body).Execute()

Sync financial connections data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    fuseVerification := "fuseVerification_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.SyncFinancialConnectionsData(context.Background()).FuseVerification(fuseVerification).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.SyncFinancialConnectionsData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncFinancialConnectionsData`: SyncFinancialConnectionsDataResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.SyncFinancialConnectionsData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncFinancialConnectionsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fuseVerification** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**SyncFinancialConnectionsDataResponse**](SyncFinancialConnectionsDataResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConsumerRiskReportCustomization

> UpdateConsumerRiskReportCustomizationResponse UpdateConsumerRiskReportCustomization(ctx, consumerRiskReportCustomizationId).UpdateConsumerRiskReportCustomizationRequest(updateConsumerRiskReportCustomizationRequest).Execute()

Update consumer risk report customization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    consumerRiskReportCustomizationId := "consumerRiskReportCustomizationId_example" // string | 
    updateConsumerRiskReportCustomizationRequest := *openapiclient.NewUpdateConsumerRiskReportCustomizationRequest() // UpdateConsumerRiskReportCustomizationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.UpdateConsumerRiskReportCustomization(context.Background(), consumerRiskReportCustomizationId).UpdateConsumerRiskReportCustomizationRequest(updateConsumerRiskReportCustomizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.UpdateConsumerRiskReportCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConsumerRiskReportCustomization`: UpdateConsumerRiskReportCustomizationResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.UpdateConsumerRiskReportCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerRiskReportCustomizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConsumerRiskReportCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConsumerRiskReportCustomizationRequest** | [**UpdateConsumerRiskReportCustomizationRequest**](UpdateConsumerRiskReportCustomizationRequest.md) |  | 

### Return type

[**UpdateConsumerRiskReportCustomizationResponse**](UpdateConsumerRiskReportCustomizationResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1FinancialConnectionsLiabilitiesPost

> GetLiabilitiesResponse V1FinancialConnectionsLiabilitiesPost(ctx).GetLiabilitiesRequest(getLiabilitiesRequest).Execute()

Get liabilities

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getLiabilitiesRequest := *openapiclient.NewGetLiabilitiesRequest("AccessToken_example") // GetLiabilitiesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FuseAPI.V1FinancialConnectionsLiabilitiesPost(context.Background()).GetLiabilitiesRequest(getLiabilitiesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FuseAPI.V1FinancialConnectionsLiabilitiesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1FinancialConnectionsLiabilitiesPost`: GetLiabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `FuseAPI.V1FinancialConnectionsLiabilitiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1FinancialConnectionsLiabilitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getLiabilitiesRequest** | [**GetLiabilitiesRequest**](GetLiabilitiesRequest.md) |  | 

### Return type

[**GetLiabilitiesResponse**](GetLiabilitiesResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

