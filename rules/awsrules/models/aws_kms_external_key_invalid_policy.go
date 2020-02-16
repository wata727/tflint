// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsKmsExternalKeyInvalidPolicyRule checks the pattern is valid
type AwsKmsExternalKeyInvalidPolicyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsKmsExternalKeyInvalidPolicyRule returns new rule with default attributes
func NewAwsKmsExternalKeyInvalidPolicyRule() *AwsKmsExternalKeyInvalidPolicyRule {
	return &AwsKmsExternalKeyInvalidPolicyRule{
		resourceType:  "aws_kms_external_key",
		attributeName: "policy",
		max:           131072,
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
	}
}

// Name returns the rule name
func (r *AwsKmsExternalKeyInvalidPolicyRule) Name() string {
	return "aws_kms_external_key_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKmsExternalKeyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKmsExternalKeyInvalidPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKmsExternalKeyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKmsExternalKeyInvalidPolicyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"policy must be 131072 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"policy must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
