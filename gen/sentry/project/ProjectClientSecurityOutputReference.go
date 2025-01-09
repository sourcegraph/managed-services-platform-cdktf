package project

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/project/internal"
)

type ProjectClientSecurityOutputReference interface {
	cdktf.ComplexObject
	AllowedDomains() *[]*string
	SetAllowedDomains(val *[]*string)
	AllowedDomainsInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ScrapeJavascript() interface{}
	SetScrapeJavascript(val interface{})
	ScrapeJavascriptInput() interface{}
	SecurityToken() *string
	SetSecurityToken(val *string)
	SecurityTokenHeader() *string
	SetSecurityTokenHeader(val *string)
	SecurityTokenHeaderInput() *string
	SecurityTokenInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VerifyTlsSsl() interface{}
	SetVerifyTlsSsl(val interface{})
	VerifyTlsSslInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAllowedDomains()
	ResetScrapeJavascript()
	ResetSecurityToken()
	ResetSecurityTokenHeader()
	ResetVerifyTlsSsl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ProjectClientSecurityOutputReference
type jsiiProxy_ProjectClientSecurityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) AllowedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) AllowedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) ScrapeJavascript() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scrapeJavascript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) ScrapeJavascriptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scrapeJavascriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) SecurityToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) SecurityTokenHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityTokenHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) SecurityTokenHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityTokenHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) SecurityTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) VerifyTlsSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyTlsSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference) VerifyTlsSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyTlsSslInput",
		&returns,
	)
	return returns
}


func NewProjectClientSecurityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ProjectClientSecurityOutputReference {
	_init_.Initialize()

	if err := validateNewProjectClientSecurityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectClientSecurityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.project.ProjectClientSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewProjectClientSecurityOutputReference_Override(p ProjectClientSecurityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.project.ProjectClientSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference)SetAllowedDomains(val *[]*string) {
	if err := j.validateSetAllowedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDomains",
		val,
	)
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference)SetScrapeJavascript(val interface{}) {
	if err := j.validateSetScrapeJavascriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scrapeJavascript",
		val,
	)
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference)SetSecurityToken(val *string) {
	if err := j.validateSetSecurityTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityToken",
		val,
	)
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference)SetSecurityTokenHeader(val *string) {
	if err := j.validateSetSecurityTokenHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityTokenHeader",
		val,
	)
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ProjectClientSecurityOutputReference)SetVerifyTlsSsl(val interface{}) {
	if err := j.validateSetVerifyTlsSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyTlsSsl",
		val,
	)
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) ResetAllowedDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) ResetScrapeJavascript() {
	_jsii_.InvokeVoid(
		p,
		"resetScrapeJavascript",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) ResetSecurityToken() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) ResetSecurityTokenHeader() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityTokenHeader",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) ResetVerifyTlsSsl() {
	_jsii_.InvokeVoid(
		p,
		"resetVerifyTlsSsl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectClientSecurityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

