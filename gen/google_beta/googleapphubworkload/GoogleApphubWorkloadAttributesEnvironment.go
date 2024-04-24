package googleapphubworkload


type GoogleApphubWorkloadAttributesEnvironment struct {
	// Environment type. Possible values: ["PRODUCTION", "STAGING", "TEST", "DEVELOPMENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_apphub_workload#type GoogleApphubWorkload#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

