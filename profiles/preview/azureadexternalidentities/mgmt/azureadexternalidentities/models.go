//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package azureadexternalidentities

import original "github.com/Azure/azure-sdk-for-go/services/azureadexternalidentities/mgmt/2021-04-01/azureadexternalidentities"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type B2CResourceSKUName = original.B2CResourceSKUName

const (
	B2CResourceSKUNamePremiumP1 B2CResourceSKUName = original.B2CResourceSKUNamePremiumP1
	B2CResourceSKUNamePremiumP2 B2CResourceSKUName = original.B2CResourceSKUNamePremiumP2
	B2CResourceSKUNameStandard  B2CResourceSKUName = original.B2CResourceSKUNameStandard
)

type B2CResourceSKUTier = original.B2CResourceSKUTier

const (
	B2CResourceSKUTierA0 B2CResourceSKUTier = original.B2CResourceSKUTierA0
)

type BillingType = original.BillingType

const (
	BillingTypeAuths BillingType = original.BillingTypeAuths
	BillingTypeMAU   BillingType = original.BillingTypeMAU
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type NameAvailabilityReasonType = original.NameAvailabilityReasonType

const (
	NameAvailabilityReasonTypeAlreadyExists NameAvailabilityReasonType = original.NameAvailabilityReasonTypeAlreadyExists
	NameAvailabilityReasonTypeInvalid       NameAvailabilityReasonType = original.NameAvailabilityReasonTypeInvalid
)

type TypeValue = original.TypeValue

const (
	TypeValueMicrosoftAzureActiveDirectoryb2cDirectories TypeValue = original.TypeValueMicrosoftAzureActiveDirectoryb2cDirectories
)

type AvailableOperations = original.AvailableOperations
type B2CResourceSKU = original.B2CResourceSKU
type B2CTenantResource = original.B2CTenantResource
type B2CTenantResourceList = original.B2CTenantResourceList
type B2CTenantResourceProperties = original.B2CTenantResourceProperties
type B2CTenantResourcePropertiesBillingConfig = original.B2CTenantResourcePropertiesBillingConfig
type B2CTenantUpdateRequest = original.B2CTenantUpdateRequest
type B2CTenantsClient = original.B2CTenantsClient
type B2CTenantsCreateFuture = original.B2CTenantsCreateFuture
type B2CTenantsDeleteFuture = original.B2CTenantsDeleteFuture
type BaseClient = original.BaseClient
type CheckNameAvailabilityRequestBody = original.CheckNameAvailabilityRequestBody
type CloudError = original.CloudError
type CreateTenantProperties = original.CreateTenantProperties
type CreateTenantRequestBody = original.CreateTenantRequestBody
type CreateTenantRequestBodyProperties = original.CreateTenantRequestBodyProperties
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type GuestUsagesClient = original.GuestUsagesClient
type GuestUsagesResource = original.GuestUsagesResource
type GuestUsagesResourceList = original.GuestUsagesResourceList
type GuestUsagesResourcePatch = original.GuestUsagesResourcePatch
type GuestUsagesResourceProperties = original.GuestUsagesResourceProperties
type NameAvailabilityResponse = original.NameAvailabilityResponse
type OperationDetail = original.OperationDetail
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type SystemData = original.SystemData

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewB2CTenantsClient(subscriptionID string) B2CTenantsClient {
	return original.NewB2CTenantsClient(subscriptionID)
}
func NewB2CTenantsClientWithBaseURI(baseURI string, subscriptionID string) B2CTenantsClient {
	return original.NewB2CTenantsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGuestUsagesClient(subscriptionID string) GuestUsagesClient {
	return original.NewGuestUsagesClient(subscriptionID)
}
func NewGuestUsagesClientWithBaseURI(baseURI string, subscriptionID string) GuestUsagesClient {
	return original.NewGuestUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleB2CResourceSKUNameValues() []B2CResourceSKUName {
	return original.PossibleB2CResourceSKUNameValues()
}
func PossibleB2CResourceSKUTierValues() []B2CResourceSKUTier {
	return original.PossibleB2CResourceSKUTierValues()
}
func PossibleBillingTypeValues() []BillingType {
	return original.PossibleBillingTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleNameAvailabilityReasonTypeValues() []NameAvailabilityReasonType {
	return original.PossibleNameAvailabilityReasonTypeValues()
}
func PossibleTypeValueValues() []TypeValue {
	return original.PossibleTypeValueValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}