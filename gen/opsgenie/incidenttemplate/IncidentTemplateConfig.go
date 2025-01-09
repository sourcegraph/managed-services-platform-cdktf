package incidenttemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IncidentTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/incident_template#message IncidentTemplate#message}.
	Message *string `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/incident_template#name IncidentTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/incident_template#priority IncidentTemplate#priority}.
	Priority *string `field:"required" json:"priority" yaml:"priority"`
	// stakeholder_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/incident_template#stakeholder_properties IncidentTemplate#stakeholder_properties}
	StakeholderProperties interface{} `field:"required" json:"stakeholderProperties" yaml:"stakeholderProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/incident_template#description IncidentTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/incident_template#details IncidentTemplate#details}.
	Details *map[string]*string `field:"optional" json:"details" yaml:"details"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/incident_template#id IncidentTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/incident_template#impacted_services IncidentTemplate#impacted_services}.
	ImpactedServices *[]*string `field:"optional" json:"impactedServices" yaml:"impactedServices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/incident_template#tags IncidentTemplate#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

