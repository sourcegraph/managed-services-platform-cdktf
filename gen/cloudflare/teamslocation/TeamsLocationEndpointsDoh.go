package teamslocation


type TeamsLocationEndpointsDoh struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#enabled TeamsLocation#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#networks TeamsLocation#networks}.
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}

