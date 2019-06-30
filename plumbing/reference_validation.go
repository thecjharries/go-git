package plumbing

import (
	"errors"
	"regexp"
)

type ActionChoice int

const (
	Validate ActionChoice = iota
	Skip
	Sanitize
)

var (
	ErrRefLeadsWithPeriod = errors.New("ref name cannot begin with .")
)

var (
	PatternLeadsWithPeriod = regexp.MustCompile(`^\.`)
)

type RefNameChecker struct {
	Name ReferenceName

	CheckRefOptions struct {
		AllowOneLevel  bool
		RefSpecPattern bool
		Normalize      bool
	}

	ActionOptions struct {
		HandleLeadingPeriods ActionChoice
	}

	PostCheckInformation struct {
		HasBeenValidated  bool
		HasBeenSanitizaed bool
	}
}

func (v *RefNameChecker) HandleLeadingPeriods() error {
	switch v.ActionOptions.HandleLeadingPeriods {
	case Validate:
		if PatternLeadsWithPeriod.MatchString(v.Name.String()) {
			return ErrRefLeadsWithPeriod
		}
		break
	case Sanitize:
		v.Name = ReferenceName(PatternLeadsWithPeriod.ReplaceAllString(v.Name.String(), ""))
	}
	return nil
}

func (v *RefNameChecker) HandleTrailingLock() error {
	return nil
}

func (v *RefNameChecker) EnsureAtLeastOneForwardSlash() error {
	return nil
}

func (v *RefNameChecker) HandleDoubleDots() error {
	return nil
}

func (v *RefNameChecker) HandleExcludedCharacters() error {
	return nil
}

func (v *RefNameChecker) HandleLeadingForwardSlash() error {
	return nil
}

func (v *RefNameChecker) HandleTrailingForwardSlash() error {
	return nil
}

func (v *RefNameChecker) HandleConsecutiveForwardSlashes() error {
	return nil
}

func (v *RefNameChecker) HandleTrailingDot() error {
	return nil
}

func (v *RefNameChecker) HandleAtOpenBrace() error {
	return nil
}
