package apphubworkload


type ApphubWorkloadAttributesCriticality struct {
	// Criticality type. Possible values: ["MISSION_CRITICAL", "HIGH", "MEDIUM", "LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/apphub_workload#type ApphubWorkload#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

