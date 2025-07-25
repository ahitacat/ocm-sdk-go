/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	api_v1 "github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1"
)

// AzureControlPlaneManagedIdentity represents the values of the 'azure_control_plane_managed_identity' type.
//
// Represents the information associated to an Azure User-Assigned
// Managed Identity belonging to the Control Plane of the cluster.
type AzureControlPlaneManagedIdentity = api_v1.AzureControlPlaneManagedIdentity

// AzureControlPlaneManagedIdentityListKind is the name of the type used to represent list of objects of
// type 'azure_control_plane_managed_identity'.
const AzureControlPlaneManagedIdentityListKind = api_v1.AzureControlPlaneManagedIdentityListKind

// AzureControlPlaneManagedIdentityListLinkKind is the name of the type used to represent links to list
// of objects of type 'azure_control_plane_managed_identity'.
const AzureControlPlaneManagedIdentityListLinkKind = api_v1.AzureControlPlaneManagedIdentityListLinkKind

// AzureControlPlaneManagedIdentityNilKind is the name of the type used to nil lists of objects of
// type 'azure_control_plane_managed_identity'.
const AzureControlPlaneManagedIdentityListNilKind = api_v1.AzureControlPlaneManagedIdentityListNilKind

type AzureControlPlaneManagedIdentityList = api_v1.AzureControlPlaneManagedIdentityList
