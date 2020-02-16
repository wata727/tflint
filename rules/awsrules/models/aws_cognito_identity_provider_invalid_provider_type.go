// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCognitoIdentityProviderInvalidProviderTypeRule checks the pattern is valid
type AwsCognitoIdentityProviderInvalidProviderTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCognitoIdentityProviderInvalidProviderTypeRule returns new rule with default attributes
func NewAwsCognitoIdentityProviderInvalidProviderTypeRule() *AwsCognitoIdentityProviderInvalidProviderTypeRule {
	return &AwsCognitoIdentityProviderInvalidProviderTypeRule{
		resourceType:  "aws_cognito_identity_provider",
		attributeName: "provider_type",
		enum: []string{
			"SAML",
			"Facebook",
			"Google",
			"LoginWithAmazon",
			"SignInWithApple",
			"OIDC",
		},
	}
}

// Name returns the rule name
func (r *AwsCognitoIdentityProviderInvalidProviderTypeRule) Name() string {
	return "aws_cognito_identity_provider_invalid_provider_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoIdentityProviderInvalidProviderTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoIdentityProviderInvalidProviderTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoIdentityProviderInvalidProviderTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoIdentityProviderInvalidProviderTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as provider_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
