/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DomainNotification struct {
	// The notification ID to be used in POST /v2/customers/{customerId}/domains/notifications to acknowledge the notification
	NotificationId string `json:"notificationId"`
	// The type of action the notification relates to
	Type_ string `json:"type"`
	// The resource the notification pertains to.
	Resource string `json:"resource"`
	// The type of resource the notification relates to
	ResourceType string `json:"resourceType"`
	// The resulting status of the action.
	Status string `json:"status"`
	// The date the notification was added
	AddedAt string `json:"addedAt"`
	// A client provided identifier (via X-Request-Id header) indicating the request this notification is for
	RequestId string `json:"requestId,omitempty"`
	// The notification data for the given type as specifed by GET /v2/customers/{customerId}/domains/notifications/schema
	Metadata *interface{} `json:"metadata,omitempty"`
}
