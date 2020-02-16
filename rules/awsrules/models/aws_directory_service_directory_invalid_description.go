// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsDirectoryServiceDirectoryInvalidDescriptionRule checks the pattern is valid
type AwsDirectoryServiceDirectoryInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDirectoryServiceDirectoryInvalidDescriptionRule returns new rule with default attributes
func NewAwsDirectoryServiceDirectoryInvalidDescriptionRule() *AwsDirectoryServiceDirectoryInvalidDescriptionRule {
	return &AwsDirectoryServiceDirectoryInvalidDescriptionRule{
		resourceType:  "aws_directory_service_directory",
		attributeName: "description",
		max:           128,
		pattern:       regexp.MustCompile(`^([a-zA-Z0-9_])[\\a-zA-Z0-9_@#%*+=:?./!\s-]*$`),
	}
}

// Name returns the rule name
func (r *AwsDirectoryServiceDirectoryInvalidDescriptionRule) Name() string {
	return "aws_directory_service_directory_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDirectoryServiceDirectoryInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDirectoryServiceDirectoryInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDirectoryServiceDirectoryInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDirectoryServiceDirectoryInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([a-zA-Z0-9_])[\\a-zA-Z0-9_@#%*+=:?./!\s-]*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
