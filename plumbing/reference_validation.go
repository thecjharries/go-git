package plumbing

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type ActionChoice int

const (
	Skip ActionChoice = iota
	Sanitize
	Validate
)

var (
	ErrRefLeadingDot                = errors.New("ref name cannot begin with a dot")
	ErrRefTrailingLock              = errors.New("ref name cannot end with .lock")
	ErrRefAtLeastOneForwardSlash    = errors.New("ref name must have at least one forward slash")
	ErrRefDoubleDots                = errors.New("ref name cannot include two consecutive dots")
	ErrRefExcludedCharacters        = errors.New("ref name cannot include many special characters")
	ErrRefLeadingForwardSlash       = errors.New("ref name cannot start with a forward slash")
	ErrRefTrailingForwardSlash      = errors.New("ref name cannot end with a forward slash")
	ErrRefConsecutiveForwardSlashes = errors.New("ref name cannot have consectutive forward slashes")
	ErrRefTrailingDot               = errors.New("ref name cannot end with a dot")
	ErrRefAtOpenBrace               = errors.New("ref name cannot include at-open-brace")
)

var (
	PatternLeadingDot                  = regexp.MustCompile(`^\.`)
	PatternTrailingLock                = regexp.MustCompile(`\.lock$`)
	PatternAtLeastOneForwardSlash      = regexp.MustCompile(`^[^/]+$`)
	PatternDoubleDots                  = regexp.MustCompile(`\.\.`)
	PatternExcludedCharacters          = regexp.MustCompile(`[\000-\037\177 ~^:?*[]+`)
	PatternLeadingForwardSlash         = regexp.MustCompile(`^/`)
	PatternTrailingForwardSlash        = regexp.MustCompile(`/$`)
	PatternConsecutiveForwardSlashes   = regexp.MustCompile(`//+`)
	PatternTrailingDot                 = regexp.MustCompile(`\.$`)
	PatternAtOpenBrace                 = regexp.MustCompile(`@{`)
	PatternExcludedCharactersAlternate = regexp.MustCompile(`[\000-\037\177 ~^:?[]+`)
	PatternOneAllowedAsterisk          = regexp.MustCompile(`^[^*]+?\*?[^*]+?$`)
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
}

type RefNameCheckerHandler func()error
type RefNameCheckerHanders []RefNameCheckerHandler


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
	switch v.ActionOptions.HandleTrailingLock {
	case Validate:
		if PatternTrailingLock.MatchString(v.Name.String()) {
			return ErrRefTrailingLock
		}
		break
	case Sanitize:
		v.Name = ReferenceName(PatternTrailingLock.ReplaceAllString(v.Name.String(), ""))
	}
	return nil
}

func (v *RefNameChecker) HandleAtLeastOneForwardSlash() error {
	if Skip == v.ActionOptions.HandleAtLeastOneForwardSlash {
		return nil
	}
	count := strings.Count(v.Name.String(), "/")
	if 1 > count {
		if v.CheckRefOptions.AllowOneLevel {
			return nil
		}
		return ErrRefAtLeastOneForwardSlash
	}
	return nil
}

func (v *RefNameChecker) HandleDoubleDots() error {
	switch v.ActionOptions.HandleDoubleDots {
	case Validate:
		if PatternDoubleDots.MatchString(v.Name.String()) {
			return ErrRefDoubleDots
		}
		break
	case Sanitize:
		v.Name = ReferenceName(PatternDoubleDots.ReplaceAllString(v.Name.String(), ""))
	}
	return nil
}

func (v *RefNameChecker) HandleExcludedCharacters() error {
	switch v.ActionOptions.HandleExcludedCharacters {
	case Validate:
		if PatternExcludedCharacters.MatchString(v.Name.String()) {
			return ErrRefExcludedCharacters
		}
		break
	case Sanitize:
		if v.CheckRefOptions.RefSpecPattern && PatternOneAllowedAsterisk.MatchString(v.Name.String()) {
			v.Name = ReferenceName(PatternExcludedCharactersAlternate.ReplaceAllString(v.Name.String(), ""))

		} else {
			v.Name = ReferenceName(PatternExcludedCharacters.ReplaceAllString(v.Name.String(), ""))
		}
	}
	return nil
}

