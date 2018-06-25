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

package reservations

import original "github.com/Azure/azure-sdk-for-go/services/reservations/mgmt/2018-06-01/reservations"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type AppliedScopeType = original.AppliedScopeType

const (
	Shared AppliedScopeType = original.Shared
	Single AppliedScopeType = original.Single
)

type AppliedScopeType1 = original.AppliedScopeType1

const (
	AppliedScopeType1Shared AppliedScopeType1 = original.AppliedScopeType1Shared
	AppliedScopeType1Single AppliedScopeType1 = original.AppliedScopeType1Single
)

type Code = original.Code

const (
	ActivateQuoteFailed                           Code = original.ActivateQuoteFailed
	AppliedScopesNotAssociatedWithCommerceAccount Code = original.AppliedScopesNotAssociatedWithCommerceAccount
	AuthorizationFailed                           Code = original.AuthorizationFailed
	BadRequest                                    Code = original.BadRequest
	BillingCustomerInputError                     Code = original.BillingCustomerInputError
	BillingError                                  Code = original.BillingError
	BillingPaymentInstrumentHardError             Code = original.BillingPaymentInstrumentHardError
	BillingPaymentInstrumentSoftError             Code = original.BillingPaymentInstrumentSoftError
	BillingScopeIDCannotBeChanged                 Code = original.BillingScopeIDCannotBeChanged
	BillingTransientError                         Code = original.BillingTransientError
	CalculatePriceFailed                          Code = original.CalculatePriceFailed
	CapacityUpdateScopesFailed                    Code = original.CapacityUpdateScopesFailed
	ClientCertificateThumbprintNotSet             Code = original.ClientCertificateThumbprintNotSet
	CreateQuoteFailed                             Code = original.CreateQuoteFailed
	Forbidden                                     Code = original.Forbidden
	FulfillmentConfigurationError                 Code = original.FulfillmentConfigurationError
	FulfillmentError                              Code = original.FulfillmentError
	FulfillmentOutOfStockError                    Code = original.FulfillmentOutOfStockError
	FulfillmentTransientError                     Code = original.FulfillmentTransientError
	HTTPMethodNotSupported                        Code = original.HTTPMethodNotSupported
	InternalServerError                           Code = original.InternalServerError
	InvalidAccessToken                            Code = original.InvalidAccessToken
	InvalidFulfillmentRequestParameters           Code = original.InvalidFulfillmentRequestParameters
	InvalidHealthCheckType                        Code = original.InvalidHealthCheckType
	InvalidLocationID                             Code = original.InvalidLocationID
	InvalidRefundQuantity                         Code = original.InvalidRefundQuantity
	InvalidRequestContent                         Code = original.InvalidRequestContent
	InvalidRequestURI                             Code = original.InvalidRequestURI
	InvalidReservationID                          Code = original.InvalidReservationID
	InvalidReservationOrderID                     Code = original.InvalidReservationOrderID
	InvalidSingleAppliedScopesCount               Code = original.InvalidSingleAppliedScopesCount
	InvalidSubscriptionID                         Code = original.InvalidSubscriptionID
	InvalidTenantID                               Code = original.InvalidTenantID
	MissingAppliedScopesForSingle                 Code = original.MissingAppliedScopesForSingle
	MissingTenantID                               Code = original.MissingTenantID
	NonsupportedAccountID                         Code = original.NonsupportedAccountID
	NotSpecified                                  Code = original.NotSpecified
	NotSupportedCountry                           Code = original.NotSupportedCountry
	NoValidReservationsToReRate                   Code = original.NoValidReservationsToReRate
	OperationCannotBePerformedInCurrentState      Code = original.OperationCannotBePerformedInCurrentState
	OperationFailed                               Code = original.OperationFailed
	PatchValuesSameAsExisting                     Code = original.PatchValuesSameAsExisting
	PaymentInstrumentNotFound                     Code = original.PaymentInstrumentNotFound
	PurchaseError                                 Code = original.PurchaseError
	ReRateOnlyAllowedForEA                        Code = original.ReRateOnlyAllowedForEA
	ReservationIDNotInReservationOrder            Code = original.ReservationIDNotInReservationOrder
	ReservationOrderCreationFailed                Code = original.ReservationOrderCreationFailed
	ReservationOrderIDAlreadyExists               Code = original.ReservationOrderIDAlreadyExists
	ReservationOrderNotEnabled                    Code = original.ReservationOrderNotEnabled
	ReservationOrderNotFound                      Code = original.ReservationOrderNotFound
	RiskCheckFailed                               Code = original.RiskCheckFailed
	RoleAssignmentCreationFailed                  Code = original.RoleAssignmentCreationFailed
	ServerTimeout                                 Code = original.ServerTimeout
	UnauthenticatedRequestsThrottled              Code = original.UnauthenticatedRequestsThrottled
	UnsupportedReservationTerm                    Code = original.UnsupportedReservationTerm
)

