package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListPoliciesForTargetRequest represents the ListPoliciesForTargetRequest schema from the OpenAPI specification
type ListPoliciesForTargetRequest struct {
	Targetid interface{} `json:"TargetId"`
	Filter interface{} `json:"Filter"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Child represents the Child schema from the OpenAPI specification
type Child struct {
	Id interface{} `json:"Id,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// CreateGovCloudAccountRequest represents the CreateGovCloudAccountRequest schema from the OpenAPI specification
type CreateGovCloudAccountRequest struct {
	Accountname interface{} `json:"AccountName"`
	Email interface{} `json:"Email"`
	Iamuseraccesstobilling interface{} `json:"IamUserAccessToBilling,omitempty"`
	Rolename interface{} `json:"RoleName,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateAccountStatus represents the CreateAccountStatus schema from the OpenAPI specification
type CreateAccountStatus struct {
	State interface{} `json:"State,omitempty"`
	Accountid interface{} `json:"AccountId,omitempty"`
	Accountname interface{} `json:"AccountName,omitempty"`
	Completedtimestamp interface{} `json:"CompletedTimestamp,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Govcloudaccountid interface{} `json:"GovCloudAccountId,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Requestedtimestamp interface{} `json:"RequestedTimestamp,omitempty"`
}

// DescribeAccountResponse represents the DescribeAccountResponse schema from the OpenAPI specification
type DescribeAccountResponse struct {
	Account interface{} `json:"Account,omitempty"`
}

// CancelHandshakeResponse represents the CancelHandshakeResponse schema from the OpenAPI specification
type CancelHandshakeResponse struct {
	Handshake interface{} `json:"Handshake,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourceid interface{} `json:"ResourceId"`
	Tagkeys interface{} `json:"TagKeys"`
}

// CreatePolicyResponse represents the CreatePolicyResponse schema from the OpenAPI specification
type CreatePolicyResponse struct {
	Policy interface{} `json:"Policy,omitempty"`
}

// ListPoliciesResponse represents the ListPoliciesResponse schema from the OpenAPI specification
type ListPoliciesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Policies interface{} `json:"Policies,omitempty"`
}

// ListParentsResponse represents the ListParentsResponse schema from the OpenAPI specification
type ListParentsResponse struct {
	Parents interface{} `json:"Parents,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// EnableAllFeaturesRequest represents the EnableAllFeaturesRequest schema from the OpenAPI specification
type EnableAllFeaturesRequest struct {
}

// Handshake represents the Handshake schema from the OpenAPI specification
type Handshake struct {
	Arn interface{} `json:"Arn,omitempty"`
	Expirationtimestamp interface{} `json:"ExpirationTimestamp,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Parties interface{} `json:"Parties,omitempty"`
	Requestedtimestamp interface{} `json:"RequestedTimestamp,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	State interface{} `json:"State,omitempty"`
	Action interface{} `json:"Action,omitempty"`
}

// DeleteOrganizationalUnitRequest represents the DeleteOrganizationalUnitRequest schema from the OpenAPI specification
type DeleteOrganizationalUnitRequest struct {
	Organizationalunitid interface{} `json:"OrganizationalUnitId"`
}

// EnablePolicyTypeResponse represents the EnablePolicyTypeResponse schema from the OpenAPI specification
type EnablePolicyTypeResponse struct {
	Root interface{} `json:"Root,omitempty"`
}

// DeregisterDelegatedAdministratorRequest represents the DeregisterDelegatedAdministratorRequest schema from the OpenAPI specification
type DeregisterDelegatedAdministratorRequest struct {
	Serviceprincipal interface{} `json:"ServicePrincipal"`
	Accountid interface{} `json:"AccountId"`
}

// PutResourcePolicyResponse represents the PutResourcePolicyResponse schema from the OpenAPI specification
type PutResourcePolicyResponse struct {
	Resourcepolicy interface{} `json:"ResourcePolicy,omitempty"`
}

// CloseAccountRequest represents the CloseAccountRequest schema from the OpenAPI specification
type CloseAccountRequest struct {
	Accountid interface{} `json:"AccountId"`
}

// CreateAccountResponse represents the CreateAccountResponse schema from the OpenAPI specification
type CreateAccountResponse struct {
	Createaccountstatus interface{} `json:"CreateAccountStatus,omitempty"`
}

// PolicyTypeSummary represents the PolicyTypeSummary schema from the OpenAPI specification
type PolicyTypeSummary struct {
	Status interface{} `json:"Status,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// DisableAWSServiceAccessRequest represents the DisableAWSServiceAccessRequest schema from the OpenAPI specification
type DisableAWSServiceAccessRequest struct {
	Serviceprincipal interface{} `json:"ServicePrincipal"`
}

// DeletePolicyRequest represents the DeletePolicyRequest schema from the OpenAPI specification
type DeletePolicyRequest struct {
	Policyid interface{} `json:"PolicyId"`
}

// ListChildrenResponse represents the ListChildrenResponse schema from the OpenAPI specification
type ListChildrenResponse struct {
	Children interface{} `json:"Children,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeOrganizationalUnitResponse represents the DescribeOrganizationalUnitResponse schema from the OpenAPI specification
type DescribeOrganizationalUnitResponse struct {
	Organizationalunit interface{} `json:"OrganizationalUnit,omitempty"`
}

// UpdateOrganizationalUnitRequest represents the UpdateOrganizationalUnitRequest schema from the OpenAPI specification
type UpdateOrganizationalUnitRequest struct {
	Name interface{} `json:"Name,omitempty"`
	Organizationalunitid interface{} `json:"OrganizationalUnitId"`
}

// DescribePolicyResponse represents the DescribePolicyResponse schema from the OpenAPI specification
type DescribePolicyResponse struct {
	Policy interface{} `json:"Policy,omitempty"`
}

// Parent represents the Parent schema from the OpenAPI specification
type Parent struct {
	TypeField interface{} `json:"Type,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// DelegatedAdministrator represents the DelegatedAdministrator schema from the OpenAPI specification
type DelegatedAdministrator struct {
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Delegationenableddate interface{} `json:"DelegationEnabledDate,omitempty"`
	Email interface{} `json:"Email,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Joinedmethod interface{} `json:"JoinedMethod,omitempty"`
	Joinedtimestamp interface{} `json:"JoinedTimestamp,omitempty"`
}

// DetachPolicyRequest represents the DetachPolicyRequest schema from the OpenAPI specification
type DetachPolicyRequest struct {
	Policyid interface{} `json:"PolicyId"`
	Targetid interface{} `json:"TargetId"`
}

// CreateOrganizationalUnitRequest represents the CreateOrganizationalUnitRequest schema from the OpenAPI specification
type CreateOrganizationalUnitRequest struct {
	Parentid interface{} `json:"ParentId"`
	Tags interface{} `json:"Tags,omitempty"`
	Name interface{} `json:"Name"`
}

// MoveAccountRequest represents the MoveAccountRequest schema from the OpenAPI specification
type MoveAccountRequest struct {
	Accountid interface{} `json:"AccountId"`
	Destinationparentid interface{} `json:"DestinationParentId"`
	Sourceparentid interface{} `json:"SourceParentId"`
}

// PolicySummary represents the PolicySummary schema from the OpenAPI specification
type PolicySummary struct {
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Awsmanaged interface{} `json:"AwsManaged,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// ListDelegatedServicesForAccountResponse represents the ListDelegatedServicesForAccountResponse schema from the OpenAPI specification
type ListDelegatedServicesForAccountResponse struct {
	Delegatedservices interface{} `json:"DelegatedServices,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// PolicyTargetSummary represents the PolicyTargetSummary schema from the OpenAPI specification
type PolicyTargetSummary struct {
	Targetid interface{} `json:"TargetId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// ListDelegatedServicesForAccountRequest represents the ListDelegatedServicesForAccountRequest schema from the OpenAPI specification
type ListDelegatedServicesForAccountRequest struct {
	Accountid interface{} `json:"AccountId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateOrganizationResponse represents the CreateOrganizationResponse schema from the OpenAPI specification
type CreateOrganizationResponse struct {
	Organization interface{} `json:"Organization,omitempty"`
}

// CreateOrganizationRequest represents the CreateOrganizationRequest schema from the OpenAPI specification
type CreateOrganizationRequest struct {
	Featureset interface{} `json:"FeatureSet,omitempty"`
}

// ListHandshakesForOrganizationResponse represents the ListHandshakesForOrganizationResponse schema from the OpenAPI specification
type ListHandshakesForOrganizationResponse struct {
	Handshakes interface{} `json:"Handshakes,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeCreateAccountStatusRequest represents the DescribeCreateAccountStatusRequest schema from the OpenAPI specification
type DescribeCreateAccountStatusRequest struct {
	Createaccountrequestid interface{} `json:"CreateAccountRequestId"`
}

// ListAccountsResponse represents the ListAccountsResponse schema from the OpenAPI specification
type ListAccountsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Accounts interface{} `json:"Accounts,omitempty"`
}

// ListHandshakesForAccountRequest represents the ListHandshakesForAccountRequest schema from the OpenAPI specification
type ListHandshakesForAccountRequest struct {
	Filter interface{} `json:"Filter,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RemoveAccountFromOrganizationRequest represents the RemoveAccountFromOrganizationRequest schema from the OpenAPI specification
type RemoveAccountFromOrganizationRequest struct {
	Accountid interface{} `json:"AccountId"`
}

// Account represents the Account schema from the OpenAPI specification
type Account struct {
	Email interface{} `json:"Email,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Joinedmethod interface{} `json:"JoinedMethod,omitempty"`
	Joinedtimestamp interface{} `json:"JoinedTimestamp,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// CreatePolicyRequest represents the CreatePolicyRequest schema from the OpenAPI specification
type CreatePolicyRequest struct {
	Content interface{} `json:"Content"`
	Description interface{} `json:"Description"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	TypeField interface{} `json:"Type"`
}

// ListHandshakesForAccountResponse represents the ListHandshakesForAccountResponse schema from the OpenAPI specification
type ListHandshakesForAccountResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Handshakes interface{} `json:"Handshakes,omitempty"`
}

// EnablePolicyTypeRequest represents the EnablePolicyTypeRequest schema from the OpenAPI specification
type EnablePolicyTypeRequest struct {
	Policytype interface{} `json:"PolicyType"`
	Rootid interface{} `json:"RootId"`
}

// HandshakeResource represents the HandshakeResource schema from the OpenAPI specification
type HandshakeResource struct {
	Resources interface{} `json:"Resources,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AttachPolicyRequest represents the AttachPolicyRequest schema from the OpenAPI specification
type AttachPolicyRequest struct {
	Policyid interface{} `json:"PolicyId"`
	Targetid interface{} `json:"TargetId"`
}

// Policy represents the Policy schema from the OpenAPI specification
type Policy struct {
	Content interface{} `json:"Content,omitempty"`
	Policysummary interface{} `json:"PolicySummary,omitempty"`
}

// ResourcePolicy represents the ResourcePolicy schema from the OpenAPI specification
type ResourcePolicy struct {
	Content interface{} `json:"Content,omitempty"`
	Resourcepolicysummary interface{} `json:"ResourcePolicySummary,omitempty"`
}

// ListCreateAccountStatusResponse represents the ListCreateAccountStatusResponse schema from the OpenAPI specification
type ListCreateAccountStatusResponse struct {
	Createaccountstatuses interface{} `json:"CreateAccountStatuses,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// InviteAccountToOrganizationRequest represents the InviteAccountToOrganizationRequest schema from the OpenAPI specification
type InviteAccountToOrganizationRequest struct {
	Notes interface{} `json:"Notes,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Target interface{} `json:"Target"`
}

// ListAccountsForParentRequest represents the ListAccountsForParentRequest schema from the OpenAPI specification
type ListAccountsForParentRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Parentid interface{} `json:"ParentId"`
}

// ListAWSServiceAccessForOrganizationResponse represents the ListAWSServiceAccessForOrganizationResponse schema from the OpenAPI specification
type ListAWSServiceAccessForOrganizationResponse struct {
	Enabledserviceprincipals interface{} `json:"EnabledServicePrincipals,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// EffectivePolicy represents the EffectivePolicy schema from the OpenAPI specification
type EffectivePolicy struct {
	Lastupdatedtimestamp interface{} `json:"LastUpdatedTimestamp,omitempty"`
	Policycontent interface{} `json:"PolicyContent,omitempty"`
	Policytype interface{} `json:"PolicyType,omitempty"`
	Targetid interface{} `json:"TargetId,omitempty"`
}

// AcceptHandshakeRequest represents the AcceptHandshakeRequest schema from the OpenAPI specification
type AcceptHandshakeRequest struct {
	Handshakeid interface{} `json:"HandshakeId"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourceid interface{} `json:"ResourceId"`
}

// Root represents the Root schema from the OpenAPI specification
type Root struct {
	Arn interface{} `json:"Arn,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Policytypes interface{} `json:"PolicyTypes,omitempty"`
}

// DescribeHandshakeResponse represents the DescribeHandshakeResponse schema from the OpenAPI specification
type DescribeHandshakeResponse struct {
	Handshake interface{} `json:"Handshake,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DescribePolicyRequest represents the DescribePolicyRequest schema from the OpenAPI specification
type DescribePolicyRequest struct {
	Policyid interface{} `json:"PolicyId"`
}

// DescribeEffectivePolicyRequest represents the DescribeEffectivePolicyRequest schema from the OpenAPI specification
type DescribeEffectivePolicyRequest struct {
	Policytype interface{} `json:"PolicyType"`
	Targetid interface{} `json:"TargetId,omitempty"`
}

// ListOrganizationalUnitsForParentResponse represents the ListOrganizationalUnitsForParentResponse schema from the OpenAPI specification
type ListOrganizationalUnitsForParentResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Organizationalunits interface{} `json:"OrganizationalUnits,omitempty"`
}

// HandshakeFilter represents the HandshakeFilter schema from the OpenAPI specification
type HandshakeFilter struct {
	Actiontype interface{} `json:"ActionType,omitempty"`
	Parenthandshakeid interface{} `json:"ParentHandshakeId,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// DeclineHandshakeRequest represents the DeclineHandshakeRequest schema from the OpenAPI specification
type DeclineHandshakeRequest struct {
	Handshakeid interface{} `json:"HandshakeId"`
}

// UpdateOrganizationalUnitResponse represents the UpdateOrganizationalUnitResponse schema from the OpenAPI specification
type UpdateOrganizationalUnitResponse struct {
	Organizationalunit interface{} `json:"OrganizationalUnit,omitempty"`
}

// ListDelegatedAdministratorsRequest represents the ListDelegatedAdministratorsRequest schema from the OpenAPI specification
type ListDelegatedAdministratorsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Serviceprincipal interface{} `json:"ServicePrincipal,omitempty"`
}

// ListChildrenRequest represents the ListChildrenRequest schema from the OpenAPI specification
type ListChildrenRequest struct {
	Parentid interface{} `json:"ParentId"`
	Childtype interface{} `json:"ChildType"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdatePolicyResponse represents the UpdatePolicyResponse schema from the OpenAPI specification
type UpdatePolicyResponse struct {
	Policy interface{} `json:"Policy,omitempty"`
}

// OrganizationalUnit represents the OrganizationalUnit schema from the OpenAPI specification
type OrganizationalUnit struct {
	Name interface{} `json:"Name,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// InviteAccountToOrganizationResponse represents the InviteAccountToOrganizationResponse schema from the OpenAPI specification
type InviteAccountToOrganizationResponse struct {
	Handshake interface{} `json:"Handshake,omitempty"`
}

// UpdatePolicyRequest represents the UpdatePolicyRequest schema from the OpenAPI specification
type UpdatePolicyRequest struct {
	Content interface{} `json:"Content,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Policyid interface{} `json:"PolicyId"`
}

// DeclineHandshakeResponse represents the DeclineHandshakeResponse schema from the OpenAPI specification
type DeclineHandshakeResponse struct {
	Handshake interface{} `json:"Handshake,omitempty"`
}

// EnableAWSServiceAccessRequest represents the EnableAWSServiceAccessRequest schema from the OpenAPI specification
type EnableAWSServiceAccessRequest struct {
	Serviceprincipal interface{} `json:"ServicePrincipal"`
}

// DelegatedService represents the DelegatedService schema from the OpenAPI specification
type DelegatedService struct {
	Delegationenableddate interface{} `json:"DelegationEnabledDate,omitempty"`
	Serviceprincipal interface{} `json:"ServicePrincipal,omitempty"`
}

// RegisterDelegatedAdministratorRequest represents the RegisterDelegatedAdministratorRequest schema from the OpenAPI specification
type RegisterDelegatedAdministratorRequest struct {
	Accountid interface{} `json:"AccountId"`
	Serviceprincipal interface{} `json:"ServicePrincipal"`
}

// ListTargetsForPolicyRequest represents the ListTargetsForPolicyRequest schema from the OpenAPI specification
type ListTargetsForPolicyRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Policyid interface{} `json:"PolicyId"`
}

// DisablePolicyTypeRequest represents the DisablePolicyTypeRequest schema from the OpenAPI specification
type DisablePolicyTypeRequest struct {
	Rootid interface{} `json:"RootId"`
	Policytype interface{} `json:"PolicyType"`
}

// ListAWSServiceAccessForOrganizationRequest represents the ListAWSServiceAccessForOrganizationRequest schema from the OpenAPI specification
type ListAWSServiceAccessForOrganizationRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListTargetsForPolicyResponse represents the ListTargetsForPolicyResponse schema from the OpenAPI specification
type ListTargetsForPolicyResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
}

// DescribeOrganizationalUnitRequest represents the DescribeOrganizationalUnitRequest schema from the OpenAPI specification
type DescribeOrganizationalUnitRequest struct {
	Organizationalunitid interface{} `json:"OrganizationalUnitId"`
}

// ListAccountsForParentResponse represents the ListAccountsForParentResponse schema from the OpenAPI specification
type ListAccountsForParentResponse struct {
	Accounts interface{} `json:"Accounts,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeResourcePolicyResponse represents the DescribeResourcePolicyResponse schema from the OpenAPI specification
type DescribeResourcePolicyResponse struct {
	Resourcepolicy interface{} `json:"ResourcePolicy,omitempty"`
}

// CancelHandshakeRequest represents the CancelHandshakeRequest schema from the OpenAPI specification
type CancelHandshakeRequest struct {
	Handshakeid interface{} `json:"HandshakeId"`
}

// ListPoliciesForTargetResponse represents the ListPoliciesForTargetResponse schema from the OpenAPI specification
type ListPoliciesForTargetResponse struct {
	Policies interface{} `json:"Policies,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListParentsRequest represents the ListParentsRequest schema from the OpenAPI specification
type ListParentsRequest struct {
	Childid interface{} `json:"ChildId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AcceptHandshakeResponse represents the AcceptHandshakeResponse schema from the OpenAPI specification
type AcceptHandshakeResponse struct {
	Handshake interface{} `json:"Handshake,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourceid interface{} `json:"ResourceId"`
	Tags interface{} `json:"Tags"`
}

// DescribeOrganizationResponse represents the DescribeOrganizationResponse schema from the OpenAPI specification
type DescribeOrganizationResponse struct {
	Organization interface{} `json:"Organization,omitempty"`
}

// DescribeAccountRequest represents the DescribeAccountRequest schema from the OpenAPI specification
type DescribeAccountRequest struct {
	Accountid interface{} `json:"AccountId"`
}

// DescribeHandshakeRequest represents the DescribeHandshakeRequest schema from the OpenAPI specification
type DescribeHandshakeRequest struct {
	Handshakeid interface{} `json:"HandshakeId"`
}

// CreateOrganizationalUnitResponse represents the CreateOrganizationalUnitResponse schema from the OpenAPI specification
type CreateOrganizationalUnitResponse struct {
	Organizationalunit interface{} `json:"OrganizationalUnit,omitempty"`
}

// ResourcePolicySummary represents the ResourcePolicySummary schema from the OpenAPI specification
type ResourcePolicySummary struct {
	Arn interface{} `json:"Arn,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// DescribeEffectivePolicyResponse represents the DescribeEffectivePolicyResponse schema from the OpenAPI specification
type DescribeEffectivePolicyResponse struct {
	Effectivepolicy interface{} `json:"EffectivePolicy,omitempty"`
}

// ListCreateAccountStatusRequest represents the ListCreateAccountStatusRequest schema from the OpenAPI specification
type ListCreateAccountStatusRequest struct {
	States interface{} `json:"States,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateGovCloudAccountResponse represents the CreateGovCloudAccountResponse schema from the OpenAPI specification
type CreateGovCloudAccountResponse struct {
	Createaccountstatus CreateAccountStatus `json:"CreateAccountStatus,omitempty"` // Contains the status about a <a>CreateAccount</a> or <a>CreateGovCloudAccount</a> request to create an Amazon Web Services account or an Amazon Web Services GovCloud (US) account in an organization.
}

// HandshakeParty represents the HandshakeParty schema from the OpenAPI specification
type HandshakeParty struct {
	Id interface{} `json:"Id"`
	TypeField interface{} `json:"Type"`
}

// PutResourcePolicyRequest represents the PutResourcePolicyRequest schema from the OpenAPI specification
type PutResourcePolicyRequest struct {
	Content interface{} `json:"Content"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateAccountRequest represents the CreateAccountRequest schema from the OpenAPI specification
type CreateAccountRequest struct {
	Accountname interface{} `json:"AccountName"`
	Email interface{} `json:"Email"`
	Iamuseraccesstobilling interface{} `json:"IamUserAccessToBilling,omitempty"`
	Rolename interface{} `json:"RoleName,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListPoliciesRequest represents the ListPoliciesRequest schema from the OpenAPI specification
type ListPoliciesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filter interface{} `json:"Filter"`
}

// EnabledServicePrincipal represents the EnabledServicePrincipal schema from the OpenAPI specification
type EnabledServicePrincipal struct {
	Dateenabled interface{} `json:"DateEnabled,omitempty"`
	Serviceprincipal interface{} `json:"ServicePrincipal,omitempty"`
}

// DescribeCreateAccountStatusResponse represents the DescribeCreateAccountStatusResponse schema from the OpenAPI specification
type DescribeCreateAccountStatusResponse struct {
	Createaccountstatus interface{} `json:"CreateAccountStatus,omitempty"`
}

// ListRootsResponse represents the ListRootsResponse schema from the OpenAPI specification
type ListRootsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Roots interface{} `json:"Roots,omitempty"`
}

// Organization represents the Organization schema from the OpenAPI specification
type Organization struct {
	Availablepolicytypes interface{} `json:"AvailablePolicyTypes,omitempty"`
	Featureset interface{} `json:"FeatureSet,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Masteraccountarn interface{} `json:"MasterAccountArn,omitempty"`
	Masteraccountemail interface{} `json:"MasterAccountEmail,omitempty"`
	Masteraccountid interface{} `json:"MasterAccountId,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// ListOrganizationalUnitsForParentRequest represents the ListOrganizationalUnitsForParentRequest schema from the OpenAPI specification
type ListOrganizationalUnitsForParentRequest struct {
	Parentid interface{} `json:"ParentId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListHandshakesForOrganizationRequest represents the ListHandshakesForOrganizationRequest schema from the OpenAPI specification
type ListHandshakesForOrganizationRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filter interface{} `json:"Filter,omitempty"`
}

// ListAccountsRequest represents the ListAccountsRequest schema from the OpenAPI specification
type ListAccountsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListDelegatedAdministratorsResponse represents the ListDelegatedAdministratorsResponse schema from the OpenAPI specification
type ListDelegatedAdministratorsResponse struct {
	Delegatedadministrators interface{} `json:"DelegatedAdministrators,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListRootsRequest represents the ListRootsRequest schema from the OpenAPI specification
type ListRootsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DisablePolicyTypeResponse represents the DisablePolicyTypeResponse schema from the OpenAPI specification
type DisablePolicyTypeResponse struct {
	Root interface{} `json:"Root,omitempty"`
}

// EnableAllFeaturesResponse represents the EnableAllFeaturesResponse schema from the OpenAPI specification
type EnableAllFeaturesResponse struct {
	Handshake interface{} `json:"Handshake,omitempty"`
}
