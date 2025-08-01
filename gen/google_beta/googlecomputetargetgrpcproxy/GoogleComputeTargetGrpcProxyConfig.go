package googlecomputetargetgrpcproxy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeTargetGrpcProxyConfig struct {
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
	// Name of the resource.
	//
	// Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_grpc_proxy#name GoogleComputeTargetGrpcProxy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_grpc_proxy#description GoogleComputeTargetGrpcProxy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_grpc_proxy#id GoogleComputeTargetGrpcProxy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_grpc_proxy#project GoogleComputeTargetGrpcProxy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_grpc_proxy#timeouts GoogleComputeTargetGrpcProxy#timeouts}
	Timeouts *GoogleComputeTargetGrpcProxyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	//
	// The protocol field in the BackendService
	// must be set to GRPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_grpc_proxy#url_map GoogleComputeTargetGrpcProxy#url_map}
	UrlMap *string `field:"optional" json:"urlMap" yaml:"urlMap"`
	// If true, indicates that the BackendServices referenced by the urlMap may be accessed by gRPC applications without using a sidecar proxy.
	//
	// This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_grpc_proxy#validate_for_proxyless GoogleComputeTargetGrpcProxy#validate_for_proxyless}
	ValidateForProxyless interface{} `field:"optional" json:"validateForProxyless" yaml:"validateForProxyless"`
}

