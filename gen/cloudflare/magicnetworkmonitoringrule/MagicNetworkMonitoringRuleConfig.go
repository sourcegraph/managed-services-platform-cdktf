package magicnetworkmonitoringrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicNetworkMonitoringRuleConfig struct {
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
	// Toggle on if you would like Cloudflare to automatically advertise the IP Prefixes within the rule via Magic Transit when the rule is triggered.
	//
	// Only available for users of Magic Transit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#automatic_advertisement MagicNetworkMonitoringRule#automatic_advertisement}
	AutomaticAdvertisement interface{} `field:"required" json:"automaticAdvertisement" yaml:"automaticAdvertisement"`
	// The name of the rule.
	//
	// Must be unique. Supports characters A-Z, a-z, 0-9, underscore (_), dash (-), period (.), and tilde (~). You can’t have a space in the rule name. Max 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#name MagicNetworkMonitoringRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#prefixes MagicNetworkMonitoringRule#prefixes}.
	Prefixes *[]*string `field:"required" json:"prefixes" yaml:"prefixes"`
	// MNM rule type. Available values: "threshold", "zscore", "advanced_ddos".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#type MagicNetworkMonitoringRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#account_id MagicNetworkMonitoringRule#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The number of bits per second for the rule.
	//
	// When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#bandwidth_threshold MagicNetworkMonitoringRule#bandwidth_threshold}
	BandwidthThreshold *float64 `field:"optional" json:"bandwidthThreshold" yaml:"bandwidthThreshold"`
	// The amount of time that the rule threshold must be exceeded to send an alert notification.
	//
	// The final value must be equivalent to one of the following 8 values ["1m","5m","10m","15m","20m","30m","45m","60m"].
	// Available values: "1m", "5m", "10m", "15m", "20m", "30m", "45m", "60m".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#duration MagicNetworkMonitoringRule#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// The number of packets per second for the rule.
	//
	// When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#packet_threshold MagicNetworkMonitoringRule#packet_threshold}
	PacketThreshold *float64 `field:"optional" json:"packetThreshold" yaml:"packetThreshold"`
	// Prefix match type to be applied for a prefix auto advertisement when using an advanced_ddos rule.
	//
	// Available values: "exact", "subnet", "supernet".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#prefix_match MagicNetworkMonitoringRule#prefix_match}
	PrefixMatch *string `field:"optional" json:"prefixMatch" yaml:"prefixMatch"`
	// Level of sensitivity set for zscore rules. Available values: "low", "medium", "high".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#zscore_sensitivity MagicNetworkMonitoringRule#zscore_sensitivity}
	ZscoreSensitivity *string `field:"optional" json:"zscoreSensitivity" yaml:"zscoreSensitivity"`
	// Target of the zscore rule analysis. Available values: "bits", "packets".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_network_monitoring_rule#zscore_target MagicNetworkMonitoringRule#zscore_target}
	ZscoreTarget *string `field:"optional" json:"zscoreTarget" yaml:"zscoreTarget"`
}

