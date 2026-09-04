package aisearchinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AiSearchInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#account_id AiSearchInstance#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Use your AI Search ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#id AiSearchInstance#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#source AiSearchInstance#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Available values: "r2", "web-crawler".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#type AiSearchInstance#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#ai_gateway_id AiSearchInstance#ai_gateway_id}.
	AiGatewayId *string `field:"optional" json:"aiGatewayId" yaml:"aiGatewayId"`
	// Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#aisearch_model AiSearchInstance#aisearch_model}
	AisearchModel *string `field:"optional" json:"aisearchModel" yaml:"aisearchModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#cache AiSearchInstance#cache}.
	Cache interface{} `field:"optional" json:"cache" yaml:"cache"`
	// Available values: "super_strict_match", "close_enough", "flexible_friend", "anything_goes".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#cache_threshold AiSearchInstance#cache_threshold}
	CacheThreshold *string `field:"optional" json:"cacheThreshold" yaml:"cacheThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#chunk AiSearchInstance#chunk}.
	Chunk interface{} `field:"optional" json:"chunk" yaml:"chunk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#chunk_overlap AiSearchInstance#chunk_overlap}.
	ChunkOverlap *float64 `field:"optional" json:"chunkOverlap" yaml:"chunkOverlap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#chunk_size AiSearchInstance#chunk_size}.
	ChunkSize *float64 `field:"optional" json:"chunkSize" yaml:"chunkSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#custom_metadata AiSearchInstance#custom_metadata}.
	CustomMetadata interface{} `field:"optional" json:"customMetadata" yaml:"customMetadata"`
	// Available values: "@cf/qwen/qwen3-embedding-0.6b", "@cf/baai/bge-m3", "@cf/baai/bge-large-en-v1.5", "@cf/google/embeddinggemma-300m", "google-ai-studio/gemini-embedding-001", "openai/text-embedding-3-small", "openai/text-embedding-3-large", "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#embedding_model AiSearchInstance#embedding_model}
	EmbeddingModel *string `field:"optional" json:"embeddingModel" yaml:"embeddingModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#hybrid_search_enabled AiSearchInstance#hybrid_search_enabled}.
	HybridSearchEnabled interface{} `field:"optional" json:"hybridSearchEnabled" yaml:"hybridSearchEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#max_num_results AiSearchInstance#max_num_results}.
	MaxNumResults *float64 `field:"optional" json:"maxNumResults" yaml:"maxNumResults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#metadata AiSearchInstance#metadata}.
	Metadata *AiSearchInstanceMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#paused AiSearchInstance#paused}.
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#public_endpoint_params AiSearchInstance#public_endpoint_params}.
	PublicEndpointParams *AiSearchInstancePublicEndpointParams `field:"optional" json:"publicEndpointParams" yaml:"publicEndpointParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#reranking AiSearchInstance#reranking}.
	Reranking interface{} `field:"optional" json:"reranking" yaml:"reranking"`
	// Available values: "@cf/baai/bge-reranker-base", "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#reranking_model AiSearchInstance#reranking_model}
	RerankingModel *string `field:"optional" json:"rerankingModel" yaml:"rerankingModel"`
	// Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#rewrite_model AiSearchInstance#rewrite_model}
	RewriteModel *string `field:"optional" json:"rewriteModel" yaml:"rewriteModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#rewrite_query AiSearchInstance#rewrite_query}.
	RewriteQuery interface{} `field:"optional" json:"rewriteQuery" yaml:"rewriteQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#score_threshold AiSearchInstance#score_threshold}.
	ScoreThreshold *float64 `field:"optional" json:"scoreThreshold" yaml:"scoreThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#source_params AiSearchInstance#source_params}.
	SourceParams *AiSearchInstanceSourceParams `field:"optional" json:"sourceParams" yaml:"sourceParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#summarization AiSearchInstance#summarization}.
	Summarization interface{} `field:"optional" json:"summarization" yaml:"summarization"`
	// Available values: "@cf/meta/llama-3.3-70b-instruct-fp8-fast", "@cf/meta/llama-3.1-8b-instruct-fast", "@cf/meta/llama-3.1-8b-instruct-fp8", "@cf/meta/llama-4-scout-17b-16e-instruct", "@cf/qwen/qwen3-30b-a3b-fp8", "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", "@cf/moonshotai/kimi-k2-instruct", "anthropic/claude-3-7-sonnet", "anthropic/claude-sonnet-4", "anthropic/claude-opus-4", "anthropic/claude-3-5-haiku", "cerebras/qwen-3-235b-a22b-instruct", "cerebras/qwen-3-235b-a22b-thinking", "cerebras/llama-3.3-70b", "cerebras/llama-4-maverick-17b-128e-instruct", "cerebras/llama-4-scout-17b-16e-instruct", "cerebras/gpt-oss-120b", "google-ai-studio/gemini-2.5-flash", "google-ai-studio/gemini-2.5-pro", "grok/grok-4", "groq/llama-3.3-70b-versatile", "groq/llama-3.1-8b-instant", "openai/gpt-5", "openai/gpt-5-mini", "openai/gpt-5-nano", "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#summarization_model AiSearchInstance#summarization_model}
	SummarizationModel *string `field:"optional" json:"summarizationModel" yaml:"summarizationModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#system_prompt_aisearch AiSearchInstance#system_prompt_aisearch}.
	SystemPromptAisearch *string `field:"optional" json:"systemPromptAisearch" yaml:"systemPromptAisearch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#system_prompt_index_summarization AiSearchInstance#system_prompt_index_summarization}.
	SystemPromptIndexSummarization *string `field:"optional" json:"systemPromptIndexSummarization" yaml:"systemPromptIndexSummarization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#system_prompt_rewrite_query AiSearchInstance#system_prompt_rewrite_query}.
	SystemPromptRewriteQuery *string `field:"optional" json:"systemPromptRewriteQuery" yaml:"systemPromptRewriteQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#token_id AiSearchInstance#token_id}.
	TokenId *string `field:"optional" json:"tokenId" yaml:"tokenId"`
}

