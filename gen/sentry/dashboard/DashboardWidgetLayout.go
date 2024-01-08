package dashboard


type DashboardWidgetLayout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#h Dashboard#h}.
	H *float64 `field:"required" json:"h" yaml:"h"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#min_h Dashboard#min_h}.
	MinH *float64 `field:"required" json:"minH" yaml:"minH"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#w Dashboard#w}.
	W *float64 `field:"required" json:"w" yaml:"w"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#x Dashboard#x}.
	X *float64 `field:"required" json:"x" yaml:"x"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#y Dashboard#y}.
	Y *float64 `field:"required" json:"y" yaml:"y"`
}

