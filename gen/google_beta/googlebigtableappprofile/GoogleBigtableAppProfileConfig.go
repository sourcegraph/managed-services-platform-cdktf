package googlebigtableappprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigtableAppProfileConfig struct {
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
	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#app_profile_id GoogleBigtableAppProfile#app_profile_id}
	AppProfileId *string `field:"required" json:"appProfileId" yaml:"appProfileId"`
	// data_boost_isolation_read_only block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#data_boost_isolation_read_only GoogleBigtableAppProfile#data_boost_isolation_read_only}
	DataBoostIsolationReadOnly *GoogleBigtableAppProfileDataBoostIsolationReadOnly `field:"optional" json:"dataBoostIsolationReadOnly" yaml:"dataBoostIsolationReadOnly"`
	// Long form description of the use case for this app profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#description GoogleBigtableAppProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#id GoogleBigtableAppProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, ignore safety checks when deleting/updating the app profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#ignore_warnings GoogleBigtableAppProfile#ignore_warnings}
	IgnoreWarnings interface{} `field:"optional" json:"ignoreWarnings" yaml:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#instance GoogleBigtableAppProfile#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// The set of clusters to route to.
	//
	// The order is ignored; clusters will be tried in order of distance. If left empty, all clusters are eligible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#multi_cluster_routing_cluster_ids GoogleBigtableAppProfile#multi_cluster_routing_cluster_ids}
	MultiClusterRoutingClusterIds *[]*string `field:"optional" json:"multiClusterRoutingClusterIds" yaml:"multiClusterRoutingClusterIds"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available in the event of transient errors or delays.
	//
	// Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#multi_cluster_routing_use_any GoogleBigtableAppProfile#multi_cluster_routing_use_any}
	MultiClusterRoutingUseAny interface{} `field:"optional" json:"multiClusterRoutingUseAny" yaml:"multiClusterRoutingUseAny"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#project GoogleBigtableAppProfile#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Must be used with multi-cluster routing.
	//
	// If true, then this app profile will use row affinity sticky routing. With row affinity, Bigtable will route single row key requests based on the row key, rather than randomly. Instead, each row key will be assigned to a cluster by Cloud Bigtable, and will stick to that cluster. Choosing this option improves read-your-writes consistency for most requests under most circumstances, without sacrificing availability. Consistency is not guaranteed, as requests may still fail over between clusters in the event of errors or latency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#row_affinity GoogleBigtableAppProfile#row_affinity}
	RowAffinity interface{} `field:"optional" json:"rowAffinity" yaml:"rowAffinity"`
	// single_cluster_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#single_cluster_routing GoogleBigtableAppProfile#single_cluster_routing}
	SingleClusterRouting *GoogleBigtableAppProfileSingleClusterRouting `field:"optional" json:"singleClusterRouting" yaml:"singleClusterRouting"`
	// standard_isolation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#standard_isolation GoogleBigtableAppProfile#standard_isolation}
	StandardIsolation *GoogleBigtableAppProfileStandardIsolation `field:"optional" json:"standardIsolation" yaml:"standardIsolation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_app_profile#timeouts GoogleBigtableAppProfile#timeouts}
	Timeouts *GoogleBigtableAppProfileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

