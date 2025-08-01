package gkehubscope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeHubScopeConfig struct {
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
	// The client-provided identifier of the scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_scope#scope_id GkeHubScope#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_scope#id GkeHubScope#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels for this Scope.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_scope#labels GkeHubScope#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Scope-level cluster namespace labels.
	//
	// For the member clusters bound
	// to the Scope, these labels are applied to each namespace under the
	// Scope. Scope-level labels take precedence over Namespace-level
	// labels ('namespace_labels' in the Fleet Namespace resource) if they
	// share a key. Keys and values must be Kubernetes-conformant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_scope#namespace_labels GkeHubScope#namespace_labels}
	NamespaceLabels *map[string]*string `field:"optional" json:"namespaceLabels" yaml:"namespaceLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_scope#project GkeHubScope#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_scope#timeouts GkeHubScope#timeouts}
	Timeouts *GkeHubScopeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

