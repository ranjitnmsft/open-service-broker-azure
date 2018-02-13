// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package documentdb

import original "github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type CollectionClient = original.CollectionClient

func NewCollectionClient(subscriptionID string) CollectionClient {
	return original.NewCollectionClient(subscriptionID)
}
func NewCollectionClientWithBaseURI(baseURI string, subscriptionID string) CollectionClient {
	return original.NewCollectionClientWithBaseURI(baseURI, subscriptionID)
}

type CollectionPartitionClient = original.CollectionPartitionClient

func NewCollectionPartitionClient(subscriptionID string) CollectionPartitionClient {
	return original.NewCollectionPartitionClient(subscriptionID)
}
func NewCollectionPartitionClientWithBaseURI(baseURI string, subscriptionID string) CollectionPartitionClient {
	return original.NewCollectionPartitionClientWithBaseURI(baseURI, subscriptionID)
}

type CollectionPartitionRegionClient = original.CollectionPartitionRegionClient

func NewCollectionPartitionRegionClient(subscriptionID string) CollectionPartitionRegionClient {
	return original.NewCollectionPartitionRegionClient(subscriptionID)
}
func NewCollectionPartitionRegionClientWithBaseURI(baseURI string, subscriptionID string) CollectionPartitionRegionClient {
	return original.NewCollectionPartitionRegionClientWithBaseURI(baseURI, subscriptionID)
}

type CollectionRegionClient = original.CollectionRegionClient

func NewCollectionRegionClient(subscriptionID string) CollectionRegionClient {
	return original.NewCollectionRegionClient(subscriptionID)
}
func NewCollectionRegionClientWithBaseURI(baseURI string, subscriptionID string) CollectionRegionClient {
	return original.NewCollectionRegionClientWithBaseURI(baseURI, subscriptionID)
}

type DatabaseClient = original.DatabaseClient

func NewDatabaseClient(subscriptionID string) DatabaseClient {
	return original.NewDatabaseClient(subscriptionID)
}
func NewDatabaseClientWithBaseURI(baseURI string, subscriptionID string) DatabaseClient {
	return original.NewDatabaseClientWithBaseURI(baseURI, subscriptionID)
}

type DatabaseAccountRegionClient = original.DatabaseAccountRegionClient

func NewDatabaseAccountRegionClient(subscriptionID string) DatabaseAccountRegionClient {
	return original.NewDatabaseAccountRegionClient(subscriptionID)
}
func NewDatabaseAccountRegionClientWithBaseURI(baseURI string, subscriptionID string) DatabaseAccountRegionClient {
	return original.NewDatabaseAccountRegionClientWithBaseURI(baseURI, subscriptionID)
}

type DatabaseAccountsClient = original.DatabaseAccountsClient

func NewDatabaseAccountsClient(subscriptionID string) DatabaseAccountsClient {
	return original.NewDatabaseAccountsClient(subscriptionID)
}
func NewDatabaseAccountsClientWithBaseURI(baseURI string, subscriptionID string) DatabaseAccountsClient {
	return original.NewDatabaseAccountsClientWithBaseURI(baseURI, subscriptionID)
}

type DatabaseAccountKind = original.DatabaseAccountKind

const (
	GlobalDocumentDB DatabaseAccountKind = original.GlobalDocumentDB
	MongoDB          DatabaseAccountKind = original.MongoDB
	Parse            DatabaseAccountKind = original.Parse
)

type DatabaseAccountOfferType = original.DatabaseAccountOfferType

const (
	Standard DatabaseAccountOfferType = original.Standard
)

type DefaultConsistencyLevel = original.DefaultConsistencyLevel

const (
	BoundedStaleness DefaultConsistencyLevel = original.BoundedStaleness
	ConsistentPrefix DefaultConsistencyLevel = original.ConsistentPrefix
	Eventual         DefaultConsistencyLevel = original.Eventual
	Session          DefaultConsistencyLevel = original.Session
	Strong           DefaultConsistencyLevel = original.Strong
)

type KeyKind = original.KeyKind

const (
	Primary           KeyKind = original.Primary
	PrimaryReadonly   KeyKind = original.PrimaryReadonly
	Secondary         KeyKind = original.Secondary
	SecondaryReadonly KeyKind = original.SecondaryReadonly
)

type PrimaryAggregationType = original.PrimaryAggregationType

const (
	Average   PrimaryAggregationType = original.Average
	Last      PrimaryAggregationType = original.Last
	Maximum   PrimaryAggregationType = original.Maximum
	Minimimum PrimaryAggregationType = original.Minimimum
	None      PrimaryAggregationType = original.None
	Total     PrimaryAggregationType = original.Total
)

type UnitType = original.UnitType

const (
	Bytes          UnitType = original.Bytes
	BytesPerSecond UnitType = original.BytesPerSecond
	Count          UnitType = original.Count
	CountPerSecond UnitType = original.CountPerSecond
	Milliseconds   UnitType = original.Milliseconds
	Percent        UnitType = original.Percent
	Seconds        UnitType = original.Seconds
)

