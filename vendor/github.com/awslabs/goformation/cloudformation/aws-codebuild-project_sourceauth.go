package cloudformation

// AWSCodeBuildProject_SourceAuth AWS CloudFormation Resource (AWS::CodeBuild::Project.SourceAuth)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html
type AWSCodeBuildProject_SourceAuth struct {

	// Resource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html#cfn-codebuild-project-sourceauth-resource
	Resource *StringIntrinsic `json:"Resource,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html#cfn-codebuild-project-sourceauth-type
	Type *StringIntrinsic `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_SourceAuth) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.SourceAuth"
}
