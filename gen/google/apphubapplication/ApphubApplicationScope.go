package apphubapplication


type ApphubApplicationScope struct {
	// Required. Scope Type.   Possible values: REGIONAL GLOBAL Possible values: ["REGIONAL", "GLOBAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/apphub_application#type ApphubApplication#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

