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
	ErrRefLeadingDot = errors.New("ref name cannot begin with .")
)

var (
	PatternLeadingDot = regexp.MustCompile(`^\.`)
)

type RefNameChecker struct {
	Name ReferenceName

	CheckRefOptions struct {
		AllowOneLevel  bool
		RefSpecPattern bool
		Normalize      bool
	}

	ActionOptions struct {
		HandleLeadingDot                ActionChoice
		HandleTrailingLock              ActionChoice
		HandleAtLeastOneForwardSlash    ActionChoice
		HandleDoubleDots                ActionChoice
		HandleExcludedCharacters        ActionChoice
		HandleLeadingForwardSlash       ActionChoice
		HandleTrailingForwardSlash      ActionChoice
		HandleConsecutiveForwardSlashes ActionChoice
		HandleTrailingDot               ActionChoice
		HandleAtOpenBrace               ActionChoice
	}

	PostCheckInformation struct {
		HasBeenValidated  bool
		HasBeenSanitizaed bool
	}
}

func (v *RefNameChecker) HandleLeadingDot() error {
	switch v.ActionOptions.HandleLeadingDot {
	case Validate:
		if PatternLeadingDot.MatchString(v.Name.String()) {
			return ErrRefLeadingDot
		}
		break
	case Sanitize:
		v.Name = ReferenceName(PatternLeadingDot.ReplaceAllString(v.Name.String(), ""))
	}
	return nil
}

func (v *RefNameChecker) HandleTrailingLock() error {
	return nil
}

func (v *RefNameChecker) HandleAtLeastOneForwardSlash() error {
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
