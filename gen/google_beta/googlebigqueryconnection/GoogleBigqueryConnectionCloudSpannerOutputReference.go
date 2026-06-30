package googlebigqueryconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlebigqueryconnection/internal"
)

type GoogleBigqueryConnectionCloudSpannerOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	DatabaseRole() *string
	SetDatabaseRole(val *string)
	DatabaseRoleInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBigqueryConnectionCloudSpanner
	SetInternalValue(val *GoogleBigqueryConnectionCloudSpanner)
	MaxParallelism() *float64
	SetMaxParallelism(val *float64)
	MaxParallelismInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseDataBoost() any
	SetUseDataBoost(val any)
	UseDataBoostInput() any
	UseParallelism() any
	SetUseParallelism(val any)
	UseParallelismInput() any
	UseServerlessAnalytics() any
	SetUseServerlessAnalytics(val any)
	UseServerlessAnalyticsInput() any
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDatabaseRole()
	ResetMaxParallelism()
	ResetUseDataBoost()
	ResetUseParallelism()
	ResetUseServerlessAnalytics()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryConnectionCloudSpannerOutputReference
type jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) DatabaseRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) DatabaseRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) InternalValue() *GoogleBigqueryConnectionCloudSpanner {
	var returns *GoogleBigqueryConnectionCloudSpanner
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) MaxParallelism() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) MaxParallelismInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) UseDataBoost() any {
	var returns any
	_jsii_.Get(
		j,
		"useDataBoost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) UseDataBoostInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useDataBoostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) UseParallelism() any {
	var returns any
	_jsii_.Get(
		j,
		"useParallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) UseParallelismInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useParallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) UseServerlessAnalytics() any {
	var returns any
	_jsii_.Get(
		j,
		"useServerlessAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) UseServerlessAnalyticsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useServerlessAnalyticsInput",
		&returns,
	)
	return returns
}

func NewGoogleBigqueryConnectionCloudSpannerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryConnectionCloudSpannerOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryConnectionCloudSpannerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleBigqueryConnection.GoogleBigqueryConnectionCloudSpannerOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryConnectionCloudSpannerOutputReference_Override(g GoogleBigqueryConnectionCloudSpannerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleBigqueryConnection.GoogleBigqueryConnectionCloudSpannerOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetDatabaseRole(val *string) {
	if err := j.validateSetDatabaseRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseRole",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetInternalValue(val *GoogleBigqueryConnectionCloudSpanner) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetMaxParallelism(val *float64) {
	if err := j.validateSetMaxParallelismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelism",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetUseDataBoost(val any) {
	if err := j.validateSetUseDataBoostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDataBoost",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetUseParallelism(val any) {
	if err := j.validateSetUseParallelismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useParallelism",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) SetUseServerlessAnalytics(val any) {
	if err := j.validateSetUseServerlessAnalyticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useServerlessAnalytics",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) ResetDatabaseRole() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseRole",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) ResetMaxParallelism() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxParallelism",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) ResetUseDataBoost() {
	_jsii_.InvokeVoid(
		g,
		"resetUseDataBoost",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) ResetUseParallelism() {
	_jsii_.InvokeVoid(
		g,
		"resetUseParallelism",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) ResetUseServerlessAnalytics() {
	_jsii_.InvokeVoid(
		g,
		"resetUseServerlessAnalytics",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionCloudSpannerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