type Capability = original.Capability
type ConsistencyPolicy = original.ConsistencyPolicy
type DatabaseAccount = original.DatabaseAccount
type DatabaseAccountConnectionString = original.DatabaseAccountConnectionString
type DatabaseAccountCreateUpdateParameters = original.DatabaseAccountCreateUpdateParameters
type DatabaseAccountCreateUpdateProperties = original.DatabaseAccountCreateUpdateProperties
type DatabaseAccountListConnectionStringsResult = original.DatabaseAccountListConnectionStringsResult
type DatabaseAccountListKeysResult = original.DatabaseAccountListKeysResult
type DatabaseAccountListReadOnlyKeysResult = original.DatabaseAccountListReadOnlyKeysResult
type DatabaseAccountPatchParameters = original.DatabaseAccountPatchParameters
type DatabaseAccountPatchProperties = original.DatabaseAccountPatchProperties
type DatabaseAccountProperties = original.DatabaseAccountProperties
type DatabaseAccountRegenerateKeyParameters = original.DatabaseAccountRegenerateKeyParameters
type DatabaseAccountsCreateOrUpdateFuture = original.DatabaseAccountsCreateOrUpdateFuture
type DatabaseAccountsDeleteFuture = original.DatabaseAccountsDeleteFuture
type DatabaseAccountsFailoverPriorityChangeFuture = original.DatabaseAccountsFailoverPriorityChangeFuture
type DatabaseAccountsListResult = original.DatabaseAccountsListResult
type DatabaseAccountsPatchFuture = original.DatabaseAccountsPatchFuture
type DatabaseAccountsRegenerateKeyFuture = original.DatabaseAccountsRegenerateKeyFuture
type FailoverPolicies = original.FailoverPolicies
type FailoverPolicy = original.FailoverPolicy
type Location = original.Location
type Metric = original.Metric
type MetricAvailability = original.MetricAvailability
type MetricDefinition = original.MetricDefinition
type MetricDefinitionsListResult = original.MetricDefinitionsListResult
type MetricListResult = original.MetricListResult
type MetricName = original.MetricName
type MetricValue = original.MetricValue
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type PartitionMetric = original.PartitionMetric
type PartitionMetricListResult = original.PartitionMetricListResult
type PartitionUsage = original.PartitionUsage
type PartitionUsagesResult = original.PartitionUsagesResult
type PercentileMetric = original.PercentileMetric
type PercentileMetricListResult = original.PercentileMetricListResult
type PercentileMetricValue = original.PercentileMetricValue
type Resource = original.Resource
type Usage = original.Usage
type UsagesResult = original.UsagesResult
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type PartitionKeyRangeIDClient = original.PartitionKeyRangeIDClient

func NewPartitionKeyRangeIDClient(subscriptionID string) PartitionKeyRangeIDClient {
	return original.NewPartitionKeyRangeIDClient(subscriptionID)
}
func NewPartitionKeyRangeIDClientWithBaseURI(baseURI string, subscriptionID string) PartitionKeyRangeIDClient {
	return original.NewPartitionKeyRangeIDClientWithBaseURI(baseURI, subscriptionID)
}

type PartitionKeyRangeIDRegionClient = original.PartitionKeyRangeIDRegionClient

func NewPartitionKeyRangeIDRegionClient(subscriptionID string) PartitionKeyRangeIDRegionClient {
	return original.NewPartitionKeyRangeIDRegionClient(subscriptionID)
}
func NewPartitionKeyRangeIDRegionClientWithBaseURI(baseURI string, subscriptionID string) PartitionKeyRangeIDRegionClient {
	return original.NewPartitionKeyRangeIDRegionClientWithBaseURI(baseURI, subscriptionID)
}

type PercentileClient = original.PercentileClient

func NewPercentileClient(subscriptionID string) PercentileClient {
	return original.NewPercentileClient(subscriptionID)
}
func NewPercentileClientWithBaseURI(baseURI string, subscriptionID string) PercentileClient {
	return original.NewPercentileClientWithBaseURI(baseURI, subscriptionID)
}

type PercentileSourceTargetClient = original.PercentileSourceTargetClient

func NewPercentileSourceTargetClient(subscriptionID string) PercentileSourceTargetClient {
	return original.NewPercentileSourceTargetClient(subscriptionID)
}
func NewPercentileSourceTargetClientWithBaseURI(baseURI string, subscriptionID string) PercentileSourceTargetClient {
	return original.NewPercentileSourceTargetClientWithBaseURI(baseURI, subscriptionID)
}

type PercentileTargetClient = original.PercentileTargetClient

func NewPercentileTargetClient(subscriptionID string) PercentileTargetClient {
	return original.NewPercentileTargetClient(subscriptionID)
}
func NewPercentileTargetClientWithBaseURI(baseURI string, subscriptionID string) PercentileTargetClient {
	return original.NewPercentileTargetClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
