# \LocalizationsApi

All URIs are relative to *https://api.dyspatch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalizationsLocalizationIdGet**](LocalizationsApi.md#LocalizationsLocalizationIdGet) | **Get** /localizations/{localizationId} | Get Localization Object by ID


# **LocalizationsLocalizationIdGet**
> LocalizationRead LocalizationsLocalizationIdGet(ctx, localizationId, accept)
Get Localization Object by ID

Returns a specific localization object with a matching ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **localizationId** | **string**| A localization ID | 
  **accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2018.08\&quot;, set the value to \&quot;application/vnd.dyspatch.2018.08+json\&quot; | 

### Return type

[**LocalizationRead**](LocalizationRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.dyspatch.2018.08+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

