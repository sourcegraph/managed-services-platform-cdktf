package datagooglednsmanagedzones

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleDnsManagedZonesConfig struct {
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
	// managed_zones block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/data-sources/google_dns_managed_zones#managed_zones DataGoogleDnsManagedZones#managed_zones}
	ManagedZones interface{} `field:"optional" json:"managedZones" yaml:"managedZones"`
	// The ID of the project for the Google Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/data-sources/google_dns_managed_zones#project DataGoogleDnsManagedZones#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}
