package cloudformation

// AWSWAFXssMatchSet_FieldToMatch AWS CloudFormation Resource (AWS::WAF::XssMatchSet.FieldToMatch)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html
type AWSWAFXssMatchSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch-data
	Data *StringIntrinsic `json:"Data,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch-type
	Type *StringIntrinsic `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFXssMatchSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAF::XssMatchSet.FieldToMatch"
}
