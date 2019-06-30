package plumbing

import (
	"fmt"
	. "gopkg.in/check.v1"
)

type ReferenceValidationSuite struct {
	Checker RefNameChecker
}

var _ = Suite(&ReferenceValidationSuite{})

var (
	LeadingPeriodNames = []string{
		".a/name",
		"a/name",
	}
)

func (s *ReferenceValidationSuite) TestValidateHandleLeadingPeriods(c *C) {
	s.Checker.ActionOptions.HandleLeadingPeriods = Validate
	s.Checker.Name = ReferenceName(LeadingPeriodNames[0])
	err := s.Checker.HandleLeadingPeriods()
	c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefLeadsWithPeriod))
	s.Checker.Name = ReferenceName(LeadingPeriodNames[1])
	err = s.Checker.HandleLeadingPeriods()
	c.Assert(err, IsNil)
}

func (s *ReferenceValidationSuite) TestSanitizeHandleLeadingPeriods(c *C) {
	s.Checker.ActionOptions.HandleLeadingPeriods = Sanitize
	s.Checker.Name = ReferenceName(LeadingPeriodNames[0])
	err := s.Checker.HandleLeadingPeriods()
	c.Assert(err, IsNil)
	c.Assert(s.Checker.Name.String(), Equals, LeadingPeriodNames[1])
}