type InstanceFlexibility = original.InstanceFlexibility

const (
	NotSupported InstanceFlexibility = original.NotSupported
	Off          InstanceFlexibility = original.Off
	On           InstanceFlexibility = original.On
)

type InstanceFlexibility1 = original.InstanceFlexibility1

const (
	InstanceFlexibility1NotSupported InstanceFlexibility1 = original.InstanceFlexibility1NotSupported
	InstanceFlexibility1Off          InstanceFlexibility1 = original.InstanceFlexibility1Off
	InstanceFlexibility1On           InstanceFlexibility1 = original.InstanceFlexibility1On
)

type Location = original.Location

const (
	Australiac         Location = original.Australiac
	Australiac2        Location = original.Australiac2
	Australiaeast      Location = original.Australiaeast
	Australiasoutheast Location = original.Australiasoutheast
	Brazilsouth        Location = original.Brazilsouth
	Canadacentral      Location = original.Canadacentral
	Canadaeast         Location = original.Canadaeast
	Centralindia       Location = original.Centralindia
	Centralus          Location = original.Centralus
	Eastasia           Location = original.Eastasia
	Eastus             Location = original.Eastus
	Eastus2            Location = original.Eastus2
	Francecentral      Location = original.Francecentral
	Francesouth        Location = original.Francesouth
	Japaneast          Location = original.Japaneast
	Japanwest          Location = original.Japanwest
	Northcentralus     Location = original.Northcentralus
	Northeurope        Location = original.Northeurope
	Southcentralus     Location = original.Southcentralus
	Southeastasia      Location = original.Southeastasia
	Southindia         Location = original.Southindia
	Uksouth            Location = original.Uksouth
	Ukwest             Location = original.Ukwest
	Westcentralus      Location = original.Westcentralus
	Westeurope         Location = original.Westeurope
	Westindia          Location = original.Westindia
	Westus             Location = original.Westus
	Westus2            Location = original.Westus2
)

type ProvisioningState = original.ProvisioningState

const (
	BillingFailed         ProvisioningState = original.BillingFailed
	Cancelled             ProvisioningState = original.Cancelled
	ConfirmedBilling      ProvisioningState = original.ConfirmedBilling
	ConfirmedResourceHold ProvisioningState = original.ConfirmedResourceHold
	Created               ProvisioningState = original.Created
	Creating              ProvisioningState = original.Creating
	Expired               ProvisioningState = original.Expired
	Failed                ProvisioningState = original.Failed
	Merged                ProvisioningState = original.Merged
	PendingBilling        ProvisioningState = original.PendingBilling
	PendingResourceHold   ProvisioningState = original.PendingResourceHold
	Split                 ProvisioningState = original.Split
	Succeeded             ProvisioningState = original.Succeeded
)

type ProvisioningState1 = original.ProvisioningState1

const (
	ProvisioningState1BillingFailed         ProvisioningState1 = original.ProvisioningState1BillingFailed
	ProvisioningState1Cancelled             ProvisioningState1 = original.ProvisioningState1Cancelled
	ProvisioningState1ConfirmedBilling      ProvisioningState1 = original.ProvisioningState1ConfirmedBilling
	ProvisioningState1ConfirmedResourceHold ProvisioningState1 = original.ProvisioningState1ConfirmedResourceHold
	ProvisioningState1Created               ProvisioningState1 = original.ProvisioningState1Created
	ProvisioningState1Creating              ProvisioningState1 = original.ProvisioningState1Creating
	ProvisioningState1Expired               ProvisioningState1 = original.ProvisioningState1Expired
	ProvisioningState1Failed                ProvisioningState1 = original.ProvisioningState1Failed
	ProvisioningState1Merged                ProvisioningState1 = original.ProvisioningState1Merged
	ProvisioningState1PendingBilling        ProvisioningState1 = original.ProvisioningState1PendingBilling
	ProvisioningState1PendingResourceHold   ProvisioningState1 = original.ProvisioningState1PendingResourceHold
	ProvisioningState1Split                 ProvisioningState1 = original.ProvisioningState1Split
	ProvisioningState1Succeeded             ProvisioningState1 = original.ProvisioningState1Succeeded
)

