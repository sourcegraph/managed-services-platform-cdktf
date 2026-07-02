package emailsecuritytrusteddomains

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/emailsecuritytrusteddomains/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/email_security_trusted_domains cloudflare_email_security_trusted_domains}.
type EmailSecurityTrustedDomains interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Body() EmailSecurityTrustedDomainsBodyList
	BodyInput() any
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comments() *string
	SetComments(val *string)
	CommentsInput() *string
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *float64
	IsRecent() any
	SetIsRecent(val any)
	IsRecentInput() any
	IsRegex() any
	SetIsRegex(val any)
	IsRegexInput() any
	IsSimilarity() any
	SetIsSimilarity(val any)
	IsSimilarityInput() any
	LastModified() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Pattern() *string
	SetPattern(val *string)
	PatternInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() any
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBody(value any)
	ResetBody()
	ResetComments()
	ResetIsRecent()
	ResetIsRegex()
	ResetIsSimilarity()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPattern()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for EmailSecurityTrustedDomains
type jsiiProxy_EmailSecurityTrustedDomains struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Body() EmailSecurityTrustedDomainsBodyList {
	var returns EmailSecurityTrustedDomainsBodyList
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) BodyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Comments() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) CommentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) IsRecent() any {
	var returns any
	_jsii_.Get(
		j,
		"isRecent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) IsRecentInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isRecentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) IsRegex() any {
	var returns any
	_jsii_.Get(
		j,
		"isRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) IsRegexInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) IsSimilarity() any {
	var returns any
	_jsii_.Get(
		j,
		"isSimilarity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) IsSimilarityInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isSimilarityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/email_security_trusted_domains cloudflare_email_security_trusted_domains} Resource.
func NewEmailSecurityTrustedDomains(scope constructs.Construct, id *string, config *EmailSecurityTrustedDomainsConfig) EmailSecurityTrustedDomains {
	_init_.Initialize()

	if err := validateNewEmailSecurityTrustedDomainsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmailSecurityTrustedDomains{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.emailSecurityTrustedDomains.EmailSecurityTrustedDomains",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/email_security_trusted_domains cloudflare_email_security_trusted_domains} Resource.
func NewEmailSecurityTrustedDomains_Override(e EmailSecurityTrustedDomains, scope constructs.Construct, id *string, config *EmailSecurityTrustedDomainsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.emailSecurityTrustedDomains.EmailSecurityTrustedDomains",
		[]any{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetComments(val *string) {
	if err := j.validateSetCommentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comments",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetIsRecent(val any) {
	if err := j.validateSetIsRecentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRecent",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetIsRegex(val any) {
	if err := j.validateSetIsRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRegex",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetIsSimilarity(val any) {
	if err := j.validateSetIsSimilarityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSimilarity",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EmailSecurityTrustedDomains) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a EmailSecurityTrustedDomains resource upon running "cdktf plan <stack-name>".
func EmailSecurityTrustedDomains_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEmailSecurityTrustedDomains_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.emailSecurityTrustedDomains.EmailSecurityTrustedDomains",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EmailSecurityTrustedDomains_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateEmailSecurityTrustedDomains_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.emailSecurityTrustedDomains.EmailSecurityTrustedDomains",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmailSecurityTrustedDomains_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateEmailSecurityTrustedDomains_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.emailSecurityTrustedDomains.EmailSecurityTrustedDomains",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmailSecurityTrustedDomains_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateEmailSecurityTrustedDomains_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.emailSecurityTrustedDomains.EmailSecurityTrustedDomains",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func EmailSecurityTrustedDomains_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.emailSecurityTrustedDomains.EmailSecurityTrustedDomains",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) AddOverride(path *string, value any) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]any{path, value},
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]any{id, provider},
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]any{id},
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) MoveTo(moveTarget *string, index any) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]any{id},
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) PutBody(value any) {
	if err := e.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBody",
		[]any{value},
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ResetBody() {
	_jsii_.InvokeVoid(
		e,
		"resetBody",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ResetComments() {
	_jsii_.InvokeVoid(
		e,
		"resetComments",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ResetIsRecent() {
	_jsii_.InvokeVoid(
		e,
		"resetIsRecent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ResetIsRegex() {
	_jsii_.InvokeVoid(
		e,
		"resetIsRegex",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ResetIsSimilarity() {
	_jsii_.InvokeVoid(
		e,
		"resetIsSimilarity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ResetPattern() {
	_jsii_.InvokeVoid(
		e,
		"resetPattern",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailSecurityTrustedDomains) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
