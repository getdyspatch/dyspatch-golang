# \LocalizationsApi

All URIs are relative to *https://api.dyspatch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalizationsLocalizationIdGet**](LocalizationsApi.md#LocalizationsLocalizationIdGet) | **Get** /localizations/{localizationId} | Get Localization Object by ID


# **LocalizationsLocalizationIdGet**
> LocalizationRead LocalizationsLocalizationIdGet(ctx, localizationId, optional)
Get Localization Object by ID

Returns a specific localization object with a matching ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **localizationId** | **string**| A localization ID | 
 **optional** | ***LocalizationsLocalizationIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationsLocalizationIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetLanguage** | **optional.String**| The type of templating language to compile as. Should only be used for visual templates. | 

### Return type

[**LocalizationRead**](LocalizationRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.dyspatch.2019.03+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

