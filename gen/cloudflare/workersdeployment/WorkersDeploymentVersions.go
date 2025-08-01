package workersdeployment


type WorkersDeploymentVersions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_deployment#percentage WorkersDeployment#percentage}.
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_deployment#version_id WorkersDeployment#version_id}.
	VersionId *string `field:"required" json:"versionId" yaml:"versionId"`
}

