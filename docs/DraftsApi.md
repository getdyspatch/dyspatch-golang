# \DraftsApi

All URIs are relative to *https://api.dyspatch.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DraftsDraftIdGet**](DraftsApi.md#DraftsDraftIdGet) | **Get** /drafts/{draftId} | Get Draft by ID
[**DraftsDraftIdLocalizationKeysGet**](DraftsApi.md#DraftsDraftIdLocalizationKeysGet) | **Get** /drafts/{draftId}/localizationKeys | Get localization keys
[**DraftsDraftIdLocalizationsGet**](DraftsApi.md#DraftsDraftIdLocalizationsGet) | **Get** /drafts/{draftId}/localizations | Get localizations on a draft
[**DraftsDraftIdLocalizationsLanguageIdDelete**](DraftsApi.md#DraftsDraftIdLocalizationsLanguageIdDelete) | **Delete** /drafts/{draftId}/localizations/{languageId} | Remove a localization
[**DraftsDraftIdLocalizationsLanguageIdPut**](DraftsApi.md#DraftsDraftIdLocalizationsLanguageIdPut) | **Put** /drafts/{draftId}/localizations/{languageId} | Create or update a localization
[**DraftsDraftIdLocalizationsLanguageIdTranslationsPut**](DraftsApi.md#DraftsDraftIdLocalizationsLanguageIdTranslationsPut) | **Put** /drafts/{draftId}/localizations/{languageId}/translations | Set translations for language
[**DraftsDraftIdPublishRequestPost**](DraftsApi.md#DraftsDraftIdPublishRequestPost) | **Post** /drafts/{draftId}/publishRequest | Submit the draft for approval
[**DraftsGet**](DraftsApi.md#DraftsGet) | **Get** /drafts | List Drafts



## DraftsDraftIdGet

> DraftRead DraftsDraftIdGet(ctx, draftId, targetLanguage)

Get Draft by ID

Gets a draft object with the matching ID. The \"compiled\" field will contain the template in the default, unlocalized form.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**targetLanguage** | **string**| The type of templating language to compile as. Should only be used for visual templates. | 

### Return type

[**DraftRead**](DraftRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2019.10+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DraftsDraftIdLocalizationKeysGet

> []LocalizationKeyRead DraftsDraftIdLocalizationKeysGet(ctx, draftId, optional)

Get localization keys

Returns the list of values that need to be translated for the draft. Set the `Accept` header to `application/vnd.dyspatch.2019.10+json` to get a JSON object, or `text/vnd.dyspatch.2019.10+x-gettext-translation` to get the POT file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
 **optional** | ***DraftsDraftIdLocalizationKeysGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DraftsDraftIdLocalizationKeysGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **optional.String**| A version of the API that should be used for the request. For example, to use version \&quot;2019.10\&quot;, set the value to \&quot;application/vnd.dyspatch.2019.10+json\&quot; | 

### Return type

[**[]LocalizationKeyRead**](LocalizationKeyRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2019.10+json, text/vnd.dyspatch.2019.10+x-gettext-translation

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DraftsDraftIdLocalizationsGet

> []LocalizationMetaRead DraftsDraftIdLocalizationsGet(ctx, draftId)

Get localizations on a draft

Returns localization metadata for the draft

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 

### Return type

[**[]LocalizationMetaRead**](LocalizationMetaRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2019.10+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DraftsDraftIdLocalizationsLanguageIdDelete

> DraftsDraftIdLocalizationsLanguageIdDelete(ctx, draftId, languageId)

Remove a localization

Deletes the localization with the given language ID if it exists

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**languageId** | **string**| A language ID (eg: en-US) | 

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


## DraftsDraftIdLocalizationsLanguageIdPut

> DraftsDraftIdLocalizationsLanguageIdPut(ctx, draftId, languageId, body)

Create or update a localization

Inserts a localization or sets the name on an existing localization that already uses the languageId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**languageId** | **string**| A language ID (eg: en-US) | 
**body** | [**InlineObject**](InlineObject.md)|  | 

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


## DraftsDraftIdLocalizationsLanguageIdTranslationsPut

> DraftsDraftIdLocalizationsLanguageIdTranslationsPut(ctx, draftId, languageId, body)

Set translations for language

Completely replaces any existing translations for the given language with those provided in request body. Variables embedded in keys or values are expected to be in the format `%(my_variable)s` and will automatically convert to the correct Dyspatch format depending on the type of template. Accepts key/value pairs in JSON format or in gettext PO file format. For JSON set `Content-Type` header to `application/json`. For gettext PO format set `Content-Type` header to `text/x-gettext-translation`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 
**languageId** | **string**| A language ID (eg: en-US) | 
**body** | [**InlineObject1**](InlineObject1.md)|  | 

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


## DraftsDraftIdPublishRequestPost

> DraftsDraftIdPublishRequestPost(ctx, draftId)

Submit the draft for approval

Moves the draft into submitted state.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string**| A draft ID | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2019.10+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DraftsGet

> DraftsRead DraftsGet(ctx, optional)

List Drafts

Returns all drafts for your organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DraftsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DraftsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| Filter the list of drafts by a particular status | 

### Return type

[**DraftsRead**](DraftsRead.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.dyspatch.2019.10+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

