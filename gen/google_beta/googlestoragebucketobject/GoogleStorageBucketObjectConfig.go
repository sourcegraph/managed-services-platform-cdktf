package googlestoragebucketobject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleStorageBucketObjectConfig struct {
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
	// The name of the containing bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#bucket GoogleStorageBucketObject#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#name GoogleStorageBucketObject#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Cache-Control directive to specify caching behavior of object data.
	//
	// If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#cache_control GoogleStorageBucketObject#cache_control}
	CacheControl *string `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// Data as string to be uploaded.
	//
	// Must be defined if source is not. Note: The content field is marked as sensitive. To view the raw contents of the object, please define an output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#content GoogleStorageBucketObject#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Content-Disposition of the object data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#content_disposition GoogleStorageBucketObject#content_disposition}
	ContentDisposition *string `field:"optional" json:"contentDisposition" yaml:"contentDisposition"`
	// Content-Encoding of the object data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#content_encoding GoogleStorageBucketObject#content_encoding}
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	// Content-Language of the object data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#content_language GoogleStorageBucketObject#content_language}
	ContentLanguage *string `field:"optional" json:"contentLanguage" yaml:"contentLanguage"`
	// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#content_type GoogleStorageBucketObject#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// customer_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#customer_encryption GoogleStorageBucketObject#customer_encryption}
	CustomerEncryption *GoogleStorageBucketObjectCustomerEncryption `field:"optional" json:"customerEncryption" yaml:"customerEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#detect_md5hash GoogleStorageBucketObject#detect_md5hash}.
	DetectMd5Hash *string `field:"optional" json:"detectMd5Hash" yaml:"detectMd5Hash"`
	// Whether an object is under event-based hold.
	//
	// Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#event_based_hold GoogleStorageBucketObject#event_based_hold}
	EventBasedHold interface{} `field:"optional" json:"eventBasedHold" yaml:"eventBasedHold"`
	// Flag to set empty Content-Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#force_empty_content_type GoogleStorageBucketObject#force_empty_content_type}
	ForceEmptyContentType interface{} `field:"optional" json:"forceEmptyContentType" yaml:"forceEmptyContentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#id GoogleStorageBucketObject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource name of the Cloud KMS key that will be used to encrypt the object.
	//
	// Overrides the object metadata's kmsKeyName value, if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#kms_key_name GoogleStorageBucketObject#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// User-provided metadata, in key/value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#metadata GoogleStorageBucketObject#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#retention GoogleStorageBucketObject#retention}
	Retention *GoogleStorageBucketObjectRetention `field:"optional" json:"retention" yaml:"retention"`
	// A path to the data you want to upload. Must be defined if content is not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#source GoogleStorageBucketObject#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// User-provided md5hash, Base 64 MD5 hash of the object data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#source_md5hash GoogleStorageBucketObject#source_md5hash}
	SourceMd5Hash *string `field:"optional" json:"sourceMd5Hash" yaml:"sourceMd5Hash"`
	// The StorageClass of the new bucket object.
	//
	// Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#storage_class GoogleStorageBucketObject#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Whether an object is under temporary hold.
	//
	// While this flag is set to true, the object is protected against deletion and overwrites.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#temporary_hold GoogleStorageBucketObject#temporary_hold}
	TemporaryHold interface{} `field:"optional" json:"temporaryHold" yaml:"temporaryHold"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_bucket_object#timeouts GoogleStorageBucketObject#timeouts}
	Timeouts *GoogleStorageBucketObjectTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

