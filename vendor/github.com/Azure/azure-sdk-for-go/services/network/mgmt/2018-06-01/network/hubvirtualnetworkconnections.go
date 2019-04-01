package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// HubVirtualNetworkConnectionsClient is the network Client
type HubVirtualNetworkConnectionsClient struct {
	BaseClient
}

// NewHubVirtualNetworkConnectionsClient creates an instance of the HubVirtualNetworkConnectionsClient client.
func NewHubVirtualNetworkConnectionsClient(subscriptionID string) HubVirtualNetworkConnectionsClient {
	return NewHubVirtualNetworkConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewHubVirtualNetworkConnectionsClientWithBaseURI creates an instance of the HubVirtualNetworkConnectionsClient
// client.
func NewHubVirtualNetworkConnectionsClientWithBaseURI(baseURI string, subscriptionID string) HubVirtualNetworkConnectionsClient {
	return HubVirtualNetworkConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieves the details of a HubVirtualNetworkConnection.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
// connectionName - the name of the vpn connection.
func (client HubVirtualNetworkConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string) (result HubVirtualNetworkConnection, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualHubName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client HubVirtualNetworkConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client HubVirtualNetworkConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client HubVirtualNetworkConnectionsClient) GetResponder(resp *http.Response) (result HubVirtualNetworkConnection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieves the details of all HubVirtualNetworkConnections.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
func (client HubVirtualNetworkConnectionsClient) List(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListHubVirtualNetworkConnectionsResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualHubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lhvncr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result.lhvncr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client HubVirtualNetworkConnectionsClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client HubVirtualNetworkConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client HubVirtualNetworkConnectionsClient) ListResponder(resp *http.Response) (result ListHubVirtualNetworkConnectionsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client HubVirtualNetworkConnectionsClient) listNextResults(lastResults ListHubVirtualNetworkConnectionsResult) (result ListHubVirtualNetworkConnectionsResult, err error) {
	req, err := lastResults.listHubVirtualNetworkConnectionsResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client HubVirtualNetworkConnectionsClient) ListComplete(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListHubVirtualNetworkConnectionsResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, virtualHubName)
	return
}
