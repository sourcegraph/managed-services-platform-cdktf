package dashboard


type DashboardWidget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/dashboard#display_type Dashboard#display_type}.
	DisplayType *string `field:"required" json:"displayType" yaml:"displayType"`
	// layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/dashboard#layout Dashboard#layout}
	Layout *DashboardWidgetLayout `field:"required" json:"layout" yaml:"layout"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/dashboard#query Dashboard#query}
	Query interface{} `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/dashboard#title Dashboard#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/dashboard#interval Dashboard#interval}.
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/dashboard#limit Dashboard#limit}.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/dashboard#widget_type Dashboard#widget_type}.
	WidgetType *string `field:"optional" json:"widgetType" yaml:"widgetType"`
}

