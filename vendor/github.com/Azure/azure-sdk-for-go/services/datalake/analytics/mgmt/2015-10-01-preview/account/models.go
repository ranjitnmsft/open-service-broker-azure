package account

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// DataLakeAnalyticsAccountState enumerates the values for data lake analytics account state.
type DataLakeAnalyticsAccountState string

const (
	// Active ...
	Active DataLakeAnalyticsAccountState = "active"
	// Suspended ...
	Suspended DataLakeAnalyticsAccountState = "suspended"
)

// DataLakeAnalyticsAccountStatus enumerates the values for data lake analytics account status.
type DataLakeAnalyticsAccountStatus string

const (
	// Creating ...
	Creating DataLakeAnalyticsAccountStatus = "Creating"
	// Deleted ...
	Deleted DataLakeAnalyticsAccountStatus = "Deleted"
	// Deleting ...
	Deleting DataLakeAnalyticsAccountStatus = "Deleting"
	// Failed ...
	Failed DataLakeAnalyticsAccountStatus = "Failed"
	// Patching ...
	Patching DataLakeAnalyticsAccountStatus = "Patching"
	// Resuming ...
	Resuming DataLakeAnalyticsAccountStatus = "Resuming"
	// Running ...
	Running DataLakeAnalyticsAccountStatus = "Running"
	// Succeeded ...
	Succeeded DataLakeAnalyticsAccountStatus = "Succeeded"
	// Suspending ...
	Suspending DataLakeAnalyticsAccountStatus = "Suspending"
)

// OperationStatus enumerates the values for operation status.
type OperationStatus string

const (
	// OperationStatusFailed ...
	OperationStatusFailed OperationStatus = "Failed"
	// OperationStatusInProgress ...
	OperationStatusInProgress OperationStatus = "InProgress"
	// OperationStatusSucceeded ...
	OperationStatusSucceeded OperationStatus = "Succeeded"
)

// AddDataLakeStoreParameters additional Data Lake Store parameters.
type AddDataLakeStoreParameters struct {
	// Properties - the properties for the Data Lake Store account being added.
	Properties *DataLakeStoreAccountInfoProperties `json:"properties,omitempty"`
}

// AddStorageAccountParameters additional Azure Storage account parameters.
type AddStorageAccountParameters struct {
	// Properties - the properties for the Azure Storage account being added.
	Properties *StorageAccountProperties `json:"properties,omitempty"`
}

// AzureAsyncOperationResult the response body contains the status of the specified asynchronous operation,
// indicating whether it has succeeded, is inprogress, or has failed. Note that this status is distinct from the
// HTTP status code returned for the Get Operation Status operation itself. If the asynchronous operation
// succeeded, the response body includes the HTTP status code for the successful request. If the asynchronous
// operation failed, the response body includes the HTTP status code for the failed request and error information
// regarding the failure.
type AzureAsyncOperationResult struct {
	// Status - the status of the AzureAsuncOperation. Possible values include: 'OperationStatusInProgress', 'OperationStatusSucceeded', 'OperationStatusFailed'
	Status OperationStatus `json:"status,omitempty"`
	Error  *Error          `json:"error,omitempty"`
}

// BlobContainer azure Storage blob container information.
type BlobContainer struct {
	autorest.Response `json:"-"`
	// Name - the name of the blob container.
	Name *string `json:"name,omitempty"`
	// ID - the unique identifier of the blob container.
	ID *string `json:"id,omitempty"`
	// Type - the type of the blob container.
	Type *string `json:"type,omitempty"`
	// Properties - the properties of the blob container.
	Properties *BlobContainerProperties `json:"properties,omitempty"`
}

// BlobContainerProperties azure Storage blob container properties information.
type BlobContainerProperties struct {
	// LastModifiedTime - the last modified time of the blob container.
	LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
}

// CreateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type CreateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future CreateFuture) Result(client Client) (dlaa DataLakeAnalyticsAccount, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.CreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return dlaa, azure.NewAsyncOpIncompleteError("account.CreateFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		dlaa, err = client.CreateResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "account.CreateFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.CreateFuture", "Result", resp, "Failure sending request")
		return
	}
	dlaa, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.CreateFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// DataLakeAnalyticsAccount a Data Lake Analytics account object, containing all information associated with the
