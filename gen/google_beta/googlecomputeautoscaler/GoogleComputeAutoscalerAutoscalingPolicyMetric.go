package googlecomputeautoscaler


type GoogleComputeAutoscalerAutoscalingPolicyMetric struct {
	// The identifier (type) of the Stackdriver Monitoring metric. The metric cannot have negative values.
	//
	// The metric must have a value type of INT64 or DOUBLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_autoscaler#name GoogleComputeAutoscaler#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A filter string to be used as the filter string for a Stackdriver Monitoring TimeSeries.list API call. This filter is used to select a specific TimeSeries for the purpose of autoscaling and to determine whether the metric is exporting per-instance or per-group data.
	//
	// You can only use the AND operator for joining selectors.
	// You can only use direct equality comparison operator (=) without
	// any functions for each selector.
	// You can specify the metric in both the filter string and in the
	// metric field. However, if specified in both places, the metric must
	// be identical.
	//
	// The monitored resource type determines what kind of values are
	// expected for the metric. If it is a gce_instance, the autoscaler
	// expects the metric to include a separate TimeSeries for each
	// instance in a group. In such a case, you cannot filter on resource
	// labels.
	//
	// If the resource type is any other value, the autoscaler expects
	// this metric to contain values that apply to the entire autoscaled
	// instance group and resource label filtering can be performed to
	// point autoscaler at the correct TimeSeries to scale upon.
	// This is called a per-group metric for the purpose of autoscaling.
	//
	// If not specified, the type defaults to gce_instance.
	//
	// You should provide a filter that is selective enough to pick just
	// one TimeSeries for the autoscaled group or for each of the instances
	// (if you are using gce_instance resource type). If multiple
	// TimeSeries are returned upon the query execution, the autoscaler
	// will sum their respective values to obtain its scaling value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_autoscaler#filter GoogleComputeAutoscaler#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// If scaling is based on a per-group metric value that represents the total amount of work to be done or resource usage, set this value to an amount assigned for a single instance of the scaled group.
	//
	// The autoscaler will keep the number of instances proportional to the
	// value of this metric, the metric itself should not change value due
	// to group resizing.
	//
	// For example, a good metric to use with the target is
	// 'pubsub.googleapis.com/subscription/num_undelivered_messages'
	// or a custom metric exporting the total number of requests coming to
	// your instances.
	//
	// A bad example would be a metric exporting an average or median
	// latency, since this value can't include a chunk assignable to a
	// single instance, it could be better used with utilization_target
	// instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_autoscaler#single_instance_assignment GoogleComputeAutoscaler#single_instance_assignment}
	SingleInstanceAssignment *float64 `field:"optional" json:"singleInstanceAssignment" yaml:"singleInstanceAssignment"`
	// The target value of the metric that autoscaler should maintain.
	//
	// This must be a positive value. A utilization
	// metric scales number of virtual machines handling requests
	// to increase or decrease proportionally to the metric.
	//
	// For example, a good metric to use as a utilizationTarget is
	// www.googleapis.com/compute/instance/network/received_bytes_count.
	// The autoscaler will work to keep this value constant for each
	// of the instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_autoscaler#target GoogleComputeAutoscaler#target}
	Target *float64 `field:"optional" json:"target" yaml:"target"`
	// Defines how target utilization value is expressed for a Stackdriver Monitoring metric. Possible values: ["GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_autoscaler#type GoogleComputeAutoscaler#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

