//go:build !no_runtime_type_checking

package pagesproject

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewVectorizeBindingsMap) validateGetParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewVectorizeBindingsMap) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewVectorizeBindingsMap) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewVectorizeBindingsMap) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *map[string]*PagesProjectDeploymentConfigsPreviewVectorizeBindings:
		val := val.(*map[string]*PagesProjectDeploymentConfigsPreviewVectorizeBindings)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case map[string]*PagesProjectDeploymentConfigsPreviewVectorizeBindings:
		val_ := val.(map[string]*PagesProjectDeploymentConfigsPreviewVectorizeBindings)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *map[string]*PagesProjectDeploymentConfigsPreviewVectorizeBindings; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewVectorizeBindingsMap) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewVectorizeBindingsMap) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPagesProjectDeploymentConfigsPreviewVectorizeBindingsMapParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

