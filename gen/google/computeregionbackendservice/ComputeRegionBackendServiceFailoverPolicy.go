package computeregionbackendservice


type ComputeRegionBackendServiceFailoverPolicy struct {
	// On failover or failback, this field indicates whether connection drain will be honored.
	//
	// Setting this to true has the following effect: connections
	// to the old active pool are not drained. Connections to the new active pool
	// use the timeout of 10 min (currently fixed). Setting to false has the
	// following effect: both old and new connections will have a drain timeout
	// of 10 min.
	// This can be set to true only if the protocol is TCP.
	// The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_backend_service#disable_connection_drain_on_failover ComputeRegionBackendService#disable_connection_drain_on_failover}
	DisableConnectionDrainOnFailover interface{} `field:"optional" json:"disableConnectionDrainOnFailover" yaml:"disableConnectionDrainOnFailover"`
	// This option is used only when no healthy VMs are detected in the primary and backup instance groups.
	//
	// When set to true, traffic is dropped. When
	// set to false, new connections are sent across all VMs in the primary group.
	// The default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_backend_service#drop_traffic_if_unhealthy ComputeRegionBackendService#drop_traffic_if_unhealthy}
	DropTrafficIfUnhealthy interface{} `field:"optional" json:"dropTrafficIfUnhealthy" yaml:"dropTrafficIfUnhealthy"`
	// The value of the field must be in [0, 1].
	//
	// If the ratio of the healthy
	// VMs in the primary backend is at or below this number, traffic arriving
	// at the load-balanced IP will be directed to the failover backend.
	// In case where 'failoverRatio' is not set or all the VMs in the backup
	// backend are unhealthy, the traffic will be directed back to the primary
	// backend in the "force" mode, where traffic will be spread to the healthy
	// VMs with the best effort, or to all VMs when no VM is healthy.
	// This field is only used with l4 load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_backend_service#failover_ratio ComputeRegionBackendService#failover_ratio}
	FailoverRatio *float64 `field:"optional" json:"failoverRatio" yaml:"failoverRatio"`
}

