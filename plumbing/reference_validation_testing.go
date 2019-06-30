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
	TrailingLockNames = []string{
		"a/name.lock",
		"a/name",
	}
	AtLeastOneForwardSlashNames = []string{
		"aname",
		"a/name",
	}
	DoubleDotsNames = []string{
		"a..name",
		"aname",
	}
	ExcludedCharactersNames = []string{
		`an^ame`,
		"aname",
	}
	LeadingForwardSlashNames = []string{
		"/a/name",
		"a/name",
	}
	TrailingForwardSlashNames = []string{
		"a/name/",
		"a/name",
	}
	ConsecutiveForwardSlashesNames = []string{
		"a//name",
		"a/name",
		"a///longer///name",
		"a/longer/name",
	}
	TrailingDotNames = []string{
		"a/name.",
		"a/name",
	}
	AtOpenBraceNames = []string{
		`a/na@{me`,
		`a/name`,
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
func (s *ReferenceValidationSuite) TestValidateHandleTrailingLock(c *C) {
	s.Checker.ActionOptions.HandleTrailingLock = Validate
	s.Checker.Name = ReferenceName(TrailingLockNames[0])
	err := s.Checker.HandleTrailingLock()
	c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefTrailingLock))
	s.Checker.Name = ReferenceName(TrailingLockNames[1])
	err = s.Checker.HandleTrailingLock()
	c.Assert(err, IsNil)
}

func (s *ReferenceValidationSuite) TestSanitizeHandleTrailingLock(c *C) {
	s.Checker.ActionOptions.HandleTrailingLock = Sanitize
	s.Checker.Name = ReferenceName(TrailingLockNames[0])
	err := s.Checker.HandleTrailingLock()
	c.Assert(err, IsNil)
	c.Assert(s.Checker.Name.String(), Equals, TrailingLockNames[1])
}

func (s *ReferenceValidationSuite) TestValidateAtLeastOneForwardSlash(c *C) {
	for _, setting := range []ActionChoice{Validate, Sanitize} {
		s.Checker.CheckRefOptions.AllowOneLevel = false
		s.Checker.ActionOptions.HandleAtLeastOneForwardSlash = setting
		s.Checker.Name = ReferenceName(AtLeastOneForwardSlashNames[0])
		err := s.Checker.HandleAtLeastOneForwardSlash()
		c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefAtLeastOneForwardSlash))
		s.Checker.Name = ReferenceName(AtLeastOneForwardSlashNames[1])
		err = s.Checker.HandleAtLeastOneForwardSlash()
		c.Assert(err, IsNil)
		s.Checker.Name = ReferenceName(AtLeastOneForwardSlashNames[0])
		s.Checker.CheckRefOptions.AllowOneLevel = true
		err = s.Checker.HandleAtLeastOneForwardSlash()
		c.Assert(err, IsNil)
	}
}

func (s *ReferenceValidationSuite) TestValidateHandleDoubleDots(c *C) {
	s.Checker.ActionOptions.HandleDoubleDots = Validate
	s.Checker.Name = ReferenceName(DoubleDotsNames[0])
	err := s.Checker.HandleDoubleDots()
	c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefDoubleDots))
	s.Checker.Name = ReferenceName(DoubleDotsNames[1])
	err = s.Checker.HandleDoubleDots()
	c.Assert(err, IsNil)
}

func (s *ReferenceValidationSuite) TestSanitizeHandleDoubleDots(c *C) {
	s.Checker.ActionOptions.HandleDoubleDots = Sanitize
	s.Checker.Name = ReferenceName(DoubleDotsNames[0])
	err := s.Checker.HandleDoubleDots()
	c.Assert(err, IsNil)
	c.Assert(s.Checker.Name.String(), Equals, DoubleDotsNames[1])
}

func (s *ReferenceValidationSuite) TestValidateHandleExcludedCharacters(c *C) {
	s.Checker.ActionOptions.HandleExcludedCharacters = Validate
	s.Checker.Name = ReferenceName(ExcludedCharactersNames[0])
	err := s.Checker.HandleExcludedCharacters()
	c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefExcludedCharacters))
	s.Checker.Name = ReferenceName(ExcludedCharactersNames[1])
	err = s.Checker.HandleExcludedCharacters()
	c.Assert(err, IsNil)
}

func (s *ReferenceValidationSuite) TestSanitizeHandleExcludedCharacters(c *C) {
	s.Checker.ActionOptions.HandleExcludedCharacters = Sanitize
	s.Checker.Name = ReferenceName(ExcludedCharactersNames[0])
	err := s.Checker.HandleExcludedCharacters()
	c.Assert(err, IsNil)
	c.Assert(s.Checker.Name.String(), Equals, ExcludedCharactersNames[1])
}

func (s *ReferenceValidationSuite) TestValidateHandleLeadingForwardSlash(c *C) {
	s.Checker.ActionOptions.HandleLeadingForwardSlash = Validate
	s.Checker.Name = ReferenceName(LeadingForwardSlashNames[0])
	err := s.Checker.HandleLeadingForwardSlash()
	c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefLeadingForwardSlash))
	s.Checker.Name = ReferenceName(LeadingForwardSlashNames[1])
	err = s.Checker.HandleLeadingForwardSlash()
	c.Assert(err, IsNil)
}

