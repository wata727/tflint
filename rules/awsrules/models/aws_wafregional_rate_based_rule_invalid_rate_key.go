// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsWafregionalRateBasedRuleInvalidRateKeyRule checks the pattern is valid
type AwsWafregionalRateBasedRuleInvalidRateKeyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsWafregionalRateBasedRuleInvalidRateKeyRule returns new rule with default attributes
func NewAwsWafregionalRateBasedRuleInvalidRateKeyRule() *AwsWafregionalRateBasedRuleInvalidRateKeyRule {
	return &AwsWafregionalRateBasedRuleInvalidRateKeyRule{
		resourceType:  "aws_wafregional_rate_based_rule",
		attributeName: "rate_key",
		enum: []string{
			"IP",
		},
	}
}

// Name returns the rule name
func (r *AwsWafregionalRateBasedRuleInvalidRateKeyRule) Name() string {
	return "aws_wafregional_rate_based_rule_invalid_rate_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafregionalRateBasedRuleInvalidRateKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafregionalRateBasedRuleInvalidRateKeyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafregionalRateBasedRuleInvalidRateKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafregionalRateBasedRuleInvalidRateKeyRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as rate_key`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
