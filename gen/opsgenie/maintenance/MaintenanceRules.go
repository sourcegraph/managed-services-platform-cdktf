package maintenance


type MaintenanceRules struct {
	// entity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/maintenance#entity Maintenance#entity}
	Entity interface{} `field:"required" json:"entity" yaml:"entity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/maintenance#state Maintenance#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

