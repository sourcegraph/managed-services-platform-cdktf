package teamslist


type TeamsListItemsWithDescription struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_list#description TeamsList#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_list#value TeamsList#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

