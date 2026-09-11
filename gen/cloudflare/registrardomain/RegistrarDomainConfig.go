package registrardomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RegistrarDomainConfig struct {
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
	// Fully qualified domain name (FQDN) including the extension (e.g., `example.com`, `mybrand.app`). The domain name uniquely identifies a registration — the same domain cannot be registered twice, making it a natural idempotency key for registration requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/registrar_domain#domain_name RegistrarDomain#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/registrar_domain#account_id RegistrarDomain#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Auto-renew controls whether subscription is automatically renewed upon domain expiration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/registrar_domain#auto_renew RegistrarDomain#auto_renew}
	AutoRenew interface{} `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// Shows whether a registrar lock is in place for a domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/registrar_domain#locked RegistrarDomain#locked}
	Locked interface{} `field:"optional" json:"locked" yaml:"locked"`
	// Privacy option controls redacting WHOIS information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/registrar_domain#privacy RegistrarDomain#privacy}
	Privacy interface{} `field:"optional" json:"privacy" yaml:"privacy"`
}

