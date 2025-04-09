package googlecomputenetworkpeeringroutesconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeNetworkPeeringRoutesConfigConfig struct {
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
	// Whether to export the custom routes to the peer network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_network_peering_routes_config#export_custom_routes GoogleComputeNetworkPeeringRoutesConfig#export_custom_routes}
	ExportCustomRoutes interface{} `field:"required" json:"exportCustomRoutes" yaml:"exportCustomRoutes"`
	// Whether to import the custom routes to the peer network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_network_peering_routes_config#import_custom_routes GoogleComputeNetworkPeeringRoutesConfig#import_custom_routes}
	ImportCustomRoutes interface{} `field:"required" json:"importCustomRoutes" yaml:"importCustomRoutes"`
	// The name of the primary network for the peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_network_peering_routes_config#network GoogleComputeNetworkPeeringRoutesConfig#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Name of the peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_network_peering_routes_config#peering GoogleComputeNetworkPeeringRoutesConfig#peering}
	Peering *string `field:"required" json:"peering" yaml:"peering"`
	// Whether subnet routes with public IP range are exported.
	//
	// IPv4 special-use ranges are always exported to peers and
	// are not controlled by this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_network_peering_routes_config#export_subnet_routes_with_public_ip GoogleComputeNetworkPeeringRoutesConfig#export_subnet_routes_with_public_ip}
	ExportSubnetRoutesWithPublicIp interface{} `field:"optional" json:"exportSubnetRoutesWithPublicIp" yaml:"exportSubnetRoutesWithPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_network_peering_routes_config#id GoogleComputeNetworkPeeringRoutesConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether subnet routes with public IP range are imported.
	//
	// IPv4 special-use ranges are always imported from peers and
	// are not controlled by this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_network_peering_routes_config#import_subnet_routes_with_public_ip GoogleComputeNetworkPeeringRoutesConfig#import_subnet_routes_with_public_ip}
	ImportSubnetRoutesWithPublicIp interface{} `field:"optional" json:"importSubnetRoutesWithPublicIp" yaml:"importSubnetRoutesWithPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_network_peering_routes_config#project GoogleComputeNetworkPeeringRoutesConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_network_peering_routes_config#timeouts GoogleComputeNetworkPeeringRoutesConfig#timeouts}
	Timeouts *GoogleComputeNetworkPeeringRoutesConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

