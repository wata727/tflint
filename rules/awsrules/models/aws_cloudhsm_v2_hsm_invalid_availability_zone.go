// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCloudhsmV2HsmInvalidAvailabilityZoneRule checks the pattern is valid
type AwsCloudhsmV2HsmInvalidAvailabilityZoneRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloudhsmV2HsmInvalidAvailabilityZoneRule returns new rule with default attributes
func NewAwsCloudhsmV2HsmInvalidAvailabilityZoneRule() *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule {
	return &AwsCloudhsmV2HsmInvalidAvailabilityZoneRule{
		resourceType:  "aws_cloudhsm_v2_hsm",
		attributeName: "availability_zone",
		pattern:       regexp.MustCompile(`^[a-z]{2}(-(gov))?-(east|west|north|south|central){1,2}-\d[a-z]$`),
	}
}

// Name returns the rule name
func (r *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule) Name() string {
	return "aws_cloudhsm_v2_hsm_invalid_availability_zone"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z]{2}(-(gov))?-(east|west|north|south|central){1,2}-\d[a-z]$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
