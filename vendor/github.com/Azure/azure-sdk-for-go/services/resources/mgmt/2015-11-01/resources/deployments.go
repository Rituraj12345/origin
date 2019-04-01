package resources

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// DeploymentsClient is the client for the Deployments methods of the Resources service.
type DeploymentsClient struct {
	BaseClient
}

// NewDeploymentsClient creates an instance of the DeploymentsClient client.
func NewDeploymentsClient(subscriptionID string) DeploymentsClient {
	return NewDeploymentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDeploymentsClientWithBaseURI creates an instance of the DeploymentsClient client.
func NewDeploymentsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsClient {
	return DeploymentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancel a currently running template deployment.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// deploymentName - the name of the deployment.
func (client DeploymentsClient) Cancel(ctx context.Context, resourceGroupName string, deploymentName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentsClient", "Cancel", err.Error())
	}

	req, err := client.CancelPreparer(ctx, resourceGroupName, deploymentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Cancel", resp, "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client DeploymentsClient) CancelPreparer(ctx context.Context, resourceGroupName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) CancelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CheckExistence checks whether deployment exists.
// Parameters:
// resourceGroupName - the name of the resource group to check. The name is case insensitive.
// deploymentName - the name of the deployment.
func (client DeploymentsClient) CheckExistence(ctx context.Context, resourceGroupName string, deploymentName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentsClient", "CheckExistence", err.Error())
	}

	req, err := client.CheckExistencePreparer(ctx, resourceGroupName, deploymentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "CheckExistence", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckExistenceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "CheckExistence", resp, "Failure sending request")
		return
	}

	result, err = client.CheckExistenceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "CheckExistence", resp, "Failure responding to request")
	}

	return
}

// CheckExistencePreparer prepares the CheckExistence request.
func (client DeploymentsClient) CheckExistencePreparer(ctx context.Context, resourceGroupName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckExistenceSender sends the CheckExistence request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) CheckExistenceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CheckExistenceResponder handles the response to the CheckExistence request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) CheckExistenceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create a named template deployment using a template.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// deploymentName - the name of the deployment.
// parameters - additional parameters supplied to the operation.
func (client DeploymentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, deploymentName string, parameters Deployment) (result DeploymentsCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Properties.TemplateLink", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Properties.TemplateLink.URI", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.Properties.ParametersLink", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Properties.ParametersLink.URI", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, deploymentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DeploymentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, deploymentName string, parameters Deployment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) CreateOrUpdateSender(req *http.Request) (future DeploymentsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) CreateOrUpdateResponder(resp *http.Response) (result DeploymentExtended, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete begin deleting deployment.To determine whether the operation has finished processing the request, call
// GetLongRunningOperationStatus.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// deploymentName - the name of the deployment to be deleted.
func (client DeploymentsClient) Delete(ctx context.Context, resourceGroupName string, deploymentName string) (result DeploymentsDeleteFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, deploymentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DeploymentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) DeleteSender(req *http.Request) (future DeploymentsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a deployment.
// Parameters:
// resourceGroupName - the name of the resource group to get. The name is case insensitive.
// deploymentName - the name of the deployment.
func (client DeploymentsClient) Get(ctx context.Context, resourceGroupName string, deploymentName string) (result DeploymentExtended, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, deploymentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeploymentsClient) GetPreparer(ctx context.Context, resourceGroupName string, deploymentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) GetResponder(resp *http.Response) (result DeploymentExtended, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get a list of deployments.
// Parameters:
// resourceGroupName - the name of the resource group to filter by. The name is case insensitive.
// filter - the filter to apply on the operation.
// top - query parameters. If null is passed returns all deployments.
func (client DeploymentsClient) List(ctx context.Context, resourceGroupName string, filter string, top *int32) (result DeploymentListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "List", resp, "Failure sending request")
		return
	}

	result.dlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DeploymentsClient) ListPreparer(ctx context.Context, resourceGroupName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) ListResponder(resp *http.Response) (result DeploymentListResult, err error) {
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
func (client DeploymentsClient) listNextResults(lastResults DeploymentListResult) (result DeploymentListResult, err error) {
	req, err := lastResults.deploymentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.DeploymentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.DeploymentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeploymentsClient) ListComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result DeploymentListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, filter, top)
	return
}

// Validate validate a deployment template.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// deploymentName - the name of the deployment.
// parameters - deployment to validate.
func (client DeploymentsClient) Validate(ctx context.Context, resourceGroupName string, deploymentName string, parameters Deployment) (result DeploymentValidateResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Properties.TemplateLink", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Properties.TemplateLink.URI", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.Properties.ParametersLink", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Properties.ParametersLink.URI", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentsClient", "Validate", err.Error())
	}

	req, err := client.ValidatePreparer(ctx, resourceGroupName, deploymentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Validate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Validate", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentsClient", "Validate", resp, "Failure responding to request")
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client DeploymentsClient) ValidatePreparer(ctx context.Context, resourceGroupName string, deploymentName string, parameters Deployment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}/validate", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentsClient) ValidateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client DeploymentsClient) ValidateResponder(resp *http.Response) (result DeploymentValidateResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
