/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ErrorLimit struct {
	// Short identifier for the error, suitable for indicating the specific error within client code
	Code string `json:"code"`
	// List of the specific fields, and the errors found with their contents
	Fields []ErrorField `json:"fields,omitempty"`
	// Human-readable, English description of the error
	Message string `json:"message,omitempty"`
	// Number of seconds to wait before attempting a similar request
	RetryAfterSec int32 `json:"retryAfterSec"`
}