func (v *RefNameChecker) HandleLeadingForwardSlash() error {
	switch v.ActionOptions.HandleLeadingForwardSlash {
	case Validate:
		if PatternLeadingForwardSlash.MatchString(v.Name.String()) {
			return ErrRefLeadingForwardSlash
		}
		break
	case Sanitize:
		v.Name = ReferenceName(PatternLeadingForwardSlash.ReplaceAllString(v.Name.String(), ""))
	}
	return nil
}

func (v *RefNameChecker) HandleTrailingForwardSlash() error {
	switch v.ActionOptions.HandleTrailingForwardSlash {
	case Validate:
		if PatternTrailingForwardSlash.MatchString(v.Name.String()) {
			return ErrRefTrailingForwardSlash
		}
		break
	case Sanitize:
		v.Name = ReferenceName(PatternTrailingForwardSlash.ReplaceAllString(v.Name.String(), ""))
	}
	return nil
}

func (v *RefNameChecker) HandleConsecutiveForwardSlashes() error {
	if Skip == v.ActionOptions.HandleConsecutiveForwardSlashes {
		return nil
	}
	if PatternConsecutiveForwardSlashes.MatchString(v.Name.String()) {
		if Sanitize == v.ActionOptions.HandleConsecutiveForwardSlashes {
			if v.CheckRefOptions.Normalize {
				v.Name = ReferenceName(PatternConsecutiveForwardSlashes.ReplaceAllString(v.Name.String(), "/"))
				return nil
			}
		}
		return ErrRefConsecutiveForwardSlashes
	}
	return nil
}

func (v *RefNameChecker) HandleTrailingDot() error {
	switch v.ActionOptions.HandleTrailingDot {
	case Validate:
		if PatternTrailingDot.MatchString(v.Name.String()) {
			return ErrRefTrailingDot
		}
		break
	case Sanitize:
		v.Name = ReferenceName(PatternTrailingDot.ReplaceAllString(v.Name.String(), ""))
	}
	return nil
}

func (v *RefNameChecker) HandleAtOpenBrace() error {
	switch v.ActionOptions.HandleAtOpenBrace {
	case Validate:
		if PatternAtOpenBrace.MatchString(v.Name.String()) {
			return ErrRefAtOpenBrace
		}
		break
	case Sanitize:
		v.Name = ReferenceName(PatternAtOpenBrace.ReplaceAllString(v.Name.String(), ""))
	}
	return nil
}

func (v *RefNameChecker) BuildHandleList() RefNameCheckerHanders {
	var handles []func()error
	handles = append(handles, v.HandleLeadingDot)
handles = append(handles, v.HandleTrailingLock)
handles = append(handles, v.HandleAtLeastOneForwardSlash)
handles = append(handles, v.HandleDoubleDots)
handles = append(handles, v.HandleExcludedCharacters)
handles = append(handles, v.HandleLeadingForwardSlash)
handles = append(handles, v.HandleTrailingForwardSlash)
handles = append(handles, v.HandleConsecutiveForwardSlashes)
handles = append(handles, v.HandleTrailingDot)
handles = append(handles, v.HandleAtOpenBrace)
for _, handle := range handles {
	handle()
}
return nil
}



func (v *RefNameChecker) CheckRefName() error {
	return nil
}

func (v *RefNameChecker) CheckRefNameReflect() error {
	v.ActionOptions.HandleLeadingDot = Validate
	actions := reflect.ValueOf(v.ActionOptions)
	for index := 0; index < actions.NumField(); index++ {
		element := actions.Field(index)
		if int64(Skip) != element.Int() {
			func_name := actions.Type().Field(index).Name
			function := reflect.ValueOf(v).MethodByName(func_name)
			methodType := function.Type()
			numIn := methodType.NumIn()
			in := make([]reflect.Value,numIn)
			err := function.Call(in)[0]
			fmt.Println(err)
		}
	}
	return nil
}
