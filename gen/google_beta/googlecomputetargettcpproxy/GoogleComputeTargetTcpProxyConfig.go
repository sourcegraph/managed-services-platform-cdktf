package googlecomputetargettcpproxy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeTargetTcpProxyConfig struct {
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
	// A reference to the BackendService resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_tcp_proxy#backend_service GoogleComputeTargetTcpProxy#backend_service}
	BackendService *string `field:"required" json:"backendService" yaml:"backendService"`
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_tcp_proxy#name GoogleComputeTargetTcpProxy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_tcp_proxy#description GoogleComputeTargetTcpProxy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_tcp_proxy#id GoogleComputeTargetTcpProxy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_tcp_proxy#project GoogleComputeTargetTcpProxy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_tcp_proxy#proxy_bind GoogleComputeTargetTcpProxy#proxy_bind}
	ProxyBind interface{} `field:"optional" json:"proxyBind" yaml:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend.
	//
	// Default value: "NONE" Possible values: ["NONE", "PROXY_V1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_tcp_proxy#proxy_header GoogleComputeTargetTcpProxy#proxy_header}
	ProxyHeader *string `field:"optional" json:"proxyHeader" yaml:"proxyHeader"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_tcp_proxy#timeouts GoogleComputeTargetTcpProxy#timeouts}
	Timeouts *GoogleComputeTargetTcpProxyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

