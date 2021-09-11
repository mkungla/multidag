package multidag

import (
	"regexp"
	"strings"
	"unicode"
)

const (
	// NamespaceMustCompile against following expression.
	NamespaceMustCompile = "^[a-z][a-z0-9_-]*[a-z0-9]$"
)

// Replacement structure.
type replacement struct {
	re *regexp.Regexp
	ch string
}

// Build regexps and replacements.
var (
	alnum = &unicode.RangeTable{
		R16: []unicode.Range16{
			{'0', '9', 1},
			{'a', 'z', 1},
		},
	}
	rExps = []replacement{
		{re: regexp.MustCompile(`[\xC0-\xC6]`), ch: "a"},
		{re: regexp.MustCompile(`[\xE0-\xE6]`), ch: "a"},
		{re: regexp.MustCompile(`[\xC8-\xCB]`), ch: "e"},
		{re: regexp.MustCompile(`[\xE8-\xEB]`), ch: "e"},
		{re: regexp.MustCompile(`[\xCC-\xCF]`), ch: "i"},
		{re: regexp.MustCompile(`[\xEC-\xEF]`), ch: "i"},
		{re: regexp.MustCompile(`[\xD2-\xD6]`), ch: "o"},
		{re: regexp.MustCompile(`[\xF2-\xF6]`), ch: "o"},
		{re: regexp.MustCompile(`[\xD9-\xDC]`), ch: "u"},
		{re: regexp.MustCompile(`[\xF9-\xFC]`), ch: "u"},
		{re: regexp.MustCompile(`[\xC7-\xE7]`), ch: "c"},
		{re: regexp.MustCompile(`[\xD1]`), ch: "n"},
		{re: regexp.MustCompile(`[\xF1]`), ch: "n"},
	}
	spacereg       = regexp.MustCompile(`\s+`)
	noncharreg     = regexp.MustCompile(`[^a-z0-9-]`)
	minusrepeatreg = regexp.MustCompile(`\-{2,}`)
)

// NewNamespace converts provided string into valid namspace
// e.g. Authority namespace
func NewNamespace(s string) string {
	for _, r := range rExps {
		s = r.re.ReplaceAllString(s, r.ch)
	}

	s = strings.ToLower(s)
	s = spacereg.ReplaceAllString(s, "-")
	s = noncharreg.ReplaceAllString(s, "")
	s = minusrepeatreg.ReplaceAllString(s, "-")
	return s
}

// Valid returns true if s is string which is valid name space.
func NamespaceValid(s string) bool {
	re, err := regexp.Compile(NamespaceMustCompile)
	return re.MatchString(s) && err == nil
}
