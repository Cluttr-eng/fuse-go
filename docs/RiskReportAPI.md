# \RiskReportAPI

All URIs are relative to *https://sandbox-api.letsfuse.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConsumerRiskReport**](RiskReportAPI.md#DeleteConsumerRiskReport) | **Delete** /v1/risk_report/consumer/{consumer_risk_report_id} | Delete consumer risk report
[**GetConsumerRiskReportCustomization**](RiskReportAPI.md#GetConsumerRiskReportCustomization) | **Get** /v1/risk_report/consumer/customization/{consumer_risk_report_customization_id} | Get consumer risk report customization



## DeleteConsumerRiskReport

> DeleteConsumerRiskReportResponse DeleteConsumerRiskReport(ctx, consumerRiskReportId).FuseClientId(fuseClientId).FuseApiKey(fuseApiKey).Execute()

Delete consumer risk report

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
    fuseClientId := "fuseClientId_example" // string | 
    fuseApiKey := "fuseApiKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskReportAPI.DeleteConsumerRiskReport(context.Background(), consumerRiskReportId).FuseClientId(fuseClientId).FuseApiKey(fuseApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskReportAPI.DeleteConsumerRiskReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConsumerRiskReport`: DeleteConsumerRiskReportResponse
    fmt.Fprintf(os.Stdout, "Response from `RiskReportAPI.DeleteConsumerRiskReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerRiskReportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsumerRiskReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fuseClientId** | **string** |  | 
 **fuseApiKey** | **string** |  | 

### Return type

[**DeleteConsumerRiskReportResponse**](DeleteConsumerRiskReportResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsumerRiskReportCustomization

> GetConsumerRiskReportCustomizationResponse GetConsumerRiskReportCustomization(ctx, consumerRiskReportCustomizationId).FuseClientId(fuseClientId).FuseApiKey(fuseApiKey).Execute()

Get consumer risk report customization

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
    fuseClientId := "fuseClientId_example" // string | 
    fuseApiKey := "fuseApiKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskReportAPI.GetConsumerRiskReportCustomization(context.Background(), consumerRiskReportCustomizationId).FuseClientId(fuseClientId).FuseApiKey(fuseApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskReportAPI.GetConsumerRiskReportCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerRiskReportCustomization`: GetConsumerRiskReportCustomizationResponse
    fmt.Fprintf(os.Stdout, "Response from `RiskReportAPI.GetConsumerRiskReportCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerRiskReportCustomizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerRiskReportCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fuseClientId** | **string** |  | 
 **fuseApiKey** | **string** |  | 

### Return type

[**GetConsumerRiskReportCustomizationResponse**](GetConsumerRiskReportCustomizationResponse.md)

### Authorization

[fuseApiKey](../README.md#fuseApiKey), [fuseClientId](../README.md#fuseClientId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

