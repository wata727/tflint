// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsMskClusterInvalidEnhancedMonitoringRule checks the pattern is valid
type AwsMskClusterInvalidEnhancedMonitoringRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsMskClusterInvalidEnhancedMonitoringRule returns new rule with default attributes
func NewAwsMskClusterInvalidEnhancedMonitoringRule() *AwsMskClusterInvalidEnhancedMonitoringRule {
	return &AwsMskClusterInvalidEnhancedMonitoringRule{
		resourceType:  "aws_msk_cluster",
		attributeName: "enhanced_monitoring",
		enum: []string{
			"DEFAULT",
			"PER_BROKER",
			"PER_TOPIC_PER_BROKER",
		},
	}
}

// Name returns the rule name
func (r *AwsMskClusterInvalidEnhancedMonitoringRule) Name() string {
	return "aws_msk_cluster_invalid_enhanced_monitoring"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMskClusterInvalidEnhancedMonitoringRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMskClusterInvalidEnhancedMonitoringRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMskClusterInvalidEnhancedMonitoringRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMskClusterInvalidEnhancedMonitoringRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as enhanced_monitoring`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
