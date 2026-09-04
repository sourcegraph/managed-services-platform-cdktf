package workersscript


type WorkersScriptBindings struct {
	// A JavaScript variable name for the binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#name WorkersScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The kind of resource that the binding provides.
	//
	// Available values: "ai", "analytics_engine", "assets", "browser", "d1", "data_blob", "dispatch_namespace", "durable_object_namespace", "hyperdrive", "inherit", "images", "json", "kv_namespace", "mtls_certificate", "plain_text", "pipelines", "queue", "r2_bucket", "secret_text", "send_email", "service", "tail_consumer", "text_blob", "vectorize", "version_metadata", "secrets_store_secret", "secret_key", "workflow", "wasm_module".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#type WorkersScript#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Algorithm-specific key parameters. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#algorithm WorkersScript#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// List of allowed destination addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#allowed_destination_addresses WorkersScript#allowed_destination_addresses}
	AllowedDestinationAddresses *[]*string `field:"optional" json:"allowedDestinationAddresses" yaml:"allowedDestinationAddresses"`
	// List of allowed sender addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#allowed_sender_addresses WorkersScript#allowed_sender_addresses}
	AllowedSenderAddresses *[]*string `field:"optional" json:"allowedSenderAddresses" yaml:"allowedSenderAddresses"`
	// R2 bucket to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#bucket_name WorkersScript#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Identifier of the certificate to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#certificate_id WorkersScript#certificate_id}
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// The exported class name of the Durable Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#class_name WorkersScript#class_name}
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// The name of the dataset to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#dataset WorkersScript#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// Destination address for the email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#destination_address WorkersScript#destination_address}
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// The environment of the script_name to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#environment WorkersScript#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Data format of the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format). Available values: "raw", "pkcs8", "spki", "jwk".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#format WorkersScript#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Identifier of the D1 database to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#id WorkersScript#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the Vectorize index to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#index_name WorkersScript#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// JSON data to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#json WorkersScript#json}
	Json *string `field:"optional" json:"json" yaml:"json"`
	// The [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions) of the R2 bucket. Available values: "eu", "fedramp".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#jurisdiction WorkersScript#jurisdiction}
	Jurisdiction *string `field:"optional" json:"jurisdiction" yaml:"jurisdiction"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#key_base64 WorkersScript#key_base64}
	KeyBase64 *string `field:"optional" json:"keyBase64" yaml:"keyBase64"`
	// Key data in [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key) format. Required if `format` is "jwk".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#key_jwk WorkersScript#key_jwk}
	KeyJwk *string `field:"optional" json:"keyJwk" yaml:"keyJwk"`
	// The name of the dispatch namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#namespace WorkersScript#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Namespace identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#namespace_id WorkersScript#namespace_id}
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// The old name of the inherited binding.
	//
	// If set, the binding will be renamed from `old_name` to `name` in the new version. If not set, the binding will keep the same name between versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#old_name WorkersScript#old_name}
	OldName *string `field:"optional" json:"oldName" yaml:"oldName"`
	// Outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#outbound WorkersScript#outbound}
	Outbound *WorkersScriptBindingsOutbound `field:"optional" json:"outbound" yaml:"outbound"`
	// The name of the file containing the data content. Only accepted for `service worker syntax` Workers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#part WorkersScript#part}
	Part *string `field:"optional" json:"part" yaml:"part"`
	// Name of the Pipeline to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#pipeline WorkersScript#pipeline}
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// Name of the Queue to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#queue_name WorkersScript#queue_name}
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// The script where the Durable Object is defined, if it is external to this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#script_name WorkersScript#script_name}
	ScriptName *string `field:"optional" json:"scriptName" yaml:"scriptName"`
	// Name of the secret in the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#secret_name WorkersScript#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// Name of Worker to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#service WorkersScript#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// A simple rate limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#simple WorkersScript#simple}
	Simple *WorkersScriptBindingsSimple `field:"optional" json:"simple" yaml:"simple"`
	// ID of the store containing the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#store_id WorkersScript#store_id}
	StoreId *string `field:"optional" json:"storeId" yaml:"storeId"`
	// The text value to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#text WorkersScript#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// Allowed operations with the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#usages WorkersScript#usages}
	Usages *[]*string `field:"optional" json:"usages" yaml:"usages"`
	// Identifier for the version to inherit the binding from, which can be the version ID or the literal "latest" to inherit from the latest version.
	//
	// Defaults to inheriting the binding from the latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#version_id WorkersScript#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// Name of the Workflow to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#workflow_name WorkersScript#workflow_name}
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

