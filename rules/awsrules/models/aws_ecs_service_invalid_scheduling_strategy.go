// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsEcsServiceInvalidSchedulingStrategyRule checks the pattern is valid
type AwsEcsServiceInvalidSchedulingStrategyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEcsServiceInvalidSchedulingStrategyRule returns new rule with default attributes
func NewAwsEcsServiceInvalidSchedulingStrategyRule() *AwsEcsServiceInvalidSchedulingStrategyRule {
	return &AwsEcsServiceInvalidSchedulingStrategyRule{
		resourceType:  "aws_ecs_service",
		attributeName: "scheduling_strategy",
		enum: []string{
			"REPLICA",
			"DAEMON",
		},
	}
}

// Name returns the rule name
func (r *AwsEcsServiceInvalidSchedulingStrategyRule) Name() string {
	return "aws_ecs_service_invalid_scheduling_strategy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcsServiceInvalidSchedulingStrategyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcsServiceInvalidSchedulingStrategyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcsServiceInvalidSchedulingStrategyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcsServiceInvalidSchedulingStrategyRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as scheduling_strategy`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
