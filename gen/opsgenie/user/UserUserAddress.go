package user


type UserUserAddress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/user#city User#city}.
	City *string `field:"required" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/user#country User#country}.
	Country *string `field:"required" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/user#line User#line}.
	Line *string `field:"required" json:"line" yaml:"line"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/user#state User#state}.
	State *string `field:"required" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/user#zipcode User#zipcode}.
	Zipcode *string `field:"required" json:"zipcode" yaml:"zipcode"`
}

