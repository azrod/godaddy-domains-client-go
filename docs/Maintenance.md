# Maintenance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | Date and time (UTC) when this maintenance was created | [default to null]
**EndsAt** | **string** | Date and time (UTC) when this maintenance will complete | [default to null]
**Environment** | **string** | The environment on which the maintenance will be performed&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;OTE&lt;/strong&gt; - The Operational Testing Environment.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PRODUCTION&lt;/strong&gt; - The Live Production Environment.&lt;/li&gt;&lt;/ul&gt; | [default to null]
**MaintenanceId** | **string** | The identifier for the system maintenance | [default to null]
**ModifiedAt** | **string** | Date and time (UTC) when this maintenance was last modified | [default to null]
**Reason** | **string** | The reason for the maintenance being performed&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;EMERGENCY&lt;/strong&gt; - Unexpected Emergency maintenance.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PLANNED&lt;/strong&gt; - Planned system maintenance.&lt;/li&gt;&lt;/ul&gt; | [default to null]
**StartsAt** | **string** | Date and time (UTC) when this maintenance will start | [default to null]
**Status** | **string** | The status of maintenance&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;ACTIVE&lt;/strong&gt; - The upcoming maintenance is active.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;CANCELLED&lt;/strong&gt; - The upcoming maintenance has been cancelled.&lt;/li&gt;&lt;/ul&gt; | [default to null]
**Summary** | **string** | A brief description of what is being performed | [default to null]
**Tlds** | **[]string** | List of tlds that are in maintenance.  Generally only applies when &#x60;type&#x60; is REGISTRY | [optional] [default to null]
**Type_** | **string** | The type of maintenance being performed&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;API&lt;/strong&gt; - Programmatic Api components.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;REGISTRY&lt;/strong&gt; - The underlying Registry providing the tld(s).&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;UI&lt;/strong&gt; - User Interface components.&lt;/li&gt;&lt;/ul&gt; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

