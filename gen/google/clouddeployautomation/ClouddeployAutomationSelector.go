package clouddeployautomation


type ClouddeployAutomationSelector struct {
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/clouddeploy_automation#targets ClouddeployAutomation#targets}
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
}

