package googlecloudrunv2service

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecloudrunv2service/internal"
)

type GoogleCloudRunV2ServiceTemplateContainersOutputReference interface {
	cdktf.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	BaseImageUri() *string
	SetBaseImageUri(val *string)
	BaseImageUriInput() *string
	BuildInfo() GoogleCloudRunV2ServiceTemplateContainersBuildInfoList
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DependsOnInput() *[]*string
	Env() GoogleCloudRunV2ServiceTemplateContainersEnvList
	EnvInput() any
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() any
	SetInternalValue(val any)
	LivenessProbe() GoogleCloudRunV2ServiceTemplateContainersLivenessProbeOutputReference
	LivenessProbeInput() *GoogleCloudRunV2ServiceTemplateContainersLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Ports() GoogleCloudRunV2ServiceTemplateContainersPortsOutputReference
	PortsInput() *GoogleCloudRunV2ServiceTemplateContainersPorts
	Resources() GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference
	ResourcesInput() *GoogleCloudRunV2ServiceTemplateContainersResources
	StartupProbe() GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference
	StartupProbeInput() *GoogleCloudRunV2ServiceTemplateContainersStartupProbe
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeMounts() GoogleCloudRunV2ServiceTemplateContainersVolumeMountsList
	VolumeMountsInput() any
	WorkingDir() *string
	SetWorkingDir(val *string)
	WorkingDirInput() *string
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
	PutEnv(value any)
	PutLivenessProbe(value *GoogleCloudRunV2ServiceTemplateContainersLivenessProbe)
	PutPorts(value *GoogleCloudRunV2ServiceTemplateContainersPorts)
	PutResources(value *GoogleCloudRunV2ServiceTemplateContainersResources)
	PutStartupProbe(value *GoogleCloudRunV2ServiceTemplateContainersStartupProbe)
	PutVolumeMounts(value any)
	ResetArgs()
	ResetBaseImageUri()
	ResetCommand()
	ResetDependsOn()
	ResetEnv()
	ResetLivenessProbe()
	ResetName()
	ResetPorts()
	ResetResources()
	ResetStartupProbe()
	ResetVolumeMounts()
	ResetWorkingDir()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudRunV2ServiceTemplateContainersOutputReference
type jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) BaseImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseImageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) BaseImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseImageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) BuildInfo() GoogleCloudRunV2ServiceTemplateContainersBuildInfoList {
	var returns GoogleCloudRunV2ServiceTemplateContainersBuildInfoList
	_jsii_.Get(
		j,
		"buildInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) DependsOnInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) Env() GoogleCloudRunV2ServiceTemplateContainersEnvList {
	var returns GoogleCloudRunV2ServiceTemplateContainersEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) EnvInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) LivenessProbe() GoogleCloudRunV2ServiceTemplateContainersLivenessProbeOutputReference {
	var returns GoogleCloudRunV2ServiceTemplateContainersLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) LivenessProbeInput() *GoogleCloudRunV2ServiceTemplateContainersLivenessProbe {
	var returns *GoogleCloudRunV2ServiceTemplateContainersLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) Ports() GoogleCloudRunV2ServiceTemplateContainersPortsOutputReference {
	var returns GoogleCloudRunV2ServiceTemplateContainersPortsOutputReference
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) PortsInput() *GoogleCloudRunV2ServiceTemplateContainersPorts {
	var returns *GoogleCloudRunV2ServiceTemplateContainersPorts
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) Resources() GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference {
	var returns GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResourcesInput() *GoogleCloudRunV2ServiceTemplateContainersResources {
	var returns *GoogleCloudRunV2ServiceTemplateContainersResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) StartupProbe() GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference {
	var returns GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) StartupProbeInput() *GoogleCloudRunV2ServiceTemplateContainersStartupProbe {
	var returns *GoogleCloudRunV2ServiceTemplateContainersStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) VolumeMounts() GoogleCloudRunV2ServiceTemplateContainersVolumeMountsList {
	var returns GoogleCloudRunV2ServiceTemplateContainersVolumeMountsList
	_jsii_.Get(
		j,
		"volumeMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) VolumeMountsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeMountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}

func NewGoogleCloudRunV2ServiceTemplateContainersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCloudRunV2ServiceTemplateContainersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2ServiceTemplateContainersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleCloudRunV2Service.GoogleCloudRunV2ServiceTemplateContainersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCloudRunV2ServiceTemplateContainersOutputReference_Override(g GoogleCloudRunV2ServiceTemplateContainersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleCloudRunV2Service.GoogleCloudRunV2ServiceTemplateContainersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetBaseImageUri(val *string) {
	if err := j.validateSetBaseImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseImageUri",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetDependsOn(val *[]*string) {
	if err := j.validateSetDependsOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) PutEnv(value any) {
	if err := g.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnv",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) PutLivenessProbe(value *GoogleCloudRunV2ServiceTemplateContainersLivenessProbe) {
	if err := g.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLivenessProbe",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) PutPorts(value *GoogleCloudRunV2ServiceTemplateContainersPorts) {
	if err := g.validatePutPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPorts",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) PutResources(value *GoogleCloudRunV2ServiceTemplateContainersResources) {
	if err := g.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResources",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) PutStartupProbe(value *GoogleCloudRunV2ServiceTemplateContainersStartupProbe) {
	if err := g.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStartupProbe",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) PutVolumeMounts(value any) {
	if err := g.validatePutVolumeMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVolumeMounts",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetBaseImageUri() {
	_jsii_.InvokeVoid(
		g,
		"resetBaseImageUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		g,
		"resetCommand",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		g,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		g,
		"resetEnv",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		g,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		g,
		"resetPorts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		g,
		"resetResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		g,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetVolumeMounts() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumeMounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
