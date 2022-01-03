/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DnsRecordCreateType struct {
	Data string `json:"data"`
	Name string `json:"name"`
	// Service port (SRV only)
	Port int32 `json:"port,omitempty"`
	// Record priority (MX and SRV only)
	Priority int32 `json:"priority,omitempty"`
	// Service protocol (SRV only)
	Protocol string `json:"protocol,omitempty"`
	// Service type (SRV only)
	Service string `json:"service,omitempty"`
	Ttl int32 `json:"ttl,omitempty"`
	// Record weight (SRV only)
	Weight int32 `json:"weight,omitempty"`
}
