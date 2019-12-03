# \LocalizationsApi

All URIs are relative to *https://api.dyspatch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalizationsLocalizationIdGet**](LocalizationsApi.md#LocalizationsLocalizationIdGet) | **Get** /localizations/{localizationId} | Get Localization Object by ID



## LocalizationsLocalizationIdGet

> LocalizationRead LocalizationsLocalizationIdGet(ctx, localizationId, targetLanguage)

Get Localization Object by ID

Returns a specific localization object with a matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localizationId** | **string**| A localization ID | 
**targetLanguage** | **string**| The type of templating language to compile as. Should only be used for visual templates. | 

### Return type

[**LocalizationRead**](LocalizationRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2019.10+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

