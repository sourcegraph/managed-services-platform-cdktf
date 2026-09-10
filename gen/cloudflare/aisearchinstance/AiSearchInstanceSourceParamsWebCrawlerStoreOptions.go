package aisearchinstance


type AiSearchInstanceSourceParamsWebCrawlerStoreOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#storage_id AiSearchInstance#storage_id}.
	StorageId *string `field:"required" json:"storageId" yaml:"storageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#r2_jurisdiction AiSearchInstance#r2_jurisdiction}.
	R2Jurisdiction *string `field:"optional" json:"r2Jurisdiction" yaml:"r2Jurisdiction"`
	// Available values: "r2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#storage_type AiSearchInstance#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
}

