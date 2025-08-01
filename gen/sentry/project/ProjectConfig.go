package project

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectConfig struct {
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
	// The name for the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#name Project#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The organization of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#organization Project#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The slugs of the teams to create the project for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#teams Project#teams}
	Teams *[]*string `field:"required" json:"teams" yaml:"teams"`
	// Configure origin URLs which Sentry should accept events from. This is used for communication with clients like [sentry-javascript](https://github.com/getsentry/sentry-javascript).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#client_security Project#client_security}
	ClientSecurity *ProjectClientSecurity `field:"optional" json:"clientSecurity" yaml:"clientSecurity"`
	// Whether to create a default key on project creation.
	//
	// By default, Sentry will create a key for you. If you wish to manage keys manually, set this to false and create keys using the `sentry_key` resource. Note that this only takes effect on project creation, not on project update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#default_key Project#default_key}
	DefaultKey interface{} `field:"optional" json:"defaultKey" yaml:"defaultKey"`
	// Whether to create a default issue alert.
	//
	// Defaults to true where the behavior is to alert the user on every new issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#default_rules Project#default_rules}
	DefaultRules interface{} `field:"optional" json:"defaultRules" yaml:"defaultRules"`
	// The maximum amount of time (in seconds) to wait between scheduling digests for delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#digests_max_delay Project#digests_max_delay}
	DigestsMaxDelay *float64 `field:"optional" json:"digestsMaxDelay" yaml:"digestsMaxDelay"`
	// The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#digests_min_delay Project#digests_min_delay}
	DigestsMinDelay *float64 `field:"optional" json:"digestsMinDelay" yaml:"digestsMinDelay"`
	// Custom filters for this project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#filters Project#filters}
	Filters *ProjectFilters `field:"optional" json:"filters" yaml:"filters"`
	// This can be used to modify the fingerprint rules on the server with custom rules.
	//
	// Rules follow the pattern `matcher:glob -> fingerprint, values`. To learn more about fingerprint rules, [read the docs](https://docs.sentry.io/concepts/data-management/event-grouping/fingerprint-rules/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#fingerprinting_rules Project#fingerprinting_rules}
	FingerprintingRules *string `field:"optional" json:"fingerprintingRules" yaml:"fingerprintingRules"`
	// This can be used to enhance the grouping algorithm with custom rules.
	//
	// Rules follow the pattern `matcher:glob [v^]?[+-]flag`. To learn more about stack trace rules, [read the docs](https://docs.sentry.io/concepts/data-management/event-grouping/stack-trace-rules/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#grouping_enhancements Project#grouping_enhancements}
	GroupingEnhancements *string `field:"optional" json:"groupingEnhancements" yaml:"groupingEnhancements"`
	// The platform for this project.
	//
	// Use `other` for platforms not listed. Valid values are: `other`, `android`, `apple`, `apple-ios`, `apple-macos`, `bun`, `capacitor`, `cordova`, `dart`, `deno`, `dotnet`, `dotnet-aspnet`, `dotnet-aspnetcore`, `dotnet-awslambda`, `dotnet-gcpfunctions`, `dotnet-maui`, `dotnet-uwp`, `dotnet-winforms`, `dotnet-wpf`, `dotnet-xamarin`, `electron`, `elixir`, `flutter`, `go`, `go-echo`, `go-fasthttp`, `go-fiber`, `go-gin`, `go-http`, `go-iris`, `go-martini`, `go-negroni`, `godot`, `ionic`, `java`, `java-log4j2`, `java-logback`, `java-spring`, `java-spring-boot`, `javascript`, `javascript-angular`, `javascript-astro`, `javascript-ember`, `javascript-gatsby`, `javascript-nextjs`, `javascript-react`, `javascript-remix`, `javascript-solid`, `javascript-solidstart`, `javascript-svelte`, `javascript-sveltekit`, `javascript-tanstackstart-react`, `javascript-nuxt`, `javascript-vue`, `kotlin`, `minidump`, `native`, `native-qt`, `nintendo-switch`, `node`, `node-awslambda`, `node-azurefunctions`, `node-cloudflare-pages`, `node-cloudflare-workers`, `node-connect`, `node-express`, `node-fastify`, `node-gcpfunctions`, `node-hapi`, `node-koa`, `node-nestjs`, `php`, `php-laravel`, `php-symfony`, `powershell`, `python`, `python-aiohttp`, `python-asgi`, `python-awslambda`, `python-bottle`, `python-celery`, `python-chalice`, `python-django`, `python-falcon`, `python-fastapi`, `python-flask`, `python-gcpfunctions`, `python-pylons`, `python-pymongo`, `python-pyramid`, `python-quart`, `python-rq`, `python-sanic`, `python-serverless`, `python-starlette`, `python-tornado`, `python-tryton`, `python-wsgi`, `react-native`, `ruby`, `ruby-rack`, `ruby-rails`, `rust`, `unity`, and `unreal`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#platform Project#platform}
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Hours in which an issue is automatically resolve if not seen after this amount of time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#resolve_age Project#resolve_age}
	ResolveAge *float64 `field:"optional" json:"resolveAge" yaml:"resolveAge"`
	// The optional slug for this project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#slug Project#slug}
	Slug *string `field:"optional" json:"slug" yaml:"slug"`
}

