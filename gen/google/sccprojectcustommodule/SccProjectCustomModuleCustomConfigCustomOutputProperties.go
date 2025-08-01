package sccprojectcustommodule


type SccProjectCustomModuleCustomConfigCustomOutputProperties struct {
	// Name of the property for the custom output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/scc_project_custom_module#name SccProjectCustomModule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// value_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/scc_project_custom_module#value_expression SccProjectCustomModule#value_expression}
	ValueExpression *SccProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpression `field:"optional" json:"valueExpression" yaml:"valueExpression"`
}

