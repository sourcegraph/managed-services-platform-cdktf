package hyperdriveconfig


type HyperdriveConfigOrigin struct {
	// Set the name of your origin database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#database HyperdriveConfig#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Set the password needed to access your origin database. The API never returns this write-only value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#password HyperdriveConfig#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies the URL scheme used to connect to your origin database. Available values: "postgres", "postgresql", "mysql".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#scheme HyperdriveConfig#scheme}
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
	// Set the user of your origin database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#user HyperdriveConfig#user}
	User *string `field:"required" json:"user" yaml:"user"`
	// Defines the Client ID of the Access token to use when connecting to the origin database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#access_client_id HyperdriveConfig#access_client_id}
	AccessClientId *string `field:"optional" json:"accessClientId" yaml:"accessClientId"`
	// Defines the Client Secret of the Access Token to use when connecting to the origin database.
	//
	// The API never returns this write-only value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#access_client_secret HyperdriveConfig#access_client_secret}
	AccessClientSecret *string `field:"optional" json:"accessClientSecret" yaml:"accessClientSecret"`
	// Defines the host (hostname or IP) of your origin database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#host HyperdriveConfig#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Defines the port of your origin database. Defaults to 5432 for PostgreSQL or 3306 for MySQL if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#port HyperdriveConfig#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The identifier of the Workers VPC Service to connect through.
	//
	// Hyperdrive will egress through the specified VPC Service to reach the origin database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#service_id HyperdriveConfig#service_id}
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
}