func (s *ReferenceValidationSuite) TestSanitizeHandleLeadingForwardSlash(c *C) {
	s.Checker.ActionOptions.HandleLeadingForwardSlash = Sanitize
	s.Checker.Name = ReferenceName(LeadingForwardSlashNames[0])
	err := s.Checker.HandleLeadingForwardSlash()
	c.Assert(err, IsNil)
	c.Assert(s.Checker.Name.String(), Equals, LeadingForwardSlashNames[1])
}

func (s *ReferenceValidationSuite) TestValidateHandleTrailingForwardSlash(c *C) {
	s.Checker.ActionOptions.HandleTrailingForwardSlash = Validate
	s.Checker.Name = ReferenceName(TrailingForwardSlashNames[0])
	err := s.Checker.HandleTrailingForwardSlash()
	c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefTrailingForwardSlash))
	s.Checker.Name = ReferenceName(TrailingForwardSlashNames[1])
	err = s.Checker.HandleTrailingForwardSlash()
	c.Assert(err, IsNil)
}

func (s *ReferenceValidationSuite) TestSanitizeHandleTrailingForwardSlash(c *C) {
	s.Checker.ActionOptions.HandleTrailingForwardSlash = Sanitize
	s.Checker.Name = ReferenceName(TrailingForwardSlashNames[0])
	err := s.Checker.HandleTrailingForwardSlash()
	c.Assert(err, IsNil)
	c.Assert(s.Checker.Name.String(), Equals, TrailingForwardSlashNames[1])
}

func (s *ReferenceValidationSuite) TestValidateHandleConsecutiveForwardSlashes(c *C) {
	s.Checker.ActionOptions.HandleConsecutiveForwardSlashes = Validate
	for index, name := range ConsecutiveForwardSlashesNames {
		s.Checker.Name = ReferenceName(name)
		err := s.Checker.HandleConsecutiveForwardSlashes()
		if 1 == index%2 {
			c.Assert(err, IsNil)
		} else {
			c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefConsecutiveForwardSlashes))
		}

	}
}

func (s *ReferenceValidationSuite) TestSanitizeHandleConsecutiveForwardSlashes(c *C) {
	s.Checker.CheckRefOptions.Normalize = true
	for _, element := range []int{0, 2} {
		s.Checker.ActionOptions.HandleConsecutiveForwardSlashes = Sanitize
		s.Checker.Name = ReferenceName(ConsecutiveForwardSlashesNames[element+0])
		err := s.Checker.HandleConsecutiveForwardSlashes()
		c.Assert(err, IsNil)
		c.Assert(s.Checker.Name.String(), Equals, ConsecutiveForwardSlashesNames[element+1])
	}
}

// func (s *ReferenceValidationSuite) TestValidateHandleTrailingDot(c *C) {
// 	s.Checker.ActionOptions.HandleTrailingDot = Validate
// 	s.Checker.Name = ReferenceName(TrailingDotNames[0])
// 	err := s.Checker.HandleTrailingDot()
// 	c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefTrailingDot))
// 	s.Checker.Name = ReferenceName(TrailingDotNames[1])
// 	err = s.Checker.HandleTrailingDot()
// 	c.Assert(err, IsNil)
// }

// func (s *ReferenceValidationSuite) TestSanitizeHandleTrailingDot(c *C) {
// 	s.Checker.ActionOptions.HandleTrailingDot = Sanitize
// 	s.Checker.Name = ReferenceName(TrailingDotNames[0])
// 	err := s.Checker.HandleTrailingDot()
// 	c.Assert(err, IsNil)
// 	c.Assert(s.Checker.Name.String(), Equals, TrailingDotNames[1])
// }

// func (s *ReferenceValidationSuite) TestValidateHandleAtOpenBrace(c *C) {
// 	s.Checker.ActionOptions.HandleAtOpenBrace = Validate
// 	s.Checker.Name = ReferenceName(AtOpenBraceNames[0])
// 	err := s.Checker.HandleAtOpenBrace()
// 	c.Assert(err, ErrorMatches, fmt.Sprint(ErrRefAtOpenBrace))
// 	s.Checker.Name = ReferenceName(AtOpenBraceNames[1])
// 	err = s.Checker.HandleAtOpenBrace()
// 	c.Assert(err, IsNil)
// }

// func (s *ReferenceValidationSuite) TestSanitizeHandleAtOpenBrace(c *C) {
// 	s.Checker.ActionOptions.HandleAtOpenBrace = Sanitize
// 	s.Checker.Name = ReferenceName(AtOpenBraceNames[0])
// 	err := s.Checker.HandleAtOpenBrace()
// 	c.Assert(err, IsNil)
// 	c.Assert(s.Checker.Name.String(), Equals, AtOpenBraceNames[1])
// }
