package endpointsservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EndpointsServiceConfig struct {
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
	// The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/endpoints_service#service_name EndpointsService#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The full text of the Service Config YAML file (Example located here).
	//
	// If provided, must also provide protoc_output_base64. open_api config must not be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/endpoints_service#grpc_config EndpointsService#grpc_config}
	GrpcConfig *string `field:"optional" json:"grpcConfig" yaml:"grpcConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/endpoints_service#id EndpointsService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The full text of the OpenAPI YAML configuration as described here.
	//
	// Either this, or both of grpc_config and protoc_output_base64 must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/endpoints_service#openapi_config EndpointsService#openapi_config}
	OpenapiConfig *string `field:"optional" json:"openapiConfig" yaml:"openapiConfig"`
	// The project ID that the service belongs to. If not provided, provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/endpoints_service#project EndpointsService#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file, base64-encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/endpoints_service#protoc_output_base64 EndpointsService#protoc_output_base64}
	ProtocOutputBase64 *string `field:"optional" json:"protocOutputBase64" yaml:"protocOutputBase64"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/endpoints_service#timeouts EndpointsService#timeouts}
	Timeouts *EndpointsServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
