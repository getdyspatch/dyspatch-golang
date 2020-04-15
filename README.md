# Go API client for Dyspatch

# Introduction

The Dyspatch API is based on the REST paradigm and features resource based URLs
with standard HTTP response codes to indicate errors. We use standard HTTP
authentication and request verbs and all responses are JSON formatted. See our
[Implementation
Guide](https://docs.dyspatch.io/development/implementing_dyspatch/) for more
details on how to implement Dyspatch.

## Overview

This API client was generated by the [OpenAPI
Generator](https://openapi-generator.tech) project.  By using the
[OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily
generate an API client.

- API version: 2020.04
- Package version: 5.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://docs.dyspatch.io](https://docs.dyspatch.io)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./dyspatch"
```

Install the latest version of the client project:
```sh
go get -u github.com/getdyspatch/dyspatch-golang
```

Installing a versioned client:

```sh
go get -u gopkg.in/getdyspatch/dyspatch-golang.v3
```

Importing a versioned client:

```go
import "gopkg.in/getdyspatch/dyspatch-golang.v3"
```

## Getting Started

```golang
package main
import (
	"fmt"
	"github.com/antihax/optional"
	"gopkg.in/getdyspatch/dyspatch-golang.v3"
	"golang.org/x/net/context"
)
func main() {
	cfg := dyspatch.NewConfiguration()
    cfg.AddDefaultHeader("Accept", "application/vnd.dyspatch.2019.10+json") // Set the API version
	auth := context.WithValue(context.Background(), dyspatch.ContextAPIKey, dyspatch.APIKey{
		Key:    "DYSPATCH_API_KEY",
		Prefix: "Bearer",
	})
	api := dyspatch.NewAPIClient(cfg)
	opts := dyspatch.TemplatesGetOpts{Cursor: optional.NewString("")} // Use this to page API results
	templates, _, err := api.TemplatesApi.TemplatesGet(auth, &opts)
	if err != nil {
		fmt.Println(err)
	}
	for _, template := range templates.Data {
		fmt.Println(template.Name)
	}

    res, err := client.DraftsApi.DraftsDraftIdLocalizationsLanguageIdPut(auth, "tdft_01dkyyta7dg5frv8e01186cdzf", "fr-CA", dyspatch.InlineObject{Name: "Test"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Status)
	}
}
```

## Documentation for API Endpoints

All URIs are relative to *https://api.dyspatch.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DraftsApi* | [**DeleteLocalization**](docs/DraftsApi.md#deletelocalization) | **Delete** /drafts/{draftId}/localizations/{languageId} | Remove a localization
*DraftsApi* | [**GetDraftById**](docs/DraftsApi.md#getdraftbyid) | **Get** /drafts/{draftId} | Get Draft by ID
*DraftsApi* | [**GetDraftLocalizationKeys**](docs/DraftsApi.md#getdraftlocalizationkeys) | **Get** /drafts/{draftId}/localizationKeys | Get localization keys
*DraftsApi* | [**GetDrafts**](docs/DraftsApi.md#getdrafts) | **Get** /drafts | List Drafts
*DraftsApi* | [**GetLocalizationForDraft**](docs/DraftsApi.md#getlocalizationfordraft) | **Get** /drafts/{draftId}/localizations | Get localizations on a draft
*DraftsApi* | [**SaveLocalization**](docs/DraftsApi.md#savelocalization) | **Put** /drafts/{draftId}/localizations/{languageId} | Create or update a localization
*DraftsApi* | [**SetTranslation**](docs/DraftsApi.md#settranslation) | **Put** /drafts/{draftId}/localizations/{languageId}/translations | Set translations for language
*DraftsApi* | [**SubmitDraftForApproval**](docs/DraftsApi.md#submitdraftforapproval) | **Post** /drafts/{draftId}/publishRequest | Submit the draft for approval
*LocalizationsApi* | [**GetLocalizationById**](docs/LocalizationsApi.md#getlocalizationbyid) | **Get** /localizations/{localizationId} | Get Localization Object by ID
*TemplatesApi* | [**GetTemplateById**](docs/TemplatesApi.md#gettemplatebyid) | **Get** /templates/{templateId} | Get Template by ID
*TemplatesApi* | [**GetTemplates**](docs/TemplatesApi.md#gettemplates) | **Get** /templates | List Templates


## Documentation For Models

 - [ApiError](docs/ApiError.md)
 - [CompiledRead](docs/CompiledRead.md)
 - [Cursor](docs/Cursor.md)
 - [DraftMetaRead](docs/DraftMetaRead.md)
 - [DraftRead](docs/DraftRead.md)
 - [DraftsRead](docs/DraftsRead.md)
 - [InlineObject](docs/InlineObject.md)
 - [LocalizationKeyRead](docs/LocalizationKeyRead.md)
 - [LocalizationMetaRead](docs/LocalizationMetaRead.md)
 - [LocalizationRead](docs/LocalizationRead.md)
 - [TemplateMetaRead](docs/TemplateMetaRead.md)
 - [TemplateRead](docs/TemplateRead.md)
 - [TemplatesRead](docs/TemplatesRead.md)


## Documentation For Authorization



## Bearer

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```


## Author

support@dyspatch.io
