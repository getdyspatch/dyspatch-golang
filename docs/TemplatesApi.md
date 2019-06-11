# \TemplatesApi

All URIs are relative to *https://api.dyspatch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplatesGet**](TemplatesApi.md#TemplatesGet) | **Get** /templates | List Templates
[**TemplatesTemplateIdGet**](TemplatesApi.md#TemplatesTemplateIdGet) | **Get** /templates/{templateId} | Get Template by ID


# **TemplatesGet**
> TemplatesRead TemplatesGet(ctx, optional)
List Templates

Gets a list of Template Metadata objects for all templates. Up to 25 results returned before results are paginated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TemplatesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplatesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| A cursor value used to retrieve a specific page from a paginated result set. | 

### Return type

[**TemplatesRead**](TemplatesRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.dyspatch.2019.03+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplatesTemplateIdGet**
> TemplateRead TemplatesTemplateIdGet(ctx, templateId, optional)
Get Template by ID

Gets a template object with the matching ID. If the template has published content the \"compiled\" field will contain the template .

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**| A template ID | 
 **optional** | ***TemplatesTemplateIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplatesTemplateIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetLanguage** | **optional.String**| The type of templating language to compile as. Should only be used for visual templates. | 

### Return type

[**TemplateRead**](TemplateRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.dyspatch.2019.03+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

