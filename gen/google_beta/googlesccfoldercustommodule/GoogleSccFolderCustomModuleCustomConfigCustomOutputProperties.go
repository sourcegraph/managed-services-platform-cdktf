package googlesccfoldercustommodule


type GoogleSccFolderCustomModuleCustomConfigCustomOutputProperties struct {
	// Name of the property for the custom output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_scc_folder_custom_module#name GoogleSccFolderCustomModule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// value_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_scc_folder_custom_module#value_expression GoogleSccFolderCustomModule#value_expression}
	ValueExpression *GoogleSccFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpression `field:"optional" json:"valueExpression" yaml:"valueExpression"`
}

