// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSsmAssociationInvalidDocumentVersionRule checks the pattern is valid
type AwsSsmAssociationInvalidDocumentVersionRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsSsmAssociationInvalidDocumentVersionRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidDocumentVersionRule() *AwsSsmAssociationInvalidDocumentVersionRule {
	return &AwsSsmAssociationInvalidDocumentVersionRule{
		resourceType:  "aws_ssm_association",
		attributeName: "document_version",
		pattern:       regexp.MustCompile(`^([$]LATEST|[$]DEFAULT|^[1-9][0-9]*$)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidDocumentVersionRule) Name() string {
	return "aws_ssm_association_invalid_document_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidDocumentVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmAssociationInvalidDocumentVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidDocumentVersionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidDocumentVersionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([$]LATEST|[$]DEFAULT|^[1-9][0-9]*$)$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
