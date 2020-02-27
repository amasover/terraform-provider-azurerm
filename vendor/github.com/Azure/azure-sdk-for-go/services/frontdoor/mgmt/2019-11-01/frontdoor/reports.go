package frontdoor

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ReportsClient is the frontDoor Client
type ReportsClient struct {
	BaseClient
}

// NewReportsClient creates an instance of the ReportsClient client.
func NewReportsClient(subscriptionID string) ReportsClient {
	return NewReportsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReportsClientWithBaseURI creates an instance of the ReportsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewReportsClientWithBaseURI(baseURI string, subscriptionID string) ReportsClient {
	return ReportsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetLatencyScorecards sends the get latency scorecards request.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - the Profile identifier associated with the Tenant and Partner
// experimentName - the Experiment identifier associated with the Experiment
// aggregationInterval - the aggregation interval of the Latency Scorecard
// endDateTimeUTC - the end DateTime of the Latency Scorecard in UTC
// country - the country associated with the Latency Scorecard. Values are country ISO codes as specified here-
// https://www.iso.org/iso-3166-country-codes.html
func (client ReportsClient) GetLatencyScorecards(ctx context.Context, resourceGroupName string, profileName string, experimentName string, aggregationInterval LatencyScorecardAggregationInterval, endDateTimeUTC string, country string) (result LatencyScorecard, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReportsClient.GetLatencyScorecards")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: profileName,
			Constraints: []validation.Constraint{{Target: "profileName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: experimentName,
			Constraints: []validation.Constraint{{Target: "experimentName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.ReportsClient", "GetLatencyScorecards", err.Error())
	}

	req, err := client.GetLatencyScorecardsPreparer(ctx, resourceGroupName, profileName, experimentName, aggregationInterval, endDateTimeUTC, country)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.ReportsClient", "GetLatencyScorecards", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLatencyScorecardsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "frontdoor.ReportsClient", "GetLatencyScorecards", resp, "Failure sending request")
		return
	}

	result, err = client.GetLatencyScorecardsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.ReportsClient", "GetLatencyScorecards", resp, "Failure responding to request")
	}

	return
}

// GetLatencyScorecardsPreparer prepares the GetLatencyScorecards request.
func (client ReportsClient) GetLatencyScorecardsPreparer(ctx context.Context, resourceGroupName string, profileName string, experimentName string, aggregationInterval LatencyScorecardAggregationInterval, endDateTimeUTC string, country string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"experimentName":    autorest.Encode("path", experimentName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"aggregationInterval": autorest.Encode("query", aggregationInterval),
		"api-version":         APIVersion,
	}
	if len(endDateTimeUTC) > 0 {
		queryParameters["endDateTimeUTC"] = autorest.Encode("query", endDateTimeUTC)
	}
	if len(country) > 0 {
		queryParameters["country"] = autorest.Encode("query", country)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}/LatencyScorecard", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLatencyScorecardsSender sends the GetLatencyScorecards request. The method will close the
// http.Response Body if it receives an error.
func (client ReportsClient) GetLatencyScorecardsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetLatencyScorecardsResponder handles the response to the GetLatencyScorecards request. The method always
// closes the http.Response Body.
func (client ReportsClient) GetLatencyScorecardsResponder(resp *http.Response) (result LatencyScorecard, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTimeseries sends the get timeseries request.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - the Profile identifier associated with the Tenant and Partner
// experimentName - the Experiment identifier associated with the Experiment
// startDateTimeUTC - the start DateTime of the Timeseries in UTC
// endDateTimeUTC - the end DateTime of the Timeseries in UTC
// aggregationInterval - the aggregation interval of the Timeseries
// timeseriesType - the type of Timeseries
// endpoint - the specific endpoint
// country - the country associated with the Timeseries. Values are country ISO codes as specified here-
// https://www.iso.org/iso-3166-country-codes.html
func (client ReportsClient) GetTimeseries(ctx context.Context, resourceGroupName string, profileName string, experimentName string, startDateTimeUTC date.Time, endDateTimeUTC date.Time, aggregationInterval TimeseriesAggregationInterval, timeseriesType TimeseriesType, endpoint string, country string) (result Timeseries, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReportsClient.GetTimeseries")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: profileName,
			Constraints: []validation.Constraint{{Target: "profileName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: experimentName,
			Constraints: []validation.Constraint{{Target: "experimentName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.ReportsClient", "GetTimeseries", err.Error())
	}

	req, err := client.GetTimeseriesPreparer(ctx, resourceGroupName, profileName, experimentName, startDateTimeUTC, endDateTimeUTC, aggregationInterval, timeseriesType, endpoint, country)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.ReportsClient", "GetTimeseries", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTimeseriesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "frontdoor.ReportsClient", "GetTimeseries", resp, "Failure sending request")
		return
	}

	result, err = client.GetTimeseriesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.ReportsClient", "GetTimeseries", resp, "Failure responding to request")
	}

	return
}

// GetTimeseriesPreparer prepares the GetTimeseries request.
func (client ReportsClient) GetTimeseriesPreparer(ctx context.Context, resourceGroupName string, profileName string, experimentName string, startDateTimeUTC date.Time, endDateTimeUTC date.Time, aggregationInterval TimeseriesAggregationInterval, timeseriesType TimeseriesType, endpoint string, country string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"experimentName":    autorest.Encode("path", experimentName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"aggregationInterval": autorest.Encode("query", aggregationInterval),
		"api-version":         APIVersion,
		"endDateTimeUTC":      autorest.Encode("query", endDateTimeUTC),
		"startDateTimeUTC":    autorest.Encode("query", startDateTimeUTC),
		"timeseriesType":      autorest.Encode("query", timeseriesType),
	}
	if len(endpoint) > 0 {
		queryParameters["endpoint"] = autorest.Encode("query", endpoint)
	}
	if len(country) > 0 {
		queryParameters["country"] = autorest.Encode("query", country)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}/Timeseries", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTimeseriesSender sends the GetTimeseries request. The method will close the
// http.Response Body if it receives an error.
func (client ReportsClient) GetTimeseriesSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetTimeseriesResponder handles the response to the GetTimeseries request. The method always
// closes the http.Response Body.
func (client ReportsClient) GetTimeseriesResponder(resp *http.Response) (result Timeseries, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
