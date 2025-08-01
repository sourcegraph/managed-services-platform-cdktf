package dnszonetransfersoutgoing

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsZoneTransfersOutgoingConfig struct {
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
	// Zone name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/dns_zone_transfers_outgoing#name DnsZoneTransfersOutgoing#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of peer tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/dns_zone_transfers_outgoing#peers DnsZoneTransfersOutgoing#peers}
	Peers *[]*string `field:"required" json:"peers" yaml:"peers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/dns_zone_transfers_outgoing#zone_id DnsZoneTransfersOutgoing#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

