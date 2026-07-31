package pagesproject


type PagesProjectDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#ai_bindings PagesProject#ai_bindings}
	AiBindings interface{} `field:"optional" json:"aiBindings" yaml:"aiBindings"`
	// Whether to always use the latest compatibility date for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#always_use_latest_compatibility_date PagesProject#always_use_latest_compatibility_date}
	AlwaysUseLatestCompatibilityDate interface{} `field:"optional" json:"alwaysUseLatestCompatibilityDate" yaml:"alwaysUseLatestCompatibilityDate"`
	// Analytics Engine bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#analytics_engine_datasets PagesProject#analytics_engine_datasets}
	AnalyticsEngineDatasets interface{} `field:"optional" json:"analyticsEngineDatasets" yaml:"analyticsEngineDatasets"`
	// Browser bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#browsers PagesProject#browsers}
	Browsers interface{} `field:"optional" json:"browsers" yaml:"browsers"`
	// The major version of the build image to use for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#build_image_major_version PagesProject#build_image_major_version}
	BuildImageMajorVersion *float64 `field:"optional" json:"buildImageMajorVersion" yaml:"buildImageMajorVersion"`
	// Compatibility date used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#compatibility_date PagesProject#compatibility_date}
	CompatibilityDate *string `field:"optional" json:"compatibilityDate" yaml:"compatibilityDate"`
	// Compatibility flags used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#compatibility_flags PagesProject#compatibility_flags}
	CompatibilityFlags *[]*string `field:"optional" json:"compatibilityFlags" yaml:"compatibilityFlags"`
	// D1 databases used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#d1_databases PagesProject#d1_databases}
	D1Databases interface{} `field:"optional" json:"d1Databases" yaml:"d1Databases"`
	// Durable Object namespaces used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#durable_object_namespaces PagesProject#durable_object_namespaces}
	DurableObjectNamespaces interface{} `field:"optional" json:"durableObjectNamespaces" yaml:"durableObjectNamespaces"`
	// Environment variables used for builds and Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#env_vars PagesProject#env_vars}
	EnvVars interface{} `field:"optional" json:"envVars" yaml:"envVars"`
	// Whether to fail open when the deployment config cannot be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#fail_open PagesProject#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
	// Hyperdrive bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#hyperdrive_bindings PagesProject#hyperdrive_bindings}
	HyperdriveBindings interface{} `field:"optional" json:"hyperdriveBindings" yaml:"hyperdriveBindings"`
	// KV namespaces used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#kv_namespaces PagesProject#kv_namespaces}
	KvNamespaces interface{} `field:"optional" json:"kvNamespaces" yaml:"kvNamespaces"`
	// Limits for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#limits PagesProject#limits}
	Limits *PagesProjectDeploymentConfigsPreviewLimits `field:"optional" json:"limits" yaml:"limits"`
	// mTLS bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#mtls_certificates PagesProject#mtls_certificates}
	MtlsCertificates interface{} `field:"optional" json:"mtlsCertificates" yaml:"mtlsCertificates"`
	// Placement setting used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#placement PagesProject#placement}
	Placement *PagesProjectDeploymentConfigsPreviewPlacement `field:"optional" json:"placement" yaml:"placement"`
	// Queue Producer bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#queue_producers PagesProject#queue_producers}
	QueueProducers interface{} `field:"optional" json:"queueProducers" yaml:"queueProducers"`
	// R2 buckets used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#r2_buckets PagesProject#r2_buckets}
	R2Buckets interface{} `field:"optional" json:"r2Buckets" yaml:"r2Buckets"`
	// Services used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#services PagesProject#services}
	Services interface{} `field:"optional" json:"services" yaml:"services"`
	// The usage model for Pages Functions. Available values: "standard", "bundled", "unbound".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#usage_model PagesProject#usage_model}
	UsageModel *string `field:"optional" json:"usageModel" yaml:"usageModel"`
	// Vectorize bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#vectorize_bindings PagesProject#vectorize_bindings}
	VectorizeBindings interface{} `field:"optional" json:"vectorizeBindings" yaml:"vectorizeBindings"`
	// Hash of the Wrangler configuration used for the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#wrangler_config_hash PagesProject#wrangler_config_hash}
	WranglerConfigHash *string `field:"optional" json:"wranglerConfigHash" yaml:"wranglerConfigHash"`
}