// named Data Lake Analytics account.
type DataLakeAnalyticsAccount struct {
	autorest.Response `json:"-"`
	// Location - the account regional location.
	Location *string `json:"location,omitempty"`
	// Name - the account name.
	Name *string `json:"name,omitempty"`
	// Type - the namespace and type of the account.
	Type *string `json:"type,omitempty"`
	// ID - the account subscription ID.
	ID *string `json:"id,omitempty"`
	// Tags - the value of custom properties.
	Tags map[string]*string `json:"tags"`
	// Properties - the properties defined by Data Lake Analytics all properties are specific to each resource provider.
	Properties *DataLakeAnalyticsAccountProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for DataLakeAnalyticsAccount.
func (dlaa DataLakeAnalyticsAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dlaa.Location != nil {
		objectMap["location"] = dlaa.Location
	}
	if dlaa.Name != nil {
		objectMap["name"] = dlaa.Name
	}
	if dlaa.Type != nil {
		objectMap["type"] = dlaa.Type
	}
	if dlaa.ID != nil {
		objectMap["id"] = dlaa.ID
	}
	if dlaa.Tags != nil {
		objectMap["tags"] = dlaa.Tags
	}
	if dlaa.Properties != nil {
		objectMap["properties"] = dlaa.Properties
	}
	return json.Marshal(objectMap)
}

// DataLakeAnalyticsAccountListDataLakeStoreResult data Lake Account list information.
type DataLakeAnalyticsAccountListDataLakeStoreResult struct {
	autorest.Response `json:"-"`
	// Value - the results of the list operation
	Value *[]DataLakeStoreAccountInfo `json:"value,omitempty"`
	// Count - total number of results.
	Count *int32 `json:"count,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// DataLakeAnalyticsAccountListDataLakeStoreResultIterator provides access to a complete listing of
// DataLakeStoreAccountInfo values.
type DataLakeAnalyticsAccountListDataLakeStoreResultIterator struct {
	i    int
	page DataLakeAnalyticsAccountListDataLakeStoreResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DataLakeAnalyticsAccountListDataLakeStoreResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DataLakeAnalyticsAccountListDataLakeStoreResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DataLakeAnalyticsAccountListDataLakeStoreResultIterator) Response() DataLakeAnalyticsAccountListDataLakeStoreResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DataLakeAnalyticsAccountListDataLakeStoreResultIterator) Value() DataLakeStoreAccountInfo {
	if !iter.page.NotDone() {
		return DataLakeStoreAccountInfo{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (dlaaldlsr DataLakeAnalyticsAccountListDataLakeStoreResult) IsEmpty() bool {
	return dlaaldlsr.Value == nil || len(*dlaaldlsr.Value) == 0
}

// dataLakeAnalyticsAccountListDataLakeStoreResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dlaaldlsr DataLakeAnalyticsAccountListDataLakeStoreResult) dataLakeAnalyticsAccountListDataLakeStoreResultPreparer() (*http.Request, error) {
	if dlaaldlsr.NextLink == nil || len(to.String(dlaaldlsr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dlaaldlsr.NextLink)))
}

// DataLakeAnalyticsAccountListDataLakeStoreResultPage contains a page of DataLakeStoreAccountInfo values.
type DataLakeAnalyticsAccountListDataLakeStoreResultPage struct {
	fn        func(DataLakeAnalyticsAccountListDataLakeStoreResult) (DataLakeAnalyticsAccountListDataLakeStoreResult, error)
	dlaaldlsr DataLakeAnalyticsAccountListDataLakeStoreResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DataLakeAnalyticsAccountListDataLakeStoreResultPage) Next() error {
	next, err := page.fn(page.dlaaldlsr)
	if err != nil {
		return err
	}
	page.dlaaldlsr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DataLakeAnalyticsAccountListDataLakeStoreResultPage) NotDone() bool {
	return !page.dlaaldlsr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DataLakeAnalyticsAccountListDataLakeStoreResultPage) Response() DataLakeAnalyticsAccountListDataLakeStoreResult {
	return page.dlaaldlsr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DataLakeAnalyticsAccountListDataLakeStoreResultPage) Values() []DataLakeStoreAccountInfo {
	if page.dlaaldlsr.IsEmpty() {
		return nil
	}
	return *page.dlaaldlsr.Value
}

// DataLakeAnalyticsAccountListResult dataLakeAnalytics Account list information.
type DataLakeAnalyticsAccountListResult struct {
	autorest.Response `json:"-"`
	// Value - the results of the list operation
	Value *[]DataLakeAnalyticsAccount `json:"value,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// DataLakeAnalyticsAccountListResultIterator provides access to a complete listing of DataLakeAnalyticsAccount
// values.
type DataLakeAnalyticsAccountListResultIterator struct {
	i    int
	page DataLakeAnalyticsAccountListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DataLakeAnalyticsAccountListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DataLakeAnalyticsAccountListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DataLakeAnalyticsAccountListResultIterator) Response() DataLakeAnalyticsAccountListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DataLakeAnalyticsAccountListResultIterator) Value() DataLakeAnalyticsAccount {
	if !iter.page.NotDone() {
		return DataLakeAnalyticsAccount{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (dlaalr DataLakeAnalyticsAccountListResult) IsEmpty() bool {
	return dlaalr.Value == nil || len(*dlaalr.Value) == 0
}

// dataLakeAnalyticsAccountListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dlaalr DataLakeAnalyticsAccountListResult) dataLakeAnalyticsAccountListResultPreparer() (*http.Request, error) {
	if dlaalr.NextLink == nil || len(to.String(dlaalr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dlaalr.NextLink)))
}

// DataLakeAnalyticsAccountListResultPage contains a page of DataLakeAnalyticsAccount values.
type DataLakeAnalyticsAccountListResultPage struct {
	fn     func(DataLakeAnalyticsAccountListResult) (DataLakeAnalyticsAccountListResult, error)
	dlaalr DataLakeAnalyticsAccountListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DataLakeAnalyticsAccountListResultPage) Next() error {
	next, err := page.fn(page.dlaalr)
	if err != nil {
		return err
	}
	page.dlaalr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DataLakeAnalyticsAccountListResultPage) NotDone() bool {
	return !page.dlaalr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DataLakeAnalyticsAccountListResultPage) Response() DataLakeAnalyticsAccountListResult {
	return page.dlaalr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DataLakeAnalyticsAccountListResultPage) Values() []DataLakeAnalyticsAccount {
	if page.dlaalr.IsEmpty() {
		return nil
	}
	return *page.dlaalr.Value
}

// DataLakeAnalyticsAccountListStorageAccountsResult azure Storage Account list information.
type DataLakeAnalyticsAccountListStorageAccountsResult struct {
	autorest.Response `json:"-"`
	// Value - the results of the list operation
	Value *[]StorageAccountInfo `json:"value,omitempty"`
	// Count - total number of results.
	Count *int32 `json:"count,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// DataLakeAnalyticsAccountListStorageAccountsResultIterator provides access to a complete listing of
// StorageAccountInfo values.
type DataLakeAnalyticsAccountListStorageAccountsResultIterator struct {
	i    int
	page DataLakeAnalyticsAccountListStorageAccountsResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DataLakeAnalyticsAccountListStorageAccountsResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DataLakeAnalyticsAccountListStorageAccountsResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DataLakeAnalyticsAccountListStorageAccountsResultIterator) Response() DataLakeAnalyticsAccountListStorageAccountsResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DataLakeAnalyticsAccountListStorageAccountsResultIterator) Value() StorageAccountInfo {
	if !iter.page.NotDone() {
		return StorageAccountInfo{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (dlaalsar DataLakeAnalyticsAccountListStorageAccountsResult) IsEmpty() bool {
	return dlaalsar.Value == nil || len(*dlaalsar.Value) == 0
}

// dataLakeAnalyticsAccountListStorageAccountsResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dlaalsar DataLakeAnalyticsAccountListStorageAccountsResult) dataLakeAnalyticsAccountListStorageAccountsResultPreparer() (*http.Request, error) {
	if dlaalsar.NextLink == nil || len(to.String(dlaalsar.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dlaalsar.NextLink)))
}

// DataLakeAnalyticsAccountListStorageAccountsResultPage contains a page of StorageAccountInfo values.
type DataLakeAnalyticsAccountListStorageAccountsResultPage struct {
	fn       func(DataLakeAnalyticsAccountListStorageAccountsResult) (DataLakeAnalyticsAccountListStorageAccountsResult, error)
	dlaalsar DataLakeAnalyticsAccountListStorageAccountsResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DataLakeAnalyticsAccountListStorageAccountsResultPage) Next() error {
	next, err := page.fn(page.dlaalsar)
	if err != nil {
		return err
	}
	page.dlaalsar = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DataLakeAnalyticsAccountListStorageAccountsResultPage) NotDone() bool {
	return !page.dlaalsar.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DataLakeAnalyticsAccountListStorageAccountsResultPage) Response() DataLakeAnalyticsAccountListStorageAccountsResult {
	return page.dlaalsar
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DataLakeAnalyticsAccountListStorageAccountsResultPage) Values() []StorageAccountInfo {
	if page.dlaalsar.IsEmpty() {
		return nil
	}
	return *page.dlaalsar.Value
}

// DataLakeAnalyticsAccountProperties the account specific properties that are associated with an underlying Data
// Lake Analytics account.
type DataLakeAnalyticsAccountProperties struct {
	// ProvisioningState - the provisioning status of the Data Lake Analytics account. Possible values include: 'Failed', 'Creating', 'Running', 'Succeeded', 'Patching', 'Suspending', 'Resuming', 'Deleting', 'Deleted'
	ProvisioningState DataLakeAnalyticsAccountStatus `json:"provisioningState,omitempty"`
	// State - the state of the Data Lake Analytics account. Possible values include: 'Active', 'Suspended'
	State DataLakeAnalyticsAccountState `json:"state,omitempty"`
	// DefaultDataLakeStoreAccount - the default data lake storage account associated with this Data Lake Analytics account.
	DefaultDataLakeStoreAccount *string `json:"defaultDataLakeStoreAccount,omitempty"`
	// MaxDegreeOfParallelism - the maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism *int32 `json:"maxDegreeOfParallelism,omitempty"`
	// MaxJobCount - the maximum supported jobs running under the account at the same time.
	MaxJobCount *int32 `json:"maxJobCount,omitempty"`
	// DataLakeStoreAccounts - the list of Data Lake storage accounts associated with this account.
	DataLakeStoreAccounts *[]DataLakeStoreAccountInfo `json:"dataLakeStoreAccounts,omitempty"`
	// StorageAccounts - the list of Azure Blob storage accounts associated with this account.
	StorageAccounts *[]StorageAccountInfo `json:"storageAccounts,omitempty"`
	// CreationTime - the account creation time.
	CreationTime *date.Time `json:"creationTime,omitempty"`
	// LastModifiedTime - the account last modified time.
	LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
	// Endpoint - the full CName endpoint for this account.
	Endpoint *string `json:"endpoint,omitempty"`
}

// DataLakeStoreAccountInfo data Lake Store account information.
type DataLakeStoreAccountInfo struct {
	autorest.Response `json:"-"`
	// Name - the account name of the Data Lake Store account.
	Name *string `json:"name,omitempty"`
	// Properties - the properties associated with this Data Lake Store account.
	Properties *DataLakeStoreAccountInfoProperties `json:"properties,omitempty"`
}

// DataLakeStoreAccountInfoProperties data Lake Store account properties information.
type DataLakeStoreAccountInfoProperties struct {
	// Suffix - the optional suffix for the Data Lake Store account.
	Suffix *string `json:"suffix,omitempty"`
}

// DeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type DeleteFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future DeleteFuture) Result(client Client) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.DeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return ar, azure.NewAsyncOpIncompleteError("account.DeleteFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ar, err = client.DeleteResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "account.DeleteFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.DeleteFuture", "Result", resp, "Failure sending request")
		return
	}
	ar, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.DeleteFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// Error generic resource error information.
type Error struct {
	// Code - the HTTP status code or error code associated with this error
	Code *string `json:"code,omitempty"`
	// Message - the error message to display.
	Message *string `json:"message,omitempty"`
	// Target - the target of the error.
	Target *string `json:"target,omitempty"`
	// Details - the list of error details
	Details *[]ErrorDetails `json:"details,omitempty"`
	// InnerError - the inner exceptions or errors, if any
	InnerError *InnerError `json:"innerError,omitempty"`
}

// ErrorDetails generic resource error details information.
type ErrorDetails struct {
	// Code - the HTTP status code or error code associated with this error
	Code *string `json:"code,omitempty"`
	// Message - the error message localized based on Accept-Language
	Message *string `json:"message,omitempty"`
	// Target - the target of the particular error (for example, the name of the property in error).
	Target *string `json:"target,omitempty"`
}

// InnerError generic resource inner error information.
type InnerError struct {
	// Trace - the stack trace for the error
	Trace *string `json:"trace,omitempty"`
	// Context - the context for the error message
	Context *string `json:"context,omitempty"`
}

// ListBlobContainersResult the list of blob containers associated with the storage account attached to the Data
// Lake Analytics account.
type ListBlobContainersResult struct {
	autorest.Response `json:"-"`
	// Value - the results of the list operation
	Value *[]BlobContainer `json:"value,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListBlobContainersResultIterator provides access to a complete listing of BlobContainer values.
type ListBlobContainersResultIterator struct {
	i    int
	page ListBlobContainersResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListBlobContainersResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListBlobContainersResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListBlobContainersResultIterator) Response() ListBlobContainersResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListBlobContainersResultIterator) Value() BlobContainer {
	if !iter.page.NotDone() {
		return BlobContainer{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (lbcr ListBlobContainersResult) IsEmpty() bool {
	return lbcr.Value == nil || len(*lbcr.Value) == 0
}

// listBlobContainersResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lbcr ListBlobContainersResult) listBlobContainersResultPreparer() (*http.Request, error) {
	if lbcr.NextLink == nil || len(to.String(lbcr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lbcr.NextLink)))
}

// ListBlobContainersResultPage contains a page of BlobContainer values.
type ListBlobContainersResultPage struct {
	fn   func(ListBlobContainersResult) (ListBlobContainersResult, error)
	lbcr ListBlobContainersResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListBlobContainersResultPage) Next() error {
	next, err := page.fn(page.lbcr)
	if err != nil {
		return err
	}
	page.lbcr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListBlobContainersResultPage) NotDone() bool {
	return !page.lbcr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListBlobContainersResultPage) Response() ListBlobContainersResult {
	return page.lbcr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListBlobContainersResultPage) Values() []BlobContainer {
	if page.lbcr.IsEmpty() {
		return nil
	}
	return *page.lbcr.Value
}

// ListSasTokensResult the SAS response that contains the storage account, container and associated SAS token for
// connection use.
type ListSasTokensResult struct {
	autorest.Response `json:"-"`
	Value             *[]SasTokenInfo `json:"value,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListSasTokensResultIterator provides access to a complete listing of SasTokenInfo values.
type ListSasTokensResultIterator struct {
	i    int
	page ListSasTokensResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListSasTokensResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListSasTokensResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListSasTokensResultIterator) Response() ListSasTokensResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListSasTokensResultIterator) Value() SasTokenInfo {
	if !iter.page.NotDone() {
		return SasTokenInfo{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (lstr ListSasTokensResult) IsEmpty() bool {
	return lstr.Value == nil || len(*lstr.Value) == 0
}

// listSasTokensResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lstr ListSasTokensResult) listSasTokensResultPreparer() (*http.Request, error) {
	if lstr.NextLink == nil || len(to.String(lstr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lstr.NextLink)))
}

// ListSasTokensResultPage contains a page of SasTokenInfo values.
type ListSasTokensResultPage struct {
	fn   func(ListSasTokensResult) (ListSasTokensResult, error)
	lstr ListSasTokensResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListSasTokensResultPage) Next() error {
	next, err := page.fn(page.lstr)
	if err != nil {
		return err
	}
	page.lstr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListSasTokensResultPage) NotDone() bool {
	return !page.lstr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListSasTokensResultPage) Response() ListSasTokensResult {
	return page.lstr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListSasTokensResultPage) Values() []SasTokenInfo {
	if page.lstr.IsEmpty() {
		return nil
	}
	return *page.lstr.Value
}

// SasTokenInfo SAS token information.
type SasTokenInfo struct {
	// AccessToken - the access token for the associated Azure Storage Container.
	AccessToken *string `json:"accessToken,omitempty"`
}

// StorageAccountInfo azure Storage account information.
type StorageAccountInfo struct {
	autorest.Response `json:"-"`
	// Name - the account name associated with the Azure storage account.
	Name *string `json:"name,omitempty"`
	// Properties - the properties associated with this storage account.
	Properties *StorageAccountProperties `json:"properties,omitempty"`
}

// StorageAccountProperties azure Storage account properties information.
type StorageAccountProperties struct {
	// AccessKey - the access key associated with this Azure Storage account that will be used to connect to it.
	AccessKey *string `json:"accessKey,omitempty"`
	// Suffix - the optional suffix for the Data Lake account.
	Suffix *string `json:"suffix,omitempty"`
}

// UpdateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type UpdateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future UpdateFuture) Result(client Client) (dlaa DataLakeAnalyticsAccount, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.UpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return dlaa, azure.NewAsyncOpIncompleteError("account.UpdateFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		dlaa, err = client.UpdateResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "account.UpdateFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.UpdateFuture", "Result", resp, "Failure sending request")
		return
	}
	dlaa, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.UpdateFuture", "Result", resp, "Failure responding to request")
	}
	return
}
