package redisenterprise

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessKeyType enumerates the values for access key type.
type AccessKeyType string

const (
	// AccessKeyTypePrimary ...
	AccessKeyTypePrimary AccessKeyType = "Primary"
	// AccessKeyTypeSecondary ...
	AccessKeyTypeSecondary AccessKeyType = "Secondary"
)

// PossibleAccessKeyTypeValues returns an array of possible values for the AccessKeyType const type.
func PossibleAccessKeyTypeValues() []AccessKeyType {
	return []AccessKeyType{AccessKeyTypePrimary, AccessKeyTypeSecondary}
}

// ActionType enumerates the values for action type.
type ActionType string

const (
	// ActionTypeInternal ...
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns an array of possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{ActionTypeInternal}
}

// AofFrequency enumerates the values for aof frequency.
type AofFrequency string

const (
	// AofFrequencyAlways ...
	AofFrequencyAlways AofFrequency = "always"
	// AofFrequencyOnes ...
	AofFrequencyOnes AofFrequency = "1s"
)

// PossibleAofFrequencyValues returns an array of possible values for the AofFrequency const type.
func PossibleAofFrequencyValues() []AofFrequency {
	return []AofFrequency{AofFrequencyAlways, AofFrequencyOnes}
}

// ClusteringPolicy enumerates the values for clustering policy.
type ClusteringPolicy string

const (
	// ClusteringPolicyEnterpriseCluster ...
	ClusteringPolicyEnterpriseCluster ClusteringPolicy = "EnterpriseCluster"
	// ClusteringPolicyOSSCluster ...
	ClusteringPolicyOSSCluster ClusteringPolicy = "OSSCluster"
)

// PossibleClusteringPolicyValues returns an array of possible values for the ClusteringPolicy const type.
func PossibleClusteringPolicyValues() []ClusteringPolicy {
	return []ClusteringPolicy{ClusteringPolicyEnterpriseCluster, ClusteringPolicyOSSCluster}
}

// EvictionPolicy enumerates the values for eviction policy.
type EvictionPolicy string

const (
	// EvictionPolicyAllKeysLFU ...
	EvictionPolicyAllKeysLFU EvictionPolicy = "AllKeysLFU"
	// EvictionPolicyAllKeysLRU ...
	EvictionPolicyAllKeysLRU EvictionPolicy = "AllKeysLRU"
	// EvictionPolicyAllKeysRandom ...
	EvictionPolicyAllKeysRandom EvictionPolicy = "AllKeysRandom"
	// EvictionPolicyNoEviction ...
	EvictionPolicyNoEviction EvictionPolicy = "NoEviction"
	// EvictionPolicyVolatileLFU ...
	EvictionPolicyVolatileLFU EvictionPolicy = "VolatileLFU"
	// EvictionPolicyVolatileLRU ...
	EvictionPolicyVolatileLRU EvictionPolicy = "VolatileLRU"
	// EvictionPolicyVolatileRandom ...
	EvictionPolicyVolatileRandom EvictionPolicy = "VolatileRandom"
	// EvictionPolicyVolatileTTL ...
	EvictionPolicyVolatileTTL EvictionPolicy = "VolatileTTL"
)

// PossibleEvictionPolicyValues returns an array of possible values for the EvictionPolicy const type.
func PossibleEvictionPolicyValues() []EvictionPolicy {
	return []EvictionPolicy{EvictionPolicyAllKeysLFU, EvictionPolicyAllKeysLRU, EvictionPolicyAllKeysRandom, EvictionPolicyNoEviction, EvictionPolicyVolatileLFU, EvictionPolicyVolatileLRU, EvictionPolicyVolatileRandom, EvictionPolicyVolatileTTL}
}

// LinkState enumerates the values for link state.
type LinkState string

