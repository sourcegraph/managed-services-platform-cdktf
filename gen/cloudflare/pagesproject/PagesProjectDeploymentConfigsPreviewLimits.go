package pagesproject


type PagesProjectDeploymentConfigsPreviewLimits struct {
	// CPU time limit in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#cpu_ms PagesProject#cpu_ms}
	CpuMs *float64 `field:"required" json:"cpuMs" yaml:"cpuMs"`
}

