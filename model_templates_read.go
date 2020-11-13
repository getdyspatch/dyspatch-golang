/*
 * Dyspatch API
 *
 * # Introduction  The Dyspatch API is based on the REST paradigm, and features resource based URLs with standard HTTP response codes to indicate errors. We use standard HTTP authentication and request verbs, and all responses are JSON formatted. See our [Implementation Guide](https://docs.dyspatch.io/development/implementing_dyspatch/) for more details on how to implement Dyspatch.  ## API Client Libraries Dyspatch provides API Clients for popular languages and web frameworks.  - [Java](https://github.com/getdyspatch/dyspatch-java) - [Javascript](https://github.com/getdyspatch/dyspatch-javascript) - [Python](https://github.com/getdyspatch/dyspatch-python) - [C#](https://github.com/getdyspatch/dyspatch-dotnet) - [Go](https://github.com/getdyspatch/dyspatch-golang) - [Ruby](https://github.com/getdyspatch/dyspatch-ruby) 
 *
 * API version: 2020.11
 * Contact: support@dyspatch.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dyspatch
// TemplatesRead list of template metadata
type TemplatesRead struct {
	Cursor Cursor `json:"cursor,omitempty"`
	// A list of template metadata objects
	Data []TemplateMetaRead `json:"data,omitempty"`
}
