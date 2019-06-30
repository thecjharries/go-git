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
	LeadingDotNames = []string{
		".a/name",
		"a/name",
	}
)

func (s *ReferenceValidationSuite) TestValidateHandleLeadingDot(c *C) {
	s.Checker.ActionOptions.HandleLeadingDot = Validate
	s.Checker.Name = ReferenceName(LeadingDotNames[0])
	err := s.Checker.HandleLeadingDot()
	c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefLeadingDot))
	s.Checker.Name = ReferenceName(LeadingDotNames[1])
	err = s.Checker.HandleLeadingDot()
	c.Assert(err, IsNil)
}

func (s *ReferenceValidationSuite) TestSanitizeHandleLeadingDot(c *C) {
	s.Checker.ActionOptions.HandleLeadingDot = Sanitize
	s.Checker.Name = ReferenceName(LeadingDotNames[0])
	err := s.Checker.HandleLeadingDot()
	c.Assert(err, IsNil)
	c.Assert(s.Checker.Name.String(), Equals, LeadingDotNames[1])
}
