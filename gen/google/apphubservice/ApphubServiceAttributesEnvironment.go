package apphubservice


type ApphubServiceAttributesEnvironment struct {
	// Environment type. Possible values: ["PRODUCTION", "STAGING", "TEST", "DEVELOPMENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/apphub_service#type ApphubService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

