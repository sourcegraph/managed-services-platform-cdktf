package googleapphubworkload


type GoogleApphubWorkloadAttributesCriticality struct {
	// Criticality type. Possible values: ["MISSION_CRITICAL", "HIGH", "MEDIUM", "LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_apphub_workload#type GoogleApphubWorkload#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

