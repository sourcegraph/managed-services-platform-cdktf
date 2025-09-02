package roleattribute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleAttributeConfig struct {
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
	// The name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#name RoleAttribute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Role to switch to at login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#assume_role RoleAttribute#assume_role}
	AssumeRole *string `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// Determine whether a role bypasses every row-level security (RLS) policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#bypass_row_level_security RoleAttribute#bypass_row_level_security}
	BypassRowLevelSecurity interface{} `field:"optional" json:"bypassRowLevelSecurity" yaml:"bypassRowLevelSecurity"`
	// How many concurrent connections can be made with this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#connection_limit RoleAttribute#connection_limit}
	ConnectionLimit *float64 `field:"optional" json:"connectionLimit" yaml:"connectionLimit"`
	// Define a role's ability to create databases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#create_database RoleAttribute#create_database}
	CreateDatabase interface{} `field:"optional" json:"createDatabase" yaml:"createDatabase"`
	// Determine whether this role will be permitted to create new roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#create_role RoleAttribute#create_role}
	CreateRole interface{} `field:"optional" json:"createRole" yaml:"createRole"`
	// Control whether the password is stored encrypted in the system catalogs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#encrypted_password RoleAttribute#encrypted_password}
	EncryptedPassword interface{} `field:"optional" json:"encryptedPassword" yaml:"encryptedPassword"`
	// Map of arbitrary GUC (Grand Unified Configuration) key-value pairs to set for the role. Supports all PostgreSQL custom variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#extension_attrs RoleAttribute#extension_attrs}
	ExtensionAttrs *map[string]*string `field:"optional" json:"extensionAttrs" yaml:"extensionAttrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#id RoleAttribute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Terminate any session with an open transaction that has been idle for longer than the specified duration in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#idle_in_transaction_session_timeout RoleAttribute#idle_in_transaction_session_timeout}
	IdleInTransactionSessionTimeout *float64 `field:"optional" json:"idleInTransactionSessionTimeout" yaml:"idleInTransactionSessionTimeout"`
	// Determine whether a role "inherits" the privileges of roles it is a member of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#inherit RoleAttribute#inherit}
	Inherit interface{} `field:"optional" json:"inherit" yaml:"inherit"`
	// Determine whether a role is allowed to log in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#login RoleAttribute#login}
	Login interface{} `field:"optional" json:"login" yaml:"login"`
	// Sets the role's password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#password RoleAttribute#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Determine whether a role is allowed to initiate streaming replication or put the system in and out of backup mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#replication RoleAttribute#replication}
	Replication interface{} `field:"optional" json:"replication" yaml:"replication"`
	// Sets the role's search path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#search_path RoleAttribute#search_path}
	SearchPath *[]*string `field:"optional" json:"searchPath" yaml:"searchPath"`
	// Abort any statement that takes more than the specified number of milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#statement_timeout RoleAttribute#statement_timeout}
	StatementTimeout *float64 `field:"optional" json:"statementTimeout" yaml:"statementTimeout"`
	// Determine whether the new role is a "superuser".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#superuser RoleAttribute#superuser}
	Superuser interface{} `field:"optional" json:"superuser" yaml:"superuser"`
	// Sets a date and time after which the role's password is no longer valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.25.0-sg.4/docs/resources/role_attribute#valid_until RoleAttribute#valid_until}
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}

