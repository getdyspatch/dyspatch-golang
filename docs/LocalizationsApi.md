# \LocalizationsApi

All URIs are relative to *https://api.dyspatch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDraftLocalizationById**](LocalizationsApi.md#GetDraftLocalizationById) | **Get** /localizations/{localizationId}/drafts/{draftId} | Get Draft Localization Object by ID
[**GetPublishedLocalizationById**](LocalizationsApi.md#GetPublishedLocalizationById) | **Get** /localizations/{localizationId} | Get Localization Object by ID



## GetDraftLocalizationById

> LocalizationRead GetDraftLocalizationById(ctx, draftId, localizationId, targetLanguage, accept)

Get Draft Localization Object by ID

Returns a specific localization object of the matching draft with a matching localization ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**localizationId** | **string**| A localization ID | 
**targetLanguage** | **string**| The type of templating language to compile as. Should only be used for visual templates. | 
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2020.11\&quot;, set the value to \&quot;application/vnd.dyspatch.2020.11+json\&quot; | 

### Return type

[**LocalizationRead**](LocalizationRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2020.11+json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublishedLocalizationById

> LocalizationRead GetPublishedLocalizationById(ctx, localizationId, targetLanguage, accept)

Get Localization Object by ID

Returns the published content associated with the localization of the matching ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localizationId** | **string**| A localization ID | 
**targetLanguage** | **string**| The type of templating language to compile as. Should only be used for visual templates. | 
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2020.11\&quot;, set the value to \&quot;application/vnd.dyspatch.2020.11+json\&quot; | 

### Return type

[**LocalizationRead**](LocalizationRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2020.11+json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

