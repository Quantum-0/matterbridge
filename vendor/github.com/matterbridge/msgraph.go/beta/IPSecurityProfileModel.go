// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// IPSecurityProfile undocumented
type IPSecurityProfile struct {
	// Entity is the base model of IPSecurityProfile
	Entity
	// ActivityGroupNames undocumented
	ActivityGroupNames []string `json:"activityGroupNames,omitempty"`
	// Address undocumented
	Address *string `json:"address,omitempty"`
	// AzureSubscriptionID undocumented
	AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty"`
	// AzureTenantID undocumented
	AzureTenantID *string `json:"azureTenantId,omitempty"`
	// CountHits undocumented
	CountHits *int `json:"countHits,omitempty"`
	// CountHosts undocumented
	CountHosts *int `json:"countHosts,omitempty"`
	// FirstSeenDateTime undocumented
	FirstSeenDateTime *time.Time `json:"firstSeenDateTime,omitempty"`
	// IPCategories undocumented
	IPCategories []IPCategory `json:"ipCategories,omitempty"`
	// IPReferenceData undocumented
	IPReferenceData []IPReferenceData `json:"ipReferenceData,omitempty"`
	// LastSeenDateTime undocumented
	LastSeenDateTime *time.Time `json:"lastSeenDateTime,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
	// Tags undocumented
	Tags []string `json:"tags,omitempty"`
	// VendorInformation undocumented
	VendorInformation *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}