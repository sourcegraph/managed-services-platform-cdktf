//go:build !no_runtime_type_checking

package vertexaiendpointwithmodelgardendeployment

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validatePutReservationAffinityParameters(value *VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinity) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateSetAcceleratorCountParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateSetAcceleratorTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateSetInternalValueParameters(val *VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpec) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateSetMachineTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateSetMultihostGpuNodeCountParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReference) validateSetTpuTopologyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