const (
	// LinkStateLinked ...
	LinkStateLinked LinkState = "Linked"
	// LinkStateLinkFailed ...
	LinkStateLinkFailed LinkState = "LinkFailed"
	// LinkStateLinking ...
	LinkStateLinking LinkState = "Linking"
	// LinkStateUnlinkFailed ...
	LinkStateUnlinkFailed LinkState = "UnlinkFailed"
	// LinkStateUnlinking ...
	LinkStateUnlinking LinkState = "Unlinking"
)

// PossibleLinkStateValues returns an array of possible values for the LinkState const type.
func PossibleLinkStateValues() []LinkState {
	return []LinkState{LinkStateLinked, LinkStateLinkFailed, LinkStateLinking, LinkStateUnlinkFailed, LinkStateUnlinking}
}

// Origin enumerates the values for origin.
type Origin string

const (
	// OriginSystem ...
	OriginSystem Origin = "system"
	// OriginUser ...
	OriginUser Origin = "user"
	// OriginUsersystem ...
	OriginUsersystem Origin = "user,system"
)

// PossibleOriginValues returns an array of possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{OriginSystem, OriginUser, OriginUsersystem}
}

// PrivateEndpointConnectionProvisioningState enumerates the values for private endpoint connection
// provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// PrivateEndpointConnectionProvisioningStateCreating ...
	PrivateEndpointConnectionProvisioningStateCreating PrivateEndpointConnectionProvisioningState = "Creating"
	// PrivateEndpointConnectionProvisioningStateDeleting ...
	PrivateEndpointConnectionProvisioningStateDeleting PrivateEndpointConnectionProvisioningState = "Deleting"
	// PrivateEndpointConnectionProvisioningStateFailed ...
	PrivateEndpointConnectionProvisioningStateFailed PrivateEndpointConnectionProvisioningState = "Failed"
	// PrivateEndpointConnectionProvisioningStateSucceeded ...
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns an array of possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{PrivateEndpointConnectionProvisioningStateCreating, PrivateEndpointConnectionProvisioningStateDeleting, PrivateEndpointConnectionProvisioningStateFailed, PrivateEndpointConnectionProvisioningStateSucceeded}
}

// PrivateEndpointServiceConnectionStatus enumerates the values for private endpoint service connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	// PrivateEndpointServiceConnectionStatusApproved ...
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	// PrivateEndpointServiceConnectionStatusPending ...
	PrivateEndpointServiceConnectionStatusPending PrivateEndpointServiceConnectionStatus = "Pending"
	// PrivateEndpointServiceConnectionStatusRejected ...
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns an array of possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{PrivateEndpointServiceConnectionStatusApproved, PrivateEndpointServiceConnectionStatusPending, PrivateEndpointServiceConnectionStatusRejected}
}

// Protocol enumerates the values for protocol.
type Protocol string

const (
	// ProtocolEncrypted ...
	ProtocolEncrypted Protocol = "Encrypted"
	// ProtocolPlaintext ...
	ProtocolPlaintext Protocol = "Plaintext"
)

// PossibleProtocolValues returns an array of possible values for the Protocol const type.
func PossibleProtocolValues() []Protocol {
	return []Protocol{ProtocolEncrypted, ProtocolPlaintext}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCanceled, ProvisioningStateCreating, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateSucceeded, ProvisioningStateUpdating}
}

// RdbFrequency enumerates the values for rdb frequency.
type RdbFrequency string

const (
	// RdbFrequencyOneh ...
	RdbFrequencyOneh RdbFrequency = "1h"
	// RdbFrequencyOneTwoh ...
	RdbFrequencyOneTwoh RdbFrequency = "12h"
	// RdbFrequencySixh ...
	RdbFrequencySixh RdbFrequency = "6h"
)

// PossibleRdbFrequencyValues returns an array of possible values for the RdbFrequency const type.
func PossibleRdbFrequencyValues() []RdbFrequency {
	return []RdbFrequency{RdbFrequencyOneh, RdbFrequencyOneTwoh, RdbFrequencySixh}
}

