//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetworkfabric

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// L2IsolationDomainsClient contains the methods for the L2IsolationDomains group.
// Don't use this type directly, use NewL2IsolationDomainsClient() instead.
type L2IsolationDomainsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewL2IsolationDomainsClient creates a new instance of L2IsolationDomainsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewL2IsolationDomainsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*L2IsolationDomainsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &L2IsolationDomainsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCommitConfiguration - Commits the configuration of the given resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2IsolationDomainName - Name of the L2 Isolation Domain.
//   - options - L2IsolationDomainsClientBeginCommitConfigurationOptions contains the optional parameters for the L2IsolationDomainsClient.BeginCommitConfiguration
//     method.
func (client *L2IsolationDomainsClient) BeginCommitConfiguration(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientBeginCommitConfigurationOptions) (*runtime.Poller[L2IsolationDomainsClientCommitConfigurationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.commitConfiguration(ctx, resourceGroupName, l2IsolationDomainName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L2IsolationDomainsClientCommitConfigurationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[L2IsolationDomainsClientCommitConfigurationResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CommitConfiguration - Commits the configuration of the given resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L2IsolationDomainsClient) commitConfiguration(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientBeginCommitConfigurationOptions) (*http.Response, error) {
	var err error
	const operationName = "L2IsolationDomainsClient.BeginCommitConfiguration"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.commitConfigurationCreateRequest(ctx, resourceGroupName, l2IsolationDomainName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// commitConfigurationCreateRequest creates the CommitConfiguration request.
func (client *L2IsolationDomainsClient) commitConfigurationCreateRequest(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientBeginCommitConfigurationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/{l2IsolationDomainName}/commitConfiguration"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2IsolationDomainName == "" {
		return nil, errors.New("parameter l2IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2IsolationDomainName}", url.PathEscape(l2IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginCreate - Creates layer 2 network connectivity between compute nodes within a rack and across racks.The configuration
// is applied on the devices only after the isolation domain is enabled.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2IsolationDomainName - Name of the L2 Isolation Domain.
//   - body - Request payload.
//   - options - L2IsolationDomainsClientBeginCreateOptions contains the optional parameters for the L2IsolationDomainsClient.BeginCreate
//     method.
func (client *L2IsolationDomainsClient) BeginCreate(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, body L2IsolationDomain, options *L2IsolationDomainsClientBeginCreateOptions) (*runtime.Poller[L2IsolationDomainsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, l2IsolationDomainName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L2IsolationDomainsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[L2IsolationDomainsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates layer 2 network connectivity between compute nodes within a rack and across racks.The configuration is
// applied on the devices only after the isolation domain is enabled.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L2IsolationDomainsClient) create(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, body L2IsolationDomain, options *L2IsolationDomainsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "L2IsolationDomainsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, l2IsolationDomainName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *L2IsolationDomainsClient) createCreateRequest(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, body L2IsolationDomain, options *L2IsolationDomainsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/{l2IsolationDomainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2IsolationDomainName == "" {
		return nil, errors.New("parameter l2IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2IsolationDomainName}", url.PathEscape(l2IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes layer 2 connectivity between compute nodes by managed by named L2 Isolation name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2IsolationDomainName - Name of the L2 Isolation Domain.
//   - options - L2IsolationDomainsClientBeginDeleteOptions contains the optional parameters for the L2IsolationDomainsClient.BeginDelete
//     method.
func (client *L2IsolationDomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientBeginDeleteOptions) (*runtime.Poller[L2IsolationDomainsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, l2IsolationDomainName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L2IsolationDomainsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[L2IsolationDomainsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes layer 2 connectivity between compute nodes by managed by named L2 Isolation name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L2IsolationDomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "L2IsolationDomainsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, l2IsolationDomainName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *L2IsolationDomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/{l2IsolationDomainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2IsolationDomainName == "" {
		return nil, errors.New("parameter l2IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2IsolationDomainName}", url.PathEscape(l2IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements L2 Isolation Domain GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2IsolationDomainName - Name of the L2 Isolation Domain.
//   - options - L2IsolationDomainsClientGetOptions contains the optional parameters for the L2IsolationDomainsClient.Get method.
func (client *L2IsolationDomainsClient) Get(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientGetOptions) (L2IsolationDomainsClientGetResponse, error) {
	var err error
	const operationName = "L2IsolationDomainsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, l2IsolationDomainName, options)
	if err != nil {
		return L2IsolationDomainsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return L2IsolationDomainsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return L2IsolationDomainsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *L2IsolationDomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/{l2IsolationDomainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2IsolationDomainName == "" {
		return nil, errors.New("parameter l2IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2IsolationDomainName}", url.PathEscape(l2IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *L2IsolationDomainsClient) getHandleResponse(resp *http.Response) (L2IsolationDomainsClientGetResponse, error) {
	result := L2IsolationDomainsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L2IsolationDomain); err != nil {
		return L2IsolationDomainsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Displays L2IsolationDomains list by resource group GET method.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - L2IsolationDomainsClientListByResourceGroupOptions contains the optional parameters for the L2IsolationDomainsClient.NewListByResourceGroupPager
//     method.
func (client *L2IsolationDomainsClient) NewListByResourceGroupPager(resourceGroupName string, options *L2IsolationDomainsClientListByResourceGroupOptions) *runtime.Pager[L2IsolationDomainsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[L2IsolationDomainsClientListByResourceGroupResponse]{
		More: func(page L2IsolationDomainsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *L2IsolationDomainsClientListByResourceGroupResponse) (L2IsolationDomainsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "L2IsolationDomainsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return L2IsolationDomainsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *L2IsolationDomainsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *L2IsolationDomainsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *L2IsolationDomainsClient) listByResourceGroupHandleResponse(resp *http.Response) (L2IsolationDomainsClientListByResourceGroupResponse, error) {
	result := L2IsolationDomainsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L2IsolationDomainsListResult); err != nil {
		return L2IsolationDomainsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Displays L2IsolationDomains list by subscription GET method.
//
// Generated from API version 2023-06-15
//   - options - L2IsolationDomainsClientListBySubscriptionOptions contains the optional parameters for the L2IsolationDomainsClient.NewListBySubscriptionPager
//     method.
func (client *L2IsolationDomainsClient) NewListBySubscriptionPager(options *L2IsolationDomainsClientListBySubscriptionOptions) *runtime.Pager[L2IsolationDomainsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[L2IsolationDomainsClientListBySubscriptionResponse]{
		More: func(page L2IsolationDomainsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *L2IsolationDomainsClientListBySubscriptionResponse) (L2IsolationDomainsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "L2IsolationDomainsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return L2IsolationDomainsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *L2IsolationDomainsClient) listBySubscriptionCreateRequest(ctx context.Context, options *L2IsolationDomainsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *L2IsolationDomainsClient) listBySubscriptionHandleResponse(resp *http.Response) (L2IsolationDomainsClientListBySubscriptionResponse, error) {
	result := L2IsolationDomainsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L2IsolationDomainsListResult); err != nil {
		return L2IsolationDomainsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - API to update certain properties of the L2 Isolation Domain resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2IsolationDomainName - Name of the L2 Isolation Domain.
//   - body - API to update certain properties of the L2 Isolation Domain resource..
//   - options - L2IsolationDomainsClientBeginUpdateOptions contains the optional parameters for the L2IsolationDomainsClient.BeginUpdate
//     method.
func (client *L2IsolationDomainsClient) BeginUpdate(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, body L2IsolationDomainPatch, options *L2IsolationDomainsClientBeginUpdateOptions) (*runtime.Poller[L2IsolationDomainsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, l2IsolationDomainName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L2IsolationDomainsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[L2IsolationDomainsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - API to update certain properties of the L2 Isolation Domain resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L2IsolationDomainsClient) update(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, body L2IsolationDomainPatch, options *L2IsolationDomainsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "L2IsolationDomainsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, l2IsolationDomainName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *L2IsolationDomainsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, body L2IsolationDomainPatch, options *L2IsolationDomainsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/{l2IsolationDomainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2IsolationDomainName == "" {
		return nil, errors.New("parameter l2IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2IsolationDomainName}", url.PathEscape(l2IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateAdministrativeState - Enables isolation domain across the fabric or on specified racks.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2IsolationDomainName - Name of the L2 Isolation Domain.
//   - body - Request payload.
//   - options - L2IsolationDomainsClientBeginUpdateAdministrativeStateOptions contains the optional parameters for the L2IsolationDomainsClient.BeginUpdateAdministrativeState
//     method.
func (client *L2IsolationDomainsClient) BeginUpdateAdministrativeState(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, body UpdateAdministrativeState, options *L2IsolationDomainsClientBeginUpdateAdministrativeStateOptions) (*runtime.Poller[L2IsolationDomainsClientUpdateAdministrativeStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateAdministrativeState(ctx, resourceGroupName, l2IsolationDomainName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L2IsolationDomainsClientUpdateAdministrativeStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[L2IsolationDomainsClientUpdateAdministrativeStateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateAdministrativeState - Enables isolation domain across the fabric or on specified racks.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L2IsolationDomainsClient) updateAdministrativeState(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, body UpdateAdministrativeState, options *L2IsolationDomainsClientBeginUpdateAdministrativeStateOptions) (*http.Response, error) {
	var err error
	const operationName = "L2IsolationDomainsClient.BeginUpdateAdministrativeState"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateAdministrativeStateCreateRequest(ctx, resourceGroupName, l2IsolationDomainName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateAdministrativeStateCreateRequest creates the UpdateAdministrativeState request.
func (client *L2IsolationDomainsClient) updateAdministrativeStateCreateRequest(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, body UpdateAdministrativeState, options *L2IsolationDomainsClientBeginUpdateAdministrativeStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/{l2IsolationDomainName}/updateAdministrativeState"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2IsolationDomainName == "" {
		return nil, errors.New("parameter l2IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2IsolationDomainName}", url.PathEscape(l2IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginValidateConfiguration - Validates the configuration of the resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2IsolationDomainName - Name of the L2 Isolation Domain.
//   - options - L2IsolationDomainsClientBeginValidateConfigurationOptions contains the optional parameters for the L2IsolationDomainsClient.BeginValidateConfiguration
//     method.
func (client *L2IsolationDomainsClient) BeginValidateConfiguration(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientBeginValidateConfigurationOptions) (*runtime.Poller[L2IsolationDomainsClientValidateConfigurationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.validateConfiguration(ctx, resourceGroupName, l2IsolationDomainName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L2IsolationDomainsClientValidateConfigurationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[L2IsolationDomainsClientValidateConfigurationResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ValidateConfiguration - Validates the configuration of the resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L2IsolationDomainsClient) validateConfiguration(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientBeginValidateConfigurationOptions) (*http.Response, error) {
	var err error
	const operationName = "L2IsolationDomainsClient.BeginValidateConfiguration"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateConfigurationCreateRequest(ctx, resourceGroupName, l2IsolationDomainName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// validateConfigurationCreateRequest creates the ValidateConfiguration request.
func (client *L2IsolationDomainsClient) validateConfigurationCreateRequest(ctx context.Context, resourceGroupName string, l2IsolationDomainName string, options *L2IsolationDomainsClientBeginValidateConfigurationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/{l2IsolationDomainName}/validateConfiguration"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2IsolationDomainName == "" {
		return nil, errors.New("parameter l2IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2IsolationDomainName}", url.PathEscape(l2IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
