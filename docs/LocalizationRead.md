# LocalizationRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An opaque, unique identifier for a localization | [optional] 
**Languages** | **[]string** | a list of locale codes that are available in the localization. See [supported languages](https://docs.dyspatch.io/localization/supported_languages/) for an exhaustive list of locale codes.  | [optional] 
**Url** | **string** | The API url for a specific localization | [optional] 
**Template** | **string** | An opaque, unique identifier for a template | [optional] 
**Compiled** | [**CompiledRead**](CompiledRead.md) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The time of initial creation | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The time of last update | [optional] 
**Name** | **string** | The user-specified name of a localization | [optional] 
**LocaleGroup** | **string** | the locale group this localization belongs to, if this field is empty the localization does not belong to any locale group | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


