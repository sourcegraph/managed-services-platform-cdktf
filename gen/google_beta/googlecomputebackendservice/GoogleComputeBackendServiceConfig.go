package googlecomputebackendservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeBackendServiceConfig struct {
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
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#name GoogleComputeBackendService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE.
	//
	// If set to 0, the cookie is non-persistent and lasts
	// only until the end of the browser session (or equivalent). The
	// maximum allowed value for TTL is one day.
	//
	// When the load balancing scheme is INTERNAL, this field is not used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#affinity_cookie_ttl_sec GoogleComputeBackendService#affinity_cookie_ttl_sec}
	AffinityCookieTtlSec *float64 `field:"optional" json:"affinityCookieTtlSec" yaml:"affinityCookieTtlSec"`
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#backend GoogleComputeBackendService#backend}
	Backend interface{} `field:"optional" json:"backend" yaml:"backend"`
	// cdn_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#cdn_policy GoogleComputeBackendService#cdn_policy}
	CdnPolicy *GoogleComputeBackendServiceCdnPolicy `field:"optional" json:"cdnPolicy" yaml:"cdnPolicy"`
	// circuit_breakers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#circuit_breakers GoogleComputeBackendService#circuit_breakers}
	CircuitBreakers *GoogleComputeBackendServiceCircuitBreakers `field:"optional" json:"circuitBreakers" yaml:"circuitBreakers"`
	// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header. Possible values: ["AUTOMATIC", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#compression_mode GoogleComputeBackendService#compression_mode}
	CompressionMode *string `field:"optional" json:"compressionMode" yaml:"compressionMode"`
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#connection_draining_timeout_sec GoogleComputeBackendService#connection_draining_timeout_sec}
	ConnectionDrainingTimeoutSec *float64 `field:"optional" json:"connectionDrainingTimeoutSec" yaml:"connectionDrainingTimeoutSec"`
	// consistent_hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#consistent_hash GoogleComputeBackendService#consistent_hash}
	ConsistentHash *GoogleComputeBackendServiceConsistentHash `field:"optional" json:"consistentHash" yaml:"consistentHash"`
	// custom_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#custom_metrics GoogleComputeBackendService#custom_metrics}
	CustomMetrics interface{} `field:"optional" json:"customMetrics" yaml:"customMetrics"`
	// Headers that the HTTP/S load balancer should add to proxied requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#custom_request_headers GoogleComputeBackendService#custom_request_headers}
	CustomRequestHeaders *[]*string `field:"optional" json:"customRequestHeaders" yaml:"customRequestHeaders"`
	// Headers that the HTTP/S load balancer should add to proxied responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#custom_response_headers GoogleComputeBackendService#custom_response_headers}
	CustomResponseHeaders *[]*string `field:"optional" json:"customResponseHeaders" yaml:"customResponseHeaders"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#description GoogleComputeBackendService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dynamic_forwarding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#dynamic_forwarding GoogleComputeBackendService#dynamic_forwarding}
	DynamicForwarding *GoogleComputeBackendServiceDynamicForwarding `field:"optional" json:"dynamicForwarding" yaml:"dynamicForwarding"`
	// The resource URL for the edge security policy associated with this backend service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#edge_security_policy GoogleComputeBackendService#edge_security_policy}
	EdgeSecurityPolicy *string `field:"optional" json:"edgeSecurityPolicy" yaml:"edgeSecurityPolicy"`
	// If true, enable Cloud CDN for this BackendService.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#enable_cdn GoogleComputeBackendService#enable_cdn}
	EnableCdn interface{} `field:"optional" json:"enableCdn" yaml:"enableCdn"`
	// Specifies the canary migration state. Possible values are PREPARE, TEST_BY_PERCENTAGE, and TEST_ALL_TRAFFIC.
	//
	// To begin the migration from EXTERNAL to EXTERNAL_MANAGED, the state must be changed to
	// PREPARE. The state must be changed to TEST_ALL_TRAFFIC before the loadBalancingScheme can be
	// changed to EXTERNAL_MANAGED. Optionally, the TEST_BY_PERCENTAGE state can be used to migrate
	// traffic by percentage using externalManagedMigrationTestingPercentage.
	//
	// Rolling back a migration requires the states to be set in reverse order. So changing the
	// scheme from EXTERNAL_MANAGED to EXTERNAL requires the state to be set to TEST_ALL_TRAFFIC at
	// the same time. Optionally, the TEST_BY_PERCENTAGE state can be used to migrate some traffic
	// back to EXTERNAL or PREPARE can be used to migrate all traffic back to EXTERNAL. Possible values: ["PREPARE", "TEST_BY_PERCENTAGE", "TEST_ALL_TRAFFIC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#external_managed_migration_state GoogleComputeBackendService#external_managed_migration_state}
	ExternalManagedMigrationState *string `field:"optional" json:"externalManagedMigrationState" yaml:"externalManagedMigrationState"`
	// Determines the fraction of requests that should be processed by the Global external Application Load Balancer.
	//
	// The value of this field must be in the range [0, 100].
	//
	// Session affinity options will slightly affect this routing behavior, for more details,
	// see: Session Affinity.
	//
	// This value can only be set if the loadBalancingScheme in the backend service is set to
	// EXTERNAL (when using the Classic ALB) and the migration state is TEST_BY_PERCENTAGE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#external_managed_migration_testing_percentage GoogleComputeBackendService#external_managed_migration_testing_percentage}
	ExternalManagedMigrationTestingPercentage *float64 `field:"optional" json:"externalManagedMigrationTestingPercentage" yaml:"externalManagedMigrationTestingPercentage"`
	// The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource for health checking this BackendService.
	//
	// Currently at most one health
	// check can be specified.
	//
	// A health check must be specified unless the backend service uses an internet
	// or serverless NEG as a backend.
	//
	// For internal load balancing, a URL to a HealthCheck resource must be specified instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#health_checks GoogleComputeBackendService#health_checks}
	HealthChecks *[]*string `field:"optional" json:"healthChecks" yaml:"healthChecks"`
	// iap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#iap GoogleComputeBackendService#iap}
	Iap *GoogleComputeBackendServiceIap `field:"optional" json:"iap" yaml:"iap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#id GoogleComputeBackendService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies preference of traffic to the backend (from the proxy and from the client for proxyless gRPC).
	//
	// Possible values: ["IPV4_ONLY", "PREFER_IPV6", "IPV6_ONLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#ip_address_selection_policy GoogleComputeBackendService#ip_address_selection_policy}
	IpAddressSelectionPolicy *string `field:"optional" json:"ipAddressSelectionPolicy" yaml:"ipAddressSelectionPolicy"`
	// Indicates whether the backend service will be used with internal or external load balancing.
	//
	// A backend service created for one type of
	// load balancing cannot be used with the other. For more information, refer to
	// [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service). Default value: "EXTERNAL" Possible values: ["EXTERNAL", "INTERNAL_SELF_MANAGED", "INTERNAL_MANAGED", "EXTERNAL_MANAGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#load_balancing_scheme GoogleComputeBackendService#load_balancing_scheme}
	LoadBalancingScheme *string `field:"optional" json:"loadBalancingScheme" yaml:"loadBalancingScheme"`
	// locality_lb_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#locality_lb_policies GoogleComputeBackendService#locality_lb_policies}
	LocalityLbPolicies interface{} `field:"optional" json:"localityLbPolicies" yaml:"localityLbPolicies"`
	// The load balancing algorithm used within the scope of the locality. The possible values are:.
	//
	// * 'ROUND_ROBIN': This is a simple policy in which each healthy backend
	//                  is selected in round robin order.
	//
	// * 'LEAST_REQUEST': An O(1) algorithm which selects two random healthy
	//                    hosts and picks the host which has fewer active requests.
	//
	// * 'RING_HASH': The ring/modulo hash load balancer implements consistent
	//                hashing to backends. The algorithm has the property that the
	//                addition/removal of a host from a set of N hosts only affects
	//                1/N of the requests.
	//
	// * 'RANDOM': The load balancer selects a random healthy host.
	//
	// * 'ORIGINAL_DESTINATION': Backend host is selected based on the client
	//                           connection metadata, i.e., connections are opened
	//                           to the same address as the destination address of
	//                           the incoming connection before the connection
	//                           was redirected to the load balancer.
	//
	// * 'MAGLEV': used as a drop in replacement for the ring hash load balancer.
	//             Maglev is not as stable as ring hash but has faster table lookup
	//             build times and host selection times. For more information about
	//             Maglev, refer to https://ai.google/research/pubs/pub44824
	//
	// * 'WEIGHTED_MAGLEV': Per-instance weighted Load Balancing via health check
	//                      reported weights. Only applicable to loadBalancingScheme
	//                      EXTERNAL. If set, the Backend Service must
	//                      configure a non legacy HTTP-based Health Check, and
	//                      health check replies are expected to contain
	//                      non-standard HTTP response header field
	//                      X-Load-Balancing-Endpoint-Weight to specify the
	//                      per-instance weights. If set, Load Balancing is weight
	//                      based on the per-instance weights reported in the last
	//                      processed health check replies, as long as every
	//                      instance either reported a valid weight or had
	//                      UNAVAILABLE_WEIGHT. Otherwise, Load Balancing remains
	//                      equal-weight.
	//
	// * 'WEIGHTED_ROUND_ROBIN': Per-endpoint weighted round-robin Load Balancing using weights computed
	//                           from Backend reported Custom Metrics. If set, the Backend Service
	//                           responses are expected to contain non-standard HTTP response header field
	//                           X-Endpoint-Load-Metrics. The reported metrics
	//                           to use for computing the weights are specified via the
	//                           backends[].customMetrics fields.
	//
	// locality_lb_policy is applicable to either:
	//
	// * A regional backend service with the service_protocol set to HTTP, HTTPS, HTTP2 or H2C,
	//   and loadBalancingScheme set to INTERNAL_MANAGED.
	// * A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	// * A regional backend service with loadBalancingScheme set to EXTERNAL (External Network
	//   Load Balancing). Only MAGLEV and WEIGHTED_MAGLEV values are possible for External
	//   Network Load Balancing. The default is MAGLEV.
	//
	// If session_affinity is not NONE, and locality_lb_policy is not set to MAGLEV, WEIGHTED_MAGLEV,
	// or RING_HASH, session affinity settings will not take effect.
	//
	// Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced
	// by a URL map that is bound to target gRPC proxy that has validate_for_proxyless
	// field set to true. Possible values: ["ROUND_ROBIN", "LEAST_REQUEST", "RING_HASH", "RANDOM", "ORIGINAL_DESTINATION", "MAGLEV", "WEIGHTED_MAGLEV", "WEIGHTED_ROUND_ROBIN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#locality_lb_policy GoogleComputeBackendService#locality_lb_policy}
	LocalityLbPolicy *string `field:"optional" json:"localityLbPolicy" yaml:"localityLbPolicy"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#log_config GoogleComputeBackendService#log_config}
	LogConfig *GoogleComputeBackendServiceLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// max_stream_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#max_stream_duration GoogleComputeBackendService#max_stream_duration}
	MaxStreamDuration *GoogleComputeBackendServiceMaxStreamDuration `field:"optional" json:"maxStreamDuration" yaml:"maxStreamDuration"`
	// network_pass_through_lb_traffic_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#network_pass_through_lb_traffic_policy GoogleComputeBackendService#network_pass_through_lb_traffic_policy}
	NetworkPassThroughLbTrafficPolicy *GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicy `field:"optional" json:"networkPassThroughLbTrafficPolicy" yaml:"networkPassThroughLbTrafficPolicy"`
	// outlier_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#outlier_detection GoogleComputeBackendService#outlier_detection}
	OutlierDetection *GoogleComputeBackendServiceOutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// Name of backend port.
	//
	// The same name should appear in the instance
	// groups referenced by this service. Required when the load balancing
	// scheme is EXTERNAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#port_name GoogleComputeBackendService#port_name}
	PortName *string `field:"optional" json:"portName" yaml:"portName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#project GoogleComputeBackendService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The protocol this BackendService uses to communicate with backends.
	//
	// The default is HTTP. Possible values are HTTP, HTTPS, HTTP2, H2C, TCP, SSL, UDP
	// or GRPC. Refer to the documentation for the load balancers or for Traffic Director
	// for more information. Must be set to GRPC when the backend service is referenced
	// by a URL map that is bound to target gRPC proxy. Possible values: ["HTTP", "HTTPS", "HTTP2", "TCP", "SSL", "UDP", "GRPC", "UNSPECIFIED", "H2C"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#protocol GoogleComputeBackendService#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The security policy associated with this backend service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#security_policy GoogleComputeBackendService#security_policy}
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// security_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#security_settings GoogleComputeBackendService#security_settings}
	SecuritySettings *GoogleComputeBackendServiceSecuritySettings `field:"optional" json:"securitySettings" yaml:"securitySettings"`
	// URL to networkservices.ServiceLbPolicy resource. Can only be set if load balancing scheme is EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED or INTERNAL_SELF_MANAGED and the scope is global.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#service_lb_policy GoogleComputeBackendService#service_lb_policy}
	ServiceLbPolicy *string `field:"optional" json:"serviceLbPolicy" yaml:"serviceLbPolicy"`
	// Type of session affinity to use.
	//
	// The default is NONE. Session affinity is
	// not applicable if the protocol is UDP. Possible values: ["NONE", "CLIENT_IP", "CLIENT_IP_PORT_PROTO", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "HEADER_FIELD", "HTTP_COOKIE", "STRONG_COOKIE_AFFINITY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#session_affinity GoogleComputeBackendService#session_affinity}
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// strong_session_affinity_cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#strong_session_affinity_cookie GoogleComputeBackendService#strong_session_affinity_cookie}
	StrongSessionAffinityCookie *GoogleComputeBackendServiceStrongSessionAffinityCookie `field:"optional" json:"strongSessionAffinityCookie" yaml:"strongSessionAffinityCookie"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#timeouts GoogleComputeBackendService#timeouts}
	Timeouts *GoogleComputeBackendServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The backend service timeout has a different meaning depending on the type of load balancer.
	//
	// For more information see, [Backend service settings](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices).
	// The default is 30 seconds.
	// The full range of timeout values allowed goes from 1 through 2,147,483,647 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#timeout_sec GoogleComputeBackendService#timeout_sec}
	TimeoutSec *float64 `field:"optional" json:"timeoutSec" yaml:"timeoutSec"`
	// tls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#tls_settings GoogleComputeBackendService#tls_settings}
	TlsSettings *GoogleComputeBackendServiceTlsSettings `field:"optional" json:"tlsSettings" yaml:"tlsSettings"`
}

