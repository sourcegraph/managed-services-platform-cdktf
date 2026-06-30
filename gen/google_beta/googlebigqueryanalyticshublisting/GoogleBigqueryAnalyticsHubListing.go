package googlebigqueryanalyticshublisting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlebigqueryanalyticshublisting/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_listing google_bigquery_analytics_hub_listing}.
type GoogleBigqueryAnalyticsHubListing interface {
	cdktf.TerraformResource
	BigqueryDataset() GoogleBigqueryAnalyticsHubListingBigqueryDatasetOutputReference
	BigqueryDatasetInput() *GoogleBigqueryAnalyticsHubListingBigqueryDataset
	Categories() *[]*string
	SetCategories(val *[]*string)
	CategoriesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DataExchangeId() *string
	SetDataExchangeId(val *string)
	DataExchangeIdInput() *string
	DataProvider() GoogleBigqueryAnalyticsHubListingDataProviderOutputReference
	DataProviderInput() *GoogleBigqueryAnalyticsHubListingDataProvider
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Documentation() *string
	SetDocumentation(val *string)
	DocumentationInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Icon() *string
	SetIcon(val *string)
	IconInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListingId() *string
	SetListingId(val *string)
	ListingIdInput() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogLinkedDatasetQueryUserEmail() any
	SetLogLinkedDatasetQueryUserEmail(val any)
	LogLinkedDatasetQueryUserEmailInput() any
	Name() *string
	// The tree node.
	Node() constructs.Node
	PrimaryContact() *string
	SetPrimaryContact(val *string)
	PrimaryContactInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	Publisher() GoogleBigqueryAnalyticsHubListingPublisherOutputReference
	PublisherInput() *GoogleBigqueryAnalyticsHubListingPublisher
	PubsubTopic() GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference
	PubsubTopicInput() *GoogleBigqueryAnalyticsHubListingPubsubTopic
	// Experimental.
	RawOverrides() any
	RequestAccess() *string
	SetRequestAccess(val *string)
	RequestAccessInput() *string
	RestrictedExportConfig() GoogleBigqueryAnalyticsHubListingRestrictedExportConfigOutputReference
	RestrictedExportConfigInput() *GoogleBigqueryAnalyticsHubListingRestrictedExportConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleBigqueryAnalyticsHubListingTimeoutsOutputReference
	TimeoutsInput() any
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
	PutBigqueryDataset(value *GoogleBigqueryAnalyticsHubListingBigqueryDataset)
	PutDataProvider(value *GoogleBigqueryAnalyticsHubListingDataProvider)
	PutPublisher(value *GoogleBigqueryAnalyticsHubListingPublisher)
	PutPubsubTopic(value *GoogleBigqueryAnalyticsHubListingPubsubTopic)
	PutRestrictedExportConfig(value *GoogleBigqueryAnalyticsHubListingRestrictedExportConfig)
	PutTimeouts(value *GoogleBigqueryAnalyticsHubListingTimeouts)
	ResetBigqueryDataset()
	ResetCategories()
	ResetDataProvider()
	ResetDescription()
	ResetDocumentation()
	ResetIcon()
	ResetId()
	ResetLogLinkedDatasetQueryUserEmail()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaryContact()
	ResetProject()
	ResetPublisher()
	ResetPubsubTopic()
	ResetRequestAccess()
	ResetRestrictedExportConfig()
	ResetTimeouts()
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

// The jsii proxy struct for GoogleBigqueryAnalyticsHubListing
type jsiiProxy_GoogleBigqueryAnalyticsHubListing struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) BigqueryDataset() GoogleBigqueryAnalyticsHubListingBigqueryDatasetOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingBigqueryDatasetOutputReference
	_jsii_.Get(
		j,
		"bigqueryDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) BigqueryDatasetInput() *GoogleBigqueryAnalyticsHubListingBigqueryDataset {
	var returns *GoogleBigqueryAnalyticsHubListingBigqueryDataset
	_jsii_.Get(
		j,
		"bigqueryDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Categories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) CategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DataExchangeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataExchangeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DataExchangeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataExchangeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DataProvider() GoogleBigqueryAnalyticsHubListingDataProviderOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingDataProviderOutputReference
	_jsii_.Get(
		j,
		"dataProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DataProviderInput() *GoogleBigqueryAnalyticsHubListingDataProvider {
	var returns *GoogleBigqueryAnalyticsHubListingDataProvider
	_jsii_.Get(
		j,
		"dataProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Documentation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DocumentationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Icon() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) IconInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iconInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ListingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ListingIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) LogLinkedDatasetQueryUserEmail() any {
	var returns any
	_jsii_.Get(
		j,
		"logLinkedDatasetQueryUserEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) LogLinkedDatasetQueryUserEmailInput() any {
	var returns any
	_jsii_.Get(
		j,
		"logLinkedDatasetQueryUserEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PrimaryContact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PrimaryContactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Publisher() GoogleBigqueryAnalyticsHubListingPublisherOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingPublisherOutputReference
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PublisherInput() *GoogleBigqueryAnalyticsHubListingPublisher {
	var returns *GoogleBigqueryAnalyticsHubListingPublisher
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PubsubTopic() GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference
	_jsii_.Get(
		j,
		"pubsubTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PubsubTopicInput() *GoogleBigqueryAnalyticsHubListingPubsubTopic {
	var returns *GoogleBigqueryAnalyticsHubListingPubsubTopic
	_jsii_.Get(
		j,
		"pubsubTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RequestAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RequestAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RestrictedExportConfig() GoogleBigqueryAnalyticsHubListingRestrictedExportConfigOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingRestrictedExportConfigOutputReference
	_jsii_.Get(
		j,
		"restrictedExportConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RestrictedExportConfigInput() *GoogleBigqueryAnalyticsHubListingRestrictedExportConfig {
	var returns *GoogleBigqueryAnalyticsHubListingRestrictedExportConfig
	_jsii_.Get(
		j,
		"restrictedExportConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Timeouts() GoogleBigqueryAnalyticsHubListingTimeoutsOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_listing google_bigquery_analytics_hub_listing} Resource.
func NewGoogleBigqueryAnalyticsHubListing(scope constructs.Construct, id *string, config *GoogleBigqueryAnalyticsHubListingConfig) GoogleBigqueryAnalyticsHubListing {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryAnalyticsHubListingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryAnalyticsHubListing{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_listing google_bigquery_analytics_hub_listing} Resource.
func NewGoogleBigqueryAnalyticsHubListing_Override(g GoogleBigqueryAnalyticsHubListing, scope constructs.Construct, id *string, config *GoogleBigqueryAnalyticsHubListingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		[]any{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetCategories(val *[]*string) {
	if err := j.validateSetCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"categories",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetDataExchangeId(val *string) {
	if err := j.validateSetDataExchangeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataExchangeId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetDocumentation(val *string) {
	if err := j.validateSetDocumentationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentation",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetIcon(val *string) {
	if err := j.validateSetIconParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icon",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetListingId(val *string) {
	if err := j.validateSetListingIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listingId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetLogLinkedDatasetQueryUserEmail(val any) {
	if err := j.validateSetLogLinkedDatasetQueryUserEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLinkedDatasetQueryUserEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetPrimaryContact(val *string) {
	if err := j.validateSetPrimaryContactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryContact",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SetRequestAccess(val *string) {
	if err := j.validateSetRequestAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestAccess",
		val,
	)
}

// Generates CDKTF code for importing a GoogleBigqueryAnalyticsHubListing resource upon running "cdktf plan <stack-name>".
func GoogleBigqueryAnalyticsHubListing_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleBigqueryAnalyticsHubListing_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
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
func GoogleBigqueryAnalyticsHubListing_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryAnalyticsHubListing_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryAnalyticsHubListing_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryAnalyticsHubListing_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryAnalyticsHubListing_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryAnalyticsHubListing_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func GoogleBigqueryAnalyticsHubListing_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) AddOverride(path *string, value any) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]any{path, value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]any{id, provider},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]any{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) MoveTo(moveTarget *string, index any) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]any{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutBigqueryDataset(value *GoogleBigqueryAnalyticsHubListingBigqueryDataset) {
	if err := g.validatePutBigqueryDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBigqueryDataset",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutDataProvider(value *GoogleBigqueryAnalyticsHubListingDataProvider) {
	if err := g.validatePutDataProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataProvider",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutPublisher(value *GoogleBigqueryAnalyticsHubListingPublisher) {
	if err := g.validatePutPublisherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublisher",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutPubsubTopic(value *GoogleBigqueryAnalyticsHubListingPubsubTopic) {
	if err := g.validatePutPubsubTopicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPubsubTopic",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutRestrictedExportConfig(value *GoogleBigqueryAnalyticsHubListingRestrictedExportConfig) {
	if err := g.validatePutRestrictedExportConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRestrictedExportConfig",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutTimeouts(value *GoogleBigqueryAnalyticsHubListingTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetBigqueryDataset() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryDataset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetCategories() {
	_jsii_.InvokeVoid(
		g,
		"resetCategories",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetDataProvider() {
	_jsii_.InvokeVoid(
		g,
		"resetDataProvider",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetDocumentation() {
	_jsii_.InvokeVoid(
		g,
		"resetDocumentation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetIcon() {
	_jsii_.InvokeVoid(
		g,
		"resetIcon",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetLogLinkedDatasetQueryUserEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetLogLinkedDatasetQueryUserEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetPrimaryContact() {
	_jsii_.InvokeVoid(
		g,
		"resetPrimaryContact",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetPublisher() {
	_jsii_.InvokeVoid(
		g,
		"resetPublisher",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetPubsubTopic() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubTopic",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetRequestAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetRestrictedExportConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRestrictedExportConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
