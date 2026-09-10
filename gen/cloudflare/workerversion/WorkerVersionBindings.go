package workerversion


type WorkerVersionBindings struct {
	// A JavaScript variable name for the binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#name WorkerVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The kind of resource that the binding provides.
	//
	// Available values: "ai", "ai_search", "ai_search_namespace", "analytics_engine", "assets", "browser", "d1", "data_blob", "dispatch_namespace", "durable_object_namespace", "hyperdrive", "inherit", "images", "json", "kv_namespace", "media", "mtls_certificate", "plain_text", "pipelines", "queue", "ratelimit", "r2_bucket", "secret_text", "send_email", "service", "text_blob", "vectorize", "version_metadata", "secrets_store_secret", "flagship", "secret_key", "workflow", "wasm_module", "vpc_service", "vpc_network".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#type WorkerVersion#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Algorithm-specific key parameters. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#algorithm WorkerVersion#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// List of allowed destination addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#allowed_destination_addresses WorkerVersion#allowed_destination_addresses}
	AllowedDestinationAddresses *[]*string `field:"optional" json:"allowedDestinationAddresses" yaml:"allowedDestinationAddresses"`
	// List of allowed sender addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#allowed_sender_addresses WorkerVersion#allowed_sender_addresses}
	AllowedSenderAddresses *[]*string `field:"optional" json:"allowedSenderAddresses" yaml:"allowedSenderAddresses"`
	// ID of the Flagship app to bind to for feature flag evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#app_id WorkerVersion#app_id}
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// R2 bucket to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#bucket_name WorkerVersion#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Identifier of the certificate to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#certificate_id WorkerVersion#certificate_id}
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// The exported class name of the Durable Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#class_name WorkerVersion#class_name}
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// Identifier of the D1 database to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#database_id WorkerVersion#database_id}
	DatabaseId *string `field:"optional" json:"databaseId" yaml:"databaseId"`
	// The name of the dataset to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#dataset WorkerVersion#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// Destination address for the email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#destination_address WorkerVersion#destination_address}
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// The dispatch namespace the Durable Object script belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#dispatch_namespace WorkerVersion#dispatch_namespace}
	DispatchNamespace *string `field:"optional" json:"dispatchNamespace" yaml:"dispatchNamespace"`
	// Entrypoint to invoke on the target Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#entrypoint WorkerVersion#entrypoint}
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The environment of the script_name to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#environment WorkerVersion#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Data format of the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format). Available values: "raw", "pkcs8", "spki", "jwk".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#format WorkerVersion#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Identifier of the D1 database to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#id WorkerVersion#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the Vectorize index to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#index_name WorkerVersion#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The user-chosen instance name.
	//
	// Must exist at deploy time. The worker can search, chat, update, and manage items/jobs on this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#instance_name WorkerVersion#instance_name}
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// JSON data to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#json WorkerVersion#json}
	Json *string `field:"optional" json:"json" yaml:"json"`
	// The [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions) of the R2 bucket. Available values: "eu", "fedramp", "fedramp-high".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#jurisdiction WorkerVersion#jurisdiction}
	Jurisdiction *string `field:"optional" json:"jurisdiction" yaml:"jurisdiction"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#key_base64 WorkerVersion#key_base64}
	KeyBase64 *string `field:"optional" json:"keyBase64" yaml:"keyBase64"`
	// Key data in [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key) format. Required if `format` is "jwk".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#key_jwk WorkerVersion#key_jwk}
	KeyJwk *string `field:"optional" json:"keyJwk" yaml:"keyJwk"`
	// The namespace the instance belongs to.
	//
	// Defaults to "default" if omitted. Customers who don't use namespaces can simply omit this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#namespace WorkerVersion#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Namespace identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#namespace_id WorkerVersion#namespace_id}
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported. Mutually exclusive with tunnel_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#network_id WorkerVersion#network_id}
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// The old name of the inherited binding.
	//
	// If set, the binding will be renamed from `old_name` to `name` in the new version. If not set, the binding will keep the same name between versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#old_name WorkerVersion#old_name}
	OldName *string `field:"optional" json:"oldName" yaml:"oldName"`
	// Outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#outbound WorkerVersion#outbound}
	Outbound *WorkerVersionBindingsOutbound `field:"optional" json:"outbound" yaml:"outbound"`
	// The name of the file containing the data content. Only accepted for `service worker syntax` Workers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#part WorkerVersion#part}
	Part *string `field:"optional" json:"part" yaml:"part"`
	// Name of the Pipeline to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#pipeline WorkerVersion#pipeline}
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// Name of the Queue to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#queue_name WorkerVersion#queue_name}
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// The script where the Durable Object is defined, if it is external to this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#script_name WorkerVersion#script_name}
	ScriptName *string `field:"optional" json:"scriptName" yaml:"scriptName"`
	// Name of the secret in the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#secret_name WorkerVersion#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// Name of Worker to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#service WorkerVersion#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// Identifier of the VPC service to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#service_id WorkerVersion#service_id}
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
	// The rate limit configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#simple WorkerVersion#simple}
	Simple *WorkerVersionBindingsSimple `field:"optional" json:"simple" yaml:"simple"`
	// ID of the store containing the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#store_id WorkerVersion#store_id}
	StoreId *string `field:"optional" json:"storeId" yaml:"storeId"`
	// The text value to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#text WorkerVersion#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#tunnel_id WorkerVersion#tunnel_id}
	TunnelId *string `field:"optional" json:"tunnelId" yaml:"tunnelId"`
	// Allowed operations with the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#usages WorkerVersion#usages}
	Usages *[]*string `field:"optional" json:"usages" yaml:"usages"`
	// Identifier for the version to inherit the binding from, which can be the version ID or the literal "latest" to inherit from the latest version.
	//
	// Defaults to inheriting the binding from the latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#version_id WorkerVersion#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// Name of the Workflow to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#workflow_name WorkerVersion#workflow_name}
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

