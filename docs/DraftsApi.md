# \DraftsApi

All URIs are relative to *https://api.dyspatch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLocalization**](DraftsApi.md#DeleteLocalization) | **Delete** /drafts/{draftId}/localizations/{languageId} | Remove a localization
[**GetDraftById**](DraftsApi.md#GetDraftById) | **Get** /drafts/{draftId} | Get Draft by ID
[**GetDraftLocalizationKeys**](DraftsApi.md#GetDraftLocalizationKeys) | **Get** /drafts/{draftId}/localizationKeys | Get localization keys
[**GetDrafts**](DraftsApi.md#GetDrafts) | **Get** /drafts | List Drafts
[**GetLocalizationForDraft**](DraftsApi.md#GetLocalizationForDraft) | **Get** /drafts/{draftId}/localizations | Get localizations on a draft
[**SaveLocalization**](DraftsApi.md#SaveLocalization) | **Put** /drafts/{draftId}/localizations/{languageId} | Create or update a localization
[**SetTranslation**](DraftsApi.md#SetTranslation) | **Put** /drafts/{draftId}/localizations/{languageId}/translations | Set translations for language
[**SubmitDraftForApproval**](DraftsApi.md#SubmitDraftForApproval) | **Post** /drafts/{draftId}/publishRequest | Submit the draft for approval



## DeleteLocalization

> DeleteLocalization(ctx, draftId, languageId, accept)

Remove a localization

Deletes the localization with the given language ID if it exists

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**languageId** | **string**| A language ID (eg: en-US) | 
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2020.08\&quot;, set the value to \&quot;application/vnd.dyspatch.2020.08+json\&quot; | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftById

> DraftRead GetDraftById(ctx, draftId, targetLanguage, accept)

Get Draft by ID

Gets a draft object with the matching ID. The \"compiled\" field will contain the template in the default, unlocalized form.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**targetLanguage** | **string**| The type of templating language to compile as. Should only be used for visual templates. | 
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2020.08\&quot;, set the value to \&quot;application/vnd.dyspatch.2020.08+json\&quot; | 

### Return type

[**DraftRead**](DraftRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2020.08+json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftLocalizationKeys

> []LocalizationKeyRead GetDraftLocalizationKeys(ctx, draftId, accept)

Get localization keys

Returns the list of values that need to be translated for the draft. Set the `Accept` header to `application/vnd.dyspatch.2020.08+json` to get a JSON object, or `text/vnd.dyspatch.2020.08+x-gettext-translation` to get the POT file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2020.08\&quot;, set the value to \&quot;application/vnd.dyspatch.2020.08+json\&quot; | 

### Return type

[**[]LocalizationKeyRead**](LocalizationKeyRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2020.08+json, text/vnd.dyspatch.2020.08+x-gettext-translation

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDrafts

> DraftsRead GetDrafts(ctx, accept, optional)

List Drafts

Returns all drafts for your organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2020.08\&quot;, set the value to \&quot;application/vnd.dyspatch.2020.08+json\&quot; | 
 **optional** | ***GetDraftsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDraftsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| A cursor value used to retrieve a specific page from a paginated result set. | 
 **status** | **optional.String**| Filter the list of drafts by a particular status | 

### Return type

[**DraftsRead**](DraftsRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2020.08+json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationForDraft

> []LocalizationMetaRead GetLocalizationForDraft(ctx, draftId, accept)

Get localizations on a draft

Returns localization metadata for the draft

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2020.08\&quot;, set the value to \&quot;application/vnd.dyspatch.2020.08+json\&quot; | 

### Return type

[**[]LocalizationMetaRead**](LocalizationMetaRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2020.08+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveLocalization

> SaveLocalization(ctx, draftId, languageId, accept, inlineObject)

Create or update a localization

Inserts a localization or sets the name on an existing localization that already uses the languageId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**languageId** | **string**| A language ID (eg: en-US) | 
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2020.08\&quot;, set the value to \&quot;application/vnd.dyspatch.2020.08+json\&quot; | 
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTranslation

> SetTranslation(ctx, draftId, languageId, accept, requestBody)

Set translations for language

Completely replaces any existing translations for the given language with those provided in request body. Variables embedded in keys or values are expected to be in the format `%(my_variable)s` and will automatically convert to the correct Dyspatch format depending on the type of template. Accepts key/value pairs in JSON format or in gettext PO file format. For JSON set `Content-Type` header to `application/json`. For gettext PO format set `Content-Type` header to `text/x-gettext-translation`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**languageId** | **string**| A language ID (eg: en-US) | 
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2020.08\&quot;, set the value to \&quot;application/vnd.dyspatch.2020.08+json\&quot; | 
**requestBody** | [**map[string]string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitDraftForApproval

> SubmitDraftForApproval(ctx, draftId, accept)

Submit the draft for approval

Moves the draft into submitted state.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**accept** | **string**| A version of the API that should be used for the request. For example, to use version \&quot;2020.08\&quot;, set the value to \&quot;application/vnd.dyspatch.2020.08+json\&quot; | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

