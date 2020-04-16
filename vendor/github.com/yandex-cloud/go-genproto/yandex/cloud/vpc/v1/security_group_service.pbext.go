// Code generated by protoc-gen-goext. DO NOT EDIT.

package vpc

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetSecurityGroupRequest) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *ListSecurityGroupsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListSecurityGroupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListSecurityGroupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListSecurityGroupsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListSecurityGroupsResponse) SetSecurityGroups(v []*SecurityGroup) {
	m.SecurityGroups = v
}

func (m *ListSecurityGroupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateSecurityGroupRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateSecurityGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateSecurityGroupRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateSecurityGroupRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateSecurityGroupRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *CreateSecurityGroupRequest) SetRuleSpecs(v []*SecurityGroupRuleSpec) {
	m.RuleSpecs = v
}

type SecurityGroupRuleSpec_Protocol = isSecurityGroupRuleSpec_Protocol

func (m *SecurityGroupRuleSpec) SetProtocol(v SecurityGroupRuleSpec_Protocol) {
	m.Protocol = v
}

type SecurityGroupRuleSpec_Target = isSecurityGroupRuleSpec_Target

func (m *SecurityGroupRuleSpec) SetTarget(v SecurityGroupRuleSpec_Target) {
	m.Target = v
}

func (m *SecurityGroupRuleSpec) SetDescription(v string) {
	m.Description = v
}

func (m *SecurityGroupRuleSpec) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *SecurityGroupRuleSpec) SetDirection(v SecurityGroupRule_Direction) {
	m.Direction = v
}

func (m *SecurityGroupRuleSpec) SetPorts(v *PortRange) {
	m.Ports = v
}

func (m *SecurityGroupRuleSpec) SetProtocolName(v string) {
	m.Protocol = &SecurityGroupRuleSpec_ProtocolName{
		ProtocolName: v,
	}
}

func (m *SecurityGroupRuleSpec) SetProtocolNumber(v int64) {
	m.Protocol = &SecurityGroupRuleSpec_ProtocolNumber{
		ProtocolNumber: v,
	}
}

func (m *SecurityGroupRuleSpec) SetCidrBlocks(v *CidrBlocks) {
	m.Target = &SecurityGroupRuleSpec_CidrBlocks{
		CidrBlocks: v,
	}
}

func (m *CreateSecurityGroupMetadata) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *UpdateSecurityGroupRequest) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *UpdateSecurityGroupRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateSecurityGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateSecurityGroupRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateSecurityGroupRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateSecurityGroupRequest) SetRuleSpecs(v []*SecurityGroupRuleSpec) {
	m.RuleSpecs = v
}

func (m *UpdateSecurityGroupMetadata) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *UpdateSecurityGroupRulesRequest) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *UpdateSecurityGroupRulesRequest) SetDeletionRuleIds(v []string) {
	m.DeletionRuleIds = v
}

func (m *UpdateSecurityGroupRulesRequest) SetAdditionRuleSpecs(v []*SecurityGroupRuleSpec) {
	m.AdditionRuleSpecs = v
}

func (m *UpdateSecurityGroupRuleRequest) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *UpdateSecurityGroupRuleRequest) SetRuleId(v string) {
	m.RuleId = v
}

func (m *UpdateSecurityGroupRuleRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateSecurityGroupRuleRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateSecurityGroupRuleRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateSecurityGroupRuleMetadata) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *UpdateSecurityGroupRuleMetadata) SetRuleId(v string) {
	m.RuleId = v
}

func (m *DeleteSecurityGroupRequest) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *DeleteSecurityGroupMetadata) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *ListSecurityGroupOperationsRequest) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *ListSecurityGroupOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListSecurityGroupOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListSecurityGroupOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListSecurityGroupOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *MoveSecurityGroupRequest) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}

func (m *MoveSecurityGroupRequest) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *MoveSecurityGroupMetadata) SetSecurityGroupId(v string) {
	m.SecurityGroupId = v
}