type ReservedResourceType = original.ReservedResourceType

const (
	SQLDatabases    ReservedResourceType = original.SQLDatabases
	SuseLinux       ReservedResourceType = original.SuseLinux
	VirtualMachines ReservedResourceType = original.VirtualMachines
)

type StatusCode = original.StatusCode

const (
	StatusCodeActive                 StatusCode = original.StatusCodeActive
	StatusCodeExpired                StatusCode = original.StatusCodeExpired
	StatusCodeMerged                 StatusCode = original.StatusCodeMerged
	StatusCodeNone                   StatusCode = original.StatusCodeNone
	StatusCodePaymentInstrumentError StatusCode = original.StatusCodePaymentInstrumentError
	StatusCodePending                StatusCode = original.StatusCodePending
	StatusCodePurchaseError          StatusCode = original.StatusCodePurchaseError
	StatusCodeSplit                  StatusCode = original.StatusCodeSplit
	StatusCodeSucceeded              StatusCode = original.StatusCodeSucceeded
)

type Term = original.Term

const (
	P1Y Term = original.P1Y
	P3Y Term = original.P3Y
)

type AppliedReservationList = original.AppliedReservationList
type AppliedReservations = original.AppliedReservations
type AppliedReservationsProperties = original.AppliedReservationsProperties
type Catalog = original.Catalog
type Error = original.Error
type ExtendedErrorInfo = original.ExtendedErrorInfo
type ExtendedStatusInfo = original.ExtendedStatusInfo
type List = original.List
type ListCatalog = original.ListCatalog
type ListIterator = original.ListIterator
type ListPage = original.ListPage
type ListResponse = original.ListResponse
type MergeProperties = original.MergeProperties
type MergePropertiesType = original.MergePropertiesType
type MergeRequest = original.MergeRequest
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationResponse = original.OperationResponse
type OrderList = original.OrderList
type OrderListIterator = original.OrderListIterator
type OrderListPage = original.OrderListPage
type OrderProperties = original.OrderProperties
type OrderResponse = original.OrderResponse
type Patch = original.Patch
type PatchProperties = original.PatchProperties
type Properties = original.Properties
type ReservationMergeFuture = original.ReservationMergeFuture
type ReservationUpdateFuture = original.ReservationUpdateFuture
type Response = original.Response
type SkuName = original.SkuName
type SkuProperty = original.SkuProperty
type SkuRestriction = original.SkuRestriction
type SplitFuture = original.SplitFuture
type SplitProperties = original.SplitProperties
type SplitPropertiesType = original.SplitPropertiesType
type SplitRequest = original.SplitRequest
type OperationClient = original.OperationClient
type OrderClient = original.OrderClient
type Client = original.Client

func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleAppliedScopeTypeValues() []AppliedScopeType {
	return original.PossibleAppliedScopeTypeValues()
}
func PossibleAppliedScopeType1Values() []AppliedScopeType1 {
	return original.PossibleAppliedScopeType1Values()
}
func PossibleCodeValues() []Code {
	return original.PossibleCodeValues()
}
func PossibleInstanceFlexibilityValues() []InstanceFlexibility {
	return original.PossibleInstanceFlexibilityValues()
}
func PossibleInstanceFlexibility1Values() []InstanceFlexibility1 {
	return original.PossibleInstanceFlexibility1Values()
}
func PossibleLocationValues() []Location {
	return original.PossibleLocationValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleProvisioningState1Values() []ProvisioningState1 {
	return original.PossibleProvisioningState1Values()
}
func PossibleReservedResourceTypeValues() []ReservedResourceType {
	return original.PossibleReservedResourceTypeValues()
}
func PossibleStatusCodeValues() []StatusCode {
	return original.PossibleStatusCodeValues()
}
func PossibleTermValues() []Term {
	return original.PossibleTermValues()
}
func NewOperationClient() OperationClient {
	return original.NewOperationClient()
}
func NewOperationClientWithBaseURI(baseURI string) OperationClient {
	return original.NewOperationClientWithBaseURI(baseURI)
}
func NewOrderClient() OrderClient {
	return original.NewOrderClient()
}
func NewOrderClientWithBaseURI(baseURI string) OrderClient {
	return original.NewOrderClientWithBaseURI(baseURI)
}
func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}