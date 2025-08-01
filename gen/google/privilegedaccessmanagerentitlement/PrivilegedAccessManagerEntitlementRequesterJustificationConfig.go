package privilegedaccessmanagerentitlement


type PrivilegedAccessManagerEntitlementRequesterJustificationConfig struct {
	// not_mandatory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/privileged_access_manager_entitlement#not_mandatory PrivilegedAccessManagerEntitlement#not_mandatory}
	NotMandatory *PrivilegedAccessManagerEntitlementRequesterJustificationConfigNotMandatory `field:"optional" json:"notMandatory" yaml:"notMandatory"`
	// unstructured block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/privileged_access_manager_entitlement#unstructured PrivilegedAccessManagerEntitlement#unstructured}
	Unstructured *PrivilegedAccessManagerEntitlementRequesterJustificationConfigUnstructured `field:"optional" json:"unstructured" yaml:"unstructured"`
}

