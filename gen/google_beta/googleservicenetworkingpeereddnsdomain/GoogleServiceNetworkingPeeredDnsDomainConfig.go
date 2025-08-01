package googleservicenetworkingpeereddnsdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleServiceNetworkingPeeredDnsDomainConfig struct {
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
	// The DNS domain name suffix of the peered DNS domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_service_networking_peered_dns_domain#dns_suffix GoogleServiceNetworkingPeeredDnsDomain#dns_suffix}
	DnsSuffix *string `field:"required" json:"dnsSuffix" yaml:"dnsSuffix"`
	// Name of the peered DNS domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_service_networking_peered_dns_domain#name GoogleServiceNetworkingPeeredDnsDomain#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Network in the consumer project to peer with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_service_networking_peered_dns_domain#network GoogleServiceNetworkingPeeredDnsDomain#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_service_networking_peered_dns_domain#id GoogleServiceNetworkingPeeredDnsDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the project that the service account will be created in. Defaults to the provider project configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_service_networking_peered_dns_domain#project GoogleServiceNetworkingPeeredDnsDomain#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The name of the service to create a peered DNS domain for, e.g. servicenetworking.googleapis.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_service_networking_peered_dns_domain#service GoogleServiceNetworkingPeeredDnsDomain#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_service_networking_peered_dns_domain#timeouts GoogleServiceNetworkingPeeredDnsDomain#timeouts}
	Timeouts *GoogleServiceNetworkingPeeredDnsDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

