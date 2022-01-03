/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type DomainDetail struct {
	// Authorization code for transferring the Domain
	AuthCode string `json:"authCode"`
	ContactAdmin *Contact `json:"contactAdmin"`
	ContactBilling *Contact `json:"contactBilling"`
	ContactRegistrant *Contact `json:"contactRegistrant"`
	ContactTech *Contact `json:"contactTech"`
	// Date and time when this domain was created
	CreatedAt time.Time `json:"createdAt"`
	// Date and time when this domain was deleted
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	// Date and time when this domain is eligible to transfer
	TransferAwayEligibleAt time.Time `json:"transferAwayEligibleAt,omitempty"`
	// Name of the domain
	Domain string `json:"domain"`
	// Unique identifier for this Domain
	DomainId float64 `json:"domainId"`
	// Whether or not the domain is protected from expiration
	ExpirationProtected bool `json:"expirationProtected"`
	// Date and time when this domain will expire
	Expires time.Time `json:"expires,omitempty"`
	// Whether or not the domain contact details should be shown in the WHOIS
	ExposeWhois bool `json:"exposeWhois,omitempty"`
	// Whether or not the domain is on-hold by the registrar
	HoldRegistrar bool `json:"holdRegistrar"`
	// Whether or not the domain is locked to prevent transfers
	Locked bool `json:"locked"`
	// Fully-qualified domain names for DNS servers
	NameServers []string `json:"nameServers"`
	// Whether or not the domain has privacy protection
	Privacy bool `json:"privacy"`
	// Whether or not the domain is configured to automatically renew
	RenewAuto bool `json:"renewAuto"`
	// Date the domain must renew on
	RenewDeadline time.Time `json:"renewDeadline"`
	// Processing status of the domain<br/><ul> <li><strong style='margin-left: 12px;'>ACTIVE</strong> - All is well</li> <li><strong style='margin-left: 12px;'>AWAITING*</strong> - System is waiting for the end-user to complete an action</li> <li><strong style='margin-left: 12px;'>CANCELLED*</strong> - Domain has been cancelled, and may or may not be reclaimable</li> <li><strong style='margin-left: 12px;'>CONFISCATED</strong> - Domain has been confiscated, usually for abuse, chargeback, or fraud</li> <li><strong style='margin-left: 12px;'>DISABLED*</strong> - Domain has been disabled</li> <li><strong style='margin-left: 12px;'>EXCLUDED*</strong> - Domain has been excluded from Firehose registration</li> <li><strong style='margin-left: 12px;'>EXPIRED*</strong> - Domain has expired</li> <li><strong style='margin-left: 12px;'>FAILED*</strong> - Domain has failed a required action, and the system is no longer retrying</li> <li><strong style='margin-left: 12px;'>HELD*</strong> - Domain has been placed on hold, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>LOCKED*</strong> - Domain has been locked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>PARKED*</strong> - Domain has been parked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>PENDING*</strong> - Domain is working its way through an automated workflow</li> <li><strong style='margin-left: 12px;'>RESERVED*</strong> - Domain is reserved, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>REVERTED</strong> - Domain has been reverted, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>SUSPENDED*</strong> - Domain has been suspended, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>TRANSFERRED*</strong> - Domain has been transferred out</li> <li><strong style='margin-left: 12px;'>UNKNOWN</strong> - Domain is in an unknown state</li> <li><strong style='margin-left: 12px;'>UNLOCKED*</strong> - Domain has been unlocked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>UNPARKED*</strong> - Domain has been unparked, and likely requires intervention from Support</li> <li><strong style='margin-left: 12px;'>UPDATED*</strong> - Domain ownership has been transferred to another account</li> </ul>
	Status string `json:"status"`
	// Reseller subaccount shopperid who can manage the domain
	SubaccountId string `json:"subaccountId,omitempty"`
	// Whether or not the domain is protected from transfer
	TransferProtected bool `json:"transferProtected"`
	Verifications *VerificationsDomain `json:"verifications,omitempty"`
}
