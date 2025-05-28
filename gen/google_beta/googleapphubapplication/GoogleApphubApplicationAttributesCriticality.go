package googleapphubapplication


type GoogleApphubApplicationAttributesCriticality struct {
	// Criticality type. Possible values: ["MISSION_CRITICAL", "HIGH", "MEDIUM", "LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_apphub_application#type GoogleApphubApplication#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