// ResourceState enumerates the values for resource state.
type ResourceState string

const (
	// ResourceStateCreateFailed ...
	ResourceStateCreateFailed ResourceState = "CreateFailed"
	// ResourceStateCreating ...
	ResourceStateCreating ResourceState = "Creating"
	// ResourceStateDeleteFailed ...
	ResourceStateDeleteFailed ResourceState = "DeleteFailed"
	// ResourceStateDeleting ...
	ResourceStateDeleting ResourceState = "Deleting"
	// ResourceStateDisabled ...
	ResourceStateDisabled ResourceState = "Disabled"
	// ResourceStateDisableFailed ...
	ResourceStateDisableFailed ResourceState = "DisableFailed"
	// ResourceStateDisabling ...
	ResourceStateDisabling ResourceState = "Disabling"
	// ResourceStateEnableFailed ...
	ResourceStateEnableFailed ResourceState = "EnableFailed"
	// ResourceStateEnabling ...
	ResourceStateEnabling ResourceState = "Enabling"
	// ResourceStateRunning ...
	ResourceStateRunning ResourceState = "Running"
	// ResourceStateUpdateFailed ...
	ResourceStateUpdateFailed ResourceState = "UpdateFailed"
	// ResourceStateUpdating ...
	ResourceStateUpdating ResourceState = "Updating"
)

// PossibleResourceStateValues returns an array of possible values for the ResourceState const type.
func PossibleResourceStateValues() []ResourceState {
	return []ResourceState{ResourceStateCreateFailed, ResourceStateCreating, ResourceStateDeleteFailed, ResourceStateDeleting, ResourceStateDisabled, ResourceStateDisableFailed, ResourceStateDisabling, ResourceStateEnableFailed, ResourceStateEnabling, ResourceStateRunning, ResourceStateUpdateFailed, ResourceStateUpdating}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// SkuNameEnterpriseE10 ...
	SkuNameEnterpriseE10 SkuName = "Enterprise_E10"
	// SkuNameEnterpriseE100 ...
	SkuNameEnterpriseE100 SkuName = "Enterprise_E100"
	// SkuNameEnterpriseE20 ...
	SkuNameEnterpriseE20 SkuName = "Enterprise_E20"
	// SkuNameEnterpriseE50 ...
	SkuNameEnterpriseE50 SkuName = "Enterprise_E50"
	// SkuNameEnterpriseFlashF1500 ...
	SkuNameEnterpriseFlashF1500 SkuName = "EnterpriseFlash_F1500"
	// SkuNameEnterpriseFlashF300 ...
	SkuNameEnterpriseFlashF300 SkuName = "EnterpriseFlash_F300"
	// SkuNameEnterpriseFlashF700 ...
	SkuNameEnterpriseFlashF700 SkuName = "EnterpriseFlash_F700"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{SkuNameEnterpriseE10, SkuNameEnterpriseE100, SkuNameEnterpriseE20, SkuNameEnterpriseE50, SkuNameEnterpriseFlashF1500, SkuNameEnterpriseFlashF300, SkuNameEnterpriseFlashF700}
}

// TLSVersion enumerates the values for tls version.
type TLSVersion string

const (
	// TLSVersionOneFullStopOne ...
	TLSVersionOneFullStopOne TLSVersion = "1.1"
	// TLSVersionOneFullStopTwo ...
	TLSVersionOneFullStopTwo TLSVersion = "1.2"
	// TLSVersionOneFullStopZero ...
	TLSVersionOneFullStopZero TLSVersion = "1.0"
)

// PossibleTLSVersionValues returns an array of possible values for the TLSVersion const type.
func PossibleTLSVersionValues() []TLSVersion {
	return []TLSVersion{TLSVersionOneFullStopOne, TLSVersionOneFullStopTwo, TLSVersionOneFullStopZero}
}