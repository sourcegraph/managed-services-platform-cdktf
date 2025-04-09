package googleapphubservice


type GoogleApphubServiceAttributesEnvironment struct {
	// Environment type. Possible values: ["PRODUCTION", "STAGING", "TEST", "DEVELOPMENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_apphub_service#type GoogleApphubService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

