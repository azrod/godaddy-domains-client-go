# DomainDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | **string** | Authorization code for transferring the Domain | [default to null]
**ContactAdmin** | [***Contact**](Contact.md) |  | [default to null]
**ContactBilling** | [***Contact**](Contact.md) |  | [default to null]
**ContactRegistrant** | [***Contact**](Contact.md) |  | [default to null]
**ContactTech** | [***Contact**](Contact.md) |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Date and time when this domain was created | [default to null]
**DeletedAt** | [**time.Time**](time.Time.md) | Date and time when this domain was deleted | [optional] [default to null]
**TransferAwayEligibleAt** | [**time.Time**](time.Time.md) | Date and time when this domain is eligible to transfer | [optional] [default to null]
**Domain** | **string** | Name of the domain | [default to null]
**DomainId** | **float64** | Unique identifier for this Domain | [default to null]
**ExpirationProtected** | **bool** | Whether or not the domain is protected from expiration | [default to null]
**Expires** | [**time.Time**](time.Time.md) | Date and time when this domain will expire | [optional] [default to null]
**ExposeWhois** | **bool** | Whether or not the domain contact details should be shown in the WHOIS | [optional] [default to null]
**HoldRegistrar** | **bool** | Whether or not the domain is on-hold by the registrar | [default to null]
**Locked** | **bool** | Whether or not the domain is locked to prevent transfers | [default to null]
**NameServers** | **[]string** | Fully-qualified domain names for DNS servers | [default to null]
**Privacy** | **bool** | Whether or not the domain has privacy protection | [default to null]
**RegistrarCreatedAt** | **string** | Date and time when this domain was created by the registrar | [optional] [default to null]
**RenewAuto** | **bool** | Whether or not the domain is configured to automatically renew | [default to null]
**RenewDeadline** | [**time.Time**](time.Time.md) | Date the domain must renew on | [default to null]
**Status** | **string** | Processing status of the domain&lt;br/&gt;&lt;ul&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;ACTIVE&lt;/strong&gt; - All is well&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;AWAITING*&lt;/strong&gt; - System is waiting for the end-user to complete an action&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;CANCELLED*&lt;/strong&gt; - Domain has been cancelled, and may or may not be reclaimable&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;CONFISCATED&lt;/strong&gt; - Domain has been confiscated, usually for abuse, chargeback, or fraud&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;DISABLED*&lt;/strong&gt; - Domain has been disabled&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;EXCLUDED*&lt;/strong&gt; - Domain has been excluded from Firehose registration&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;EXPIRED*&lt;/strong&gt; - Domain has expired&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;FAILED*&lt;/strong&gt; - Domain has failed a required action, and the system is no longer retrying&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;HELD*&lt;/strong&gt; - Domain has been placed on hold, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;LOCKED*&lt;/strong&gt; - Domain has been locked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PARKED*&lt;/strong&gt; - Domain has been parked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PENDING*&lt;/strong&gt; - Domain is working its way through an automated workflow&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;RESERVED*&lt;/strong&gt; - Domain is reserved, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;REVERTED&lt;/strong&gt; - Domain has been reverted, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;SUSPENDED*&lt;/strong&gt; - Domain has been suspended, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;TRANSFERRED*&lt;/strong&gt; - Domain has been transferred out&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;UNKNOWN&lt;/strong&gt; - Domain is in an unknown state&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;UNLOCKED*&lt;/strong&gt; - Domain has been unlocked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;UNPARKED*&lt;/strong&gt; - Domain has been unparked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;UPDATED*&lt;/strong&gt; - Domain ownership has been transferred to another account&lt;/li&gt; &lt;/ul&gt; | [default to null]
**SubaccountId** | **string** | Reseller subaccount shopperid who can manage the domain | [optional] [default to null]
**TransferProtected** | **bool** | Whether or not the domain is protected from transfer | [default to null]
**Verifications** | [***VerificationsDomain**](VerificationsDomain.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

