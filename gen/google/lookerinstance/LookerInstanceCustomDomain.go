package lookerinstance


type LookerInstanceCustomDomain struct {
	// Domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/looker_instance#domain LookerInstance#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

