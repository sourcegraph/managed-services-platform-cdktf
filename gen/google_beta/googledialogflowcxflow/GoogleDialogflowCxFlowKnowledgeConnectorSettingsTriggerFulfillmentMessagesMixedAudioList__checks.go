//go:build !no_runtime_type_checking

package googledialogflowcxflow

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

