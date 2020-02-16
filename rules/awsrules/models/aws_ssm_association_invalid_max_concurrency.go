// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSsmAssociationInvalidMaxConcurrencyRule checks the pattern is valid
type AwsSsmAssociationInvalidMaxConcurrencyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsmAssociationInvalidMaxConcurrencyRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidMaxConcurrencyRule() *AwsSsmAssociationInvalidMaxConcurrencyRule {
	return &AwsSsmAssociationInvalidMaxConcurrencyRule{
		resourceType:  "aws_ssm_association",
		attributeName: "max_concurrency",
		max:           7,
		min:           1,
		pattern:       regexp.MustCompile(`^([1-9][0-9]*|[1-9][0-9]%|[1-9]%|100%)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidMaxConcurrencyRule) Name() string {
	return "aws_ssm_association_invalid_max_concurrency"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidMaxConcurrencyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmAssociationInvalidMaxConcurrencyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidMaxConcurrencyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidMaxConcurrencyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"max_concurrency must be 7 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"max_concurrency must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([1-9][0-9]*|[1-9][0-9]%|[1-9]%|100%)$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
