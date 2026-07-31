package zerotrustaccessservicetoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessServiceTokenConfig struct {
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
	// The name of the service token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_service_token#name ZeroTrustAccessServiceToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_service_token#account_id ZeroTrustAccessServiceToken#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// A version number identifying the current `client_secret` associated with the service token.
	//
	// Incrementing it triggers a rotation; the previous secret will still be accepted until the time indicated by `previous_client_secret_expires_at`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_service_token#client_secret_version ZeroTrustAccessServiceToken#client_secret_version}
	ClientSecretVersion *float64 `field:"optional" json:"clientSecretVersion" yaml:"clientSecretVersion"`
	// The duration for how long the service token will be valid.
	//
	// Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The default is 1 year in hours (8760h).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_service_token#duration ZeroTrustAccessServiceToken#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// The expiration of the previous `client_secret`.
	//
	// This can be modified at any point after a rotation. For example, you may extend it further into the future if you need more time to update services with the new secret; or move it into the past to immediately invalidate the previous token in case of compromise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_service_token#previous_client_secret_expires_at ZeroTrustAccessServiceToken#previous_client_secret_expires_at}
	PreviousClientSecretExpiresAt *string `field:"optional" json:"previousClientSecretExpiresAt" yaml:"previousClientSecretExpiresAt"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_service_token#zone_id ZeroTrustAccessServiceToken#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

