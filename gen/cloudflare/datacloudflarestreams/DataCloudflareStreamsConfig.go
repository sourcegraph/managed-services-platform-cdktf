package datacloudflarestreams

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareStreamsConfig struct {
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
	// The account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#account_id DataCloudflareStreams#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Alias for 'start'. Returns videos created after this date/time (RFC 3339 format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#after DataCloudflareStreams#after}
	After *string `field:"optional" json:"after" yaml:"after"`
	// Lists videos in ascending order of creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#asc DataCloudflareStreams#asc}
	Asc interface{} `field:"optional" json:"asc" yaml:"asc"`
	// Alias for 'end'. Returns videos created before this date/time (RFC 3339 format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#before DataCloudflareStreams#before}
	Before *string `field:"optional" json:"before" yaml:"before"`
	// A user-defined identifier for the media creator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#creator DataCloudflareStreams#creator}
	Creator *string `field:"optional" json:"creator" yaml:"creator"`
	// Lists videos created before the specified date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#end DataCloudflareStreams#end}
	End *string `field:"optional" json:"end" yaml:"end"`
	// Filter by video ID(s). Can be a single ID or a comma-separated list of IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#id DataCloudflareStreams#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Includes the total number of videos associated with the submitted query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#include_counts DataCloudflareStreams#include_counts}
	IncludeCounts interface{} `field:"optional" json:"includeCounts" yaml:"includeCounts"`
	// Maximum number of videos to return (default 1000, max 1000).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#limit DataCloudflareStreams#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Filter by live input ID to find videos associated with a specific live stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#live_input_id DataCloudflareStreams#live_input_id}
	LiveInputId *string `field:"optional" json:"liveInputId" yaml:"liveInputId"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#max_items DataCloudflareStreams#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Filter by video name/UID(s). Can be a single name or a comma-separated list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#name DataCloudflareStreams#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Provides a partial word match of the `name` key in the `meta` field.
	//
	// Slow for medium to large video libraries. May be unavailable for very large libraries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#search DataCloudflareStreams#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Lists videos created after the specified date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#start DataCloudflareStreams#start}
	Start *string `field:"optional" json:"start" yaml:"start"`
	// Specifies the processing status for all quality levels for a video. Available values: "pendingupload", "downloading", "queued", "inprogress", "ready", "error", "live-inprogress".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#status DataCloudflareStreams#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Specifies whether the video is `vod` or `live`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#type DataCloudflareStreams#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Provides a fast, exact string match on the `name` key in the `meta` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/streams#video_name DataCloudflareStreams#video_name}
	VideoName *string `field:"optional" json:"videoName" yaml:"videoName"`
}

