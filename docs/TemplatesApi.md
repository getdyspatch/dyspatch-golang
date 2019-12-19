# \TemplatesApi

All URIs are relative to *https://api.dyspatch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTemplateById**](TemplatesApi.md#GetTemplateById) | **Get** /templates/{templateId} | Get Template by ID
[**GetTemplates**](TemplatesApi.md#GetTemplates) | **Get** /templates | List Templates



## GetTemplateById

> TemplateRead GetTemplateById(ctx, templateId, targetLanguage, accept)

Get Template by ID

Gets a template object with the matching ID. If the template has published content the \"compiled\" field will contain the template .

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string**| A template ID | 
**targetLanguage** | **string**| The type of templating language to compile as. Should only be used for visual templates. | 
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2019.10\&quot;, set the value to \&quot;application/vnd.dyspatch.2019.10+json\&quot; | 

### Return type

[**TemplateRead**](TemplateRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2019.10+json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplates

> TemplatesRead GetTemplates(ctx, accept, optional)

List Templates

Gets a list of Template Metadata objects for all templates. Up to 25 results returned before results are paginated.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2019.10\&quot;, set the value to \&quot;application/vnd.dyspatch.2019.10+json\&quot; | 
 **optional** | ***GetTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| A cursor value used to retrieve a specific page from a paginated result set. | 

### Return type

[**TemplatesRead**](TemplatesRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2019.10+json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

