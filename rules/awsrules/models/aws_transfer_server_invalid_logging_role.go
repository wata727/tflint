// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsTransferServerInvalidLoggingRoleRule checks the pattern is valid
type AwsTransferServerInvalidLoggingRoleRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsTransferServerInvalidLoggingRoleRule returns new rule with default attributes
func NewAwsTransferServerInvalidLoggingRoleRule() *AwsTransferServerInvalidLoggingRoleRule {
	return &AwsTransferServerInvalidLoggingRoleRule{
		resourceType:  "aws_transfer_server",
		attributeName: "logging_role",
		max:           2048,
		min:           20,
		pattern:       regexp.MustCompile(`^arn:.*role/.*$`),
	}
}

// Name returns the rule name
func (r *AwsTransferServerInvalidLoggingRoleRule) Name() string {
	return "aws_transfer_server_invalid_logging_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferServerInvalidLoggingRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferServerInvalidLoggingRoleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferServerInvalidLoggingRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferServerInvalidLoggingRoleRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"logging_role must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"logging_role must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:.*role/.*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
