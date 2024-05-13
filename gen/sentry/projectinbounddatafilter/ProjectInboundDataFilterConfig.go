package projectinbounddatafilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectInboundDataFilterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The type of filter toggle to update. See the [Sentry documentation](https://docs.sentry.io/api/projects/update-an-inbound-data-filter/) for a list of available filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.3/docs/resources/project_inbound_data_filter#filter_id ProjectInboundDataFilter#filter_id}
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// The slug of the organization the project belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.3/docs/resources/project_inbound_data_filter#organization ProjectInboundDataFilter#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The slug of the project to create the filter for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.3/docs/resources/project_inbound_data_filter#project ProjectInboundDataFilter#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Toggle the browser-extensions, localhost, filtered-transaction, or web-crawlers filter on or off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.3/docs/resources/project_inbound_data_filter#active ProjectInboundDataFilter#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Specifies which legacy browser filters should be active.
	//
	// Anything excluded from the list will be disabled. See the [Sentry documentation](https://docs.sentry.io/api/projects/update-an-inbound-data-filter/) for a list of available subfilters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.3/docs/resources/project_inbound_data_filter#subfilters ProjectInboundDataFilter#subfilters}
	Subfilters *[]*string `field:"optional" json:"subfilters" yaml:"subfilters"`
}

