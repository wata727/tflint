// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCodepipelineInvalidNameRule checks the pattern is valid
type AwsCodepipelineInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodepipelineInvalidNameRule returns new rule with default attributes
func NewAwsCodepipelineInvalidNameRule() *AwsCodepipelineInvalidNameRule {
	return &AwsCodepipelineInvalidNameRule{
		resourceType:  "aws_codepipeline",
		attributeName: "name",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^[A-Za-z0-9.@\-_]+$`),
	}
}

// Name returns the rule name
func (r *AwsCodepipelineInvalidNameRule) Name() string {
	return "aws_codepipeline_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodepipelineInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodepipelineInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodepipelineInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodepipelineInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9.@\-_]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
