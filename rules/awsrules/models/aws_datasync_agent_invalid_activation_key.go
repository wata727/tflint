// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsDatasyncAgentInvalidActivationKeyRule checks the pattern is valid
type AwsDatasyncAgentInvalidActivationKeyRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncAgentInvalidActivationKeyRule returns new rule with default attributes
func NewAwsDatasyncAgentInvalidActivationKeyRule() *AwsDatasyncAgentInvalidActivationKeyRule {
	return &AwsDatasyncAgentInvalidActivationKeyRule{
		resourceType:  "aws_datasync_agent",
		attributeName: "activation_key",
		max:           29,
		pattern:       regexp.MustCompile(`^[A-Z0-9]{5}(-[A-Z0-9]{5}){4}$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncAgentInvalidActivationKeyRule) Name() string {
	return "aws_datasync_agent_invalid_activation_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncAgentInvalidActivationKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncAgentInvalidActivationKeyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncAgentInvalidActivationKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncAgentInvalidActivationKeyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"activation_key must be 29 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Z0-9]{5}(-[A-Z0-9]{5}){4}$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
