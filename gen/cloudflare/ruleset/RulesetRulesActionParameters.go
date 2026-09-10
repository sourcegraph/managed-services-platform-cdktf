package ruleset


type RulesetRulesActionParameters struct {
	// A list of additional ports that caching should be enabled on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#additional_cacheable_ports Ruleset#additional_cacheable_ports}
	AdditionalCacheablePorts *[]*float64 `field:"optional" json:"additionalCacheablePorts" yaml:"additionalCacheablePorts"`
	// Custom order for compression algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#algorithms Ruleset#algorithms}
	Algorithms interface{} `field:"optional" json:"algorithms" yaml:"algorithms"`
	// The name of a custom asset to serve as the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#asset_name Ruleset#asset_name}
	AssetName *string `field:"optional" json:"assetName" yaml:"assetName"`
	// Whether to enable Automatic HTTPS Rewrites.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#automatic_https_rewrites Ruleset#automatic_https_rewrites}
	AutomaticHttpsRewrites interface{} `field:"optional" json:"automaticHttpsRewrites" yaml:"automaticHttpsRewrites"`
	// Which file extensions to minify automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#autominify Ruleset#autominify}
	Autominify *RulesetRulesActionParametersAutominify `field:"optional" json:"autominify" yaml:"autominify"`
	// Whether to enable Browser Integrity Check (BIC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#bic Ruleset#bic}
	Bic interface{} `field:"optional" json:"bic" yaml:"bic"`
	// How long client browsers should cache the response.
	//
	// Cloudflare cache purge will not purge content cached on client browsers, so high browser TTLs may lead to stale content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#browser_ttl Ruleset#browser_ttl}
	BrowserTtl *RulesetRulesActionParametersBrowserTtl `field:"optional" json:"browserTtl" yaml:"browserTtl"`
	// Whether the request's response from the origin is eligible for caching.
	//
	// Caching itself will still depend on the cache control header and your other caching configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#cache Ruleset#cache}
	Cache interface{} `field:"optional" json:"cache" yaml:"cache"`
	// Which components of the request are included in or excluded from the cache key Cloudflare uses to store the response in cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#cache_key Ruleset#cache_key}
	CacheKey *RulesetRulesActionParametersCacheKey `field:"optional" json:"cacheKey" yaml:"cacheKey"`
	// Settings to determine whether the request's response from origin is eligible for Cache Reserve (requires a Cache Reserve add-on plan).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#cache_reserve Ruleset#cache_reserve}
	CacheReserve *RulesetRulesActionParametersCacheReserve `field:"optional" json:"cacheReserve" yaml:"cacheReserve"`
	// The response content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#content Ruleset#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Whether to enable content conversion (e.g., HTML to Markdown).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#content_converter Ruleset#content_converter}
	ContentConverter interface{} `field:"optional" json:"contentConverter" yaml:"contentConverter"`
	// The content type header to set with the error response. Available values: "application/json", "text/html", "text/plain", "text/xml".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#content_type Ruleset#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The cookie fields to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#cookie_fields Ruleset#cookie_fields}
	CookieFields interface{} `field:"optional" json:"cookieFields" yaml:"cookieFields"`
	// Whether to disable Cloudflare Apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#disable_apps Ruleset#disable_apps}
	DisableApps interface{} `field:"optional" json:"disableApps" yaml:"disableApps"`
	// Whether to disable Real User Monitoring (RUM).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#disable_rum Ruleset#disable_rum}
	DisableRum interface{} `field:"optional" json:"disableRum" yaml:"disableRum"`
	// Whether to disable Zaraz.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#disable_zaraz Ruleset#disable_zaraz}
	DisableZaraz interface{} `field:"optional" json:"disableZaraz" yaml:"disableZaraz"`
	// How long the Cloudflare edge network should cache the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#edge_ttl Ruleset#edge_ttl}
	EdgeTtl *RulesetRulesActionParametersEdgeTtl `field:"optional" json:"edgeTtl" yaml:"edgeTtl"`
	// Whether to enable Email Obfuscation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#email_obfuscation Ruleset#email_obfuscation}
	EmailObfuscation interface{} `field:"optional" json:"emailObfuscation" yaml:"emailObfuscation"`
	// An expression to generate cache tags for set_cache_tags action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#expression Ruleset#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Whether to enable Cloudflare Fonts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#fonts Ruleset#fonts}
	Fonts interface{} `field:"optional" json:"fonts" yaml:"fonts"`
	// A redirect based on a bulk list lookup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#from_list Ruleset#from_list}
	FromList *RulesetRulesActionParametersFromListStruct `field:"optional" json:"fromList" yaml:"fromList"`
	// A redirect based on the request properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#from_value Ruleset#from_value}
	FromValue *RulesetRulesActionParametersFromValue `field:"optional" json:"fromValue" yaml:"fromValue"`
	// A map of headers to rewrite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#headers Ruleset#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// A value to rewrite the HTTP host header to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#host_header Ruleset#host_header}
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Whether to enable Hotlink Protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#hotlink_protection Ruleset#hotlink_protection}
	HotlinkProtection interface{} `field:"optional" json:"hotlinkProtection" yaml:"hotlinkProtection"`
	// The ID of the ruleset to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#id Ruleset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set the immutable cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#immutable Ruleset#immutable}
	Immutable *RulesetRulesActionParametersImmutable `field:"optional" json:"immutable" yaml:"immutable"`
	// A delta to change the score by, which can be either positive or negative.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#increment Ruleset#increment}
	Increment *float64 `field:"optional" json:"increment" yaml:"increment"`
	// The configuration to use for matched data logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#matched_data Ruleset#matched_data}
	MatchedData *RulesetRulesActionParametersMatchedData `field:"optional" json:"matchedData" yaml:"matchedData"`
	// Set the max-age cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#max_age Ruleset#max_age}
	MaxAge *RulesetRulesActionParametersMaxAge `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Whether to enable Mirage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#mirage Ruleset#mirage}
	Mirage interface{} `field:"optional" json:"mirage" yaml:"mirage"`
	// Set the must-revalidate cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#must_revalidate Ruleset#must_revalidate}
	MustRevalidate *RulesetRulesActionParametersMustRevalidate `field:"optional" json:"mustRevalidate" yaml:"mustRevalidate"`
	// Set the must-understand cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#must_understand Ruleset#must_understand}
	MustUnderstand *RulesetRulesActionParametersMustUnderstand `field:"optional" json:"mustUnderstand" yaml:"mustUnderstand"`
	// Set the no-cache cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#no_cache Ruleset#no_cache}
	NoCache *RulesetRulesActionParametersNoCache `field:"optional" json:"noCache" yaml:"noCache"`
	// Set the no-store cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#no_store Ruleset#no_store}
	NoStore *RulesetRulesActionParametersNoStore `field:"optional" json:"noStore" yaml:"noStore"`
	// Set the no-transform cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#no_transform Ruleset#no_transform}
	NoTransform *RulesetRulesActionParametersNoTransform `field:"optional" json:"noTransform" yaml:"noTransform"`
	// The operation to perform for set_cache_tags action. Available values: "set", "add", "remove".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#operation Ruleset#operation}
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Whether to enable Opportunistic Encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#opportunistic_encryption Ruleset#opportunistic_encryption}
	OpportunisticEncryption interface{} `field:"optional" json:"opportunisticEncryption" yaml:"opportunisticEncryption"`
	// An origin to route to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#origin Ruleset#origin}
	Origin *RulesetRulesActionParametersOrigin `field:"optional" json:"origin" yaml:"origin"`
	// Whether Cloudflare will aim to strictly adhere to RFC 7234.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#origin_cache_control Ruleset#origin_cache_control}
	OriginCacheControl interface{} `field:"optional" json:"originCacheControl" yaml:"originCacheControl"`
	// Whether to generate Cloudflare error pages for issues from the origin server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#origin_error_page_passthru Ruleset#origin_error_page_passthru}
	OriginErrorPagePassthru interface{} `field:"optional" json:"originErrorPagePassthru" yaml:"originErrorPagePassthru"`
	// A set of overrides to apply to the target ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#overrides Ruleset#overrides}
	Overrides *RulesetRulesActionParametersOverrides `field:"optional" json:"overrides" yaml:"overrides"`
	// A list of phases to skip the execution of.
	//
	// This option is incompatible with the rulesets option.
	// Available values: "ddos_l4", "ddos_l7", "http_config_settings", "http_custom_errors", "http_log_custom_fields", "http_ratelimit", "http_request_cache_settings", "http_request_dynamic_redirect", "http_request_firewall_custom", "http_request_firewall_managed", "http_request_late_transform", "http_request_origin", "http_request_redirect", "http_request_sanitize", "http_request_sbfm", "http_request_transform", "http_response_cache_settings", "http_response_compression", "http_response_firewall_managed", "http_response_headers_transform", "magic_transit", "magic_transit_ids_managed", "magic_transit_managed", "magic_transit_ratelimit".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#phases Ruleset#phases}
	Phases *[]*string `field:"optional" json:"phases" yaml:"phases"`
	// The Polish level to configure. Available values: "off", "lossless", "lossy", "webp".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#polish Ruleset#polish}
	Polish *string `field:"optional" json:"polish" yaml:"polish"`
	// Set the private cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#private Ruleset#private}
	Private *RulesetRulesActionParametersPrivate `field:"optional" json:"private" yaml:"private"`
	// A list of legacy security products to skip the execution of. Available values: "bic", "hot", "rateLimit", "securityLevel", "uaBlock", "waf", "zoneLockdown".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#products Ruleset#products}
	Products *[]*string `field:"optional" json:"products" yaml:"products"`
	// Set the proxy-revalidate cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#proxy_revalidate Ruleset#proxy_revalidate}
	ProxyRevalidate *RulesetRulesActionParametersProxyRevalidate `field:"optional" json:"proxyRevalidate" yaml:"proxyRevalidate"`
	// Set the public cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#public Ruleset#public}
	Public *RulesetRulesActionParametersPublic `field:"optional" json:"public" yaml:"public"`
	// The raw response fields to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#raw_response_fields Ruleset#raw_response_fields}
	RawResponseFields interface{} `field:"optional" json:"rawResponseFields" yaml:"rawResponseFields"`
	// A timeout value between two successive read operations to use for your origin server.
	//
	// Historically, the timeout value between two read options from Cloudflare to an origin server is 100 seconds. If you are attempting to reduce HTTP 524 errors because of timeouts from an origin server, try increasing this timeout value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#read_timeout Ruleset#read_timeout}
	ReadTimeout *float64 `field:"optional" json:"readTimeout" yaml:"readTimeout"`
	// Whether to redirect verified AI training crawlers to canonical URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#redirects_for_ai_training Ruleset#redirects_for_ai_training}
	RedirectsForAiTraining interface{} `field:"optional" json:"redirectsForAiTraining" yaml:"redirectsForAiTraining"`
	// The request body buffering mode to configure. Available values: "none", "standard", "full".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#request_body_buffering Ruleset#request_body_buffering}
	RequestBodyBuffering *string `field:"optional" json:"requestBodyBuffering" yaml:"requestBodyBuffering"`
	// The raw request fields to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#request_fields Ruleset#request_fields}
	RequestFields interface{} `field:"optional" json:"requestFields" yaml:"requestFields"`
	// Whether Cloudflare should respect strong ETag (entity tag) headers.
	//
	// If false, Cloudflare converts strong ETag headers to weak ETag headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#respect_strong_etags Ruleset#respect_strong_etags}
	RespectStrongEtags interface{} `field:"optional" json:"respectStrongEtags" yaml:"respectStrongEtags"`
	// The response to show when the block is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#response Ruleset#response}
	Response *RulesetRulesActionParametersResponse `field:"optional" json:"response" yaml:"response"`
	// The response body buffering mode to configure. Available values: "none", "standard".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#response_body_buffering Ruleset#response_body_buffering}
	ResponseBodyBuffering *string `field:"optional" json:"responseBodyBuffering" yaml:"responseBodyBuffering"`
	// The transformed response fields to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#response_fields Ruleset#response_fields}
	ResponseFields interface{} `field:"optional" json:"responseFields" yaml:"responseFields"`
	// Whether to enable Rocket Loader.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#rocket_loader Ruleset#rocket_loader}
	RocketLoader interface{} `field:"optional" json:"rocketLoader" yaml:"rocketLoader"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the execution of.
	//
	// This option is incompatible with the ruleset option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#rules Ruleset#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the rulesets option. Available values: "current".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#ruleset Ruleset#ruleset}
	Ruleset *string `field:"optional" json:"ruleset" yaml:"ruleset"`
	// A list of ruleset IDs to skip the execution of.
	//
	// This option is incompatible with the ruleset and phases options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#rulesets Ruleset#rulesets}
	Rulesets *[]*string `field:"optional" json:"rulesets" yaml:"rulesets"`
	// The Security Level to configure. Available values: "off", "essentially_off", "low", "medium", "high", "under_attack".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#security_level Ruleset#security_level}
	SecurityLevel *string `field:"optional" json:"securityLevel" yaml:"securityLevel"`
	// Whether to enable Server-Side Excludes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#server_side_excludes Ruleset#server_side_excludes}
	ServerSideExcludes interface{} `field:"optional" json:"serverSideExcludes" yaml:"serverSideExcludes"`
	// When to serve stale content from cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#serve_stale Ruleset#serve_stale}
	ServeStale *RulesetRulesActionParametersServeStale `field:"optional" json:"serveStale" yaml:"serveStale"`
	// Set the s-maxage cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#s_maxage Ruleset#s_maxage}
	SMaxage *RulesetRulesActionParametersSMaxage `field:"optional" json:"sMaxage" yaml:"sMaxage"`
	// A Server Name Indication (SNI) override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#sni Ruleset#sni}
	Sni *RulesetRulesActionParametersSni `field:"optional" json:"sni" yaml:"sni"`
	// The SSL level to configure. Available values: "off", "flexible", "full", "strict", "origin_pull".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#ssl Ruleset#ssl}
	Ssl *string `field:"optional" json:"ssl" yaml:"ssl"`
	// Set the stale-if-error cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#stale_if_error Ruleset#stale_if_error}
	StaleIfError *RulesetRulesActionParametersStaleIfError `field:"optional" json:"staleIfError" yaml:"staleIfError"`
	// Set the stale-while-revalidate cache control directive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#stale_while_revalidate Ruleset#stale_while_revalidate}
	StaleWhileRevalidate *RulesetRulesActionParametersStaleWhileRevalidate `field:"optional" json:"staleWhileRevalidate" yaml:"staleWhileRevalidate"`
	// The status code to use for the error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Whether to strip the ETag header from the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#strip_etags Ruleset#strip_etags}
	StripEtags interface{} `field:"optional" json:"stripEtags" yaml:"stripEtags"`
	// Whether to strip the Last-Modified header from the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#strip_last_modified Ruleset#strip_last_modified}
	StripLastModified interface{} `field:"optional" json:"stripLastModified" yaml:"stripLastModified"`
	// Whether to strip the Set-Cookie header from the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#strip_set_cookie Ruleset#strip_set_cookie}
	StripSetCookie interface{} `field:"optional" json:"stripSetCookie" yaml:"stripSetCookie"`
	// Whether to enable Signed Exchanges (SXG).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#sxg Ruleset#sxg}
	Sxg interface{} `field:"optional" json:"sxg" yaml:"sxg"`
	// The transformed request fields to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#transformed_request_fields Ruleset#transformed_request_fields}
	TransformedRequestFields interface{} `field:"optional" json:"transformedRequestFields" yaml:"transformedRequestFields"`
	// A URI rewrite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#uri Ruleset#uri}
	Uri *RulesetRulesActionParametersUri `field:"optional" json:"uri" yaml:"uri"`
	// The cache tag values for set_cache_tags action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#values Ruleset#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

