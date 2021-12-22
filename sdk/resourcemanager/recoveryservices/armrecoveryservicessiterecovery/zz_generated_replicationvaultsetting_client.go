//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReplicationVaultSettingClient contains the methods for the ReplicationVaultSetting group.
// Don't use this type directly, use NewReplicationVaultSettingClient() instead.
type ReplicationVaultSettingClient struct {
	ep                string
	pl                runtime.Pipeline
	resourceName      string
	resourceGroupName string
	subscriptionID    string
}

// NewReplicationVaultSettingClient creates a new instance of ReplicationVaultSettingClient with the specified values.
func NewReplicationVaultSettingClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ReplicationVaultSettingClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ReplicationVaultSettingClient{resourceName: resourceName, resourceGroupName: resourceGroupName, subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - The operation to configure vault setting.
// If the operation fails it returns a generic error.
func (client *ReplicationVaultSettingClient) BeginCreate(ctx context.Context, vaultSettingName string, input VaultSettingCreationInput, options *ReplicationVaultSettingBeginCreateOptions) (ReplicationVaultSettingCreatePollerResponse, error) {
	resp, err := client.create(ctx, vaultSettingName, input, options)
	if err != nil {
		return ReplicationVaultSettingCreatePollerResponse{}, err
	}
	result := ReplicationVaultSettingCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReplicationVaultSettingClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return ReplicationVaultSettingCreatePollerResponse{}, err
	}
	result.Poller = &ReplicationVaultSettingCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - The operation to configure vault setting.
// If the operation fails it returns a generic error.
func (client *ReplicationVaultSettingClient) create(ctx context.Context, vaultSettingName string, input VaultSettingCreationInput, options *ReplicationVaultSettingBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, vaultSettingName, input, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ReplicationVaultSettingClient) createCreateRequest(ctx context.Context, vaultSettingName string, input VaultSettingCreationInput, options *ReplicationVaultSettingBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings/{vaultSettingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if vaultSettingName == "" {
		return nil, errors.New("parameter vaultSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultSettingName}", url.PathEscape(vaultSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, input)
}

// createHandleError handles the Create error response.
func (client *ReplicationVaultSettingClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the vault setting. This includes the Migration Hub connection settings.
// If the operation fails it returns a generic error.
func (client *ReplicationVaultSettingClient) Get(ctx context.Context, vaultSettingName string, options *ReplicationVaultSettingGetOptions) (ReplicationVaultSettingGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultSettingName, options)
	if err != nil {
		return ReplicationVaultSettingGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationVaultSettingGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationVaultSettingGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationVaultSettingClient) getCreateRequest(ctx context.Context, vaultSettingName string, options *ReplicationVaultSettingGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings/{vaultSettingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if vaultSettingName == "" {
		return nil, errors.New("parameter vaultSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultSettingName}", url.PathEscape(vaultSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationVaultSettingClient) getHandleResponse(resp *http.Response) (ReplicationVaultSettingGetResponse, error) {
	result := ReplicationVaultSettingGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultSetting); err != nil {
		return ReplicationVaultSettingGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ReplicationVaultSettingClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Gets the list of vault setting. This includes the Migration Hub connection settings.
// If the operation fails it returns a generic error.
func (client *ReplicationVaultSettingClient) List(options *ReplicationVaultSettingListOptions) *ReplicationVaultSettingListPager {
	return &ReplicationVaultSettingListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ReplicationVaultSettingListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VaultSettingCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ReplicationVaultSettingClient) listCreateRequest(ctx context.Context, options *ReplicationVaultSettingListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationVaultSettingClient) listHandleResponse(resp *http.Response) (ReplicationVaultSettingListResponse, error) {
	result := ReplicationVaultSettingListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultSettingCollection); err != nil {
		return ReplicationVaultSettingListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ReplicationVaultSettingClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}