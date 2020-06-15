/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha3

import clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"

const (
	// InstanceReadyCondition reports on current status of the EC2 instance. Ready indicates the instance is in a Running state.
	InstanceReadyCondition clusterv1.ConditionType = "InstanceReady"

	// InstanceNotFoundReason used when the instance couldn't be retrieved.
	InstanceNotFoundReason = "InstanceNotFound"
	// InstanceTerminatedReason instance is in a terminated state.
	InstanceTerminatedReason = "InstanceTerminated"
	// InstanceStoppedReason instance is in a stopped state.
	InstanceStoppedReason = "InstanceStopped"
	// InstanceNotReadyReason used when the instance is in a pending state.
	InstanceNotReadyReason = "InstanceNotReady"
	// InstanceProvisionFailedReason used for failures during instance provisioning.
	InstanceProvisionFailedReason = "InstanceProvisionFailed"
	// WaitingForClusterInfrastructureReason used when machine is waiting for cluster infrastructure to be ready before proceeding.
	WaitingForClusterInfrastructureReason = "WaitingForClusterInfrastructure"
	// WaitingForBootstrapDataReason used when machine is waiting for bootstrap data to be ready before proceeding.
	WaitingForBootstrapDataReason = "WaitingForBootstrapData"
)

const (
	// SecurityGroupsReadyCondition indicates the security groups are up to date on the AWSMachine.
	SecurityGroupsReadyCondition clusterv1.ConditionType = "SecurityGroupsReady"

	// SecurityGroupsFailedReason used when the security groups could not be synced
	SecurityGroupsFailedReason = "SecurityGroupsSyncFailed"
)

const (
	// Only applicable to control plane machines. ELBAttachedCondition will report true when a control plane is successfully registered with an ELB
	// When set to false, severity can be an Error if the subnet is not found or unavailable in the instance's AZ
	ELBAttachedCondition clusterv1.ConditionType = "ELBAttached"

	// ELBAttachFailedReason used when a control plane node fails to attach to the ELB
	ELBAttachFailedReason = "ELBAttachFailed"
	// ELBDetachFailedReason used when a control plane node fails to detach from an ELB
	ELBDetachFailedReason = "ELBDetachFailed"
)
