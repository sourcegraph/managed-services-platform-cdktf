package googlecomputeinstancetemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeInstanceTemplateConfig struct {
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
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#disk GoogleComputeInstanceTemplate#disk}
	Disk interface{} `field:"required" json:"disk" yaml:"disk"`
	// The machine type to create.
	//
	// To create a machine with a custom type (such as extended memory), format the value like custom-VCPUS-MEM_IN_MB like custom-6-20480 for 6 vCPU and 20GB of RAM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#machine_type GoogleComputeInstanceTemplate#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// advanced_machine_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#advanced_machine_features GoogleComputeInstanceTemplate#advanced_machine_features}
	AdvancedMachineFeatures *GoogleComputeInstanceTemplateAdvancedMachineFeatures `field:"optional" json:"advancedMachineFeatures" yaml:"advancedMachineFeatures"`
	// Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#can_ip_forward GoogleComputeInstanceTemplate#can_ip_forward}
	CanIpForward interface{} `field:"optional" json:"canIpForward" yaml:"canIpForward"`
	// confidential_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#confidential_instance_config GoogleComputeInstanceTemplate#confidential_instance_config}
	ConfidentialInstanceConfig *GoogleComputeInstanceTemplateConfidentialInstanceConfig `field:"optional" json:"confidentialInstanceConfig" yaml:"confidentialInstanceConfig"`
	// A brief description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#description GoogleComputeInstanceTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable Virtual Displays on this instance. Note: allow_stopping_for_update must be set to true in order to update this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#enable_display GoogleComputeInstanceTemplate#enable_display}
	EnableDisplay interface{} `field:"optional" json:"enableDisplay" yaml:"enableDisplay"`
	// guest_accelerator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#guest_accelerator GoogleComputeInstanceTemplate#guest_accelerator}
	GuestAccelerator interface{} `field:"optional" json:"guestAccelerator" yaml:"guestAccelerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#id GoogleComputeInstanceTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A description of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#instance_description GoogleComputeInstanceTemplate#instance_description}
	InstanceDescription *string `field:"optional" json:"instanceDescription" yaml:"instanceDescription"`
	// Action to be taken when a customer's encryption key is revoked.
	//
	// Supports "STOP" and "NONE", with "NONE" being the default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#key_revocation_action_type GoogleComputeInstanceTemplate#key_revocation_action_type}
	KeyRevocationActionType *string `field:"optional" json:"keyRevocationActionType" yaml:"keyRevocationActionType"`
	// A set of key/value label pairs to assign to instances created from this template.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// 				Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#labels GoogleComputeInstanceTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Metadata key/value pairs to make available from within instances created from this template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#metadata GoogleComputeInstanceTemplate#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// An alternative to using the startup-script metadata key, mostly to match the compute_instance resource.
	//
	// This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#metadata_startup_script GoogleComputeInstanceTemplate#metadata_startup_script}
	MetadataStartupScript *string `field:"optional" json:"metadataStartupScript" yaml:"metadataStartupScript"`
	// Specifies a minimum CPU platform.
	//
	// Applicable values are the friendly names of CPU platforms, such as Intel Haswell or Intel Skylake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#min_cpu_platform GoogleComputeInstanceTemplate#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// The name of the instance template. If you leave this blank, Terraform will auto-generate a unique name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#name GoogleComputeInstanceTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Creates a unique name beginning with the specified prefix.
	//
	// Conflicts with name. Max length is 54 characters. Prefixes with lengths longer than 37 characters will use a shortened UUID that will be more prone to collisions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#name_prefix GoogleComputeInstanceTemplate#name_prefix}
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#network_interface GoogleComputeInstanceTemplate#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// network_performance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#network_performance_config GoogleComputeInstanceTemplate#network_performance_config}
	NetworkPerformanceConfig *GoogleComputeInstanceTemplateNetworkPerformanceConfig `field:"optional" json:"networkPerformanceConfig" yaml:"networkPerformanceConfig"`
	// Partner Metadata Map made available within the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#partner_metadata GoogleComputeInstanceTemplate#partner_metadata}
	PartnerMetadata *map[string]*string `field:"optional" json:"partnerMetadata" yaml:"partnerMetadata"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#project GoogleComputeInstanceTemplate#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// An instance template is a global resource that is not bound to a zone or a region.
	//
	// However, you can still specify some regional resources in an instance template, which restricts the template to the region where that resource resides. For example, a custom subnetwork resource is tied to a specific region. Defaults to the region of the Provider if no value is given.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#region GoogleComputeInstanceTemplate#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#reservation_affinity GoogleComputeInstanceTemplate#reservation_affinity}
	ReservationAffinity *GoogleComputeInstanceTemplateReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys and values have the same definition as resource manager tags.
	// 				Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456.
	// 				The field is ignored (both PUT & PATCH) when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#resource_manager_tags GoogleComputeInstanceTemplate#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
	// A list of self_links of resource policies to attach to the instance.
	//
	// Currently a max of 1 resource policy is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#resource_policies GoogleComputeInstanceTemplate#resource_policies}
	ResourcePolicies *[]*string `field:"optional" json:"resourcePolicies" yaml:"resourcePolicies"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#scheduling GoogleComputeInstanceTemplate#scheduling}
	Scheduling *GoogleComputeInstanceTemplateScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#service_account GoogleComputeInstanceTemplate#service_account}
	ServiceAccount *GoogleComputeInstanceTemplateServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#shielded_instance_config GoogleComputeInstanceTemplate#shielded_instance_config}
	ShieldedInstanceConfig *GoogleComputeInstanceTemplateShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// Tags to attach to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#tags GoogleComputeInstanceTemplate#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#timeouts GoogleComputeInstanceTemplate#timeouts}
	Timeouts *GoogleComputeInstanceTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

