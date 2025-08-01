package googleworkbenchinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleWorkbenchInstanceConfig struct {
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
	// Part of 'parent'. See documentation of 'projectsId'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#location GoogleWorkbenchInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of this workbench instance. Format: 'projects/{project_id}/locations/{location}/instances/{instance_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#name GoogleWorkbenchInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Desired state of the Workbench Instance.
	//
	// Set this field to 'ACTIVE' to start the Instance, and 'STOPPED' to stop the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#desired_state GoogleWorkbenchInstance#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Optional. If true, the workbench instance will not register with the proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#disable_proxy_access GoogleWorkbenchInstance#disable_proxy_access}
	DisableProxyAccess interface{} `field:"optional" json:"disableProxyAccess" yaml:"disableProxyAccess"`
	// Flag that specifies that a notebook can be accessed with third party identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#enable_third_party_identity GoogleWorkbenchInstance#enable_third_party_identity}
	EnableThirdPartyIdentity interface{} `field:"optional" json:"enableThirdPartyIdentity" yaml:"enableThirdPartyIdentity"`
	// gce_setup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#gce_setup GoogleWorkbenchInstance#gce_setup}
	GceSetup *GoogleWorkbenchInstanceGceSetup `field:"optional" json:"gceSetup" yaml:"gceSetup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#id GoogleWorkbenchInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Required. User-defined unique ID of this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#instance_id GoogleWorkbenchInstance#instance_id}
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// 'Optional.
	//
	// Input only. The owner of this instance after creation. Format:
	// 'alias@example.com' Currently supports one owner only. If not specified, all of
	// the service account users of your VM instance''s service account can use the instance.
	// If specified, sets the access mode to 'Single user'. For more details, see
	// https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-access-jupyterlab'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#instance_owners GoogleWorkbenchInstance#instance_owners}
	InstanceOwners *[]*string `field:"optional" json:"instanceOwners" yaml:"instanceOwners"`
	// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#labels GoogleWorkbenchInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#project GoogleWorkbenchInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#timeouts GoogleWorkbenchInstance#timeouts}
	Timeouts *GoogleWorkbenchInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

