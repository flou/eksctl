package builder

import (
	gfn "github.com/awslabs/goformation/cloudformation"
)

const (
	cfnOutputNodeInstanceRoleARN = "NodeInstanceRoleARN"

	iamPolicyAmazonEKSServicePolicyARN = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
	iamPolicyAmazonEKSClusterPolicyARN = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"

	iamPolicyAmazonEKSWorkerNodePolicyARN           = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
	iamPolicyAmazonEKSCNIPolicyARN                  = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
	iamPolicyAmazonEC2ContainerRegistryPowerUserARN = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryPowerUser"
	iamPolicyAmazonEC2ContainerRegistryReadOnlyARN  = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
)

var (
	iamDefaultNodePolicyARNs = []string{
		iamPolicyAmazonEKSWorkerNodePolicyARN,
		iamPolicyAmazonEKSCNIPolicyARN,
	}
)

func makePolicyDocument(statement map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []interface{}{
			statement,
		},
	}
}

func makeAssumeRolePolicyDocument(service string) map[string]interface{} {
	return makePolicyDocument(map[string]interface{}{
		"Effect": "Allow",
		"Principal": map[string][]string{
			"Service": []string{service},
		},
		"Action": []string{"sts:AssumeRole"},
	})
}

func (c *resourceSet) attachAllowPolicy(name string, refRole gfn.Intrinsic, resources interface{}, actions []string) {
	c.newResource(name, &gfn.AWSIAMPolicy{
		PolicyName: makeName(name),
		Roles:      makeSlice(refRole),
		PolicyDocument: makePolicyDocument(map[string]interface{}{
			"Effect":   "Allow",
			"Resource": resources,
			"Action":   actions,
		}),
	})
}

func (c *clusterResourceSet) WithIAM() bool {
	return c.rs.withIAM
}

func (c *clusterResourceSet) addResourcesForIAM() {
	c.rs.withIAM = true

	refSR := c.newResource("ServiceRole", &gfn.AWSIAMRole{
		AssumeRolePolicyDocument: makeAssumeRolePolicyDocument("eks.amazonaws.com"),
		ManagedPolicyArns: makeStringSlice(
			iamPolicyAmazonEKSServicePolicyARN,
			iamPolicyAmazonEKSClusterPolicyARN,
		),
	})
	c.rs.attachAllowPolicy("PolicyNLB", refSR, "*", []string{
		"elasticloadbalancing:*",
		"ec2:CreateSecurityGroup",
		"ec2:Describe*",
	})
}

func (n *nodeGroupResourceSet) WithIAM() bool {
	return n.rs.withIAM
}

func (n *nodeGroupResourceSet) addResourcesForIAM() {
	n.rs.withIAM = true

	if len(n.spec.NodePolicyARNs) == 0 {
		n.spec.NodePolicyARNs = iamDefaultNodePolicyARNs
	}
	if n.spec.Addons.WithIAM.PolicyAmazonEC2ContainerRegistryPowerUser {
		n.spec.NodePolicyARNs = append(n.spec.NodePolicyARNs, iamPolicyAmazonEC2ContainerRegistryReadOnlyARN)
	} else {
		n.spec.NodePolicyARNs = append(n.spec.NodePolicyARNs, iamPolicyAmazonEC2ContainerRegistryPowerUserARN)
	}

	refIR := n.newResource("NodeInstanceRole", &gfn.AWSIAMRole{
		Path: gfn.String("/"),
		AssumeRolePolicyDocument: makeAssumeRolePolicyDocument("ec2.amazonaws.com"),
		ManagedPolicyArns:        makeStringSlice(n.spec.NodePolicyARNs...),
	})

	n.instanceProfile = n.newResource("NodeInstanceProfile", &gfn.AWSIAMInstanceProfile{
		Path:  gfn.String("/"),
		Roles: makeSlice(refIR),
	})
	n.rs.attachAllowPolicy("PolicyTagDiscovery", refIR, "*", []string{
		"ec2:DescribeTags",
	})
	n.rs.attachAllowPolicy("PolicyStackSignal", refIR,
		map[string]interface{}{
			gfn.FnJoin: []interface{}{
				":",
				[]interface{}{
					"arn:aws:cloudformation",
					map[string]string{"Ref": gfn.Region},
					map[string]string{"Ref": gfn.AccountID},
					map[string]interface{}{
						gfn.FnJoin: []interface{}{
							"/",
							[]interface{}{
								"stack",
								map[string]string{"Ref": gfn.StackName},
								"*",
							},
						},
					},
				},
			},
		},
		[]string{
			"cloudformation:SignalResource",
		},
	)

	n.rs.newOutputFromAtt(cfnOutputNodeInstanceRoleARN, "NodeInstanceRole.Arn", true)
}
