package ami

import (
	"fmt"
	"strings"
)

// ResolveAMI will resolve an AMI from the supplied region
// and instance type. It will invoke a specific resolver
// to do the actual detrminng of AMI.
func ResolveAMI(region string, instanceType string) (string, error) {
	var resolver Resolver

	if strings.HasPrefix(instanceType, "p2") ||
		strings.HasPrefix(instanceType, "p3") {
		resolver = &GpuResolver{}
	} else {
		resolver = &DefaultResolver{}
	}

	return resolver.Resolve(region, instanceType)
}

// Resolver provides an interface to enable implementing multiple
// ways to determine which AMI to use from the region/instance type.
type Resolver interface {
	Resolve(region string, instanceType string) (string, error)
}

// DefaultResolver resolves the AMi to the defaults for the region
type DefaultResolver struct {
}

// Resolve will return an AMI to use based on the default AMI for each region
// TODO: https://github.com/weaveworks/eksctl/issues/49
// currently source of truth for these is here:
// https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html
func (r *DefaultResolver) Resolve(region string, instanceType string) (string, error) {
	switch region {
	case "us-west-2":
		return "ami-08cab282f9979fc7a", nil
	case "us-east-1":
		return "ami-0b2ae3c6bda8b5c06", nil
	case "eu-west-1":
		return "ami-066110c1a7466949e", nil
	default:
		return "", fmt.Errorf("Unable to resolve AMI for region %s and instance type %s", region, instanceType)
	}
}

// GpuResolver resolves the AMI for GPU instances types.
type GpuResolver struct {
}

// Resolve will return an AMI based on the region for GPU instance types
func (r *GpuResolver) Resolve(region string, instanceType string) (string, error) {
	if !strings.HasPrefix(instanceType, "p2") &&
		!strings.HasPrefix(instanceType, "p3") {
		return "", fmt.Errorf("Cannot resolve AMI as the instance type isn'y GPU optimized")
	}

	switch region {
	case "us-west-2":
		return "ami-0d20f2404b9a1c4d1", nil
	case "us-east-1":
		return "ami-09fe6fc9106bda972", nil
	case "eu-west-1":
		return "ami-09e0c6b3d3cf906f1", nil
	default:
		return "", fmt.Errorf("Unable to resolve AMI for region %s and instance type %s", region, instanceType)
	}
}
