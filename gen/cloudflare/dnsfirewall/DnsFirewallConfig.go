package dnsfirewall

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsFirewallConfig struct {
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
	// DNS Firewall cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#name DnsFirewall#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#upstream_ips DnsFirewall#upstream_ips}.
	UpstreamIps *[]*string `field:"required" json:"upstreamIps" yaml:"upstreamIps"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#account_id DnsFirewall#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Attack mitigation settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#attack_mitigation DnsFirewall#attack_mitigation}
	AttackMitigation *DnsFirewallAttackMitigation `field:"optional" json:"attackMitigation" yaml:"attackMitigation"`
	// Whether to refuse to answer queries for the ANY type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#deprecate_any_requests DnsFirewall#deprecate_any_requests}
	DeprecateAnyRequests interface{} `field:"optional" json:"deprecateAnyRequests" yaml:"deprecateAnyRequests"`
	// Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#ecs_fallback DnsFirewall#ecs_fallback}
	EcsFallback interface{} `field:"optional" json:"ecsFallback" yaml:"ecsFallback"`
	// By default, Cloudflare attempts to cache responses for as long as indicated by the TTL received from upstream nameservers.
	//
	// This setting
	// sets an upper bound on this duration. For caching purposes, higher TTLs
	// will be decreased to the maximum value defined by this setting.
	//
	// This setting does not affect the TTL value in the DNS response
	// Cloudflare returns to clients. Cloudflare will always forward the TTL
	// value received from upstream nameservers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#maximum_cache_ttl DnsFirewall#maximum_cache_ttl}
	MaximumCacheTtl *float64 `field:"optional" json:"maximumCacheTtl" yaml:"maximumCacheTtl"`
	// By default, Cloudflare attempts to cache responses for as long as indicated by the TTL received from upstream nameservers.
	//
	// This setting
	// sets a lower bound on this duration. For caching purposes, lower TTLs
	// will be increased to the minimum value defined by this setting.
	//
	// This setting does not affect the TTL value in the DNS response
	// Cloudflare returns to clients. Cloudflare will always forward the TTL
	// value received from upstream nameservers.
	//
	// Note that, even with this setting, there is no guarantee that a
	// response will be cached for at least the specified duration. Cached
	// responses may be removed earlier for capacity or other operational
	// reasons.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#minimum_cache_ttl DnsFirewall#minimum_cache_ttl}
	MinimumCacheTtl *float64 `field:"optional" json:"minimumCacheTtl" yaml:"minimumCacheTtl"`
	// This setting controls how long DNS Firewall should cache negative responses (e.g., NXDOMAIN) from the upstream servers.
	//
	// This setting does not affect the TTL value in the DNS response
	// Cloudflare returns to clients. Cloudflare will always forward the TTL
	// value received from upstream nameservers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#negative_cache_ttl DnsFirewall#negative_cache_ttl}
	NegativeCacheTtl *float64 `field:"optional" json:"negativeCacheTtl" yaml:"negativeCacheTtl"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to the upstream nameservers configured on the cluster).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#ratelimit DnsFirewall#ratelimit}
	Ratelimit *float64 `field:"optional" json:"ratelimit" yaml:"ratelimit"`
	// Number of retries for fetching DNS responses from upstream nameservers (not counting the initial attempt).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/dns_firewall#retries DnsFirewall#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
}

